package problems

import (
	"fmt"
)

/* print 1-100
- divisible by 3 => fizz,
- divisible by 5 => buzz,
- divisible by both => fizzbuzz
*/
func Main_FizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "==>", "fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "==>", "fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "==>", "buzz")
			} else {
			fmt.Println(i, "==>", i)
		}
	}
}

func Main_FizzBuzz2() {
	for i := 1; i <= 100; i++ {
		fuzz := i%3 == 0
		buzz := i%5 == 0
		if fuzz && buzz {
			fmt.Println(i, "==>", "fizzbuzz")
		} else if fuzz {
			fmt.Println(i, "==>", "fizz")
		} else if buzz {
			fmt.Println(i, "==>", "buzz")
		} else {
			fmt.Println(i, "==>", i)
		}
	}
}