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

import "github.com/torbenschinke/rtc/math"

// Canvas holds a buffer
type Canvas struct {
	Buffer        []math.Vec4f //rgba pixels
	Width, Height int
}

// NewCanvas allocates a new canvas buffer at the heap.
func NewCanvas(w, h int) Canvas {
	return Canvas{
		Buffer: make([]math.Vec4f, w*h),
		Width:  w,
		Height: h,
	}
}

func (c *Canvas)Clear(color math.Vec4f){
	for i := range c.Buffer {
		c.Buffer[i] = color
	}
}

// Write sets the color at the pixel position.
func (c *Canvas) Write(x, y int, color *math.Vec4f) {
	c.Buffer[y*c.Width+x] = *color
}

// Read returns the pointer into the buffer at the required location. The color values are on the heap
// anyway, because the buffer has no static known size.
func (c *Canvas) Read(x, y int) *math.Vec4f {
	return &c.Buffer[y*c.Width+x]
}
