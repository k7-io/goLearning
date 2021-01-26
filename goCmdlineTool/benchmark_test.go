package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const Num = 10000

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for i := 0; i < Num; i++ {
			str = fmt.Sprintf("%s%d", str, i)
		}
	}
	b.StopTimer()
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for i := 0; i < Num; i++ {
			builder.WriteString(strconv.Itoa(i))
		}
		_ = builder.String()
	}
	b.StopTimer()
}

func BenchmarkBytesBuf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for i := 0; i < Num; i++ {
			buf.WriteString(strconv.Itoa(i))
		}
		_ = buf.String()
	}
	b.StopTimer()
}

func BenchmarkAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for i := 0; i < Num; i++ {
			str += strconv.Itoa(i)
		}
	}
	b.StopTimer()
}
