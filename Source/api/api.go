package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func initApi(apiUrl string, data interface{}) error {

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodGet, apiUrl, nil)
	if errReq != nil {
		return fmt.Errorf("Requete - Erreur lors de l'initialisation de la requête : %v", errReq)
	}

	res, errRes := httpClient.Do(req)
	if errRes != nil {
		return fmt.Errorf("Requete - Erreur lors de l'envoi de la requête : %v", errRes)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {

		errDecode := json.NewDecoder(res.Body).Decode(data)
		if errDecode != nil {
			return fmt.Errorf("Requete - Erreur lors de la lecture du JSON : %v", errDecode)
		}
	} else {
		return fmt.Errorf("Requete - Erreur code : %v message : %v", res.StatusCode, res.Status)
	}

	return nil
}

type Character struct {
	Results []struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Status  string `json:"status"`
		Species string `json:"species"`
		Gender  string `json:"gender"`
		Images  string `json:"image"`
	} `json:"results"`
}

func GetCharacters() Character {
	var allCharacters Character

	for i := 1; i < 11; i++ {
		pageNbr := strconv.Itoa(i)

		var characters Character

		err := initApi("https://rickandmortyapi.com/api/character?page="+pageNbr, &characters)
		if err != nil {
			fmt.Println("Erreur lors de la récupération des personnages :", err)
			return Character{}
		}

		allCharacters.Results = append(allCharacters.Results, characters.Results...)
	}

	return allCharacters
}

type Location struct {
	Results []struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		Type      string `json:"type"`
		Dimension string `json:"dimension"`
	} `json:"results"`
}

func GetLocation() Location {
	var allLocations Location

	for i := 1; i < 8; i++ {
		pageNbr := strconv.Itoa(i)

		var locations Location

		err := initApi("https://rickandmortyapi.com/api/location?page="+pageNbr, &locations)
		if err != nil {
			fmt.Println("Erreur lors de la récupération des lieux :", err)
			return Location{}
		}

		allLocations.Results = append(allLocations.Results, locations.Results...)
	}

	return allLocations
}

type Episode struct {
	Results []struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Episode  string `json:"episode"`
		Air_date string `json:"air_date"`
	} `json:"results"`
}

func GetEpisode() Episode {
	var allEpisodes Episode

	for i := 1; i < 4; i++ {
		pageNbr := strconv.Itoa(i)

		var episodes Episode

		err := initApi("https://rickandmortyapi.com/api/episode?page="+pageNbr, &episodes)
		if err != nil {
			fmt.Println("Erreur lors de la récupération des épisodes :", err)
			return Episode{}
		}

		allEpisodes.Results = append(allEpisodes.Results, episodes.Results...)
	}

	return allEpisodes
}
