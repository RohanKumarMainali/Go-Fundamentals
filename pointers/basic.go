package pointers

import "fmt"

func Pointer() {
	var name string = "rohan"
	fmt.Println("Initial Name")
	fmt.Println(name)
	var name2 *string = &name
	var name3 string = name

	fmt.Println("----Reference Name----")
	*name2 = "sohan"
	fmt.Println(name, " ", *name2)

	fmt.Println("----Value Name----")
	fmt.Println(name3)
}
