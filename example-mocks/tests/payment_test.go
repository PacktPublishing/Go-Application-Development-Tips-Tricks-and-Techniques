package customer_tests

import (
	"testing"

	"github.com/golang/mock/gomock"
	customer "github.com/martin-helmich/go-tips/example-mocks"
	customer_mocks "github.com/martin-helmich/go-tips/example-mocks/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNoCCWithScoreLessThan80(t *testing.T) {
	cus := customer.Customer{Name: "John", Age: 25}
	ctrl := gomock.NewController(t)
	csp := customer_mocks.NewMockCreditScoreProvider(ctrl)
	csp.EXPECT().CreditScoreForCustomer(&cus).Return(79, nil)
	sut := customer.PaymentMethodProvider{CSP: csp}

	methods, err := sut.MethodsForCustomer(&cus)

	assert.Nil(t, err)
	assert.NotContains(t, methods, "creditcard")
}

func BenchmarkPaymentMethods(b *testing.B) {
	cus := customer.Customer{Name: "John", Age: 25}
	ctrl := gomock.NewController(b)
	csp := customer_mocks.NewMockCreditScoreProvider(ctrl)
	csp.EXPECT().CreditScoreForCustomer(&cus).Return(79, nil).AnyTimes()
	sut := customer.PaymentMethodProvider{CSP: csp}

	for n := 0; n < b.N; n++ {
		sut.MethodsForCustomer(&cus)
	}
}
