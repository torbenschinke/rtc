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

// A Vec4f represents either a tuple of point or vector coordinates.
// Validate that pointer receivers do not escape using
//   go build -gcflags="-m"
// I have chosen to use pointer receivers for modificating components
// because escape analysis works usually properly for stack values and instead
// of coping all components over and over again we just need the pointer to pass.
type Vec4f struct {
	X, Y, Z, W float32 // or also R,G,B,A
}

// NewPoint creates a point tuple with a W component of 1.
func NewPoint(x, y, z float32) Vec4f {
	return Vec4f{X: x, Y: y, Z: z, W: 1}
}

// NewVector creates a point tuple with a W component of 0.
func NewVector(x, y, z float32) Vec4f {
	return Vec4f{X: x, Y: y, Z: z, W: 0}
}

// NewRGB creates a new color values with alpha 1.
func NewRGB(r, g, b float32) Vec4f {
	return NewPoint(r, g, b)
}

// NewRGBA creates a new color values.
func NewRGBA(r, g, b, a float32) Vec4f {
	return Vec4f{r, g, b, a}
}

// IsVector returns true if the W component is 0.
func (v Vec4f) IsVector() bool {
	return v.W == 0
}

// IsPoint returns true if the W component is 1.
func (v Vec4f) IsPoint() bool {
	return v.W == 1
}

// Equals compares two tuples using Equalf.
func (v *Vec4f) Equals(o *Vec4f) bool {
	return Equalf(v.X, o.X) && Equalf(v.Y, o.Y) && Equalf(v.Z, o.Z) && Equalf(v.W, o.W)
}

// Add appends all components to the receiver from the given vector.
func (v *Vec4f) Add(o *Vec4f) {
	v.X += o.X
	v.Y += o.Y
	v.Z += o.Z
	v.W += o.W
}

// Add removes all given components from the receiver.
func (v *Vec4f) Sub(o *Vec4f) {
	v.X -= o.X
	v.Y -= o.Y
	v.Z -= o.Z
	v.W -= o.W
}

// Negate inverts each component.
func (v *Vec4f) Negate() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
	v.W = -v.W
}

// Mul multiplies each component with the given scalar.
func (v *Vec4f) Mul(scalar float32) {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar
	v.W *= scalar
}

// MulVec multiplies each component with the given scalar. If v is a color, this is blending (Hadamard product
// or Schur product).
func (v *Vec4f) MulVec(o *Vec4f) {
	v.X *= o.X
	v.Y *= o.Y
	v.Z *= o.Z
	v.W *= o.W
}

// Div divides each component with the given scalar.
func (v *Vec4f) Div(scalar float32) {
	v.X /= scalar
	v.Y /= scalar
	v.Z /= scalar
	v.W /= scalar
}

// Len returns length or magnitude of a vector.
func (v *Vec4f) Len() float32 {
	return Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

// Normalize sets this vector to be a unit vector with a length of 1.
func (v *Vec4f) Normalize() {
	v.Div(v.Len())
}

// Dot calculates the dot product with the other vector. The larger the angle between the
// vectors, the smaller the dot product is. A dot product of 1 means that the vectors are
// equal and -1 means they are just opposite. Also the dot product of two unit vectors
// is cos(angle(a,b)).
func (v *Vec4f) Dot(o *Vec4f) float32 {
	return v.X*o.X +
		v.Y*o.Y +
		v.Z*o.Z +
		v.W*o.W
}

// Cross calculates the cross product with the other vector. The resulting vector
// is perpendicular to v and o.
func (v *Vec4f) Cross(o *Vec4f) {
	*v = NewVector(
		v.Y*o.Z-v.Z*o.Y,
		v.Z*o.X-v.X*o.Z,
		v.X*o.Y-v.Y*o.X,
	)
}

// Saturate clamps all components range into 0 and 1.
func (v *Vec4f) Saturate() {
	if v.X < 0 {
		v.X = 0
	}

	if v.X > 1 {
		v.X = 1
	}

	if v.Y < 0 {
		v.Y = 0
	}

	if v.Y > 1 {
		v.Y = 1
	}

	if v.Z < 0 {
		v.Z = 0
	}

	if v.Z > 1 {
		v.Z = 1
	}

	if v.W < 0 {
		v.W = 0
	}

	if v.W > 1 {
		v.W = 1
	}
}
