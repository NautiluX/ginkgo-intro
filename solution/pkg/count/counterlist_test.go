package count_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "count/pkg/count"
)

var _ = Describe("Counterlist", func() {

	var (
		counter *CounterList
	)
	Context("Single counter", func() {
		BeforeEach(func() { // Given a counter
			counter, _ = NewCounter(1)
		})

		Context("when a counter outside the range of counters is fetched", func() {
			Context("When a counter above the number of counters is fetched", func() {
				It("raises an error", func() {
					_, err := counter.Get(counter.Size() + 1)
					Expect(err).To(HaveOccurred())
				})
			})
			Context("when a negative counter index is fetched", func() {
				It("raises an error", func() {
					_, err := counter.Get(-1)
					Expect(err).To(HaveOccurred())
				})
			})
		})

		//-------- Scenario 1 --------
		Context("when the counter is increased", func() { // When the counter is increased
			It("increases the count by 1", func() { // Then the new value is 1 higher than the old value
				oldVal, err := counter.Get(0)
				Expect(err).NotTo(HaveOccurred())
				newVal, err := counter.Increment(0)
				Expect(newVal).To(Equal(oldVal + 1))
				Expect(err).NotTo(HaveOccurred())
			})
		})

		//-------- Scenario 2 --------
		Context("when the counter is decreased", func() {
			It("decrements the counter by 1", func() {
				oldVal, err := counter.Get(0)
				Expect(err).NotTo(HaveOccurred())
				newVal, err := counter.Decrement(0)
				Expect(newVal).To(Equal(oldVal - 1))
				Expect(err).NotTo(HaveOccurred())
			})

			//-------- Scenario 3 --------
			Context("when the counter drops below 0", func() {
				It("raises an error", func() {
					err := counter.Set(0, 0)
					Expect(err).NotTo(HaveOccurred())
					_, err = counter.Decrement(0)
					Expect(err).To(HaveOccurred())
				})
			})

		})

		//-------- Scenario 4 --------
		It("is initialized with count 1", func() {
			val, err := counter.Get(0)
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal(1))
		})
	})

	//-------- Scenario 5 --------
	Context("Multiple counters", func() {
		var numCounters int
		BeforeEach(func() { // Given a counter
			numCounters = 42
			counter, _ = NewCounter(numCounters)
		})

		It("can be initialized with a specific number of counters", func() {
			Expect(counter.Size()).To(Equal(numCounters))
		})

		Context("when all counters are increased at once", func() {
			It("increments all the counter values by 1", func() {
				var err error
				oldValues := make([]int, counter.Size())
				for i := range oldValues {
					oldValues[i], err = counter.Get(i)
					Expect(err).ToNot(HaveOccurred())
				}
				err = counter.IncrementAll()
				Expect(err).NotTo(HaveOccurred())
				for i, oldVal := range oldValues {
					newVal, err := counter.Get(i)
					Expect(err).NotTo(HaveOccurred())
					Expect(newVal).To(Equal(oldVal + 1))
				}
			})
		})
	})

})
