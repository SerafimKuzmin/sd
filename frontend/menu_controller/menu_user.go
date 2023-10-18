package menu_controller

import (
	"fmt"
	errors "github.com/SerafimKuzmin/sd/frontend/errors"
	handlers "github.com/SerafimKuzmin/sd/frontend/handlers"
	models "github.com/SerafimKuzmin/sd/frontend/models"
	utils "github.com/SerafimKuzmin/sd/frontend/utils"
	view "github.com/SerafimKuzmin/sd/frontend/view"
	"net/http"
)

func userMenu(client *http.Client) error {

	view.PrintUserMenu()
	var num int
	fmt.Scanf("%d", &num)

	var token string
	var err error

	switch num {
	case 0:
		return nil
	case 1:
		token, err = loginClient(client)
		if err != nil {
			return err
		}
		fmt.Printf("\nПользователь успешно авторизован!\n\n")
	case 2:
		token, err = createClient(client)
		if err != nil {
			return err
		}
		fmt.Printf("\nПользователь успешно добавлен!\n\n")
	default:
		return errors.ErrorCase
	}

	return userLoop(client, token)
}

// 0 -- выйти
// 1 -- создать новый список фильмов
// 2 -- добавить фильм в список
// 3 -- оценить фильм
// 4 -- поиск фильма по ID
// 5 -- поиск фильма по стране
// 6 -- получить свои списки фильмов
func userLoop(client *http.Client, token string) error {
	var num int = 1
	var err error

	for num != 0 {
		view.PrintUserLoop()

		fmt.Scanf("%d", &num)

		if num == 0 {
			return nil
		}

		switch num {
		case 1:
			err := createList(client, token)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				fmt.Printf("\nСписок успешно создан!\n\n")
			}
		case 2:
			err = addFilmToList(client, token)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				fmt.Printf("\nФильм успешно добавлен!\n\n")
			}
		case 3:
			err = rateFilm(client, token)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				fmt.Printf("\nФильм успешно оценен!\n\n")
			}
		case 4:
			err = getFilmByID(client, token)
			if err != nil {
				fmt.Println(err)
			}
		case 5:
			err = getFilmByCountry(client, token)
			if err != nil {
				fmt.Println(err)
			}
		case 6:
			err = getLists(client, token)
			if err != nil {
				fmt.Println(err)
			}
		case 7:
			err = getFilmByList(client, token)
			if err != nil {
				fmt.Println(err)
			}
		default:
			return errors.ErrorCase
		}
	}

	return nil
}

func loginClient(client *http.Client) (string, error) {
	login, password, err := view.InputCred()
	if err != nil {
		return "", err
	}
	newClient := models.Client{Login: login, Password: password}

	response, err := handlers.Login(client, &newClient)
	if err == errors.ErrorResponseStatus {
		return "", utils.CheckErrorInBody(response)
	} else if err != nil {
		return "", err
	}

	result, err := utils.ParseClientBody(response)
	if err != nil {
		return "", err
	}

	return result.Token, nil
}

func createClient(client *http.Client) (string, error) {
	login, password, err := view.InputCred()
	if err != nil {
		return "", err
	}
	newClient := models.User{Login: login, Password: password}

	response, err := handlers.Create(client, &newClient)
	if err == errors.ErrorResponseStatus {
		return "", utils.CheckErrorInBody(response)
	} else if err != nil {
		return "", err
	}

	result, err := utils.ParseClientBody(response)
	if err != nil {
		return "", err
	}

	return result.Token, nil
}

func createList(client *http.Client, token string) error {
	body, err := view.CreateListData()
	if err != nil {
		return err
	}

	response, err := handlers.CreateList(client, token, &body)
	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	return nil
}

func addFilmToList(client *http.Client, token string) error {
	body, err := view.AddFilmToListData()
	if err != nil {
		return err
	}

	response, err := handlers.AddFilmList(client, token, &body)
	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	return err
}

func rateFilm(client *http.Client, token string) error {
	body, err := view.RateFilmData()
	if err != nil {
		return err
	}

	response, err := handlers.RateFilm(client, token, &body)
	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	return err
}

func getFilmByCountry(client *http.Client, token string) error {
	body, err := view.GetID()
	if err != nil {
		return err
	}

	response, err := handlers.GetFilmByCountry(client, body)
	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	records, err := utils.ParseFilmsBody(response)
	if err != nil {
		return err
	}

	view.PrintFilms(records)

	return nil
}

func getFilmByID(client *http.Client, token string) error {
	body, err := view.GetID()
	if err != nil {
		return err
	}

	response, err := handlers.GetFilmByID(client, body)
	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	records, err := utils.ParseFilmBody(response)
	if err != nil {
		return err
	}

	view.PrintFilm(records)

	return nil
}

func getFilmByList(client *http.Client, token string) error {
	body, err := view.GetID()
	if err != nil {
		return err
	}

	response, err := handlers.GetFilmsByList(client, token, body)
	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	records, err := utils.ParseFilmsBody(response)
	if err != nil {
		return err
	}

	view.PrintFilms(records)

	return nil
}

func getLists(client *http.Client, token string) error {
	body, err := view.GetID()
	if err != nil {
		return err
	}

	response, err := handlers.GetLists(client, token, body)
	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	records, err := utils.ParseListsBody(response)
	if err != nil {
		return err
	}

	view.PrintLists(records)

	return nil
}
