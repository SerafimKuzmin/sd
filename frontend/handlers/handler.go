package handlers

import (
	"bytes"
	"fmt"
	errors "github.com/SerafimKuzmin/sd/frontend/errors"
	models "github.com/SerafimKuzmin/sd/frontend/models"
	"net/http"
)

const port = "8080"
const address = "localhost"

func DoRequest(client *http.Client, request *http.Request) (*http.Response, error) {
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.ErrorExecuteRequest
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return response, errors.ErrorResponseStatus
	}

	return response, nil
}

func Login(client *http.Client, newClient *models.Client) (*http.Response, error) {

	url := "http://" + address + ":" + port + "/signin"
	params := fmt.Sprintf("{\"email\": \"%s\", \"password\": \"%s\"}", newClient.Login, newClient.Password)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := DoRequest(client, request)

	return response, err
}

func Create(client *http.Client, user *models.User) (*http.Response, error) {

	url := "http://" + address + ":" + port + "/singup"
	params := fmt.Sprintf("{\"login\": \"%s\", \"password\": \"%s\", \"role\": \"%d\", \"email\": \"%s\", \"full_name\": \"%s\", \"is_active\": \"%d\", \"country_id\": \"%d\"}",
		user.Login, user.Password, user.Role, user.Email, user.FullName, user.IsActive, user.CountryID)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := DoRequest(client, request)

	return response, err
}

func GetFilmByID(client *http.Client, id uint64) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/film/"
	url = fmt.Sprintf("url%d", id)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := DoRequest(client, request)

	return response, err
}

func GetFilmByCountry(client *http.Client, id uint64) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/country/"
	url = fmt.Sprintf("url%d/films", id)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := DoRequest(client, request)

	return response, err
}

func GetLists(client *http.Client, token string, id uint64) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/user/"
	url = fmt.Sprintf("url%d/lists", id)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := DoRequest(client, request)

	return response, err
}

func GetFilmsByList(client *http.Client, token string, id uint64) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/list/"
	url = fmt.Sprintf("url%d/films", id)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := DoRequest(client, request)

	return response, err
}

func AddFilmList(client *http.Client, token string, listFilm *models.FilmList) (*http.Response, error) {

	url := "http://" + address + ":" + port + "/list/add"
	params := fmt.Sprintf("{\"list_id\": \"%d\", \"film_id\": \"%d\"}",
		listFilm.ListID, listFilm.FilmID)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := DoRequest(client, request)

	return response, err
}

func RateFilm(client *http.Client, token string, rate *models.RateFilm) (*http.Response, error) {

	url := "http://" + address + ":" + port + "/personal_rating/create"
	params := fmt.Sprintf("{\"id\": \"%d\", \"user_id\": \"%d\", \"film_id\": \"%d\", \"rate\": \"%f\"}",
		rate.ID, rate.UserID, rate.FilmID, rate.Rate)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := DoRequest(client, request)

	return response, err
}

func AddFilm(client *http.Client, token string, film *models.Film) (*http.Response, error) {

	url := "http://" + address + ":" + port + "/film/create"
	params := fmt.Sprintf("{\"id\": \"%d\", \"name\": \"%s\", \"description\": \"%s\", \"rate\": \"%f\", \"genre\": \"%s\", \"duration\": \"%d\", \"country_id\": \"%d\"}",
		film.ID, film.Name, film.Description, film.Rate, film.Genre, film.Duration, film.CountryID)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := DoRequest(client, request)

	return response, err
}

func CreateList(client *http.Client, token string, list *models.List) (*http.Response, error) {

	url := "http://" + address + ":" + port + "/list/create"
	params := fmt.Sprintf("{\"name\": \"%s\"}", list.Name)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := DoRequest(client, request)

	return response, err
}
