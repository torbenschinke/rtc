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
	"github.com/torbenschinke/rtc/math"
	"strconv"
	"testing"
)

func TestCanvas_Export(t *testing.T) {

	tests := []struct {
		factory func() *Canvas
		ppm     string
	}{
	/*	{
			factory: func() *Canvas {
				c := NewCanvas(5, 3)
				return &c
			},
			ppm:
			`P3
5 3
255
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

`,
		},

		{
			factory: func() *Canvas {
				c := NewCanvas(5, 3)
				c1 := math.NewRGB(1.5, 0, 0)
				c2 := math.NewRGB(0, 0.5, 0)
				c3 := math.NewRGB(-0.5, 0, 1)
				c.Write(0, 0, &c1)
				c.Write(2, 1, &c2)
				c.Write(4, 2, &c3)

				return &c
			},

			ppm:
			`P3
5 3
255
255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255

`,
		},
*/
		{
			factory: func() *Canvas {
				c := NewCanvas(10, 2)
				c.Clear(math.NewRGB(1, 0.8, 0.6))
				return &c
			},
			ppm:
			`P3
10 2
255
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153  
`,
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := tt.factory()
			tmp := &bytes.Buffer{}
			err := c.Export(tmp)
			if err != nil {
				t.Errorf("Export() error = %v", err)
				return
			}

			if string(tmp.Bytes()) != tt.ppm {
				t.Errorf("Export() gotW = \n%v, want \n%v", string(tmp.Bytes()), tt.ppm)
			}
		})
	}
}
