package main

import (
	"bytes"
	"testing"
)

func BenchmarkMaxLength(b *testing.B) {
	a := [][]string{
		{"123", "sdfgerwgeg", "sdf43terge", "sad", "tzutzutz"},
		{"hfghfgh", "3", "fghehjthjgfhjezthg", "sad", "tzutzutzttttttttttt"},
		{"545", "4554", "hgfhfg", "sad", "ghjtzj65hgfh"},
		{"fghfghfg", "fghfg", "tuzutzutzt", "hgjez565ztd", "tzutzutz"},
	}
	for n := 0; n < b.N; n++ {
		maxColLen(a)
	}
}

func BenchmarkBuildMarkdown(b *testing.B) {
	table := [][]string{{"colA1", "colA2", "colA3", "colA4", "colA5"}, {"colB1", "colB2", "colB3", "colB4", "colB5"}, {"colC1", "colC2", "colC3", "colC4", "colC5"}}
	var buf bytes.Buffer
	for n := 0; n < b.N; n++ {
		buildMarkdown(&buf, table, false)
	}
}
