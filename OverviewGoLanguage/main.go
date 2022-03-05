package main

import "fmt"

type Pokemons struct {
	Name  string
	Power int
}

func main() {
	animals := make(map[string]string)
	animals["dog"] = "Cesur"
	animals["cat"] = "Boncuk"

	for animalType, animal := range animals {
		fmt.Println(animalType, animal)
	}

	var personnels []string

	personnels = append(personnels, "Kemal", "Mehmet", "Aysel", "Elif")
	for i, personnel := range personnels {
		fmt.Println(i, personnel)
	}

	var pokemon []Pokemons

	pokemon1 := Pokemons{
		Name:  "Charmander",
		Power: 38,
	}
	pokemon2 := Pokemons{
		Name:  "Charizard",
		Power: 98,
	}
	pokemon3 := Pokemons{
		Name:  "Bulbasaur",
		Power: 45,
	}
	pokemon4 := Pokemons{
		Name:  "Ivysaur",
		Power: 60,
	}

	pokemon = append(pokemon, pokemon1, pokemon2, pokemon3, pokemon4)

	for _, p := range pokemon {
		fmt.Println(p.Name, " ", p.Power)
	}

}
