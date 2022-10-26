package main

import "fmt"

func main() {
	canal := make(chan string) // Vazio

	go func() {
		canal <- "Olá mundo!" // Está cheio
	}()

	msg := <-canal // Canal esvazia
	fmt.Println(msg)
}
