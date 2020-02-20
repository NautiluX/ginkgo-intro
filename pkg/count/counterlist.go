package count

type CounterList struct {
	counters []int
}

func NewCounter() *CounterList {
	return &CounterList{}
}

func (c *CounterList) Increment(index int) int {
	return 0
}

func (c *CounterList) Get(index int) int {
	return 0
}
