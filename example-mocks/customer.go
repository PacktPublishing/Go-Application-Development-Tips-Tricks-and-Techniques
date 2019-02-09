package customer

import "fmt"

type Customer struct {
	Name string
	Age  int
}

func (c *Customer) MayBuyBeer(country string) (bool, error) {
	switch country {
	case "USA":
		return c.Age >= 21, nil
	case "GER":
		return c.Age >= 16, nil
	default:
		return false, fmt.Errorf("unknown country %s", country)
	}
}
