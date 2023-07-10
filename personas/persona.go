package personas

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/garnermccloud/Crowd-GPT/external"
)

// Persona is a struct that represents a persona
type Persona struct {
	Description string
}

// Person is a struct that represents a person.  It is a specific instance of a Persona.
type Person struct {
	Name        string
	Description string
}

func CrowdPersonas() []*Persona {
	return []*Persona{
		{
			Description: "Opinions Specialists: Individuals who have professional or expert knowledge in a particular subject matter and therefore, can provide skilled insights.",
		},
		{
			Description: "Generalists: Individuals with broad, general knowledge who can offer multiple perspectives on a given topic.",
		},
		{
			Description: "Outsiders: Individuals who are not intimately familiar with the topic, lending an unbiased perspective which can offer fresh or unconventional insights.",
		},
		{
			Description: "Contrarians: Those who often hold opposing viewpoints from the majority, offering a challenging perspective that encourages the exploration of all facets of an issue.",
		},
		{
			Description: "Comparable Industry Perspective: Those who are experienced or knowledgeable in comparable industries can add a different viewpoint to the common knowledge of the crowd",
		},
		{
			Description: "Everyday Users: Depending on the topic, average individuals or consumers who will give an on-the-ground perspective.",
		},
	}
}

// NewPerson is a function that returns a new Person
func NewPerson(name string, description string) *Person {
	return &Person{
		Name:        name,
		Description: description,
	}
}

// String is a method that returns a string representation of a Person
func (p *Person) String() string {
	return fmt.Sprintf("Name: %s, Description: %s", p.Name, p.Description)
}

// GetPersonsFromAI is a function that returns new Persons from an AI based on a given Persona and question topic
func GetPersonaPersonsFromAI(persona *Persona, topic string) []*Person {
	_ = external.GetAIClient()
	ctx := context.Background()

	// ask the AI for a person
	message := fmt.Sprintf(`Create three random people who fit the following persona and subject matter.
	Persona: %s
	Subject Matter: %s
	
	 Please respond in json following this schema:
	[{"name": "Joe", "description": "Joe's description"}]
	`, persona.Description, topic)
	aiResponse, err := external.AskAIGPT35Turbo(message, ctx)

	if err != nil {
		fmt.Println(err)
	}

	// parse the AI response into a Person. The response is a json string that looks like this: [{"name": "Joe", "description": "Joe's description"}]

	persons := []*Person{}
	err = json.Unmarshal([]byte(aiResponse), &persons)
	if err != nil {
		fmt.Println(err)
	}

	return persons

}
