package bag

import "testing"

func Test_toWords(t *testing.T) {
	type args struct {
		in string
	}

	type testcase struct {
		name string
		args args
		want []string
	}

	tests := []testcase{
		{
			name: "truncating basic",
			args: args{
				in: "helllo",
			},
			want: []string{"hello"},
		},
		{
			name: "truncating long",
			args: args{
				in: "hellllllllllllllllo",
			},
			want: []string{"hello"},
		},
		{
			name: "varied",
			args: args{
				in: "what upp my duuuuuudeeeee?",
			},
			want: []string{"what", "upp", "my", "duudee"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var i int
			toWords(tt.args.in, func(s string) {
				if s != tt.want[i] {
					t.Fatalf("invalid value for index <%d>, expected <%s> and received <%s>", i, tt.want[i], s)
				}

				i++
			})
		})
	}
}
