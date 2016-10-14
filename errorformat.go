// Package errorformat provides 'errorformat' functionality of Vim. :h
// errorformat
package errorformat

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Errorformat provides errorformat feature.
type Errorformat struct {
	Efms []*Efm
}

type Scanner struct {
	*Errorformat
	source *bufio.Scanner

	qi *qfinfo
}

func NewErrorformat(efms []string) (*Errorformat, error) {
	errorformat := &Errorformat{Efms: make([]*Efm, 0, len(efms))}
	for _, efm := range efms {
		e, err := NewEfm(efm)
		if err != nil {
			return nil, err
		}
		errorformat.Efms = append(errorformat.Efms, e)
	}
	return errorformat, nil
}

func (errorformat *Errorformat) NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		Errorformat: errorformat,
		source:      bufio.NewScanner(r),
		qi:          &qfinfo{},
	}
}

type qfinfo struct {
	filestack   []string
	currfile    string
	dirstack    []string
	directory   string
	multiscan   bool
	multiline   bool
	multiignore bool

	qflist []*qfline
}

type qffields struct {
	namebuf   string
	errmsg    string
	lnum      int
	col       int
	useviscol bool
	pattern   string
	enr       int
	etype     byte
	valid     bool
}

type qfline struct {
	lnum    int
	col     int
	nr      int
	pattern string
	text    string
	viscol  bool
	etype   byte
	valid   bool
}

// Scan scans reader and returns next match.
// It returns nil if next match doesn't exist.
// func (s *Scanner) Scan() (*Match, error) {
func (s *Scanner) Scan() (bool, *qfline, error) {
	for s.source.Scan() {
		line := s.source.Text()
		status, fields, err := s.parseLine(line)
		switch status {
		case qffail:
			continue
			_ = err
			// return true, nil, err
		case qfendmultiline:
			fmt.Printf("%#v\n", s.qi.qflist[len(s.qi.qflist)-1])
			return true, s.qi.qflist[len(s.qi.qflist)-1], nil
		case qfignoreline:
			continue
		}
		qfl := &qfline{
			lnum:    fields.lnum,
			col:     fields.col,
			nr:      fields.enr,
			pattern: fields.pattern,
			text:    fields.errmsg,
			viscol:  fields.useviscol,
			valid:   fields.valid,
		}
		s.qi.qflist = append(s.qi.qflist, qfl)
		if s.qi.multiline {
			continue
		}
		return true, qfl, nil
	}
	return false, nil, nil
}

type qfstatus int

const (
	qffail qfstatus = iota
	qfignoreline
	qfendmultiline
	qfok
)

func (s *Scanner) parseLine(line string) (qfstatus, *qffields, error) {
	return s.parseLineInternal(line, 0)
}

