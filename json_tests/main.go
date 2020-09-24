package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	First       string
	Last        string
	Age         int `json:"Experience"`
	notExported int
}

type SecretAgent struct {
	Person
	First              string
	HasAssignedMission bool
}

func (p Person) fullName() string {
	return p.First + " " + p.Last
}

//method override
func (p Person) Greeting() {
	fmt.Println("I'm just a normal person.")
}

func (a SecretAgent) Greeting() {
	fmt.Println("<<I'm just a normal person>> I promise!! Wink.")
}

func main() {
	p1 := Person{"Jane", "Sanders", 31, 120}
	p2 := &Person{"John", "Scotty", 30, 125}

	fmt.Println("\n------User defined types-------")
	fmt.Println(p2)
	fmt.Printf("Type of p2: %T \n", p2)
	fmt.Printf("Value of p2: %v \n", p2)
	fmt.Println(p1)
	fmt.Printf("Type of p1: %T \n", p1)
	fmt.Printf("Value of p1: %v \n", p1)

	fmt.Println(p1.First, p1.Last, p1.Age)
	fmt.Println(p2.First, p2.Last, p2.Age)

	fmt.Println(p2.fullName())
	fmt.Println("-------------")

	fmt.Println("\n------Composition-------")

	agent001 := SecretAgent{
		Person{
			"Magenta",
			"Piggy",
			21,
			126,
		},
		"Outer Turquoise",
		true,
	}

	agent002 := SecretAgent{
		Person: Person{
			First:       "Flying",
			Last:        "Rat",
			Age:         23,
			notExported: 127,
		},
		First:              "Jumping",
		HasAssignedMission: true,
	}
	fmt.Println(agent001.First, agent001.Person.First)
	fmt.Println(agent002.First, agent002.Person.First)
	fmt.Println("------------------------------")

	fmt.Println("\n------Method override-------")

	agent001.Greeting()
	agent001.Person.Greeting()
	p1.Greeting()
	fmt.Println("------------------------------")

	fmt.Println("\n------JSON Marshal-------")
	bs, _ := json.Marshal(p1)
	fmt.Println("Marshalled p1: ", bs)
	fmt.Printf("Typeof bs: %T\n ", bs)
	fmt.Println("string(bs): ", string(bs))
	fmt.Println("---------------------------")

	fmt.Println("\n------JSON Unmarshal-------")
	var p3 Person
	fmt.Println("p3 first:", p3.First)
	fmt.Println("p3 Last:", p3.Last)
	fmt.Println("p3 Age:", p3.Age)

	bs2 := []byte(`{"First": "Gil", "Last": "Mill", "Experience": 30}`)
	json.Unmarshal(bs2, &p3)
	fmt.Println("After unmarshal:")

	fmt.Println("p3 first:", p3.First)
	fmt.Println("p3 Last:", p3.Last)
	fmt.Println("p3 Age:", p3.Age)
	fmt.Printf("Type: %T:", p3)
	fmt.Println("\n------------------------")

	fmt.Println("\n------JSON Encode-------")
	p4 := Person{"James", "Bond", 40, 007}
	json.NewEncoder(os.Stdout).Encode(p4)

	fmt.Println("------------------------")

	fmt.Println("\n------JSON Decode-------")
	var p5 Person
	rdr := strings.NewReader(`{"First": "Julia", "Last": "Boss", "Experience": 30}`)
	json.NewDecoder(rdr).Decode(&p5)
	fmt.Println("p5 first:", p5.First)
	fmt.Println("p5 Last:", p5.Last)
	fmt.Println("p5 Age:", p5.Age)
	fmt.Println("------------------------")

}
