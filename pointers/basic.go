package pointers

import "fmt"

type Creature struct {
	Species string
}

func changeCreature(creature *Creature) {
	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}

func add(a int32, b int32) int32 {
	a += 2
	return a + b
}

func addByReference(a *int32, b int32) int32 {
	*a += 2
	return *a + b
}

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

	var firstNum int32 = 3
	var secondNum int32 = 5
	ans := addByReference(&firstNum, secondNum)
	fmt.Println(ans)
	fmt.Println("first  ", firstNum)

	var creature Creature = Creature{Species: "shark"}

	fmt.Printf("1) %+v\n", creature)
	changeCreature(&creature)
	fmt.Printf("3) %+v\n", creature)
}
