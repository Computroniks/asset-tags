// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

package mysql

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"

	"github.com/Computroniks/asset-tags/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qustavo/dotsql"
)

//go:embed schema.sql
var schema string

var patches = []string{
	"create-patches",
	"create-prefixes",
}

type MySQLDB struct {
	conn *sql.DB
	dbPath string
}

func New(addr string, uname string, pwd string, dbname string, timeout string) (*MySQLDB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=%ss", uname, pwd, addr, dbname, timeout)
	conn, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to database:", addr)

	db := MySQLDB {
		conn: conn,
		dbPath: dbname,
	}

	err = db.patch()
	if err != nil {
		log.Fatalln(err)
	}

	return &db, nil
}

// Apply required patches to database
func (o *MySQLDB) patch() error {
	log.Println("Checking if database needs patching")
	var count int
	err := o.conn.QueryRow(
		"SELECT COUNT(*) FROM information_schema.tables WHERE table_schema=? AND table_name='patches' LIMIT 1;",
		util.DatabaseName,
	).Scan(&count)
	if err != nil {
		return err
	}
	var appliedPatches []string
	if count < 1 {
		// New database. Need to set up everything
		log.Println("Database has not been configured. Configuring now")
	} else {
		rows, err := o.conn.Query("SELECT name FROM patches;")

		if err != nil {
			return err
		}

		for rows.Next() {
			var patch string
			rows.Scan(&patch)
			appliedPatches = append(appliedPatches, patch)
		}
	}

	dot, err := dotsql.LoadFromString(schema)

	if err != nil {
		return err
	}

	for _, patch := range patches {
		if util.SInArray(appliedPatches, patch) {
			log.Printf("%s already applied. Skipping", patch)
		} else {
			log.Printf("Applying patch %s to database.", patch)
			_, err := dot.Exec(o.conn, patch)
			if err != nil {
				return err
			}

			_, err = o.conn.Exec("INSERT INTO patches(name) VALUES(?);", patch)
			if err != nil {
				log.Printf("Failed to insert %s into patches table. Please insert manually.", patch)
				return err
			}
		}
	}
	
	return nil
}

func (o *MySQLDB) getCount(prefix string) (int, error) {
	row := o.conn.QueryRow("SELECT count FROM prefixes WHERE prefix = ?;", prefix)
	var count int
	err := row.Scan(&count)

	if err != nil {
		return 0, err
	}
	return count, nil
}

// Get the next tag for the specified prefix
func (o *MySQLDB) GetTag(prefix string) (string, error) {
	count, err := o.getCount(prefix)
	if err != nil {
		return "", err
	}
	return util.FormatTag(prefix, count), nil
}

// Get all the prefixes
func (o *MySQLDB) GetPrefixes() ([]string, error) {
	rows, err:= o.conn.Query("SELECT prefix FROM prefixes;")
	if err != nil{
		return nil, err
	}

	var prefixes []string
	for rows.Next() {
		var prefix string
		err := rows.Scan(&prefix)
		if err != nil {
			return prefixes, err
		}
		prefixes = append(prefixes, prefix)
	}

	if err = rows.Err(); err != nil {
		return prefixes, err
	}

	return prefixes, nil
}

// Increment the count of the specified tag
func (o *MySQLDB) IncrementTag(prefix string) error {
	count, err := o.getCount(prefix)
	if err != nil {
		return err
	}

	count ++

	_, err = o.conn.Exec("UPDATE prefixes SET count=? WHERE prefix=?", count, prefix)

	// Error should already be nil so no need to verify
	return err
}

func (o *MySQLDB) AddPrefix(prefix string) error {
	_, err := o.conn.Exec("INSERT INTO prefixes(prefix) VALUES(?)", prefix)

	// Error should already be nil so no need to verify
	return err
}

// Close connection to database
func (o *MySQLDB) Close() {
	o.conn.Close()
}
