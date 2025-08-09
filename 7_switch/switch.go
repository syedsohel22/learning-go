package main

import (
	"fmt"
	"time"
)


func main(){
	// simpe switch
	// Does not need break like others and defualt is optional
	i:=5
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("not in 1 to 5")
	}

	// multipal condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
		default:
		fmt.Println("It's a weekday")
	}

	// type switch

	 whoIam := func(i interface{}){
		switch t:=i.(type){
		case string:
			fmt.Println("I'm a string",t)
		case int:
			fmt.Println("I'm an int",t)
		case bool:
			fmt.Println("I'm a bool",t)
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whoIam("hello")
	whoIam(1)
	whoIam(true)


}