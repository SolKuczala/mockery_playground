package calc

import (
	"mockery_playground/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalc_sum(t *testing.T) {
	newCalc := new(mocks.CalcInterface)
	calc := &Calc{}

	newCalc.On("sum", 1, 2).Return(3)
	n := calc.sum(1, 2)
	assert.Equal(t, 3, n)
}
