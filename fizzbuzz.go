package main

import (
		"fmt"
		"reflect"
)

func main() {
		for i := 0; i < 102; i++ {

			switch {
			case i%15 == 0:
					fmt.Println("FizzBuzz")
			case i%5 == 0:
					fmt.Println("Buzz")
			case i%3 == 0:
					fmt.Println("Fizz")
			default:
					fmt.Println(i)
			}
		}
		v := 1
		r := reflect.ValueOf(v)
		fmt.Println(r.Type())
		fmt.Println(r.Kind())
}