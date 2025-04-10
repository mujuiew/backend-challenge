package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {
	tests := []struct {
		name      string
		inputFile string
		fileName  string
		expect    int
	}{
		{
			name:      "case 1",
			inputFile: "[[59], [73, 41], [52, 40, 53], [26, 53, 6, 34]]",
			fileName:  "./hard01.json",
			expect:    237,
		},
		{
			name:      "case emtpy",
			inputFile: "[]",
			fileName:  "./hard_emtpy.json",
			expect:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, errOs := os.Create(tt.fileName)
			assert.NoError(t, errOs)
			defer os.Remove(tt.fileName)
			defer file.Close()
			_, err := file.WriteString(tt.inputFile)
			assert.NoError(t, err)

			res := Handle(tt.fileName)
			assert.Equal(t, tt.expect, res, tt.name)
		})
	}
}

func TestHandle2(t *testing.T) {
	tests := []struct {
		name      string
		inputFile string
		fileName  string
		expect    int
	}{
		{
			name:      "case hard",
			inputFile: "[]",
			fileName:  "../files/hard.json",
			expect:    7273,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Handle(tt.fileName)
			assert.Equal(t, tt.expect, res, tt.name)
		})
	}
}
