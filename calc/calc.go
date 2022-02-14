package calc

// calc example

type CalcInterface interface {
	sum(a int, b int) int
	sub(a int, b int) int
	mul(a int, b int) int
	div(a int, b int) int
}

type Calc struct{}

func (c *Calc) sum(a int, b int) int {
	return a + b
}

func (c *Calc) sub(a int, b int) int {
	return a - b
}

func (c *Calc) mul(a int, b int) int {
	return a * b
}

func (c *Calc) div(a int, b int) int {
	return a / b
}
