package strategy

func ExampleStrategy() {
	//Output:
	//pay to tom 100.00 by cash
	//pay to tom 100.00 by card 001
	payment := NewPayment(&Cash{}, 100, "tom", "")

	payment.Pay()
	payment2 := NewPayment(&Card{}, 100, "tom", "001")

	payment2.Pay()
}
