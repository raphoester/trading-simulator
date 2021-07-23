package simulator

import (
	"errors"
	"log"
	"strconv"
	"trading-simulator/internal/binance"
	"trading-simulator/internal/helper"
)

func ApplyCondition(condition Condition, candle binance.Candle) (bool, error) {
	var results []bool
	for i := 0; i < len(condition.Comparisons); i++ {
		// fmt.Println("analyse d'une comparaison : ")
		// helper.ConsoleJson(condition.Comparisons[i])
		res, err := ApplyComparison(condition.Comparisons[i], candle)
		if err != nil {
			return false, err
		}
		// fmt.Println("cette comparaison est", res)
		results = append(results, res)
	}

	for i := 0; i < len(condition.Conditions); i++ {
		// fmt.Println("sous condition trouvée")
		// helper.ConsoleJson(condition.Conditions[i])

		res, err := ApplyCondition(condition.Conditions[i], candle)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("Cette condition est", res)
		results = append(results, res)
	}

	var ret bool
	switch condition.Keyword {
	case "and":
		if helper.IsInBoolList(false, results) {
			ret = false
		} else {
			ret = true
		}
	case "or":
		if helper.IsInBoolList(true, results) {
			ret = true
		} else {
			ret = false
		}

	default:
		return false, errors.New("can't apply boolean operator " + condition.Keyword)
	}

	// fmt.Println("fin d'analyse de la condition : ", ret)
	return ret, nil
}

// todo : factoriser le code
func ApplyComparison(comparison Comparison, candle binance.Candle) (bool, error) {
	var value1 float64
	var value2 float64
	switch comparison.Indicator1.Name {
	case "SMA":
		count, err := strconv.Atoi(comparison.Indicator1.Params["UnitCount"])
		if err != nil {
			return false, errors.New("missing SMA parameter UnitCount")
		}
		value1 = candle.SMA[count]

	case "EMA":
		count, err := strconv.Atoi(comparison.Indicator1.Params["UnitCount"])
		if err != nil {
			return false, errors.New("missing EMA parameter UnitCount")
		}
		value1 = candle.EMA[count]

	case "Close":
		value1 = candle.Close
	default:
		return true, errors.New("unknown indicator : " + comparison.Indicator1.Name)
	}

	switch comparison.Indicator2.Name {
	case "SMA":
		count, err := strconv.Atoi(comparison.Indicator2.Params["UnitCount"])
		if err != nil {
			return false, errors.New("missing SMA parameter UnitCount")
		}
		value2 = candle.SMA[count]

	case "EMA":
		count, err := strconv.Atoi(comparison.Indicator2.Params["UnitCount"])
		if err != nil {
			return false, errors.New("missing EMA parameter UnitCount")
		}
		value2 = candle.EMA[count]

	case "Close":
		value2 = candle.Close
	default:
		return true, errors.New("unknown indicator : " + comparison.Indicator2.Name)
	}
	return BoolOperationFromSymbol(comparison.Operator, value1, value2), nil
}

func BoolOperationFromSymbol(symbol string, a float64, b float64) bool {
	// fmt.Println(a, symbol, b)
	switch symbol {
	case "==":
		return a == b
	case "!=":
		return a != b
	case ">":
		return a > b
	case "<":
		return a < b
	case ">=":
		return a >= b
	case "<=":
		return a <= b
	default:
		return false
	}
}