func (s *Scanner) parseLineInternal(line string, i int) (qfstatus, *qffields, error) {
	fields := &qffields{valid: true, enr: -1}
	tail := ""
	var idx byte
	nomatch := false
	var efm *Efm
	for ; i <= len(s.Efms); i++ {
		if i == len(s.Efms) {
			nomatch = true
			break
		}
		efm = s.Efms[i]

		idx = efm.prefix
		if s.qi.multiscan && strchar("OPQ", idx) {
			continue
		}

		if (idx == 'C' || idx == 'Z') && !s.qi.multiline {
			continue
		}

		r := efm.Match(line)
		if r == nil {
			continue
		}

		if strchar("EWI", idx) {
			fields.etype = idx
		}

		if r.F != "" { // %f
			fields.namebuf = r.F
			if strchar("OPQ", idx) && !fileexists(fields.namebuf) {
				continue
			}
		}
		fields.enr = r.N                     // %n
		fields.lnum = r.L                    // %l
		fields.col = r.C                     // %c
		fields.etype = r.T                   // %t
		if efm.flagplus && !s.qi.multiscan { // %+
			fields.errmsg = line
		} else if r.M != "" {
			fields.errmsg = r.M
		}
		tail = r.R     // %r
		if r.P != "" { // %p
			fields.useviscol = true
			fields.col = 0
			for _, m := range r.P {
				fields.col++
				if m == '\t' {
					fields.col += 7
					fields.col -= fields.col % 8
				}
			}
			fields.col++ // last pointer (e.g. ^)
		}
		if r.V != 0 {
			fields.useviscol = true
			fields.col = r.V
		}
		if r.S != "" {
			fields.pattern = fmt.Sprintf("^%v$", regexp.QuoteMeta(r.S))
		}
		break
	}
	s.qi.multiscan = false
	if nomatch || idx == 'D' || idx == 'X' {
		if !nomatch {
			if idx == 'D' {
				if fields.namebuf == "" {
					return qffail, nil, errors.New("E379: Missing or empty directory name")
				}
				s.qi.directory = fields.namebuf
				s.qi.dirstack = append(s.qi.dirstack, s.qi.directory)
			} else if idx == 'X' && len(s.qi.dirstack) > 0 {
				s.qi.directory = s.qi.dirstack[len(s.qi.dirstack)-1]
				s.qi.dirstack = s.qi.dirstack[:len(s.qi.dirstack)-1]
			}
		}
		fields.namebuf = ""
		fields.lnum = 0
		fields.valid = false
		fields.errmsg = line
		if nomatch {
			s.qi.multiline = false
			s.qi.multiignore = false
		}
	} else if !nomatch {
		if strchar("AEWI", idx) {
			s.qi.multiline = true    // start of a multi-line message
			s.qi.multiignore = false // reset continuation
		} else if strchar("CZ", idx) {
			// continuation of multi-line msg
			if !s.qi.multiignore {
				qfprev := s.qi.qflist[len(s.qi.qflist)-1]
				if qfprev == nil {
					return qffail, nil, errors.New("prev qfline doesn't exist")
				}
				if fields.errmsg != "" && !s.qi.multiignore {
					fmt.Println("CZ!", line)
					if qfprev.text == "" {
						qfprev.text = fields.errmsg
					} else {
						qfprev.text += "\n" + fields.errmsg
					}
				}
				if qfprev.nr == -1 {
					qfprev.nr = fields.enr
				}
				if fields.etype != 0 && qfprev.etype == 0 {
					qfprev.etype = fields.etype
				}
				if qfprev.lnum == 0 {
					qfprev.lnum = fields.lnum
				}
				if qfprev.col == 0 {
					qfprev.col = fields.col
				}
				qfprev.viscol = fields.useviscol
			}
			if idx == 'Z' {
				s.qi.multiline = false
				s.qi.multiignore = false
				return qfendmultiline, fields, nil
			}
			return qfignoreline, nil, nil
		} else if strchar("OPQ", idx) {
			// global file names
			fields.valid = false
			if fields.namebuf == "" || fileexists(fields.namebuf) {
				if fields.namebuf != "" && idx == 'P' {
					s.qi.currfile = fields.namebuf
					s.qi.filestack = append(s.qi.filestack, s.qi.currfile)
				} else if idx == 'Q' && len(s.qi.filestack) > 0 {
					s.qi.currfile = s.qi.filestack[len(s.qi.filestack)-1]
					s.qi.filestack = s.qi.filestack[:len(s.qi.filestack)-1]
				}
				fields.namebuf = ""
				if tail != "" {
					s.qi.multiscan = true
					return s.parseLineInternal(strings.TrimLeft(tail, " \t"), i)
				}
			}
		}
		if efm.flagminus { // generally exclude this line
			if s.qi.multiline { // also exclude continuation lines
				s.qi.multiignore = true
			}
			return qfignoreline, nil, nil
		}
	}
	return qfok, fields, nil
}

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
			// - do not support %>
			if re, ok := fmtpattern[efmp]; ok {
				regpat.WriteString(re)
			} else if efmp == '*' {
				incefmp()
				if efmp == '[' || efmp == '\\' {
					regpat.WriteByte(efmp)
					if efmp == '[' { // %*[^a-z0-9] etc.
						incefmp()
						for efmp != ']' {
							regpat.WriteByte(efmp)
							if i == len(errorformat)-1 {
								return nil, errors.New("E374: Missing ] in format string")
							}
							incefmp()
						}
						regpat.WriteByte(efmp)
					} else { // %*\D, %*\s etc.
						incefmp()
						regpat.WriteByte(efmp)
					}
					regpat.WriteRune('+')
				} else {
					return nil, fmt.Errorf("E375: Unsupported %%%v in format string", string(efmp))
				}
			} else if strchar(`%\.^$?+`, efmp) {
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
					return nil, fmt.Errorf("E376: Invalid %%%v in format string prefix", string(efmp))
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
	T byte   // (%t) error type
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
			match.T = m[0]
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

var fileexists = func(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
