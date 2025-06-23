package main

import (
	"fmt"
	"time" // Adaugat importul pentru pachetul time
)

func main() {
	//! -- ARRAYS
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3]) // Aceasta va afișa un slice format din elementele de la index 1 la 3 (exclusiv)

	// Aici faci reatribuirea, ceea ce este corect
	intArr = [3]int32{1, 2, 3}
	fmt.Println(intArr)

	//! -- SLICES
	var intSlice []int32 = []int32{4, 5, 6} // Declaram un slice, nu un array fix
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7) // Adaugam un element la slice
	fmt.Printf("The ---NEW-- length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	// Acum intSlice este {4, 5, 6, 7}
	fmt.Println(intSlice)

	var intSlice2 []int32 = []int32{8, 9} // Declaram un nou slice intSlice2
	intSlice = append(intSlice, intSlice2...) // Adaugam elementele din intSlice2 la intSlice
	fmt.Println(intSlice) // intSlice va fi acum {4, 5, 6, 7, 8, 9}

	var intSlice3 []int32 = make([]int32, 3, 10) //
	fmt.Println(intSlice3) // Va afisa [0 0 0]


	//! -- MAPS
	var myMap map[string]uint8 = make(map[string]uint8) // Declaram o harta (map)
	myMap["Dog"] = 1 // Adaugam elemente
	myMap["Cat"] = 2
	fmt.Println(myMap["Dog"]) // Accesam elemente

	// O alta modalitate de a declara si initializa o harta
	anotherMap := map[string]uint8{"Dog": 1, "Cat": 2}
	fmt.Println(anotherMap)

	// Verificam daca o cheie exista in harta
	value, ok := myMap["Dog"]
	fmt.Printf("Value: %v, Present: %v\n", value, ok)

	value, ok = myMap["Bird"]
	fmt.Printf("Value: %v, Present: %v\n", value, ok)

	// Stergem un element din harta
	delete(myMap, "Dog")
	fmt.Println(myMap)

	// myMap2 (din imaginea image_d08efa.jpg)
	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45} // Initializare directa
	fmt.Println(myMap2["Adam"]) // Afisam valoarea pentru cheia "Adam"

	var age, ok = myMap2["Jason"] // Verificam existenta cheii "Jason"
	if ok { // Daca cheia exista
		fmt.Printf("The age is %v\n", age)
	} else { // Daca cheia NU exista
		fmt.Println("Invalid Name")
	}

	// Iteram peste elementele din myMap2 (din imaginea image_d08efa.jpg)
	// Ordinea nu este garantata!
	for name, age := range myMap2 { //
		fmt.Printf("Name: %s, Age: %v\n", name, age) // Folosim %s pentru string si %v pentru orice valoare
	}

	//! -- LOOPS (Bucle)
	// Bucle for clasice
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Bucle while-like (for cu conditie)
	var j int = 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Bucle for-each (range) pentru slices
	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
	}

	// Bucle for-each (range) pentru maps
	animalMap := map[string]string{"Dog": "Woof", "Cat": "Meow"}
	for key, value := range animalMap {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Bucle infinite (cu break)
	k := 0
	for {
		if k >= 3 {
			break
		}
		fmt.Println("Looping infinitely until break:", k)
		k++
	}

	// Apelăm noua funcție care conține logica de performanță a slice-urilor
	fmt.Println("\n--- Demonstrating Slice Performance ---") // Adaugă un separator pentru output
	demonstrateSlicePerformance()
}

//!  este o funcție separată pentru a ilustra performanța slice-urilor.
func demonstrateSlicePerformance() {
	var n int = 1000000                                           // Declara o variabilă 'n' cu valoarea 1 milion, folosită pentru iterații.
	var testSlice = []int{}                                            // Declara un slice 'testSlice' fără a-i aloca memorie inițial.
	timeLoop(testSlice, n, "Total time without preallocation: %v\n") // Apeleză timeLoop pentru a măsura timpul adăugării elementelor fără prealocare.

	var testSlice2 = make([]int, 0, n)                            // Declara 'testSlice2', un slice cu lungimea 0 și capacitatea 'n' (1 milion), prealocând memoria.
	timeLoop(testSlice2, n, "Total time with preallocation: %v\n")  // Apeleză timeLoop pentru a măsura timpul adăugării elementelor cu prealocare.
}

//! timeLoop este o funcție auxiliară care măsoară timpul necesar pentru a adăuga 'n' elemente la un slice.
func timeLoop(slice []int, n int, message string) { // Definește funcția timeLoop care primește un slice, un număr de iterații și un mesaj.
	t0 := time.Now()                                    // Înregistrează timpul de început al execuției buclei.
	for len(slice) < n {                                // Buclează atâta timp cât lungimea slice-ului este mai mică decât 'n'.
		slice = append(slice, 1)                        // Adaugă elementul '1' la slice, potențial realocând memorie.
	}
	t1 := time.Now()                                    // Înregistrează timpul de sfârșit al execuției buclei.
	totalTime := t1.Sub(t0)                             // Calculează durata totală a buclei.
	fmt.Printf(message, totalTime)                      // Afișează mesajul de performanță și timpul total.
}