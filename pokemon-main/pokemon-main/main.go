package main

import (
	"fmt"
	"math/rand"
	"time"

	"pokemon/model" // Import the model package containing the Pokemon struct
)

// catchProbability checks if catching is successful given a percentage rate
func catchProbability(rate int) string {
	if rate < 0 || rate > 100 {
		return "Invalid rate. Please provide a rate between 0 and 100."
	}

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	chance := rand.Intn(100) + 1     // Generate a random number between 1 and 100

	if chance <= rate {
		return "SUCCESS, you caught it"
	}
	return "FAIL, it got away"
}

func main() {
	now := time.Now()

	// Create sample Pokemon
	pokemons := []model.Pokemon{
		{ID: 1, Name: "Pikachu", Type: "Electric", CatchRate: 70, IsRare: false, RegisteredDate: now},
		{ID: 2, Name: "Charmander", Type: "Fire", CatchRate: 40, IsRare: true, RegisteredDate: now.AddDate(0, -1, 0)},
		{ID: 3, Name: "Bulbasaur", Type: "Grass/Poison", CatchRate: 10, IsRare: true, RegisteredDate: now.AddDate(0, -6, 0)},
		{ID: 4, Name: "Dragonite", Type: "Dragon/Flying", CatchRate: 30, IsRare: true, RegisteredDate: now.AddDate(0, -8, 0)},
		{ID: 5, Name: "Mew", Type: "Psychic", CatchRate: 1, IsRare: true, RegisteredDate: now.AddDate(0, -10, 0)},
		{ID: 6, Name: "Niko", Type: "Psychic", CatchRate: 100, IsRare: true, RegisteredDate: now.AddDate(0, -10, 0)},
		{ID: 7, Name: "bene", Type: "air", CatchRate: 100, IsRare: true, RegisteredDate: now.AddDate(0, -10, 0)},
		{ID: 8, Name: "dd", Type: "udara", CatchRate: 100, IsRare: true, RegisteredDate: now.AddDate(0, -10, 0)},
		{ID: 9, Name: "reymond", Type: "udara", CatchRate: 100, IsRare: true, RegisteredDate: now.AddDate(0, -10, 0)},
	}

	fmt.Println("Available Pokémon:")
	for _, pokemon := range pokemons {
		fmt.Printf("ID: %d, Name: %s, Type: %s, Catch Rate: %d%%, Is Rare: %t, Registered Date: %s\n",
			pokemon.ID, pokemon.Name, pokemon.Type, pokemon.CatchRate, pokemon.IsRare, pokemon.RegisteredDate.Format(time.RFC1123))
	}

	var ids [10]int
	var selectedPokemons [10]*model.Pokemon

	for i := 0; i < 10; i++ {
		fmt.Print("Enter Pokémon ID to attempt to catch: ")
		for {
			_, err := fmt.Scanf("%d\n", &ids[i])
			if err != nil {
				fmt.Println("Error:", err)
				fmt.Print("Please enter a valid integer : ")
				continue
			}

			var found bool
			for _, pokemon := range pokemons {
				if pokemon.ID == ids[i] {
					selectedPokemons[i] = &pokemon
					found = true
					break
				}
			}

			if found {
				break
			} else {
				fmt.Printf("Invalid Pokémon ID %d. Please try again.\n", ids[i])
			}
		}
	}

	fmt.Println("\nAttempting to catch Pokémon:")
	for i := 0; i < 10; i++ {
		if selectedPokemons[i] != nil {
			result := catchProbability(selectedPokemons[i].CatchRate)
			fmt.Printf("You attempted to catch %s (%s type) with a catch rate of %d%%: %s\n",
				selectedPokemons[i].Name, selectedPokemons[i].Type, selectedPokemons[i].CatchRate, result)
		}
	}
}
