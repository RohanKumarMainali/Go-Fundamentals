package dsa

import "fmt"

func PrintMap() {
	// bad practice : because when
	// A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic
	// var m map[string]int

	m := make(map[string]int)
	m["test"] = 1
	fmt.Println(m)
	fmt.Println(m["test"])
	_, present := m["location"]
	println(present)

	// if key is not present in map,
	// it will give zero value
	// NOTE: zero value ande 0 is not same

	type Person struct {
		Name  string
		Likes []string
	}
	people := []*Person{
		{
			Name:  "rohan",
			Likes: []string{"bacon"},
		},
		{
			Name:  "anil",
			Likes: []string{"bacon"},
		},
		{
			Name:  "manish",
			Likes: []string{"cheesse"},
		},
	}

	likes := make(map[string][]*Person)

	for _, val := range people {
		for _, like := range val.Likes {
			likes[like] = append(likes[like], val)
		}
	}

	for _, like := range likes["cheesse"] {
		fmt.Println(like.Name, "likes cheese")
	}
}
