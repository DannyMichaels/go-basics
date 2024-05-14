package main

import (
	"fmt"
	"math"
)

func main() {
	
	var investmentAmount float64 = 1000

	// when type is not specified can declare variable with := operator instead of var
	expectedReturnRate := 5.5;
	var years float64 = 10;

	// when type is not specified can declare variable with := operator instead of var
	futureValue := investmentAmount * math.Pow((1 + expectedReturnRate / 100), years);
	fmt.Println(futureValue);

}