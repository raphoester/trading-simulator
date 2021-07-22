package simulator

import "fmt"

type Wallet struct {
	CUR1 float64
	CUR2 float64
}

func (w *Wallet) BuyCUR1(cur2Amount float64, rate float64) Trade {
	toBuy := cur2Amount / rate
	trade := Trade{
		BuyOrSell: true,
		Rate:      rate,
		CUR1:      toBuy,
		CUR2:      cur2Amount,
	}
	if w.CUR2 > 0 && w.CUR1 == 0 {
		toBuy := cur2Amount / rate
		fmt.Println("ACHAT de", toBuy, "actifs au taux de ", rate)
		w.CUR1 += trade.CUR1
		w.CUR2 -= trade.CUR2
		return trade
	}
	return Trade{}
}

func (w *Wallet) SellAllCUR1(rate float64) Trade {
	trade := Trade{
		BuyOrSell: false,
		Rate:      rate,
		CUR1:      w.CUR1,
		CUR2:      w.CUR1 * rate,
	}
	if w.CUR1 > 0 {
		fmt.Println("VENTE de tous les actifs au taux de ", rate)
		w.CUR2 += trade.CUR2
		w.CUR1 = 0
		return trade
	}
	return Trade{}
}

func (w *Wallet) DisplayCapital() {
	fmt.Printf("CUR1 : %f\n", w.CUR1)
	fmt.Printf("CUR2 : %f\n", w.CUR2)
}
