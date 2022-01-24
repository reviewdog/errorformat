package errorformat

import (
	"reflect"
	"strings"
	"testing"
)

func TestScanner_Scan(t *testing.T) {
	defer func(save func(string) bool) { fileexists = save }(fileexists)
	fileexists = func(string) bool { return true }

	tests := []struct {
		efm  []string
		in   string
		want []string
	}{
		{
			efm: []string{`%f:%l:%c: %m`, `%-G%.%#`},
			in: `
golint.new.go:3:5: exported var V should have comment or be unexported
golint.new.go:5:5: exported var NewError1 should have comment or be unexported
golint.new.go:7:1: comment on exported function F should be of the form "F ..."
golint.new.go:11:1: comment on exported function F2 should be of the form "F2 ..."
hoge
		`,
			want: []string{
				`golint.new.go|3 col 5| exported var V should have comment or be unexported`,
				`golint.new.go|5 col 5| exported var NewError1 should have comment or be unexported`,
				`golint.new.go|7 col 1| comment on exported function F should be of the form "F ..."`,
				`golint.new.go|11 col 1| comment on exported function F2 should be of the form "F2 ..."`,
			},
		},
		{
			efm: []string{
				`%C %.%#`,
				`%A  File "%f", line %l%.%#`,
				`%Z%\b%m`,
			},
			in: `
==============================================================
FAIL: testGetTypeIdCachesResult (dbfacadeTest.DjsDBFacadeTest)
--------------------------------------------------------------
Traceback (most recent call last):
  File "unittests/dbfacadeTest.py", line 89, in testFoo
    self.assertEquals(34, dtid)
  File "/usr/lib/python2.2/unittest.py", line 286, in
 failUnlessEqual
    raise self.failureException, \
AssertionError: 34 != 33

--------------------------------------------------------------
Ran 27 tests in 0.063s`,
			want: []string{
				`||`,
				`|| ==============================================================`,
				`|| FAIL: testGetTypeIdCachesResult (dbfacadeTest.DjsDBFacadeTest)`,
				`|| --------------------------------------------------------------`,
				`|| Traceback (most recent call last):`,
				`unittests/dbfacadeTest.py|89| AssertionError: 34 != 33`,
				`||`,
				`|| --------------------------------------------------------------`,
				`|| Ran 27 tests in 0.063s`,
			},
		},
		{
			efm: []string{
				`%A[%f]`,
				`%C%trror`,
				`%C%l,%c`,
				`%Z%m`,
			},
			in: `[~/.vimrc]
Error
12,5`,
			want: []string{"~/.vimrc|12 col 5 error|"},
		},
		{
			efm: []string{
				`%A[%f]`,
				`%C%tarning %n`,
				`%C%l,%c`,
				`%Z%m`,
			},
			in: `[~/.vimrc]
warning 14
12,5`,
			want: []string{"~/.vimrc|12 col 5 warning 14|"},
		},
		{
			efm: []string{
				`$%m`,
				`%-G%.%#`,
			},
			in: `$hoge
$foo
piyo`,
			want: []string{"|| hoge", "|| foo"},
		},
		{
			efm: []string{`(%l,%c): %m`},
			in: `(1,2): abc
(13,27): xyz
hoge`,
			want: []string{"|1 col 2| abc", "|13 col 27| xyz", "|| hoge"},
		},
		{
			efm: []string{`%[a-i]%m`},
			in: `hoge
foo
piyo`,
			want: []string{"|| oge", "|| oo", "|| piyo"},
		},
		{efm: []string{`{%l} %m`}, in: `{14} msg`, want: []string{"|14| msg"}},
		{efm: []string{`[%l] %m`}, in: `[14] msg`, want: []string{"|14| msg"}},
		{
			efm: []string{`%EError %n`, `%Cline %l`, `%Ccolumn %c`, `%Z%m`},
			in: `Error 275
line 42
column 3
' ' expected after '--'`,
			want: []string{"|42 col 3 error 275| ' ' expected after '--'"},
		},
		{ // :h errorformat-separate-filename
			efm: []string{`%+P[%f]`, `(%l,%c)%*[ ]%t%*[^:]: %m`, `%-Q`},
			in: `[a1.tt]
(1,17)  error: ';' missing
(21,2)  warning: variable 'z' not defined
(67,3)  error: end of file found before string ended

[a2.tt]

[a3.tt]
NEW compiler v1.1
(2,2)   warning: variable 'x' not defined
(67,3)  warning: 's' already defined
`,
			want: []string{
				"a1.tt|| [a1.tt]",
				"a1.tt|1 col 17 error| ';' missing",
				"a1.tt|21 col 2 warning| variable 'z' not defined",
				"a1.tt|67 col 3 error| end of file found before string ended",
				"a2.tt|| [a2.tt]",
				"a3.tt|| [a3.tt]",
				"a3.tt|| NEW compiler v1.1",
				"a3.tt|2 col 2 warning| variable 'x' not defined",
				"a3.tt|67 col 3 warning| 's' already defined",
			},
		},
		{
			efm: []string{`%-P[%f]`, `(%l,%c)%*[ ]%t%*[^:]: %m`, `%-Q`, `%-G%.%#`},
			in: `[a1.tt]
(1,17)  error: ';' missing
(21,2)  warning: variable 'z' not defined
(67,3)  error: end of file found before string ended

[a2.tt]

[a3.tt]
NEW compiler v1.1
(2,2)   warning: variable 'x' not defined
(67,3)  warning: 's' already defined
`,
			want: []string{
				"a1.tt|1 col 17 error| ';' missing",
				"a1.tt|21 col 2 warning| variable 'z' not defined",
				"a1.tt|67 col 3 error| end of file found before string ended",
				"a3.tt|2 col 2 warning| variable 'x' not defined",
				"a3.tt|67 col 3 warning| 's' already defined",
			},
		},
		{ // grep format. (ag, pt, etc...)
			efm: []string{`%l:%m`, `%-P%f`, `%-Q`},
			in: `errorformat.go
1:// Package errorformat provides 'errorformat' functionality of Vim. :h
398:// NewEfm converts a 'errorformat' string to regular expression pattern with

README.md
1:## go-errorformat - vim 'errorformat' implementation in Go
`,
			want: []string{
				"errorformat.go|1| // Package errorformat provides 'errorformat' functionality of Vim. :h",
				"errorformat.go|398| // NewEfm converts a 'errorformat' string to regular expression pattern with",
				"README.md|1| ## go-errorformat - vim 'errorformat' implementation in Go",
			},
		},
		{ // sbt
			efm: []string{
				`%E[%t%.%+] %f:%l: error: %m`,
				`%A[%t%.%+] %f:%l: %m`,
				`%Z[%.%+] %p^`,
				`%C[%.%+] %.%#`,
				`%-G%.%#`,
			},
			in: `
[error] /path/to/file:14: error: value ++ is not a member of Int
[error]   val x = 1 ++ 2
[error]             ^
[warn] /path/to/file2:14: local val in method watch is never used
[warn]    val x = 1
[warn]        ^
`,
			want: []string{
				"/path/to/file|14 col 13 error| value ++ is not a member of Int",
				"/path/to/file2|14 col 8 warning| local val in method watch is never used",
			},
		},
		{ // multiline
			efm: []string{
				`%E[%t%.%+] %f:%l: error: %m`,
				`%A[%t%.%+] %f:%l: %m`,
				`[%t%.%+] %f: error: %m`, // oneline
				`%Z[%.%+] %p^`,
				`%C[%.%+] %.%#`,
				`%-G%.%#`,
			},
			in: `
[error] /path/to/file:14: error: value ++ is not a member of Int
[error]   val x = 1 ++ 2
[error]             ^
[error] /path/to/file: error: oneline error for the file
[error] /path/to/file:14: error: multiline
[error]  without pointer
[error] /path/to/file: error: oneline error for the file 2
`,
			want: []string{
				"/path/to/file|14 col 13 error| value ++ is not a member of Int",
				"/path/to/file| error| oneline error for the file",
				"/path/to/file|14 error| multiline",
				"/path/to/file| error| oneline error for the file 2",
			},
		},
		{
			efm: []string{
				`%EMultilineError`,
				`%Z%f:%l-%e:%c-%k`,
			},
			in: `MultilineError
~/.vimrc:1-2:3-4`,
			want: []string{"~/.vimrc|1-2 col 3-4 error|"},
		},
		{ // note
			efm: []string{
				`[%tote]%f:%l:%c: %m`,
				`%-G%.%#`,
			},
			in: `
[note]/path/to/file:1:2: note msg 1
[Note]/path/to/file:1:2: note msg 2`,
			want: []string{
				"/path/to/file|1 col 2 note| note msg 1",
				"/path/to/file|1 col 2 note| note msg 2",
			},
		},
		{ // end line/column
			efm: []string{
				`%f:%l-%e:%c-%k: %m`,
				`%-G%.%#`,
			},
			in: `/path/to/file:1-2:3-4: msg`,
			want: []string{
				"/path/to/file|1-2 col 3-4| msg",
			},
		},
	}
nexttext:
	for _, tt := range tests {
		efm, err := NewErrorformat(tt.efm)
		if err != nil {
			t.Error(err)
			continue
		}
		s := efm.NewScanner(strings.NewReader(tt.in))
		i := 0
		for s.Scan() {
			got := s.Entry().String()
			efm := strings.Join(tt.efm, ",")
			if i >= len(tt.want) {
				t.Errorf("%v:%d: got %q, want is not specified", efm, i, got)
				continue nexttext
			}
			want := tt.want[i]
			if got != want {
				t.Errorf("%v:%d:\ngot:  %q\nwant: %q", efm, i, got, want)
			}
			i++
		}
	}
}

