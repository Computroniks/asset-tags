// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

package util

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// Get the specified environment variable. If it doesn't exist, return
// the fallback instead
func Getenv(key string, fallback string) string {
	val := os.Getenv(key)

	if len(val) == 0 {
		return fallback
	} else {
		return val
	}
}

// Attempt to get the environment variable. If it is not set, log error
// and exit
func Mustgetenv(key string) string {
	val := os.Getenv(key)

	if len(val) == 0 {
		log.Fatalln("Failed to get", key, ". Environment variable not set")
	}
	return val
}

func FormatTag(prefix string, count int) string {
	unique := strconv.Itoa(count)
	padding := TagLength - len(unique)
	if padding <= 0 {
		return prefix + unique
	} else {
		return prefix + strings.Repeat("0", padding) + unique
	}
}

// Linear search of array
func SInArray(arr []string, search string) bool {
    for _, x := range arr {
        if x == search {
            return true
        }
    }
	return false
}