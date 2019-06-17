package main

import (
	"fmt"
)

type Person struct {
	name   string
	job    string
	gender int32
}

func sliceAndArrayExample() {
	nhan := Person{
		name:   "Nhan DInh",
		job:    "IT",
		gender: 1,
	}

	hoang := Person{
		name:   "Hoang Bui",
		job:    "IT",
		gender: 1,
	}

	trung := Person{
		name:   "Trung Huynh",
		job:    "IT",
		gender: 1,
	}
	arrayPerson := [...]Person{
		Person{
			name:   "Nam Dam",
			job:    "IT",
			gender: 1,
		},
		hoang,
		nhan,
	}
	listPerson := arrayPerson[:]
	listPerson = append(listPerson, trung)
	for _, person := range listPerson {
		gender := "Nam"
		if person.gender == 0 {
			gender = "Nu"
		}
		per := "Name: " + person.name + " - Job: " + person.job + " - Gender: " + gender
		fmt.Println(per)
	}
	fmt.Println("------------------------")
	slicePerson := make([]Person, 3, 5)
	for i := 0; i < len(slicePerson); i++ {
		slicePerson[i] = listPerson[i]
	}
	personStrings := printPerson(slicePerson)
	fmt.Println(personStrings)
}

func printPerson(persons []Person) string {
	result := ""
	for _, person := range persons {
		gender := "Nam"
		if person.gender == 0 {
			gender = "Nu"
		}
		per := "Name: " + person.name + " - Job: " + person.job + " - Gender: " + gender
		result += per + "\n"
	}
	return result
}

func main() {
	sliceAndArrayExample()
}
