package main

import (
	"encoding/json"
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
	URL  string `json:"url"`
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
	//fmt.Println(string(responseData))

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	fmt.Println(responseObject.Pokemon[0].Species.URL)

	for i := 0; i < len(responseObject.Pokemon); i++ {
		fmt.Println(responseObject.Pokemon[i].Species.Name)
	}
}
