package main

import (
    "fmt"
    "strings"
)

func main() {
    message := "jlhzhyjpwolypzvsk" //this is the message
    
     for shiftit := 1; shiftit <= 26; shiftit++ {    //loop all shift possibilities
	    fmt.Print("shift ", shiftit, ": ")
     	fmt.Println(shift(message, shiftit))
    }
}

func shift(msg string, shift int) string {
    
    bases := ""

    for i := 97; i < 123; i++ {
        bases += string(rune(i))
    }

    res := ""

    for c := 0; c < len(msg); c++ {
        ix := strings.Index(bases, string(msg[c]))  + shift
        if(ix < 0) { ix *= -1 }
        ix = ix % len(bases)	
        res += 	string(bases[ix])
    }
    
    return res
}       
