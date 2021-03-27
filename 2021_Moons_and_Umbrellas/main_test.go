package main

import (
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		s  string
		ca int
		cb int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Simple test",
			args: args{
				s:  "AB",
				ca: 1,
				cb: 2,
			},
			want: 1,
		},
		{
			name: "Simple test 2",
			args: args{
				s:  "AB?",
				ca: 1,
				cb: 2,
			},
			want: 1,
		},
		{
			name: "Simple test 3",
			args: args{
				s:  "AB?A",
				ca: 1,
				cb: 2,
			},
			want: 3,
		},
		{
			name: "Simple test 4",
			args: args{
				s:  "AB?AA?",
				ca: 1,
				cb: 2,
			},
			want: 3,
		},
		{
			name: "Simple test 5",
			args: args{
				s:  "ABAB",
				ca: 1,
				cb: 2,
			},
			want: 4,
		},
		{
			name: "Simple test 5",
			args: args{
				s:  "ABAB",
				ca: 1,
				cb: 2,
			},
			want: 4,
		},
		{
			name: "Simple test 5",
			args: args{
				s:  "BABA",
				ca: 2,
				cb: 4,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.s, tt.args.ca, tt.args.cb); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_final(t *testing.T) {
	type args struct {
		text string
		cx   int
		cy   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "real test 1",
			args: args{
				text: "CJ?CC?",
				cx:   2,
				cy:   3,
			},
			want: 5,
		},
		{
			name: "real test 2",
			args: args{
				text: "CJCJ",
				cx:   4,
				cy:   2,
			},
			want: 10,
		},
		{
			name: "real test 3",
			args: args{
				text: "C?J",
				cx:   1,
				cy:   3,
			},
			want: 1,
		},
		{
			name: "real test 4",
			args: args{
				text: "??J???",
				cx:   2,
				cy:   5,
			},
			want: 0,
		},
		{
			name: "dum test 1",
			args: args{
				text: "??????",
				cx:   2,
				cy:   5,
			},
			want: 0,
		},
		{
			name: "real test 5",
			args: args{
				text: "??JJ??",
				cx:   2,
				cy:   -5,
			},
			want: 0,
		},
		{
			name: "dumb test 5",
			args: args{
				text: "CJC??J?",
				cx:   4,
				cy:   2,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := final(tt.args.text, tt.args.cx, tt.args.cy); got != tt.want {
				t.Errorf("final() = %v, want %v", got, tt.want)
			}
		})
	}
}
