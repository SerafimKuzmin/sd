package view

import (
	"fmt"
	errors "github.com/SerafimKuzmin/sd/frontend/errors"
	models "github.com/SerafimKuzmin/sd/frontend/models"
)

func InputCred() (string, string, error) {
	var login string
	fmt.Printf("Введите email: ")
	_, err := fmt.Scanf("%s", &login)
	if err != nil {
		return "", "", errors.ErrorInput
	}

	var password string
	fmt.Printf("Введите пароль: ")
	_, err = fmt.Scanf("%s", &password)
	if err != nil {
		return "", "", errors.ErrorInput
	}

	return login, password, nil
}

func CreateListData() (models.List, error) {
	fmt.Printf("Введите название списка: ")
	var name string
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		return models.List{}, errors.ErrorInput
	}

	return models.List{Name: name}, nil
}

func AddFilmToListData() (models.FilmList, error) {
	fmt.Printf("Введите ID списка: ")
	var listId uint64
	_, err := fmt.Scanf("%d", &listId)
	if err != nil {
		return models.FilmList{}, errors.ErrorInput
	}

	fmt.Printf("Введите ID фильма: ")
	var filmId uint64
	_, err = fmt.Scanf("%d", &filmId)
	if err != nil {
		return models.FilmList{}, errors.ErrorInput
	}

	return models.FilmList{ListID: listId, FilmID: filmId}, nil
}

func RateFilmData() (models.RateFilm, error) {

	fmt.Printf("Введите ID фильма: ")
	var filmId uint64
	_, err := fmt.Scanf("%d", &filmId)
	if err != nil {
		return models.RateFilm{}, errors.ErrorInput
	}

	fmt.Printf("Введите оценку (от 0 до 10): ")
	var rate float64
	_, err = fmt.Scanf("%f", &rate)
	if err != nil {
		return models.RateFilm{}, errors.ErrorInput
	}

	return models.RateFilm{Rate: rate, FilmID: filmId}, nil
}

func GetID() (uint64, error) {

	fmt.Printf("Введите ID : ")
	var id uint64
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		return 0, errors.ErrorInput
	}

	return id, nil
}

func CreateFilmData() (models.Film, error) {

	fmt.Printf("Введите название фильма: ")
	var name string
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		return models.Film{}, errors.ErrorInput
	}
	fmt.Printf("Введите описание фильма: ")
	var desc string
	_, err = fmt.Scanf("%s", &desc)
	if err != nil {
		return models.Film{}, errors.ErrorInput
	}
	fmt.Printf("Введите рейтинг: ")
	var rate float64
	_, err = fmt.Scanf("%f", &rate)
	if err != nil {
		return models.Film{}, errors.ErrorInput
	}
	fmt.Printf("Введите жанр: ")
	var genre string
	_, err = fmt.Scanf("%s", &genre)
	if err != nil {
		return models.Film{}, errors.ErrorInput
	}
	fmt.Printf("Введите продолжительность: ")
	var dur uint
	_, err = fmt.Scanf("%d", &dur)
	if err != nil {
		return models.Film{}, errors.ErrorInput
	}

	return models.Film{Name: name, Duration: dur, Description: desc, Rate: rate, Genre: genre}, nil
}
