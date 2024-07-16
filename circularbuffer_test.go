package bag

import (
	"reflect"
	"testing"
)

func Test_circularBuffer_Shift(t *testing.T) {
	type fields struct {
		size int
	}

	type args struct {
		values []int
	}

	type testcase struct {
		name   string
		fields fields
		args   args

		wantPopped []int
		wantSlice  []int
	}

	tests := []testcase{
		{
			name: "basic",
			fields: fields{
				size: 3,
			},
			args: args{
				values: []int{1, 2, 3},
			},
			wantPopped: []int{0, 0, 0},
			wantSlice:  []int{1, 2, 3},
		},
		{
			name: "with popped",
			fields: fields{
				size: 3,
			},
			args: args{
				values: []int{1, 2, 3, 4, 5, 6},
			},
			wantPopped: []int{0, 0, 0, 1, 2, 3},
			wantSlice:  []int{4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := newCircularBuffer[int](tt.fields.size)
			for i, arg := range tt.args.values {
				if got := b.Shift(arg); got != tt.wantPopped[i] {
					t.Fatalf("invalid value, wanted <%d> and received <%d>", tt.wantPopped[i], got)
				}
			}

			if !reflect.DeepEqual(b.s, tt.wantSlice) {
				t.Fatalf("invalid slice, wanted <%+v> and received <%+v>", tt.wantSlice, b.s)
			}
		})
	}
}

func Test_circularBuffer_ForEach(t *testing.T) {
	type fields struct {
		size int
	}

	type args struct {
		values   []int
		hasBreak bool
	}

	type testcase struct {
		name   string
		fields fields
		args   args

		want      []int
		wantBreak bool
	}

	tests := []testcase{
		{
			name: "basic",
			fields: fields{
				size: 3,
			},
			args: args{
				values: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "not full",
			fields: fields{
				size: 3,
			},
			args: args{
				values: []int{1, 2},
			},
			want: []int{1, 2},
		},
		{
			name: "with partial popped",
			fields: fields{
				size: 3,
			},
			args: args{
				values: []int{1, 2, 3, 4, 5},
			},
			want: []int{3, 4, 5},
		},
		{
			name: "with complete popped",
			fields: fields{
				size: 3,
			},
			args: args{
				values: []int{1, 2, 3, 4, 5, 6},
			},
			want: []int{4, 5, 6},
		},
		{
			name: "with has break",
			fields: fields{
				size: 3,
			},
			args: args{
				values:   []int{1, 2, 3, 4, 5, 6},
				hasBreak: true,
			},
			want:      []int{4, 5, 6},
			wantBreak: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := newCircularBuffer[int](tt.fields.size)
			for _, arg := range tt.args.values {
				b.Shift(arg)
			}

			var i int
			gotBreak := b.ForEach(func(val int) (end bool) {
				if val != tt.want[i] {
					t.Fatalf("invalid iteration value, expected %d and received %d", tt.want[i], val)
				}
				i++

				return tt.args.hasBreak
			})

			if gotBreak != tt.wantBreak {
				t.Fatalf("invalid break value, expected %v and received %v", tt.wantBreak, gotBreak)
			}
		})
	}
}

func Test_circularBuffer_Len(t *testing.T) {
	type fields struct {
		size int
	}

	type args struct {
		values []int
	}

	type testcase struct {
		name   string
		fields fields
		args   args

		want int
	}

	tests := []testcase{
		{
			name: "basic",
			fields: fields{
				size: 3,
			},
			args: args{
				values: []int{1, 2, 3},
			},
			want: 3,
		},
		{
			name: "partial",
			fields: fields{
				size: 3,
			},
			args: args{
				values: []int{1, 2},
			},
			want: 2,
		},
		{
			name: "empty",
			fields: fields{
				size: 3,
			},
			args: args{
				values: []int{},
			},
			want: 0,
		},
		{
			name: "with popped",
			fields: fields{
				size: 3,
			},
			args: args{
				values: []int{1, 2, 3, 4, 5, 6},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := newCircularBuffer[int](tt.fields.size)
			for _, arg := range tt.args.values {
				b.Shift(arg)
			}

			if got := b.Len(); got != tt.want {
				t.Fatalf("invalid length, expected %d and recieved %d", tt.want, got)
			}
		})
	}
}

func Test_newCircularBuffer(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "basic",
			args: args{
				capacity: 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newCircularBuffer[string](tt.args.capacity)
			if got := c.Cap(); got != tt.args.capacity {
				t.Errorf("newCircularBuffer().Capacity = %v, want %v", got, tt.args.capacity)
			}
		})
	}
}
