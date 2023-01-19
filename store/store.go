// SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
// SPDX-License-Identifier: MIT

package store

type Store interface {
	GetTag(prefix string) (string, error)
	GetPrefixes() ([]string, error)
	IncrementTag(prefix string) error
	AddPrefix(prefix string) error
	Close()
}