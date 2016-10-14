package errorformat

import (
	"reflect"
	"testing"
)

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
