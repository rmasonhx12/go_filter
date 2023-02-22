// simple go filterpackage main
package main

import (
	"fmt"
	"log"
)

type User struct {
    name       string
    occupation string
    country    string
}

func main() {

    users := []User{

        {"John Doe", "gardener", "USA"},
        {"Roger Roe", "driver", "UK"},
        {"Paul Smith", "programmer", "Canada"},
        {"Lucia Mala", "teacher", "Slovakia"},
        {"Patrick Connor", "shopkeeper", "USA"},
        {"Tim Welson", "programmer", "Canada"},
        {"Tomas Smutny", "programmer", "Slovakia"},
		{"Robert M Hendricks", "programmer", "USA"},
    }

    var programmers []User

    for _, user := range users {

        if isProgrammer(user) {
            programmers = append(programmers, user)
        }
    }

    fmt.Println("Programmers:")
    for _, u := range programmers {

        fmt.Println(u) // filter out programmers

		log.Println(u) // log entry of programmers
    }
}

func isProgrammer(user User) bool {

    return user.occupation == "programmer"
}


