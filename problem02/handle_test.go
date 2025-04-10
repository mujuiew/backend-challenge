package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_process(t *testing.T) {
	type args struct {
		encoded string
	}
	tests := []struct {
		name   string
		args   args
		expect string
	}{
		{
			name: "case 1",
			args: args{
				encoded: "LLRR=",
			},
			expect: "210122",
		},
		{
			name: "case 2",
			args: args{
				encoded: "==RLL",
			},
			expect: "000210",
		},
		{
			name: "case 3",
			args: args{
				encoded: "=LLRR",
			},
			expect: "221012",
		},
		{
			name: "case 4",
			args: args{
				encoded: "RRL=R",
			},
			expect: "012001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := process(tt.args.encoded)
			assert.Equal(t, tt.expect, res)
		})
	}
}
