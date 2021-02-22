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
	"strconv"
	"testing"
)

func TestT4f_IsVector(t1 *testing.T) {
	tests := []struct {
		tuple    T4f
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
			t := T4f{
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
