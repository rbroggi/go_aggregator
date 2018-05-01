package aggregator

import (
	"testing"
)

func TestIndexWiseFunc(t *testing.T) {
	testArray := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	expectedOne := []float64{0.0, 0.0, 0.0, 0.0, 0.0}
	expectedTwo := []float64{2.0, 4.0, 6.0, 8.0, 10.0}
	tables := []struct {
		arrayOne []float64
		arrayTwo []float64
		f        func(float64, float64) float64
		expected []float64
	}{
		{testArray, testArray, func(a, b float64) float64 { return a - b }, expectedOne},
		{testArray, testArray, func(a, b float64) float64 { return a + b }, expectedTwo},
	}

	for _, table := range tables {
		result := IndexWiseFunc(table.arrayOne, table.arrayTwo, table.f)
		for i, _ := range table.expected {
			if !floatEquality(table.expected[i], result[i]) {
				t.Errorf("IndexWiseFunc wrong behavior - expected %d and got %d for index %d", table.expected[i], result[i], i)
			}
		}
	}

}

func floatEquality(a, b float64) bool {
	if a-b < 0.000001 {
		return true
	}
	return false
}