func TestEntry_Types(t *testing.T) {
	tests := []struct {
		t    rune
		nr   int
		want string
	}{
		{t: 'e', want: "error"},
		{t: 'E', want: "error"},
		{t: 'e', nr: 14, want: "error 14"},
		{t: 'w', nr: 14, want: "warning 14"},
		{t: 'i', want: "info"},
		{t: 'n', want: "note"},
		{nr: 14, want: "error 14"},
	}
	for _, tt := range tests {
		e := &Entry{Type: tt.t, Nr: tt.nr}
		if got := e.Types(); got != tt.want {
			t.Errorf("(%v, %v): got %q, want %q", e.Type, e.Nr, got, tt.want)
		}
	}
}

type matchCase struct {
	in   string
	want *Match
}

func TestNewEfm(t *testing.T) {
	tests := []struct {
		errorformat string
		cases       []matchCase
	}{
		{
			errorformat: `%f:%l:%c: %m`,
			cases: []matchCase{
				{
					in: "golint.new.go:3:5: exported var V should have comment or be unexported",
					want: &Match{
						F: "golint.new.go",
						L: 3,
						C: 5,
						M: "exported var V should have comment or be unexported",
					},
				},
				{in: "", want: nil},
				{in: "golint.new.go:3:5", want: nil},
				{
					in:   `/path/t\ o/file:1:4: message`,
					want: &Match{F: `/path/t\ o/file`, L: 1, C: 4, M: "message"},
				},
				{
					in:   `\path\to\file:1:4: message`,
					want: &Match{F: `\path\to\file`, L: 1, C: 4, M: "message"},
				},
			},
		},
		{
			errorformat: ``,
			cases: []matchCase{
				{in: "", want: &Match{}},
				{in: "golint.new.go:3:5", want: nil},
			},
		},
		{
			// windows
			errorformat: `%f:%m`,
			cases: []matchCase{
				{
					in:   "c:/foo/bar.c:msg:hi:14",
					want: &Match{F: "c:/foo/bar.c", M: "msg:hi:14"},
				},
				{
					in:   "hoge/c:/foo/bar.c:msg:hi:14",
					want: &Match{F: "hoge/c", M: "/foo/bar.c:msg:hi:14"},
				},
			},
		},
	}
	for _, tt := range tests {
		efm, err := NewEfm(tt.errorformat)
		if err != nil {
			t.Errorf("NewEfm(%v) got an unexpected error: %v", tt.errorformat, err)
		}

		for _, c := range tt.cases {
			if got := efm.Match(c.in); !reflect.DeepEqual(got, c.want) {
				t.Errorf("efm: %v, efm.Match(%v) =\n%#v\nwant:\n %#v",
					tt.errorformat, c.in, got, c.want)
			}
		}
	}
}
