package math32

import (
	"math/rand"
	"testing"
)

// The tests in this file check that the assembly implementations agree
// bit-for-bit with the pure Go implementations. They are skipped on
// architectures (or build tags) where no assembly implementation is used.

// sameBits reports whether a and b are identical, treating all NaNs as equal.
func sameBits(a, b float32) bool {
	if IsNaN(a) || IsNaN(b) {
		return IsNaN(a) && IsNaN(b)
	}
	return Float32bits(a) == Float32bits(b)
}

// archEdgeValues are special values and values around the boundaries the
// assembly implementations care about.
var archEdgeValues = func() []float32 {
	bits := []uint32{
		0x00000000, // +0
		0x00000001, // smallest subnormal
		0x007fffff, // largest subnormal
		0x00800000, // smallest normal
		0x3effffff, // largest float32 below 0.5
		0x3f000000, // 0.5
		0x3f000001,
		0x3f7fffff, // largest float32 below 1
		0x3f800000, // 1
		0x3f800001,
		0x3fc00000, // 1.5
		0x40000000, // 2
		0x4afffffe, // 2^23 - 1
		0x4affffff, // 2^23 - 0.5, the largest non-integer float32
		0x4b000000, // 2^23
		0x4b000001, // 2^23 + 1
		0x4b7fffff, // 2^24 - 1
		0x4b800000, // 2^24
		0x4effffff, // largest float32 below 2^31
		0x4f000000, // 2^31
		0x5f000000, // 2^63
		0x7f7fffff, // largest finite
		0x7f800000, // +Inf
		0x7fc00000, // quiet NaN
		0x7f800001, // signaling NaN
	}
	vals := make([]float32, 0, 2*len(bits))
	for _, b := range bits {
		vals = append(vals, Float32frombits(b), Float32frombits(b|0x80000000))
	}
	return vals
}()

// randFloat32 returns a random float32 drawn uniformly over all bit patterns.
func randFloat32(r *rand.Rand) float32 {
	return Float32frombits(r.Uint32())
}

func archIterations() int {
	if testing.Short() {
		return 10000
	}
	return 200000
}

func TestArchFMA(t *testing.T) {
	if !haveArchFMA {
		t.Skip("no assembly FMA on this architecture")
	}
	errs := 0
	check := func(x, y, z float32) {
		got, want := archFMA(x, y, z), fma(x, y, z)
		if !sameBits(got, want) && errs < 20 {
			errs++
			t.Errorf("archFMA(%#08x, %#08x, %#08x) = %#08x, fma = %#08x (%g*%g%+g = %g, want %g)",
				Float32bits(x), Float32bits(y), Float32bits(z),
				Float32bits(got), Float32bits(want), x, y, z, got, want)
		}
	}

	for _, x := range archEdgeValues {
		for _, y := range archEdgeValues {
			for _, z := range archEdgeValues {
				check(x, y, z)
			}
		}
	}

	r := rand.New(rand.NewSource(1))
	n := archIterations()
	for i := 0; i < n; i++ {
		// Uniformly random bit patterns.
		x, y, z := randFloat32(r), randFloat32(r), randFloat32(r)
		check(x, y, z)

		// z nearly cancels x*y, so the result depends on the bits of the
		// exact product that a separate multiply would round away.
		x = Float32frombits(uint32(r.Intn(0x40000000)) + 0x10000000) // |x| in [2^-95, 2^33)
		y = Float32frombits(uint32(r.Intn(0x40000000)) + 0x10000000)
		if r.Intn(2) == 0 {
			y = -y
		}
		p := float32(float64(x) * float64(y))
		z = Float32frombits(Float32bits(-p) + uint32(r.Intn(5)) - 2)
		check(x, y, z)

		// Results in or near the subnormal range.
		x = Float32frombits(uint32(r.Intn(1<<29)) + 0x1c000000)            // |x| in [2^-71, 2^-7)
		y = Float32frombits(uint32(r.Intn(1<<29)) + 0x1c000000)            // so |x*y| < 2^-14
		z = Float32frombits(r.Uint32()&0x807fffff | uint32(r.Intn(4))<<23) // |z| < 2^-123
		check(x, -y, z)
	}
}

func TestArchFloorCeilTrunc(t *testing.T) {
	if !haveArchFloor || !haveArchCeil || !haveArchTrunc {
		t.Skip("no assembly Floor/Ceil/Trunc on this architecture")
	}
	errs := 0
	check := func(x float32) {
		for _, f := range []struct {
			name       string
			arch, pure func(float32) float32
		}{
			{"Floor", archFloor, floor},
			{"Ceil", archCeil, ceil},
			{"Trunc", archTrunc, trunc},
		} {
			got, want := f.arch(x), f.pure(x)
			if !sameBits(got, want) && errs < 20 {
				errs++
				t.Errorf("arch%s(%g [%#08x]) = %g [%#08x], want %g [%#08x]",
					f.name, x, Float32bits(x), got, Float32bits(got), want, Float32bits(want))
			}
		}
	}

	for _, x := range archEdgeValues {
		check(x)
		// Neighbours of each edge value, to probe both sides of every boundary.
		if b := Float32bits(x); b&0x7fffffff != 0 && b&0x7fffffff < 0x7f800000 {
			check(Float32frombits(b - 1))
			check(Float32frombits(b + 1))
		}
	}

	r := rand.New(rand.NewSource(1))
	n := archIterations()
	for i := 0; i < n; i++ {
		check(randFloat32(r))
		// |x| in [2^-27, 2^24), where the rounding logic does real work.
		check(Float32frombits(r.Uint32()&0x807fffff | uint32(100+r.Intn(51))<<23))
		// Integers and half-integers near 0.
		k := float32(r.Intn(2001) - 1000)
		check(k)
		check(k + 0.5)
	}
}

func TestArchSqrt(t *testing.T) {
	if !haveArchSqrt {
		t.Skip("no assembly Sqrt on this architecture")
	}
	errs := 0
	check := func(x float32) {
		got, want := archSqrt(x), sqrt(x)
		if !sameBits(got, want) && errs < 20 {
			errs++
			t.Errorf("archSqrt(%g [%#08x]) = %g [%#08x], want %g [%#08x]",
				x, Float32bits(x), got, Float32bits(got), want, Float32bits(want))
		}
	}
	for _, x := range archEdgeValues {
		check(x)
	}
	r := rand.New(rand.NewSource(1))
	n := archIterations()
	for i := 0; i < n; i++ {
		check(randFloat32(r))
		check(Abs(randFloat32(r)))
	}
}
