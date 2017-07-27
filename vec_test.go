package geo

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

const (
	e = 1e-10
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestVecFn(t *testing.T) {
	f := func() (x, y float64) {
		return 1, 2
	}
	got := VecFn(f())
	want := Vec{X: 1, Y: 2}
	if !got.Equals(want, e) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestVecLen(t *testing.T) {
	cases := []struct {
		v    Vec
		want float64
	}{
		{Vec{X: 3, Y: 4}, 5},
	}

	for i, c := range cases {
		got := c.v.Len()
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecLen2(t *testing.T) {
	cases := []struct {
		v    Vec
		want float64
	}{
		{Vec{X: 3, Y: 4}, 25},
	}

	for i, c := range cases {
		got := c.v.Len2()
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecSetLen(t *testing.T) {
	cases := []struct {
		v    Vec
		len  float64
		want Vec
	}{
		{Vec{X: 3, Y: 4}, 10, Vec{X: 3, Y: 4}.Normalized().Times(10)},
		{Vec{X: 3, Y: 4}, -10, Vec{X: 3, Y: 4}.Normalized().Times(-10)},
		{Vec{X: 3, Y: 4}, 0, Vec{}},
		{Vec{X: 0, Y: 0}, 1, Vec{}},
	}

	for i, c := range cases {
		c.v.SetLen(c.len)
		if !c.v.Equals(c.want, e) {
			t.Errorf("case %d: got %#v, want %#v", i, c.v, c.want)
		}
	}
}

func TestVecWithLen(t *testing.T) {
	cases := []struct {
		v    Vec
		len  float64
		want Vec
	}{
		{Vec{X: 3, Y: 4}, 10, Vec{X: 3, Y: 4}.Normalized().Times(10)},
		{Vec{X: 3, Y: 4}, -10, Vec{X: 3, Y: 4}.Normalized().Times(-10)},
		{Vec{X: 3, Y: 4}, 0, Vec{}},
		{Vec{X: 0, Y: 0}, 1, Vec{}},
	}

	for i, c := range cases {
		got := c.v.WithLen(c.len)
		if !got.Equals(c.want, e) {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecDist(t *testing.T) {
	trials := 100
	for i := 0; i < trials; i++ {
		v1 := RandVec().Times(rand.Float64() * 100)
		v2 := RandVec().Times(rand.Float64() * 100)
		got := v1.Dist(v2)
		want := v1.Minus(v2).Len()
		if math.Abs(got-want) > e {
			t.Errorf("trial %d: v1: %#v, v2: %#v, got %#v, want %#v", i, v1, v2, got, want)
		}
	}
}

func TestVecDist2(t *testing.T) {
	trials := 100
	for i := 0; i < trials; i++ {
		v1 := RandVec().Times(rand.Float64() * 100)
		v2 := RandVec().Times(rand.Float64() * 100)
		got := v1.Dist2(v2)
		want := v1.Minus(v2).Len2()
		if math.Abs(got-want) > e {
			t.Errorf("trial %d: v1: %#v, v2: %#v, got %#v, want %#v", i, v1, v2, got, want)
		}
	}
}

func TestVecAdd(t *testing.T) {
	cases := []struct {
		v1   Vec
		v2   Vec
		want Vec
	}{
		{Vec{X: 3, Y: -4}, Vec{X: -3, Y: 4}, Vec{X: 0, Y: 0}},
		{Vec{X: 3, Y: 4}, Vec{X: 3, Y: 4}, Vec{X: 6, Y: 8}},
	}

	for i, c := range cases {
		got := c.v1
		got.Add(c.v2)
		if got != c.want {
			t.Errorf("Add case %d: got %#v, want %#v", i, got, c.want)
		}
	}

	for i, c := range cases {
		got := c.v1.Plus(c.v2)
		if got != c.want {
			t.Errorf("Plus case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecSub(t *testing.T) {
	cases := []struct {
		v1   Vec
		v2   Vec
		want Vec
	}{
		{Vec{X: 3, Y: -4}, Vec{X: -3, Y: 4}, Vec{X: 6, Y: -8}},
		{Vec{X: 3, Y: 4}, Vec{X: 3, Y: 4}, Vec{X: 0, Y: 0}},
	}

	for i, c := range cases {
		got := c.v1
		got.Sub(c.v2)
		if got != c.want {
			t.Errorf("Sub case %d: got %#v, want %#v", i, got, c.want)
		}
	}

	for i, c := range cases {
		got := c.v1.Minus(c.v2)
		if got != c.want {
			t.Errorf("Minus case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecMul(t *testing.T) {
	cases := []struct {
		v    Vec
		n    float64
		want Vec
	}{
		{Vec{X: 3, Y: -4}, 2, Vec{X: 6, Y: -8}},
		{Vec{X: 3, Y: 4}, 0, Vec{X: 0, Y: 0}},
	}

	for i, c := range cases {
		got := c.v
		got.Mul(c.n)
		if got != c.want {
			t.Errorf("Mul case %d: got %#v, want %#v", i, got, c.want)
		}
	}

	for i, c := range cases {
		got := c.v.Times(c.n)
		if got != c.want {
			t.Errorf("Times case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecDiv(t *testing.T) {
	cases := []struct {
		v    Vec
		n    float64
		want Vec
	}{
		{Vec{X: 6, Y: -4}, 2, Vec{X: 3, Y: -2}},
		{Vec{X: 3, Y: 4}, 1, Vec{X: 3, Y: 4}},
	}

	for i, c := range cases {
		got := c.v
		got.Div(c.n)
		if got != c.want {
			t.Errorf("Div case %d: got %#v, want %#v", i, got, c.want)
		}
	}

	for i, c := range cases {
		got := c.v.DividedBy(c.n)
		if got != c.want {
			t.Errorf("DividedBy case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecNormalize(t *testing.T) {
	cases := []struct {
		v    Vec
		want Vec
	}{
		{Vec{X: 5, Y: 0}, Vec{X: 1, Y: 0}},
		{Vec{X: 0, Y: -4}, Vec{X: 0, Y: -1}},
	}

	for i, c := range cases {
		got := c.v
		got.Normalize()
		if got != c.want {
			t.Errorf("Normalize case %d: got %#v, want %#v", i, got, c.want)
		}
	}

	for i, c := range cases {
		got := c.v.Normalized()
		if got != c.want {
			t.Errorf("Normalized case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecDot(t *testing.T) {
	cases := []struct {
		v1, v2 Vec
		want   float64
	}{
		{Vec{X: 5, Y: 0}, Vec{X: 1, Y: 0}, 5},
		{Vec{X: 0, Y: -4}, Vec{X: 0, Y: -1}, 4},
		{Vec{X: 1, Y: 2}, Vec{X: 2, Y: -1}, 0},
	}

	for i, c := range cases {
		got := c.v1.Dot(c.v2)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecProject(t *testing.T) {
	cases := []struct {
		v1, v2, want Vec
	}{
		{Vec{X: 0, Y: 5}, Vec{X: 3, Y: 0}, Vec{}},
		{Vec{X: 3, Y: 4}, Vec{X: 7, Y: 0}, Vec{X: 3, Y: 0}},
		{Vec{X: 3, Y: -4}, Vec{X: 7, Y: 0}, Vec{X: 3, Y: 0}},
		{Vec{X: 3, Y: 4}, Vec{X: -7, Y: 0}, Vec{X: 3, Y: 0}},
		{Vec{X: 3, Y: -4}, Vec{X: -7, Y: 0}, Vec{X: 3, Y: 0}},
	}

	for i, c := range cases {
		got := c.v1
		got.Project(c.v2)
		if !got.Equals(c.want, e) {
			t.Errorf("Project case %d: got %#v, want %#v", i, got, c.want)
		}
	}

	for i, c := range cases {
		got := c.v1.Projected(c.v2)
		if !got.Equals(c.want, e) {
			t.Errorf("Projected case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecRand(t *testing.T) {
	trials := 10000
	for i := 0; i < trials; i++ {
		got := RandVec()
		if math.Abs(got.Len()-1) > e {
			t.Errorf("case %d: %#v is length %f", i, got, got.Len())
		}
	}
}

func TestVecLimit(t *testing.T) {
	cases := []struct {
		v    Vec
		len  float64
		want Vec
	}{
		{Vec{X: 5, Y: 0}, 2, Vec{X: 2, Y: 0}},
		{Vec{X: 0, Y: -4}, 2, Vec{X: 0, Y: -2}},
		{Vec{X: 3, Y: 4}, 6, Vec{X: 3, Y: 4}},
	}

	for i, c := range cases {
		got := c.v
		got.Limit(c.len)
		if got != c.want {
			t.Errorf("Limit case %d: len: %#v, got %#v, want %#v", i, c.len, got, c.want)
		}
	}

	for i, c := range cases {
		got := c.v.Limited(c.len)
		if got != c.want {
			t.Errorf("Limited case %d: len: %#v, got %#v, want %#v", i, c.len, got, c.want)
		}
	}
}

func TestVecAngle(t *testing.T) {
	cases := []struct {
		v    Vec
		want float64
	}{
		{Vec{X: 5, Y: 0}, 0},
		{Vec{X: -3, Y: 0}, -math.Pi},
		{Vec{X: 0, Y: -4}, math.Pi / 2},
		{Vec{X: 0, Y: 1}, -math.Pi / 2},
	}

	for i, c := range cases {
		got := c.v.Angle()
		if math.Abs(got-c.want) > 1e-10 {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecAngleFrom(t *testing.T) {
	cases := []struct {
		v1, v2 Vec
		want   float64
	}{
		{Vec{X: 5, Y: 0}, Vec{X: 2, Y: 0}, 0},
		{Vec{X: -5, Y: 0}, Vec{X: -2, Y: 0}, 0},
		{Vec{X: 0, Y: -4}, Vec{X: -1, Y: -1}, -math.Pi / 4},
		{Vec{X: -1, Y: 0}, Vec{X: 1, Y: 0}, -math.Pi},
		{Vec{X: -1, Y: 0}, Vec{X: 0, Y: -1}, math.Pi / 2},
		{Vec{X: -1, Y: 1}, Vec{X: -1, Y: -1}, math.Pi / 2},
		{Vec{X: -1, Y: -1}, Vec{X: -1, Y: 1}, -math.Pi / 2},
	}

	for i, c := range cases {
		got := c.v1.AngleFrom(c.v2)
		if math.Abs(got-c.want) > 1e-10 {
			t.Errorf("case %d: got %#v, want %#v", i, got/math.Pi*180, c.want/math.Pi*180)
		}
	}
}

func TestVecRotate(t *testing.T) {
	cases := []struct {
		v    Vec
		rad  float64
		want Vec
	}{
		{Vec{X: 5, Y: 0}, 0, Vec{X: 5, Y: 0}},
		{Vec{X: 3, Y: 0}, math.Pi / 2, Vec{X: 0, Y: -3}},
		{Vec{X: 3, Y: 0}, -math.Pi / 2, Vec{X: 0, Y: 3}},
		{Vec{X: 3, Y: 0}, math.Pi, Vec{X: -3, Y: 0}},
		{Vec{X: 3, Y: 0}, -math.Pi, Vec{X: -3, Y: 0}},
		{Vec{X: 0, Y: -1}, math.Pi, Vec{X: 0, Y: 1}},
		{Vec{X: 0, Y: -1}, math.Pi / 4, Vec{X: -1, Y: -1}.Normalized()},
		{Vec{X: -1, Y: 0}, -math.Pi, Vec{X: 1, Y: 0}},
		{Vec{X: -1, Y: 0}, 3 * math.Pi / 2, Vec{X: 0, Y: -1}},
	}

	for i, c := range cases {
		got := c.v
		got.Rotate(c.rad)
		if !got.Equals(c.want, e) {
			t.Errorf("Rotate case %d: got %#v, want %#v", i, got, c.want)
		}
	}

	for i, c := range cases {
		got := c.v.Rotated(c.rad)
		if !got.Equals(c.want, e) {
			t.Errorf("Rotated case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecRotateStress(t *testing.T) {
	trials := 10000
	for i := 0; i < trials; i++ {
		v1 := RandVec()
		v2 := RandVec()
		between := v1.AngleFrom(v2)
		rotated := v2.Rotated(between)
		v2rotated := v2
		v2rotated.Rotate(between)
		if !v1.Equals(rotated, e) || !v1.Equals(v2rotated, e) {
			t.Errorf("case %d: v1: %#v v2: %#v between: %#v rotated: %#v v2rotated: %#v",
				i, v1, v2, between, rotated, v2rotated)
		}
	}
}
