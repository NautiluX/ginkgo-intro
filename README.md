# Ginko Introduction

## Installing ginkgo and gomega


```bash
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega/...
```

## Scenario

A package that can be used to manage a list of number counters. The user can increase or decrease a single one or all of them.
By default, a single counter is created so the user can get started right away.


## User Scenario 1

### User Story

*As a* user,

*I want* to increase a counter by 1,

*so that* I can keep track of received items

### Behaviour description

*Given* a counter

*when* the counter is increased

*Then* the new value is by 1 higher than the old value

### Ginkgo

```go
var _ = Describe("Counterlist", func() {
	var (
		counter *CounterList
	)
	Context("Single counter", func(){
		BeforeEach(func(){                                // Given a counter
			counter = NewCounter()
		})
		Context("when the counter is increased", func(){ // When the counter is increased
			It("increases the count by 1", func(){         // Then the new value is 1 higher than the old value
				oldVal := counter.Get(0)
				newVal := counter.Increment(0)
				Expect(newVal).To(Equal(oldVal + 1))
			})
		})
	})
})
```

### Bootstrapping the tests

```
$ cd pkg/count
$ ginkgo bootstrap
$ ginkgo generate counterlist.go

find .
./counterlist.go
./count_suite_test.go
./counterlist_test.go
```

### Corner cases

* User fetches value outside the range of the list


## User Scenario 2

### User Story

*As a* user,

*I want* to decrease a counter by one

*so that* I can track removed items

### Behaviour description

*Given* a counter

*when* the counter is decreased

*Then* the new value is by 1 lower than the old value

## User Scenario 3

### User Story

*As a* user,

*I want* to get an error when a count drops below 0

*so that* I realize when I accidentally try to remove items from an empty counter

## User Scenario 4

### User Story

*As a* user,

*I want* new counters to start at 1

*so that* they are already initialized with the first item

### Behaviour description

*Given* a new counter

*when* the counter is created

*Then* the starting value is 1

## User Scenario 5

### User Story

*As a* user,

*I want* to increase all my counters at once

*so that* I can add items simulateously

### Behaviour description

*Given* multiple counters

*when* all counters are increased at once

*Then* all values are by 1 higher than the old values

