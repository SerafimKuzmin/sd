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

func adminMenu(client *http.Client) error {

	view.PrintAdminMenu()
	var num int
	fmt.Scanf("%d", &num)

	var token string
	var err error

	switch num {
	case 0:
		return nil
	case 1:
		token, err = loginAdmin(client)
		if err != nil {
			return err
		}
		fmt.Printf("\nАдмин успешно авторизован!\n\n")
	default:
		return errors.ErrorCase
	}

	return adminLoop(client, token)
}

func adminLoop(client *http.Client, token string) error {
	var num int = 1
	var err error

	for num != 0 {
		view.PrintAdminLoop()

		fmt.Scanf("%d", &num)

		if num == 0 {
			return nil
		}

		switch num {
		case 1:
			err := createListAdmin(client, token)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				fmt.Printf("\nСписок успешно создан!\n\n")
			}
		case 2:
			err = addFilmToListAdmin(client, token)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				fmt.Printf("\nФильм успешно добавлен!\n\n")
			}
		case 3:
			err = rateFilmAdmin(client, token)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				fmt.Printf("\nФильм успешно оценен!\n\n")
			}
		case 4:
			err = getFilmByIDAdmin(client)
			if err != nil {
				fmt.Println(err)
			}
		case 5:
			err = getFilmByCountryAdmin(client)
			if err != nil {
				fmt.Println(err)
			}
		case 6:
			err = getListsAdmin(client, token)
			if err != nil {
				fmt.Println(err)
			}
		case 7:
			err = getFilmsByListAdmin(client, token)
			if err != nil {
				fmt.Println(err)
			}
		case 8:
			err = createFilm(client, token)
			if err != nil {
				fmt.Println(err)
			}
		default:
			return errors.ErrorCase
		}
	}

	return nil
}

func loginAdmin(client *http.Client) (string, error) {
	login, password, err := view.InputCred()
	if err != nil {
		return "", err
	}

	newAdmin := models.Client{Login: login, Password: password}

	response, err := handlers.Login(client, &newAdmin)
	if err == errors.ErrorResponseStatus {
		return "", utils.CheckErrorInBody(response)
	} else if err != nil {
		return "", err
	}

	_, err = utils.ParseUserBody(response)
	if err != nil {
		return "", err
	}

	return utils.GetToken(), nil
}

func createListAdmin(client *http.Client, token string) error {
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

func addFilmToListAdmin(client *http.Client, token string) error {
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

func rateFilmAdmin(client *http.Client, token string) error {
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

func getFilmByCountryAdmin(client *http.Client) error {
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

func getFilmByIDAdmin(client *http.Client) error {
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

func getFilmsByListAdmin(client *http.Client, token string) error {
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

func getListsAdmin(client *http.Client, token string) error {
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

func createFilm(client *http.Client, token string) error {
	body, err := view.CreateFilmData()
	if err != nil {
		return err
	}

	response, err := handlers.AddFilm(client, token, &body)
	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	}

	return nil
}
