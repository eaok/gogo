package main

import "fmt"

func main() {
    grade := "E"

    switch marks := 90; {
        case marks >= 90:
            grade = "A"
        case marks >= 80:
            grade = "B"
        default:
            grade = "E"
    }

    fmt.Printf("grade=%s\n", grade)

    switch {
    case false:
            fmt.Println("The integer was <= 4")
            fallthrough
    case true:
            fmt.Println("The integer was <= 5")
            fallthrough
    case false:
            fmt.Println("The integer was <= 6")
            fallthrough
    case true:
            fmt.Println("The integer was <= 7")
    case false:
            fmt.Println("The integer was <= 8")
            fallthrough
    default:
            fmt.Println("default case")
    }
}
