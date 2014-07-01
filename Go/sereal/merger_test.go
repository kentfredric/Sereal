package sereal

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func BenchmarkMerger(b *testing.B) {
	files, _ := filepath.Glob("data/*.srl")
	if files == nil {
		b.Fatal("no files found")
	}

	var data [][]byte
	for _, file := range files {
		buf, ok := ioutil.ReadFile(file)
		if ok != nil {
			b.Fatal("failed to read file: " + file)
		}

		data = append(data, buf)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var m = NewMerger()
		for _, buf := range data {
			err := m.Append(buf)
			if err != nil {
				panic(err)
			}
		}
	}
}
