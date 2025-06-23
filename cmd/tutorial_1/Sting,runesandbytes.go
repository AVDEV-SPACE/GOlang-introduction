package main

import (
    "fmt"
    "strings"
)

func main() {
    // Exemplu 1: Lucrul cu rune-uri
    var myString = []rune("résumé") // Convertim string-ul în slice de rune pentru a lucra corect cu Unicode
    var indexed = myString[1]       // Indexăm al doilea rune (é)
    fmt.Printf("%v, %T\n", indexed, indexed) // Afișăm valoarea și tipul

    // Parcurgem rune-urile
    for i, v := range myString {
        fmt.Println(i, v)
    }

    // Afișăm lungimea în rune-uri
    fmt.Printf("\nThe length of 'myString' is %v", len(myString))

    // Exemplu 2: Manipularea șirurilor
    var myRune = 'a'
    fmt.Printf("\nmyRune = %v", myRune)

    var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
    var catStr = ""
    for i := range strSlice {
        catStr += strSlice[i]
    }

    fmt.Printf("\n%v", catStr) // Afișăm șirul concatenat ("subscribe")
}