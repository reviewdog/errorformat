package errorformat

import (
	"reflect"
	"strings"
	"testing"
)

func TestScanner_Scan(t *testing.T) {
	defer func(save func(string) bool) { fileexists = save }(fileexists)
	fileexists = func(string) bool { return true }

	result := `
golint.new.go:3:5: exported var V should have comment or be unexported
golint.new.go:5:5: exported var NewError1 should have comment or be unexported
golint.new.go:7:1: comment on exported function F should be of the form "F ..."
golint.new.go:11:1: comment on exported function F2 should be of the form "F2 ..."
hoge
`
	efm, err := NewErrorformat([]string{`%f:%l:%c: %m`, `%-G%.%#`})
	if err != nil {
		t.Fatal(err)
	}
	s := efm.NewScanner(strings.NewReader(result))
	for s.Scan() {
		e := s.Entry()
		t.Errorf("%#v", e)
	}
}

func TestScanner_Scan_multi(t *testing.T) {
	defer func(save func(string) bool) { fileexists = save }(fileexists)
	fileexists = func(string) bool { return true }
	result := `
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
Ran 27 tests in 0.063s
`
	efms := []string{
		`%C %.%#`,
		`%A  File "%f", line %l%.%#`,
		`%Z%\b%m`,
	}
	efm, err := NewErrorformat(efms)
	if err != nil {
		t.Fatal(err)
	}
	s := efm.NewScanner(strings.NewReader(result))
	for s.Scan() {
		e := s.Entry()
		t.Errorf("%#v", e)
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
