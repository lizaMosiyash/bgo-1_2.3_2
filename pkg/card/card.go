package card

import (
	"fmt"
	"sync"
	"time"
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

func GroupTransactionsByMonth(start, finish time.Time, transactions []*Transaction) []*Part {
	months := make([]*Part, 0)
	parts := make([]*Part, 0)

	next := start
	for next.Before(finish) {
		months = append(months, &Part{
			MonthTimestamp: next.Unix(),
		})
		next = next.AddDate(0, 1, 0)
	}
	months = append(months, &Part{
		MonthTimestamp: finish.Unix(),
	})
	for _, transaction := range transactions {
		count := len(months)
		if transaction != nil {
			for i := 0; i < count; i++ {
				month := months[i]
				t1 := time.Unix(months[i].MonthTimestamp, 0)
				t2 := time.Unix(transaction.TimeStamp, 0)
				if t1.Year() == t2.Year() && t1.Month() == t2.Month() {
					month.PartTransactions = append(month.PartTransactions, transaction)
				}
				parts = append(parts, month)}
		}
	}
return parts
}
