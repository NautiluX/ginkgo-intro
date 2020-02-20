package count

import "fmt"

type CounterList struct {
	counters []int
}

func NewCounter(numCounters int) (*CounterList, error) {
	counterList := CounterList{make([]int, numCounters)}
	for i := range counterList.counters {
		err := counterList.Set(i, 1)
		if err != nil {
			return nil, err
		}
	}
	return &counterList, nil
}

func (c *CounterList) Increment(index int) (int, error) {
	c.counters[index] = c.counters[index] + 1
	return c.Get(index)
}

func (c *CounterList) Get(index int) (int, error) {
	if index >= c.Size() || index < 0 {
		return -1, fmt.Errorf("woops, I don't have that counter")
	}
	return c.counters[index], nil
}

func (c *CounterList) Decrement(index int) (int, error) {
	if c.counters[index] == 0 {
		return 0, fmt.Errorf("counter already at 0, can't decrement further.")
	}
	c.counters[index] = c.counters[index] - 1
	return c.Get(index)
}

func (c *CounterList) Size() int {
	return len(c.counters)
}

func (c *CounterList) Set(index int, value int) error {
	if index >= c.Size() || index < 0 {
		return fmt.Errorf("woops, I don't have that counter")
	}
	c.counters[index] = value
	return nil
}

func (c *CounterList) IncrementAll() error {
	for i := range c.counters {
		_, err := c.Increment(i)
		if err != nil {
			return err
		}
	}
	return nil
}
