package main

import (
	"fmt"
	"github.com/thanthtooaung-coding/jcf-go/collection"
)

func main() {
	var names collection.List[string] = collection.NewArrayList[string]()
	names.Add("Bunny")

	name, err := names.Get(0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Name at index 0: %s\n", name)
	}

	name, err = names.Get(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Name at index 1: %s\n", name)
	}

	err = names.AddAt(1, "Vinn")
	if err != nil {
		return
	}

	name, err = names.Get(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Name at index 1: %s\n", name)
	}

	removed, _ := names.Remove(1)
	fmt.Printf("Removed name '%s'.\n", removed)
	fmt.Println("Final List:", names.ToSlice())

	numbers := collection.NewComparableArrayList[int]()
	numbers.Add(10)
	numbers.Add(20)
	numbers.Add(30)

	fmt.Println("Numbers list:", numbers.ToSlice())
	fmt.Printf("contain 20? %t\n", numbers.Contains(20))
	fmt.Printf("contain 99? %t\n", numbers.Contains(99))
}
