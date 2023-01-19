// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Computroniks/asset-tags/util"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	conn *sql.DB
	dbPath string
}

func New(addr string, uname string, pwd string, dbname string) (*MySQLDB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s", uname, pwd, addr, dbname)
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

	return &db, nil
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

	_, err = o.conn.Exec("UPDATE prefix SET count=? WHERE prefix=?", count, prefix)

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
