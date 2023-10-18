package main

import (
	"fmt"
	menu "github.com/SerafimKuzmin/sd/frontend/menu_controller"
	"net/http"
)

func main() {
	client := &http.Client{}

	err := menu.RunMenu(client)
	if err != nil {
		fmt.Println(err)
	}
}
