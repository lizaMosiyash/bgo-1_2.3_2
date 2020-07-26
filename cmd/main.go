package main

import (
	"fmt"
	"github.com/lizaMosiyash/bgo-1_2.3_2/pkg/card"
	"time"
)

func main() {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	finish := time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local)

	months := make([]*card.Part, 0)

	next := start
	for next.Before(finish) {
		months = append(months, &card.Part{
			MonthTimestamp: next.Unix(),
		})
		next = next.AddDate(0, 1, 0)
	}
	months = append(months, &card.Part{
		MonthTimestamp: finish.Unix(),
	})

	transactions := make([]*card.Transaction, 10_000)
	t := &card.Transaction{
		TimeStamp: time.Date(2020, 04, 03, 15, 36, 00, 00, time.Local).Unix(),
		Sum:       3_000_99,
	}
	transactions = append(transactions, t)
	t = &card.Transaction{
		TimeStamp: time.Date(2020, 04, 03, 15, 36, 00, 00, time.Local).Unix(),
		Sum:       10_688_00,
	}
	transactions = append(transactions, t)
	t = &card.Transaction{
		TimeStamp: time.Date(2020, 05, 03, 15, 36, 00, 00, time.Local).Unix(),
		Sum:       359_00,
	}
	transactions = append(transactions, t)
	t = &card.Transaction{
		TimeStamp: time.Date(2020, 05, 13, 16, 36, 00, 00, time.Local).Unix(),
		Sum:       7_000_00,
	}
	transactions = append(transactions, t)
	t = &card.Transaction{
		TimeStamp: time.Date(2020, 06, 9, 15, 36, 00, 00, time.Local).Unix(),
		Sum:       3_457_00,
	}
	transactions = append(transactions, t)
	t = &card.Transaction{
		TimeStamp: time.Date(2020, 07, 23, 22, 06, 00, 00, time.Local).Unix(),
		Sum:       999_99,
	}
	transactions = append(transactions, t)
	for _, transaction := range transactions {
		count := len(months)
		if transaction != nil {
			for i := 0; i < count; i++ {
				c := i+1
				month := months[i]
				if transaction.TimeStamp >= months[i].MonthTimestamp && transaction.TimeStamp < months[c].MonthTimestamp{
					month.PartTransactions = append(month.PartTransactions, transaction)
				}
			}
		}
	}
	fmt.Println(card.SumConcurrently(months, 4))

}