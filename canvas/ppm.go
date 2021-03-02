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
	"fmt"
	"io"
	"math"
	"strconv"
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
	w                 io.Writer
	err               error
	maxLineLength     int
	currentLineLength int
	hadPixelInLine    bool
}

func newPPM(w io.Writer) *ppm {
	return &ppm{
		w:             w,
		maxLineLength: 70,
	}
}

// Error returns nil or the first occurred error.
func (p *ppm) Error() error {
	return p.err
}

// Printf uses fmt.Sprintf to render the string.
func (p *ppm) Printf(format string, args ...interface{}) {
	if p.err != nil {
		return
	}

	_, p.err = p.w.Write([]byte(fmt.Sprintf(format, args...)))
}

func (p *ppm) writeNum(f float32) {
	v := strconv.Itoa(int(math.RoundToEven(float64(f))))
	if p.currentLineLength+len(v) >= p.maxLineLength {
		p.currentLineLength = 0
		p.Printf("\n")
	} else {
		if p.hadPixelInLine {
			p.Printf(" ")
			p.currentLineLength++
		}
	}

	p.currentLineLength += len(v)
	p.Printf(v)
	p.hadPixelInLine = true
}

// WriteHeader emits the header bytes.
func (p *ppm) WriteHeader(width, height, max int) {
	p.Printf("P3\n")
	p.Printf("%d %d\n", width, height)
	p.Printf("255\n")
}

func (p *ppm) WritePixel(r, g, b float32) {
	p.writeNum(r)
	p.writeNum(g)
	p.writeNum(b)
}

func (p *ppm) EndRow() {
	p.Printf("\n")
	p.currentLineLength = 0
	p.hadPixelInLine = false
}

func (p *ppm) Close() error {
	if p.hadPixelInLine {
		p.EndRow()
	}

	p.Printf("\n")
	return p.Error()
}
