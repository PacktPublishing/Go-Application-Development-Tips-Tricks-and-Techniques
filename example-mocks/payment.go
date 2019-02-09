package customer

type CreditScoreProvider interface {
	CreditScoreForCustomer(c *Customer) (int, error)
}

type PaymentMethodProvider struct {
	CSP CreditScoreProvider
}

func (p *PaymentMethodProvider) MethodsForCustomer(c *Customer) ([]string, error) {
	methods := []string{"prepay", "paypal"}
	score, err := p.CSP.CreditScoreForCustomer(c)
	if err != nil {
		return nil, err
	}

	if score >= 80 {
		methods = append(methods, "creditcard")
	}

	if score >= 90 {
		methods = append(methods, "invoice")
	}

	return methods, nil
}
