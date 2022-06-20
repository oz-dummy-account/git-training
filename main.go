package main

import (
	"fmt"
)

func main() {
	fmt.Println("Git Demo")

	fmt.Println("Resolving git merge while rebase")

	ch := make(chan string, 3)

	ch <- "Veda"
	ch <- "Omega"
	ch <- "Heyjay"
	close(ch)

	for s := range ch {
		fmt.Println("Employee Name: ", s)
	}
}
