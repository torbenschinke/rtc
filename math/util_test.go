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

func TestEqual(t *testing.T) {
	tests := []struct {
		a    float32
		b    float32
		want bool
	}{
		{1, 1, true},
		{2, 1, false},
		{-1, -1, true},
		{-1, 1, false},
		{1 + Epsilon, 1 + Epsilon*0.1, true},
		{1 + Epsilon, 1 - Epsilon*0.1, false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Equalf(tt.a, tt.b); got != tt.want {
				t.Errorf("Equalf() = %v, want %v", got, tt.want)
			}
		})
	}
}
