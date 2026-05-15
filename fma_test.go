package math32_test

import (
	"testing"

	. "github.com/chewxy/math32"
)

// TestFMASpecialValues verifies FMA with IEEE 754 special values.
// These cases follow from the IEEE 754-2008 specification for fusedMultiplyAdd.
func TestFMASpecialValues(t *testing.T) {
	nan := NaN()
	inf := Inf(1)
	ninf := Inf(-1)

	cases := []struct {
		name    string
		x, y, z float32
		want    float32
	}{
		{"zero", 0, 0, 0, 0},
		{"one", 1, 1, 0, 1},
		// NaN propagation
		{"nan*1+0", nan, 1, 0, nan},
		{"1*nan+0", 1, nan, 0, nan},
		{"1*1+nan", 1, 1, nan, nan},
		// Inf propagation
		{"+inf*1+0", inf, 1, 0, inf},
		{"-inf*1+0", ninf, 1, 0, ninf},
		{"inf*2+0", inf, 2, 0, inf},
		// Inf - Inf = NaN
		{"inf*1+(-inf)", inf, 1, ninf, nan},
		{"(-inf)*1+inf", ninf, 1, inf, nan},
		// 0 * Inf = NaN
		{"0*inf+0", 0, inf, 0, nan},
		{"inf*0+0", inf, 0, 0, nan},
		// Inf * Inf
		{"inf*inf+0", inf, inf, 0, inf},
		{"(-inf)*(-inf)+0", ninf, ninf, 0, inf},
		{"inf*(-inf)+0", inf, ninf, 0, ninf},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := FMA(c.x, c.y, c.z)
			if !alike(got, c.want) {
				t.Errorf("FMA(%v, %v, %v) = %v, want %v", c.x, c.y, c.z, got, c.want)
			}
		})
	}
}

// TestFMASingleRounding tests that FMA rounds only once (after the full x*y+z
// computation), not twice (once after x*y, then again after adding z). All
// cases below are analytically constructed so that the naive float32(x*y)+z
// gives a different answer than the correctly rounded x*y+z.
//
// Each case is of the form: the exact product p = x*y cannot be represented
// exactly in float32, and the nearest-float32(p) differs from p by exactly
// half an ULP (round-to-even), so that when z nearly cancels p the error
// from the premature rounding becomes visible in the result.
func TestFMASingleRounding(t *testing.T) {
	// All expected values are derived analytically:
	//
	// Case A: x = y = 4097 = 2^12 + 1 (exactly representable in float32).
	//   Exact product: 4097² = 16785409 = 2^24 + 2^13 + 1.
	//   At scale 2^24 the float32 ULP is 2, so 16785409 lies exactly halfway
	//   between 16785408 and 16785410; round-to-even chooses 16785408.
	//   Naive implementation (x*y then +z):
	//     float32(16785409) = 16785408 → 16785408 + z
	//   True FMA (one rounding at the end):
	//     round_float32(16785409 + z)
	//
	// Case B: x = y = 8388609 = 2^23 + 1 (exactly representable in float32).
	//   Exact product: 8388609² = 2^46 + 2^24 + 1.
	//   At scale 2^46 the float32 ULP is 2^23, so +1 is rounded away → 2^46 + 2^24.
	//   True FMA keeps the +1 until the final rounding.

	cases := []struct {
		name      string
		x, y, z   float32
		want      float32 // correctly rounded FMA result
		naiveWant float32 // what the naive x*y+z returns (wrong)
	}{
		// --- Case A variants ---
		// z cancels 4097² exactly → naive loses the +1 rounding error.
		{
			"4097²-16785408: want 1, naive 0",
			4097, 4097, -16785408,
			1, 0,
		},
		// z cancels 2^24 → naive misses the low bit of the product.
		{
			"4097²-16777216: want 8193, naive 8192",
			4097, 4097, -16777216,
			8193, 8192,
		},
		// z cancels 2^24+2^12 → same pattern, different residual.
		{
			"4097²-16781312: want 4097, naive 4096",
			4097, 4097, -16781312,
			4097, 4096,
		},
		// Negative x*y: same 4097² structure with sign flip.
		{
			"4097*(-4097)+16785408: want -1, naive 0",
			4097, -4097, 16785408,
			-1, 0,
		},
		// --- Case B ---
		// z = -(2^46 + 2^23); exactly representable (= (2^23+1) ULPs above 2^46).
		// Naive: (2^46+2^24)-(2^46+2^23) = 2^23 = 8388608
		// FMA:   (2^46+2^24+1)-(2^46+2^23) = 2^23+1 = 8388609
		{
			"8388609²-(2^46+2^23): want 8388609, naive 8388608",
			8388609, 8388609, -70368752566272,
			8388609, 8388608,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Confirm the test is set up correctly: naive impl gives the wrong answer.
			naive := c.x*c.y + c.z
			if naive != c.naiveWant {
				t.Fatalf("test setup error: naive x*y+z = %v, expected naiveWant %v",
					naive, c.naiveWant)
			}

			got := FMA(c.x, c.y, c.z)
			if got != c.want {
				t.Errorf("FMA(%v, %v, %v) = %v, want %v (naive x*y+z = %v)",
					c.x, c.y, c.z, got, c.want, naive)
			}
		})
	}
}
