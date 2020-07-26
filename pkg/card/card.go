package card

import (
	"fmt"
	"sync"
)

type Transaction struct {
	TimeStamp int64
	Sum int64
}

type Part struct {
	MonthTimestamp int64
	PartTransactions   []*Transaction
}

func Sum(transactions []*Part) int64 {
	result := int64(0)
	for _, transaction := range transactions {
		for i := range transaction.PartTransactions {
			result += transaction.PartTransactions[i].Sum
		}}
	return result
}

func SumConcurrently(transactions []*Part, goroutines int) int64 {
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	total := int64(0)
	partSize := len(transactions)/goroutines
	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			sum := Sum(part)
			fmt.Println(sum)
			wg.Done()
		}()
	}

	wg.Wait()
	return total
}