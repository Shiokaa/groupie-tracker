package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

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

func GetChar() Character {
	var characters Character

	err := initApi("https://rickandmortyapi.com/api/character", &characters)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des personnages :", err)
		return Character{}
	}

	return characters
}
