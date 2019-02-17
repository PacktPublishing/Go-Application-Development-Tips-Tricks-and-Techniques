package customer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerMayNotBuyBeerInUSAIf20(t *testing.T) {
	sut := Customer{"John", 20}
	mayBuyBeer, err := sut.MayBuyBeer("USA")

	assert.Nilf(t, err, "unexpected error: %s", err)
	assert.Falsef(t, mayBuyBeer, "should not be able to buy beer at %d", sut.Age)
}

func TestCustomerMayBuyBeerInGERIf20(t *testing.T) {
	sut := Customer{"John", 20}
	mayBuyBeer, err := sut.MayBuyBeer("GER")

	assert.Nilf(t, err, "unexpected error: %s", err)
	assert.Truef(t, mayBuyBeer, "should be able to buy beer at %d", sut.Age)
}

func TestCustomerMayBuyBeerInUSAIf21(t *testing.T) {
	sut := Customer{"John", 21}
	mayBuyBeer, err := sut.MayBuyBeer("USA")

	assert.Nilf(t, err, "unexpected error: %s", err)
	assert.Truef(t, mayBuyBeer, "should be able to buy beer at %d", sut.Age)
}

func TestCustomerMayBuyBeerInGERIf16(t *testing.T) {
	sut := Customer{"John", 16}
	mayBuyBeer, err := sut.MayBuyBeer("GER")

	assert.Nilf(t, err, "unexpected error: %s", err)
	assert.Truef(t, mayBuyBeer, "should be able to buy beer at %d", sut.Age)
}

func TestCustomerMayNotBuyBeerInUSAIf15(t *testing.T) {
	sut := Customer{"John", 15}
	mayBuyBeer, err := sut.MayBuyBeer("USA")

	assert.Nilf(t, err, "unexpected error: %s", err)
	assert.Falsef(t, mayBuyBeer, "should not be able to buy beer at %d", sut.Age)
}

func TestCustomerMayNotBuyBeerInGERIf15(t *testing.T) {
	sut := Customer{"John", 15}
	mayBuyBeer, err := sut.MayBuyBeer("GER")

	assert.Nilf(t, err, "unexpected error: %s", err)
	assert.Falsef(t, mayBuyBeer, "should not be able to buy beer at %d", sut.Age)
}

func TestUnknownCountryCodeReturnsError(t *testing.T) {
	sut := Customer{"John", 23}
	mayBuyBeer, err := sut.MayBuyBeer("FRA")

	assert.NotNil(t, err)
	assert.False(t, mayBuyBeer)
}
