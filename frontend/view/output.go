package view

import (
	"fmt"
	models "github.com/SerafimKuzmin/sd/frontend/models"
	"os"
	"text/tabwriter"
)

func PrintRunMenu() {
	startPosition :=
		`Кто вы?
0 -- пользователь
1 -- админ
2 -- гость
Выберите роль: `

	fmt.Printf("%s", startPosition)
}

func PrintUserMenu() {
	startMenu :=
		`Меню: 
0 -- выйти
1 -- войти
2 -- зарегестрироваться
Выберите действие: `

	fmt.Printf("%s", startMenu)
}

func PrintUserLoop() {
	startMenu := `Меню пользователя: 
0 -- выйти
1 -- создать новый список фильмов
2 -- добавить фильм в список
3 -- оценить фильм
4 -- поиск фильма по ID
5 -- поиск фильма по стране
6 -- получить свои списки фильмов
7 -- получить фильмы по списку
Выберите действие: `

	fmt.Printf("%s", startMenu)
}

func PrintAdminMenu() {
	startMenu :=
		`Меню: 
0 -- выйти
1 -- войти
Выберите действие: `

	fmt.Printf("%s", startMenu)
}

func PrintAdminLoop() {
	startMenu :=
		`Меню админа: 
0 -- выйти
1 -- создать новый список фильмов
2 -- добавить фильм в список
3 -- оценить фильм
4 -- поиск фильма по ID
5 -- поиск фильма по стране
6 -- получить свои списки фильмов
7 -- получить фильмы по списку
8 -- добавить новый фильм
Выберите действие: `

	fmt.Printf("%s", startMenu)
}

func PrintGuestMenu() {
	startMenu :=
		`Меню: 
0 -- выйти
Выберите действие: `

	fmt.Printf("%s", startMenu)
}

func PrintGuestLoop() {
	startMenu := `Меню пользователя: 
0 -- выйти
1 -- поиск фильма по ID
2 -- поиск фильма по стране
Выберите действие: `

	fmt.Printf("%s", startMenu)
}

func PrintFilms(films models.Films) {
	fmt.Printf("\n Фильм:\n")
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 9, 8, 0, '\t', 0)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
		"№", "Id записи", "Название", "Описание", "Рейтинг", "Дата выхода", "Хронометраж")

	for i, r := range films.Films {
		fmt.Fprintf(w, "\n %d\t%d\t%s\t%s\t%f\t%d\t%d\t\n",
			i+1, r.ID, r.Name, r.Description, r.Rate, r.ReleaseDT.Year(), r.Duration)
	}
	w.Flush()

	fmt.Printf("\nКонец записей!\n\n")
}

func PrintFilm(film models.Film) {
	fmt.Printf("\n Фильм:\n")
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 9, 8, 0, '\t', 0)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
		"№", "Id записи", "Название", "Описание", "Рейтинг", "Дата выхода", "Хронометраж")

	fmt.Fprintf(w, "\n %d\t%d\t%s\t%s\t%f\t%d\t%d\t\n",
		film.ID, film.Name, film.Description, film.Rate, film.ReleaseDT.Year(), film.Duration)
	w.Flush()

	fmt.Printf("\nКонец записей!\n\n")
}

func PrintLists(lists models.Lists) {
	fmt.Printf("\n Фильм:\n")
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 9, 8, 0, '\t', 0)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t\n",
		"№", "Id записи", "Название", "Дата создания")

	for i, r := range lists.Lists {
		fmt.Fprintf(w, "\n %d\t%d\t%s\t%s\t\n",
			i+1, r.ID, r.Name, r.CreateDT.String())
	}
	w.Flush()
	fmt.Printf("\nКонец записей!\n\n")
}
