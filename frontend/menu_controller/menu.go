package menu_controller

import (
	"fmt"
	errors "github.com/SerafimKuzmin/sd/frontend/errors"
	view "github.com/SerafimKuzmin/sd/frontend/view"
	"net/http"
)

func RunMenu(client *http.Client) error {
	view.PrintRunMenu()

	var who int
	fmt.Scanf("%d", &who)

	switch who {
	case 0:

		err := userMenu(client)
		if err != nil {
			return err
		}
	case 1:

		err := doctorMenu(client)
		if err != nil {
			return err
		}
	case 2:
		!!!!!!!!!!!!!!!!
		err := getDoctors(client)
		if err != nil {
			return err
		}
	default:
		return errors.ErrorCase
	}

	return nil
}
