package main

import (
	"fmt"

	"github.com/garnermccloud/Crowd-GPT/personas"
)

func main() {

	pSlice := personas.CrowdPersonas()
	topic := "What is the meaning of life?"
	persons := make([]*personas.Person, 0)
	for _, p := range pSlice {
		personSlice := personas.GetPersonaPersonsFromAI(p, topic)
		for _, person := range personSlice {

			fmt.Println(person)
		}

		persons = append(persons, personSlice...)
	}
	fmt.Println(persons)
}
