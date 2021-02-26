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
	"github.com/torbenschinke/rtc/math"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	black := math.NewRGBA(0, 0, 0, 0)
	c := NewCanvas(10, 20)
	for w := 0; w < c.Width; w++ {
		for h := 0; h < c.Height; h++ {
			if !c.Read(w, h).Equals(&black) {
				t.Errorf("expected black color")
			}
		}
	}
}

func TestWriteCanvas(t *testing.T) {
	red := math.NewRGB(1, 0, 0)
	c := NewCanvas(10, 20)
	c.Write(2, 3, &red)
	if !c.Read(2, 3).Equals(&red) {
		t.Errorf("expected red")
	}
}
