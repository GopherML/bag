package bag

import (
	"reflect"
	"testing"
)

func Test_toCharacterNGrams(t *testing.T) {
	type args struct {
		in   string
		size int
	}

	type testcase struct {
		name   string
		args   args
		wantNs []string
	}

	tests := []testcase{
		{
			name: "basic",
			args: args{
				in:   "This is my test string!",
				size: 3,
			},
			wantNs: []string{
				"thi",
				"his",
				"is ",
				"s i",
				" is",
				"is ",
				"s m",
				" my",
				"my ",
				"y t",
				" te",
				"tes",
				"est",
				"st ",
				"t s",
				" st",
				"str",
				"tri",
				"rin",
				"ing",
			},
		},
		{
			name: "short",
			args: args{
				in:   "ya",
				size: 3,
			},
			wantNs: []string{
				"ya",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNs := toCharacterNGrams(tt.args.in, tt.args.size); !reflect.DeepEqual(gotNs, tt.wantNs) {
				t.Errorf("toCharacterNGrams() = \n%v\n, want \n%v", gotNs, tt.wantNs)
			}
		})
	}
}
