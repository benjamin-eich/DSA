package greedy

import "sort"

/*
Fractional Knapsack Problem
The fractional knapsack problem asks: Given a set of items, each with a Weight and a Value, how do you maximize the
Value in the knapsack while staying within a Weight limit? In the fractional version, you can take fractions of items.
*/

/*
FractionalKnapsack calculates the maximum Value in the Knapsack for the given parameters
*/
func FractionalKnapsack(weights, values []float64, capacity float64) float64 {
	type RatioItem struct {
		Ratio  float64
		Weight float64
		Value  float64
	}
	ratioItems := make([]RatioItem, len(weights))

	for i := 0; i < len(weights); i++ {
		ratioItems[i] = RatioItem{float64(values[i]) / float64(weights[i]), weights[i], values[i]}
	}

	sort.Slice(ratioItems, func(i, j int) bool {
		return ratioItems[i].Ratio > ratioItems[j].Ratio
	})

	value := 0.0
	for _, ratioItem := range ratioItems {
		if capacity <= 0 {
			return value
		}

		weightToTake := min(ratioItem.Weight, capacity)

		value += weightToTake * ratioItem.Ratio
		capacity -= weightToTake
	}
	return value
}
