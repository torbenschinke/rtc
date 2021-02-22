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

package math

// A T4f represents either a tuple of point or vector coordinates.
type T4f struct {
	X, Y, Z, W float32
}

// NewPoint creates a point tuple with a W component of 1.
func NewPoint(x, y, z float32) T4f {
	return T4f{X: x, Y: y, Z: z, W: 1}
}

// NewVector creates a point tuple with a W component of 0.
func NewVector(x, y, z float32) T4f {
	return T4f{X: x, Y: y, Z: z, W: 0}
}

// IsVector returns true if the W component is 0.
func (t T4f) IsVector() bool {
	return t.W == 0
}

// IsPoint returns true if the W component is 1.
func (t T4f) IsPoint() bool {
	return t.W == 1
}

// Equals compares two tuples using Equalf.
func EqualsT4f(a, b T4f) bool {
	return Equalf(a.X, b.X) && Equalf(a.Y, b.Y) && Equalf(a.Z, b.Z) && Equalf(a.W, b.W)
}

// TODO proof non-escape? go build -gcflags="-m"
func (t *T4f) AddT4f(b *T4f) {
	t.X += b.X
}
