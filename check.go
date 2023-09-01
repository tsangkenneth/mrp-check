package main

import (
	"fmt"
	"os"
	"strconv"
)

func weight(position int) int {
	modulo := position % 3

	if modulo == 0 {
		return 7
	} else if modulo == 1 {
		return 3
	} else {
		return 1
	}
}

func main() {
	var passportNum string = os.Args[1]
	var positionWeights []int

	for index, character := range passportNum {
		number, _ := strconv.ParseInt(string(character), 10, 0) // TODO: Handle non-number characters
		positionWeights = append(positionWeights, int(number)*weight(index))
	}

	sumOfWeights := 0

	for i := 0; i < len(positionWeights); i++ {
		sumOfWeights = sumOfWeights + positionWeights[i]
	}

	fmt.Println("Value: ", passportNum)
	fmt.Println("Weights: ", positionWeights)
	fmt.Println("Check number: ", sumOfWeights%10)
}
