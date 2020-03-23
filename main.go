package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Response armazenará o retorno completo da chamada
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

//Pokemon mapeará todos os pokémons
type Pokemon struct {
	EntryNumber int            `json:"entry_number"`
	Species     PokemonSpecies `json:"pokemon_species"`
}

//PokemonSpecies armazena o nome das espécies
type PokemonSpecies struct {
	Name string `json:"name"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
}
