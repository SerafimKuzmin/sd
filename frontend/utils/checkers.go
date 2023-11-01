package utils

import (
	"encoding/json"
	basicErrors "errors"
	errors "github.com/SerafimKuzmin/sd/frontend/errors"
	models "github.com/SerafimKuzmin/sd/frontend/models"
	"io"
	"math/rand"
	"net/http"
)

func CheckErrorInBody(response *http.Response) error {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.ErrorReadBody
	}

	var result models.ErrorBody
	if err := json.Unmarshal(body, &result); err != nil {
		return errors.ErrorParseBody
	}

	return basicErrors.New(result.Err)
}

func DoAndCheckRequest(client *http.Client, request *http.Request) (*http.Response, error) {
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return &http.Response{}, CheckErrorInBody(response)
	}

	return response, nil
}

func ParseClientBody(response *http.Response) (models.Client, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.Client{}, errors.ErrorReadBody
	}

	var result models.Client
	if err := json.Unmarshal(body, &result); err != nil {
		return models.Client{}, errors.ErrorParseBody
	}

	return result, nil
}

func ParseUserBody(response *http.Response) (models.User, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.User{}, errors.ErrorReadBody
	}

	var result models.User
	if err := json.Unmarshal(body, &result); err != nil {
		return models.User{}, errors.ErrorParseBody
	}

	return result, nil
}

func ParseFilmBody(response *http.Response) (models.Film, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.Film{}, errors.ErrorReadBody
	}

	var result models.FilmResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return models.Film{}, errors.ErrorParseBody
	}

	return result.Body, nil
}

func ParseFilmsBody(response *http.Response) (models.FilmsResponse, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.FilmsResponse{}, errors.ErrorReadBody
	}

	var result models.FilmsResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return models.FilmsResponse{}, errors.ErrorParseBody
	}

	return result, nil
}

func ParseListsBody(response *http.Response) (models.Lists, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.Lists{}, errors.ErrorReadBody
	}

	var result models.Lists
	if err := json.Unmarshal(body, &result); err != nil {
		return models.Lists{}, errors.ErrorParseBody
	}

	return result, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetToken() string {
	n := 10
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
