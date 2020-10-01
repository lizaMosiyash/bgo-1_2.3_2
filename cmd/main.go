package main

import (
	"fmt"
	"github.com/lizaMosiyash/bgo-1_2.3_2/pkg/card"
	"sync"
	"time"
)

func main() {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	finish := time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local)


	transactions := make([]*card.Transaction, 10_000)
	t := &card.Transaction{
		TimeStamp: time.Date(2020, 4, 3, 15, 36, 0, 0, time.Local).Unix(),
		Sum:       3_000_99,
	}
	transactions = append(transactions, t)
	t = &card.Transaction{
		TimeStamp: time.Date(2020, 4, 3, 15, 36, 0, 0, time.Local).Unix(),
		Sum:       10_688_00,
	}
	transactions = append(transactions, t)
	t = &card.Transaction{
		TimeStamp: time.Date(2020, 5, 3, 15, 36, 0, 0, time.Local).Unix(),
		Sum:       359_00,
	}
	transactions = append(transactions, t)
	t = &card.Transaction{
		TimeStamp: time.Date(2020, 5, 13, 16, 36, 0, 0, time.Local).Unix(),
		Sum:       7_000_00,
	}
	transactions = append(transactions, t)
	t = &card.Transaction{
		TimeStamp: time.Date(2020, 6, 9, 15, 36, 0, 0, time.Local).Unix(),
		Sum:       3_457_00,
	}
	transactions = append(transactions, t)
	t = &card.Transaction{
		TimeStamp: time.Date(2020, 7, 23, 22, 6, 0, 0, time.Local).Unix(),
		Sum:       999_99,
	}
	transactions = append(transactions, t)


	parts := card.GroupTransactionsByMonth(start, finish, transactions)

	var wg sync.WaitGroup
	for i := range parts {
		wg.Add(1)
		go func(part *card.Part) {
			defer wg.Done()

			sum := card.Sum(parts)
			fmt.Println(time.Unix(part.MonthTimestamp, 0).Format("2006-01"), sum)
		}(parts[i])
	}

	wg.Wait()

}