// Copyright 2024 Ivan Korobkov. All rights reserved.
// Use of this software is governed by the MIT License
// that can be found in the LICENSE file.

package mpx

import (
	"testing"

	"github.com/basecomplextech/baselibrary/alloc"
	"github.com/basecomplextech/baselibrary/bin"
	"github.com/basecomplextech/spec/proto/pmpx"
)

func BenchmarkMessageBuild(b *testing.B) {
	buf := alloc.NewBuffer()
	input := pmpx.MessageInput{
		Id:     bin.Random128(),
		Data:   make([]byte, 128),
		Window: 16 * 1024 * 1024,
		Open:   true,
	}

	for i := 0; i < b.N; i++ {
		buf.Reset()

		msg, err := pmpx.BuildChannelMessage(buf, input)
		if err != nil {
			b.Fatal(err)
		}
		_ = msg
	}

	sec := b.Elapsed().Seconds()
	ops := float64(b.N) / sec

	b.ReportMetric(ops/1000_000, "mops")
}

func BenchmarkBatchBuild(b *testing.B) {
	buf := alloc.NewBuffer()
	id := bin.Random128()
	data := make([]byte, 128)
	window := int32(16 * 1024 * 1024)

	for i := 0; i < b.N; i++ {
		buf.Reset()

		x := pmpx.NewBatchBuilder(buf)
		x, err := x.Open(id, data, window)
		if err != nil {
			b.Fatal(err)
		}
		x, err = x.Close(id, nil)
		if err != nil {
			b.Fatal(err)
		}

		msg, err := x.Build()
		if err != nil {
			b.Fatal(err)
		}
		_ = msg
	}

	sec := b.Elapsed().Seconds()
	ops := float64(b.N) / sec

	b.ReportMetric(ops/1000_000, "mops")
}
