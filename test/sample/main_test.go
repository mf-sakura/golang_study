package main

import (
	"testing"
)

func TestCounter(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "when x is < 99", args: args{x: 1}, want: "1"},
		{name: "when x is >= 99", args: args{x: 100}, want: "99+"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Counter(tt.args.x); got != tt.want {
				t.Errorf("Counter() = %v, want %v", got, tt.want)
			}
		})
	}
}

// gotestsでの自動生成やってみる
func TestSum(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "when sum is = 999", args: args{x: 1, y: 998}, want: 999},
		{name: "when sum is = 1000", args: args{x: 1, y: 999}, want: 0},
		{name: "when sum is > 1000", args: args{x: 1, y: 1000}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
