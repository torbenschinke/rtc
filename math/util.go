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

import "math"

// Epsilon is the margin by which floats are treated equal.
const Epsilon = 0.00001

// Equalf returns true if the difference between the given two floats is smaller than epsilon.
// Take a look at Knuths book, because this is not necessarily correct and only works
// reasonable for values near zero.
func Equalf(a, b float32) bool {
	if Abs(a-b) < Epsilon {
		return true
	} else {
		return false
	}
}

// Abs is just like math.Abs but for float32, so we don't need a type conversion.
// Don't know if this is actually correct, because I don't have seen that anywhere else and this
// is just a translation from the 64bit math from the stdlib.
func Abs(x float32) float32 {
	return math.Float32frombits(math.Float32bits(x) &^ (1 << 31))
}

// Just like math.Sqrt but with float32.
func Sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}
