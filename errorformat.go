// Package errorformat provides 'errorformat' functionality of Vim. :h
// errorformat
package errorformat

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

// Efm represents a errorformat.
type Efm struct {
	regex *regexp.Regexp

	flagplus  bool
	flagminus bool
	prefix    byte
}

var fmtpattern = map[byte]string{
	'f': `(?P<f>.+)`, // only used when at end
	'n': `(?P<n>\d+)`,
	'l': `(?P<l>\d+)`,
	'c': `(?P<c>\d+)`,
	't': `(?P<t>.)`,
	'm': `(?P<m>.+)`,
	'r': `(?P<r>.*)`,
	'p': `(?P<p>[- 	.]*)`,
	'v': `(?P<v>\d+)`,
	's': `(?P<s>.+)`,
}

// NewEfm converts a 'errorformat' string to regular expression pattern with
// flags and returns Efm.
//
// quickfix.c: efm_to_regpat
func NewEfm(errorformat string) (*Efm, error) {
	var regpat bytes.Buffer
	var efmp byte
	var i = 0
	var incefmp = func() {
		i++
		efmp = errorformat[i]
	}
	efm := &Efm{}
	regpat.WriteRune('^')
	for ; i < len(errorformat); i++ {
		efmp = errorformat[i]
		if efmp == '%' {
			incefmp()
			// - do not support scanf()-like "%*[]" notation
			// - do not support %>
			if re, ok := fmtpattern[efmp]; ok {
				regpat.WriteString(re)
			} else if strchar(`%\.^$`, efmp) {
				// regexp magic characters
				// XXX: ? -> ~[
				regpat.WriteByte(efmp)
			} else if efmp == '#' {
				regpat.WriteRune('*')
			} else {
				if efmp == '+' {
					efm.flagplus = true
					incefmp()
				} else if efmp == '-' {
					efm.flagminus = true
					incefmp()
				}
				if strchar("DXAEWICZGOPQ", efmp) {
					efm.prefix = efmp
				} else {
					return nil, fmt.Errorf("E376: Invalid %%%v in format string prefix", efmp)
				}
			}
		} else { // copy normal character
			regpat.WriteString(regexp.QuoteMeta(string(efmp)))
		}
	}
	regpat.WriteRune('$')
	re, err := regexp.Compile(regpat.String())
	if err != nil {
		return nil, err
	}
	efm.regex = re
	return efm, nil
}

// Match represents match of Efm. ref: Basic items in :h errorformat
type Match struct {
	F string // (%f) file name
	N int    // (%n) error number
	L int    // (%l) line number
	C int    // (%c) column number
	T string // (%t) error type
	M string // (%m) error message
	R string // (%r) the "rest" of a single-line file message
	P string // (%p) pointer line
	V int    // (%v) virtual column number
	S string // (%s) search text
}

// Match returns match against given string.
func (efm *Efm) Match(s string) *Match {
	ms := efm.regex.FindStringSubmatch(s)
	if len(ms) == 0 {
		return nil
	}
	match := &Match{}
	names := efm.regex.SubexpNames()
	for i, name := range names {
		if i == 0 {
			continue
		}
		m := ms[i]
		switch name {
		case "f":
			match.F = m
		case "n":
			match.N = mustAtoI(m)
		case "l":
			match.L = mustAtoI(m)
		case "c":
			match.C = mustAtoI(m)
		case "t":
			match.T = m
		case "m":
			match.M = m
		case "r":
			match.R = m
		case "p":
			match.P = m
		case "v":
			match.V = mustAtoI(m)
		case "s":
			match.S = m
		}
	}
	return match
}

func strchar(chars string, c byte) bool {
	return bytes.IndexAny([]byte{c}, chars) != -1
}

func mustAtoI(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
