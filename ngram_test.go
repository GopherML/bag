package bag

import (
	"reflect"
	"testing"
)

var ngramsSink []string

func Test_toNGrams(t *testing.T) {
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
				in:   "hello world! This is really cool, wowo",
				size: 3,
			},
			wantNs: []string{
				"hello world this",
				"world this is",
				"this is really",
				"is really cool",
				"really cool wowo",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNs := toNGrams(tt.args.in, tt.args.size); !reflect.DeepEqual(gotNs, tt.wantNs) {
				t.Errorf("toNGrams() = \n%v,\n want \n%v", gotNs, tt.wantNs)
			}
		})
	}
}

func Benchmark_toNGrams(b *testing.B) {
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
				in:   "hello world! This is really cool, wowo",
				size: 3,
			},
			wantNs: []string{
				"hello world this",
				"world this is",
				"this is really",
				"is really cool",
				"really cool wowo",
			},
		},
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range tests {
			ngramsSink = toNGrams(tc.args.in, tc.args.size)
		}
	}
}
