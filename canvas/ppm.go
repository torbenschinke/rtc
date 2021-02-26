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
	"io"
	"math"
	"strconv"
)

// Export writes the buffer into a ppm (Portable Pixmap) format in plain PPM.
func (c *Canvas) Export(w io.Writer) error {
	// write magic header
	if _, err := w.Write([]byte("P3\n")); err != nil {
		return err
	}

	// write width and height
	if _, err := w.Write([]byte(strconv.Itoa(c.Width) + " " + strconv.Itoa(c.Height) + "\n")); err != nil {
		return err
	}

	// max color value
	if _, err := w.Write([]byte("255\n")); err != nil {
		return err
	}

	// actual pixel values
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			v := *c.Read(x, y)
			v.Saturate()
			v.Mul(255)
			if _, err := w.Write([]byte(strconv.Itoa(int(math.RoundToEven(float64(v.X)))))); err != nil {
				return err
			}

			if _, err := w.Write([]byte(" ")); err != nil {
				return err
			}

			if _, err := w.Write([]byte(strconv.Itoa(int(math.RoundToEven(float64(v.Y)))))); err != nil {
				return err
			}

			if _, err := w.Write([]byte(" ")); err != nil {
				return err
			}

			if _, err := w.Write([]byte(strconv.Itoa(int(math.RoundToEven(float64(v.Z)))))); err != nil {
				return err
			}

			if _, err := w.Write([]byte(" ")); err != nil {
				return err
			}
		}

		if _, err := w.Write([]byte("\n")); err != nil {
			return err
		}
	}

	return nil
}
