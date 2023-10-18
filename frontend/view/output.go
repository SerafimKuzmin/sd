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
7 -- добавить новый фильм
Выберите действие: `

	fmt.Printf("%s", startMenu)
}

func PrintRecords(records models.Records) {
	fmt.Printf("\n Ваши записи:\n")
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 9, 8, 0, '\t', 0)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
		"№", "Id записи", "Питомец", "Клиент", "Доктор", "Год", "Месяц", "День", "Начало", "Конец")

	for i, r := range records.Records {
		fmt.Fprintf(w, "\n %d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\n",
			i+1, r.RecordId, r.PetId, r.ClientId, r.DoctorId, r.DatetimeStart.Year(), r.DatetimeStart.Month(),
			r.DatetimeStart.Day(), r.DatetimeStart.Hour(),
			r.DatetimeEnd.Hour())
	}
	w.Flush()

	fmt.Printf("\nКонец записей!\n\n")
}

func PrintPets(pets models.Pets) {
	fmt.Printf("\nВаши питомцы:\n")
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 9, 8, 0, '\t', 0)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\n",
		"№", "Id питомца", "Кличка", "Тип", "Возраст", "Уровень здоровья")

	for i, p := range pets.Pets {
		fmt.Fprintf(w, "\n %d\t%d\t%s\t%s\t%d\t%d\n",
			i+1, p.PetId, p.Name, p.Type, p.Age, p.Health)
	}
	w.Flush()

	fmt.Printf("\nКонец!\n\n")
}

func PrintDoctors(doctors models.Doctors) {
	fmt.Printf("\nДоктора клиники:\n")
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 9, 8, 0, '\t', 0)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\n",
		"№", "Id доктора", "Логин доктора", "Начало приема", "Конец приема")

	for i, d := range doctors.Doctors {
		fmt.Fprintf(w, "\n %d\t%d\t%s\t%d\t%d\n",
			i+1, d.DoctorId, d.Login, d.StartTime, d.EndTime)
	}
	w.Flush()

	fmt.Printf("\nКонец!\n\n")
}

func PrintClientInfo(client models.Client) {
	fmt.Printf("\nВаш логин: %s\nВаш Id: %d\n\n", client.Login, client.ClientId)
}

func PrintDoctorInfo(doctor models.Doctor) {
	fmt.Printf("\nВаш логин: %s\nВаш Id: %d\nЕжедневное время начала приема: %d\nЕжедневное время конца приема: %d\n\n",
		doctor.Login, doctor.DoctorId, doctor.StartTime, doctor.EndTime)
}
