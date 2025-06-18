package greedy

import (
	"math"
	"testing"
)

func floatEquals(a, b float64) bool {
	const eps = 1e-9
	return math.Abs(a-b) < eps
}

func TestFractionalKnapsack_Basic(t *testing.T) {
	weights := []float64{10, 20, 30}
	values := []float64{60, 100, 120}
	capacity := 50.0
	got := FractionalKnapsack(weights, values, capacity)
	want := 240.0
	if !floatEquals(got, want) {
		t.Errorf("FractionalKnapsack() = %v, want %v", got, want)
	}
}

func TestFractionalKnapsack_ZeroCapacity(t *testing.T) {
	weights := []float64{10, 20, 30}
	values := []float64{60, 100, 120}
	capacity := 0.0
	got := FractionalKnapsack(weights, values, capacity)
	want := 0.0
	if !floatEquals(got, want) {
		t.Errorf("FractionalKnapsack() = %v, want %v", got, want)
	}
}

func TestFractionalKnapsack_EmptyItems(t *testing.T) {
	weights := []float64{}
	values := []float64{}
	capacity := 100.0
	got := FractionalKnapsack(weights, values, capacity)
	want := 0.0
	if !floatEquals(got, want) {
		t.Errorf("FractionalKnapsack() = %v, want %v", got, want)
	}
}

func TestFractionalKnapsack_CapacityGreaterThanTotalWeight(t *testing.T) {
	weights := []float64{5, 10}
	values := []float64{10, 30}
	capacity := 20.0
	got := FractionalKnapsack(weights, values, capacity)
	want := 40.0
	if !floatEquals(got, want) {
		t.Errorf("FractionalKnapsack() = %v, want %v", got, want)
	}
}

func TestFractionalKnapsack_TakeFraction(t *testing.T) {
	weights := []float64{5, 10}
	values := []float64{15, 20}
	capacity := 7.0
	got := FractionalKnapsack(weights, values, capacity)
	want := 15.0 + 2.0*2.0 // 5kg (15) + 2kg (2)
	if !floatEquals(got, want) {
		t.Errorf("FractionalKnapsack() = %v, want %v", got, want)
	}
}

func TestFractionalKnapsack_AllZeroValues(t *testing.T) {
	weights := []float64{5, 10, 15}
	values := []float64{0, 0, 0}
	capacity := 20.0
	got := FractionalKnapsack(weights, values, capacity)
	want := 0.0
	if !floatEquals(got, want) {
		t.Errorf("FractionalKnapsack() = %v, want %v", got, want)
	}
}

func TestFractionalKnapsack_AllZeroWeights(t *testing.T) {
	weights := []float64{0, 0, 0}
	values := []float64{10, 20, 30}
	capacity := 10.0
	got := FractionalKnapsack(weights, values, capacity)
	if !math.IsNaN(got) {
		t.Errorf("FractionalKnapsack() = %v, want NaN", got)
	}
}
