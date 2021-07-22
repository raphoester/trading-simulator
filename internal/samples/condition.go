package simulator

import "trading-simulator/internal/simulator"

func ConditionTest() simulator.Condition {
	// Représentation d'une condition de rang 2 :
	// Si la moyenne mobile en 10 est supérieure au prix
	// ET QUE
	// la moyenne mobile en 50 est supérieure à la moyenne mobile en 20
	cnd := simulator.Condition{
		Keyword: "and",
		Comparisons: []simulator.Comparison{
			{
				Operator: "==",
				Indicator1: simulator.Indicator{
					Name:   "movingAverage",
					Params: map[string]string{"UnitCount": "10"},
				},
				Indicator2: simulator.Indicator{
					Name: "price",
					// pas de paramètres : attribut natif
				},
			},
			{
				Operator: ">",
				Indicator1: simulator.Indicator{
					Name:   "movingAverage",
					Params: map[string]string{"UnitCount": "20"},
				},
				Indicator2: simulator.Indicator{
					Name:   "movingAverage",
					Params: map[string]string{"UnitCount": "50"},
				},
			},
		},
	}
	return cnd
}
