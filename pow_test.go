package math32_test

import (
	"testing"

	. "github.com/chewxy/math32"
)

func TestPowLargeIntegerExponents(t *testing.T) {
	cases := []struct {
		name       string
		x, y, want float32
	}{
		{"negative_one", -1, 1 << 31, 1},
		{"negative_one_reciprocal", -1, -(1 << 31), 1},
		{"negative_one_max_exponent", -1, MaxFloat32, 1},
		{"negative_one_min_exponent", -1, -MaxFloat32, 1},
		{"overflow", -2, 1 << 31, Inf(1)},
		{"underflow", -2, -(1 << 31), 0},
		{"fraction_underflow", -0.5, 1 << 31, 0},
		{"fraction_overflow", -0.5, -(1 << 31), Inf(1)},
		{"below_one", Nextafter(-1, 0), 1 << 31, 0},
		{"above_one", Nextafter(-1, Inf(-1)), 1 << 31, Inf(1)},
		{"negative_zero_largest_odd", Copysign(0, -1), (1 << 24) - 1, Copysign(0, -1)},
		{"negative_zero_even_threshold", Copysign(0, -1), 1 << 24, 0},
		{"negative_zero", Copysign(0, -1), 1 << 31, 0},
		{"negative_zero_reciprocal", Copysign(0, -1), -(1 << 31), Inf(1)},
		{"negative_zero_infinite_exponent", Copysign(0, -1), Inf(1), 0},
		{"negative_zero_negative_infinite_exponent", Copysign(0, -1), Inf(-1), Inf(1)},
		{"negative_infinity", Inf(-1), 1 << 31, Inf(1)},
		{"negative_infinity_reciprocal", Inf(-1), -(1 << 31), 0},
		{"below_threshold", -1, Nextafter(1<<31, 0), 1},
		{"positive_base", 2, 1 << 31, Inf(1)},
		{"fractional_exponent", -2, 1.5, NaN()},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Pow(tc.x, tc.y); !alike(got, tc.want) {
				t.Errorf("Pow(%g, %g) = %g, want %g", tc.x, tc.y, got, tc.want)
			}
		})
	}
}
