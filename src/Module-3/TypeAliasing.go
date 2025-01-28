package Module3

import "fmt"

func TypeAliasing() {
	type MyInt int
	type MyString string
	type MyFloat float64
	type MyBool bool

	var a MyInt = 10
	var b MyString = "Hello"
	var c MyFloat = 10.5
	var d MyBool = true

	fmt.Println(a, b, c, d)

}
