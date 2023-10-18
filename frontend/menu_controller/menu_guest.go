package menu_controller

import (
	"fmt"
	errors "github.com/SerafimKuzmin/sd/frontend/errors"
	handlers "github.com/SerafimKuzmin/sd/frontend/handlers"
	utils "github.com/SerafimKuzmin/sd/frontend/utils"
	view "github.com/SerafimKuzmin/sd/frontend/view"
	"net/http"
)

func guestMenu(client *http.Client) error {

	view.PrintUserMenu()
	var num int
	fmt.Scanf("%d", &num)

	var token string

	switch num {
	case 0:
		return nil
	default:
		return errors.ErrorCase
	}

	return guestLoop(client, token)
}

func guestLoop(client *http.Client, token string) error {
	var num int = 1
	var err error

	for num != 0 {
		view.PrintGuestLoop()

		fmt.Scanf("%d", &num)

		if num == 0 {
			return nil
		}

		switch num {
		case 1:
			err = getFilmByIDGuest(client, token)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			err = getFilmByCountryGuest(client, token)
			if err != nil {
				fmt.Println(err)
			}
		default:
			return errors.ErrorCase
		}
	}

	return nil
}

func getFilmByCountryGuest(client *http.Client, token string) error {
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

func getFilmByIDGuest(client *http.Client, token string) error {
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
