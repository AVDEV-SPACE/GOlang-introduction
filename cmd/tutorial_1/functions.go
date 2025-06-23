package main

import (
	"errors"
	"fmt"
)

// Corectat: eliminat '/' rătăcit de la sfârșitul liniei
func intDivision(numerator int, denominator int) (int, int, error) {
	var err error // Aceste linii trebuie să fie în interiorul funcției

	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	
	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, nil
}


func main() {
	var printValue string = "Func with spec. type"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)

	if err != nil {
		fmt.Printf(err.Error() + "\n")
	} else {
		fmt.Printf("test result of the integer division is %v with remainder %v\n", result, remainder)
	}

	// Corectat: adaugat ':' dupa conditia case
	switch {
	case err != nil:
		fmt.Printf(err.Error() + "\n") // Adaugat \n pentru consistenta
	case remainder == 0: // <-- Aici era problema! Lipsea ':'
		fmt.Printf("SWITCH FUNC --The result of the integer division is %v\n", result) // Adaugat \n pentru consistenta
	default:
		fmt.Printf("SWITCH FUNC --test result of the integer division is %v\n with remainder %v\n", result, remainder)
	}

	switch remainder{
	case 0:
		fmt.Printf("REMAINDER SWITCH -- The division was exact")
	case 1,2:
		fmt.Printf("REMAINDER SWITCH -- The division was exact")
	default:
		fmt.Printf("REMAINDER SWITCH -- The division was not close")
	}


	result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error() + "\n")
	} else {
		fmt.Printf("test result of the integer division is %v\n with remainder %v\n", result, remainder)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

