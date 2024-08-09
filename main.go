package main

import "fmt"

func main() {
	// Reto 2: Encriptar mensaje
	fmt.Println("Reto 2: Encriptar mensaje")

	key := "dcj"
	message := "I love prOgrAmming!"

	fmt.Println(encryptMessage(key, message))

	// Reto 3: Suma a cero
	fmt.Println("Reto 3: Suma a cero")

	arr := []int{3, 4, -7, 5, -6, 2, 5, -1, 8}
	fmt.Println(removeZeroSum(arr))
}
