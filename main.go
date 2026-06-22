package main

import "fmt"

//sub1
func square(n int) int {
	return n * n

}


//sub2
func maxNumber(f, d int) int {
	if f > d {
		return f
	} else {
		return d
	}

}

//sub3

func isEven(s int) bool {
	if s%2 == 0 {
		return true
	} else {
		return false
	}
}

//sub4

func greetUser(x string) {
	fmt.Println("Привет,", x)

}

//sub5

func sumSlice(l []int) int {
	summ := 0
	for _, value := range l {
		summ += value
	}
	return summ
}

// sub 6

func checkLogin(login, password string) bool {
	var login1 = map[string]string{
		"login": "admin", "password": "1234",
	}

	if login1["login"] == login && login1["password"] == password {
		return true
	} else {
		return false
	}
}

// sub 7

func increaseBalance(b *int, g int) {
	*b += g
}

// sub 8

func resetAttempts(r *int) {
	if *r > 3 {
		*r = 0
	}
}

func main() {
	result := square(5)
	result2 := square(10)
	result3 := square(15)

	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)

	result4 := maxNumber(101, 100)

	fmt.Println(result4)

	result5 := isEven(10)
	fmt.Println(result5)

	greetUser("sdasdasd")

	greetUser("fffff")

	fmt.Println(sumSlice([]int{5, 6, 5, 6, 4, 5, 4}))

	fmt.Println(checkLogin("admin", "1234"))
	balance := 0
	increaseBalance(&balance, 100)
	fmt.Println(balance)
	attempts := 5
	resetAttempts(&attempts)
	fmt.Println(attempts)
}
