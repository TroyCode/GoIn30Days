package main

import (
	"fmt"
	"github.com/TroyCode/GoIn30Days/day6/alien"
	"github.com/TroyCode/GoIn30Days/day6/person"
)

func main() {
	fmt.Println(person.SayHelloTo("George"))
	fmt.Println(person.MyName)
	fmt.Println(alien.AlienName)
}
