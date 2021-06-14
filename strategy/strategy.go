package strategy

import "fmt"

type PayContext struct {
	To     string
	Money  float64
	CardId string
}

type Strategy interface {
	Pay(*PayContext)
}

type Payment struct {
	Context  *PayContext
	Strategy Strategy
}

type Cash struct{}

type Card struct{}

func (c *Cash) Pay(context *PayContext) {
	fmt.Printf("pay to %s %.2f by cash\n", context.To, context.Money)
}
func (c *Card) Pay(context *PayContext) {
	fmt.Printf("pay to %s %.2f by card %s\n", context.To, context.Money, context.CardId)
}

func NewPayment(s Strategy, money float64, to string, cardId string) *Payment {
	return &Payment{
		Strategy: s,
		Context:  &PayContext{To: to, Money: money, CardId: cardId},
	}
}

func (p *Payment) Pay() {
	p.Strategy.Pay(p.Context)
}
