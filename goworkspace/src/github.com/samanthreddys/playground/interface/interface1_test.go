package main

import "testing"

func Test_info(t *testing.T) {
	type args struct {
		z Shape
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info(tt.args.z)
		})
	}
}
