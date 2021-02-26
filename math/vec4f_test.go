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

import (
	"math"
	"strconv"
	"testing"
)

func TestT4f_IsVector(t1 *testing.T) {
	tests := []struct {
		tuple    Vec4f
		isVector bool
		isPoint  bool
	}{
		{
			NewPoint(4, -4, 3),
			false,
			true,
		},
		{
			NewVector(4, -4, 3),
			true,
			false,
		},
	}
	for i, tt := range tests {
		t1.Run(strconv.Itoa(i), func(t1 *testing.T) {
			t := Vec4f{
				X: tt.tuple.X,
				Y: tt.tuple.Y,
				Z: tt.tuple.Z,
				W: tt.tuple.W,
			}

			if got := t.IsVector(); got != tt.isVector {
				t1.Errorf("IsVector() = %v, want %v", got, tt.isVector)
			}

			if got := t.IsPoint(); got != tt.isPoint {
				t1.Errorf("IsVector() = %v, want %v", got, tt.isPoint)
			}
		})
	}
}

func TestVec4f_Equals(t *testing.T) {
	tests := []struct {
		v    Vec4f
		args Vec4f
		want bool
	}{
		{NewVector(1, 2, 3), NewVector(1, 2, 3), true},
		{NewVector(1, 2, 3), NewPoint(1, 2, 3), false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := tt.v.Equals(&tt.args); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec4f_Add(t *testing.T) {
	tests := []struct {
		a, b Vec4f
		res  Vec4f
		want bool
	}{
		{Vec4f{3, -2, 5, 1}, Vec4f{-2, 3, 1, 0}, Vec4f{1, 1, 6, 1}, true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tmp := tt.a
			tmp.Add(&tt.b)

			if got := tmp.Equals(&tt.res); got != tt.want {
				t.Errorf("%v + %v = %v, want %v", tt.a, tt.b, tt.res, tt.want)
			}
		})
	}
}

func TestVec4f_Sub(t *testing.T) {
	tests := []struct {
		a, b Vec4f
		res  Vec4f
		want bool
	}{
		{NewPoint(3, 2, 1), NewPoint(5, 6, 7), NewVector(-2, -4, -6), true},
		{NewPoint(3, 2, 1), NewVector(5, 6, 7), NewPoint(-2, -4, -6), true},
		{NewVector(3, 2, 1), NewVector(5, 6, 7), NewVector(-2, -4, -6), true},
		{NewVector(0, 0, 0), NewVector(1, -2, 3), NewVector(-1, 2, -3), true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tmp := tt.a
			tmp.Sub(&tt.b)

			if got := tmp.Equals(&tt.res); got != tt.want {
				t.Errorf("%v - %v = %v, want %v", tt.a, tt.b, tt.res, tt.want)
			}
		})
	}
}

func TestVec4f_Negate(t *testing.T) {
	tests := []struct {
		a    Vec4f
		res  Vec4f
		want bool
	}{
		{Vec4f{1, -2, 3, -4}, Vec4f{-1, 2, -3, 4}, true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tmp := tt.a
			tmp.Negate()

			if got := tmp.Equals(&tt.res); got != tt.want {
				t.Errorf("- %v = %v, want %v", tt.a, tt.res, tt.want)
			}
		})
	}
}

func TestVec4f_Mul(t *testing.T) {
	tests := []struct {
		a    Vec4f
		b    float32
		res  Vec4f
		want bool
	}{
		{Vec4f{1, -2, 3, -4}, 3.5, Vec4f{3.5, -7, 10.5, -14}, true},
		{Vec4f{1, -2, 3, -4}, 0.5, Vec4f{0.5, -1, 1.5, -2}, true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tmp := tt.a
			tmp.Mul(tt.b)

			if got := tmp.Equals(&tt.res); got != tt.want {
				t.Errorf("%v * %v = %v, want %v", tt.a, tt.b, tt.res, tt.want)
			}
		})
	}
}

func TestVec4f_Div(t *testing.T) {
	tests := []struct {
		a    Vec4f
		b    float32
		res  Vec4f
		want bool
	}{
		{Vec4f{1, -2, 3, -4}, 2, Vec4f{0.5, -1, 1.5, -2}, true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tmp := tt.a
			tmp.Div(tt.b)

			if got := tmp.Equals(&tt.res); got != tt.want {
				t.Errorf("%v / %v = %v, want %v", tt.a, tt.b, tt.res, tt.want)
			}
		})
	}
}

func TestVec4f_Len(t *testing.T) {
	tests := []struct {
		a    Vec4f
		res  float32
		want bool
	}{
		{NewVector(0, 1, 0), 1, true},
		{NewVector(0, 0, 1), 1, true},
		{NewVector(1, 2, 3), float32(math.Sqrt(14)), true},
		{NewVector(-1, -2, -3), float32(math.Sqrt(14)), true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res := tt.a.Len()

			if got := res == tt.res; got != tt.want {
				t.Errorf("len(%v) = %v, want %v", tt.a, tt.res, tt.want)
			}
		})
	}
}

func TestVec4f_Normalize(t *testing.T) {
	tests := []struct {
		a    Vec4f
		res  Vec4f
		want bool
	}{
		{NewVector(4, 0, 0), NewVector(1, 0, 0), true},
		{NewVector(1, 2, 3), NewVector(1/Sqrt(14), 2/Sqrt(14), 3/Sqrt(14)), true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tmp := tt.a
			tmp.Normalize()

			if got := tmp.Equals(&tt.res); got != tt.want {
				t.Errorf("normalize(%v) = %v, want %v", tt.a, tt.res, tt.want)
			}
		})
	}

	vec := NewVector(1, 2, 3)
	vec.Normalize()
	if !Equalf(vec.Len(), 1) {
		t.Errorf("len %v, want 1", vec.Len())
	}
}

func TestVec4f_Dot(t *testing.T) {
	tests := []struct {
		a    Vec4f
		b    Vec4f
		res  float32
		want bool
	}{
		{NewVector(1, 2, 3), NewVector(2, 3, 4), 20, true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tmp := tt.a.Dot(&tt.b)

			if got := tt.res == tmp; got != tt.want {
				t.Errorf("%v dot %v = %v, want %v", tt.a, tt.b, tt.res, tt.want)
			}
		})
	}
}

func TestVec4f_Cross(t *testing.T) {
	tests := []struct {
		a    Vec4f
		b    Vec4f
		res  Vec4f
		want bool
	}{
		{NewVector(1, 2, 3), NewVector(2, 3, 4), NewVector(-1, 2, -1), true},
		{NewVector(2, 3, 4), NewVector(1, 2, 3), NewVector(1, -2, 1), true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tmp := tt.a
			tmp.Cross(&tt.b)

			if got := tt.res == tmp; got != tt.want {
				t.Errorf("%v cross %v = %v, want %v", tt.a, tt.b, tt.res, tt.want)
			}
		})
	}
}

func TestVec4f_MulVec(t *testing.T) {
	tests := []struct {
		a, b Vec4f
		res  Vec4f
		want bool
	}{
		{NewRGB(1, 0.2, 0.4), NewPoint(0.9, 1, 0.1), NewPoint(0.9, 0.2, 0.04), true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tmp := tt.a
			tmp.MulVec(&tt.b)

			if got := tmp.Equals(&tt.res); got != tt.want {
				t.Errorf("%v * %v = %v, want %v", tt.a, tt.b, tt.res, tt.want)
			}
		})
	}
}
