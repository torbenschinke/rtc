// Copyright 2021 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package canvas

import (
	"bytes"
	"fmt"
	"io"
	"math"
)

// Export writes the buffer into a ppm (Portable Pixmap) format in plain PPM.
func (c *Canvas) Export(w io.Writer) error {
	ppm := newPPM(w)
	ppm.WriteHeader(c.Width, c.Height, 255)

	// actual pixel values
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			v := *c.Read(x, y)
			v.Saturate()
			v.Mul(255)
			ppm.WritePixel(v.X, v.Y, v.Z)
		}

		ppm.EndRow()
	}

	return ppm.Close()
}

type ppm struct {
	w             io.Writer
	tmp           *bytes.Buffer
	err           error
	maxLineLength int
	hadPixel      bool
}

func newPPM(w io.Writer) *ppm {
	return &ppm{
		w:             w,
		tmp:           &bytes.Buffer{},
		maxLineLength: 70,
	}
}

// Error returns nil or the first occured error.
func (p *ppm) Error() error {
	return p.err
}

// Write delegates and captures any error. Subsequent calls after an error are a no-op.
func (p *ppm) Write(buf []byte) (int, error) {
	if p.err != nil {
		return 0, p.err
	}

	delta := p.maxLineLength - (p.tmp.Len() + len(buf))
	if delta >= 0 {
		_, _ = p.tmp.Write(buf)
		return len(buf), nil
	}

	fragLen := len(buf) + delta
	_, _ = p.tmp.Write(buf[:fragLen])
	_ = p.Flush()
	_, _ = p.tmp.Write(buf[fragLen:])

	return len(buf), p.err
}

// Flush emits the internal buffer to the writer.
func (p *ppm) Flush() error {
	if p.err != nil {
		return p.err
	}

	if p.tmp.Len() > 0 {
		_, err := p.w.Write(p.tmp.Bytes())
		if err != nil {
			p.err = err
		} else {
			p.tmp.Reset()
		}
	}

	return p.err
}

// Printf uses fmt.Sprintf to render the string.
func (p *ppm) Printf(format string, args ...interface{}) {
	_, _ = p.Write([]byte(fmt.Sprintf(format, args...)))
}

// WriteHeader emits the header bytes.
func (p *ppm) WriteHeader(width, height, max int) {
	p.Printf("P3\n")
	p.Printf("%d %d\n", width, height)
	p.Printf("255\n")
}

func (p *ppm) WritePixel(r, g, b float32) {
	if p.hadPixel {
		p.Printf(" ")
	}
	ri := int(math.RoundToEven(float64(r)))
	gi := int(math.RoundToEven(float64(g)))
	bi := int(math.RoundToEven(float64(b)))
	p.Printf("%d %d %d", ri, gi, bi)
	p.hadPixel = true
}

func (p *ppm) EndRow() {
	p.Printf("\n")
	p.hadPixel = false
}

func (p *ppm) Close() error {
	if p.hadPixel {
		p.EndRow()
	}

	p.Printf("\n")
	return p.Flush()
}
