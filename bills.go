package main

import (
	"fmt"
	"os"
)

func incSlice(s []bill) {
	for i := range s {
		if s[i].amount == 2 {
			s[i].amount = 0
		} else {
			s[i].amount++
			return
		}
	}
}

type bill struct {
	nom    int
	amount int
}

func main() {
	var n, m, sum int
	var mid bill
	var bills []bill

	fmt.Fscan(os.Stdin, &n)
	fmt.Fscan(os.Stdin, &m)

	mid.amount = 0
	for i := 0; i < m; i++ {
		fmt.Fscan(os.Stdin, &mid.nom)
		bills = append(bills, mid)
	}

	for {
		incSlice(bills)
		sum = 0
		for _, mid := range bills {
			sum += mid.nom * mid.amount
		}
		if sum == 0 {
			fmt.Println(-1)
			break
		}
		if sum == n {
			//прописать вывод(успешный)
			amount := 0
			var output []int
			for _, mid := range bills {
				amount += mid.amount
				for i := 0; i < mid.amount; i++ {
					output = append(output, mid.nom)
				}
			}
			fmt.Println(amount)

			for _, j := range output {
				fmt.Print(j)

			}
			break
		}
	}

}
