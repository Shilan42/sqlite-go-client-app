package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// Client представляет структуру данных клиента
type Client struct {
	ID       int    // Уникальный идентификатор клиента
	FIO      string // ФИО клиента
	Login    string // Логин пользователя
	Birthday string // Дата рождения в формате ГГГГММДД
	Email    string // Электронная почта
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Client.
// Теперь, если передать объект Client в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (c Client) String() string {
	return fmt.Sprintf("ID: %d FIO: %s Login: %s Birthday: %s Email: %s",
		c.ID, c.FIO, c.Login, c.Birthday, c.Email)
}

func main() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// добавление нового клиента
	newClient := Client{
		FIO:      "Шариков Полиграф Полиграфович", // укажите ФИО
		Login:    "Sharikov_PP ",                  // укажите логин
		Birthday: "20250811",                      // укажите день рождения
		Email:    "Sharikov_PP@yandex.ru",         // укажите почту
	}

	id, err := insertClient(db, newClient)
	if err != nil {
		fmt.Println(err)
		return
	}

	// получение клиента по идентификатору и вывод на консоль
	client, err := selectClient(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client)

	// обновление логина клиента
	newLogin := "Sharikov_2025" // укажите новый логин
	err = updateClientLogin(db, newLogin, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	// получение клиента по идентификатору и вывод на консоль
	client, err = selectClient(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client)

	// удаление клиента
	err = deleteClient(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	// получение клиента по идентификатору и вывод на консоль
	_, err = selectClient(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
insertClient добавляет нового клиента в базу данных
Возвращает ID добавленного клиента и ошибку (если есть)
*/
func insertClient(db *sql.DB, client Client) (int64, error) {
	// Выполнение SQL-запроса на вставку данных
	res, err := db.Exec("INSERT INTO clients (fio, login, birthday, email) VALUES (:fio, :login, :birthday, :email)",
		sql.Named("fio", client.FIO),
		sql.Named("login", client.Login),
		sql.Named("birthday", client.Birthday),
		sql.Named("email", client.Email))
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	// Получение ID последней вставленной записи
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

/*
updateClientLogin обновляет логин клиента в базе данных
Принимает новый логин и ID клиента
*/
func updateClientLogin(db *sql.DB, login string, id int64) error {
	// Выполнение SQL-запроса на обновление
	_, err := db.Exec("UPDATE clients SET login = :login WHERE id = :id",
		sql.Named("login", login),
		sql.Named("id", id))
	return err
}

// deleteClient удаляет клиента из базы данных по ID
func deleteClient(db *sql.DB, id int64) error {
	// Выполнение SQL-запроса на удаление
	_, err := db.Exec("DELETE FROM clients WHERE id = :id", sql.Named("id", id))
	return err
}

/*
selectClient получает данные клиента из базы данных по ID
Возвращает объект Client и ошибку (если есть)
*/
func selectClient(db *sql.DB, id int64) (Client, error) {
	client := Client{}
	// Выполнение SQL-запроса на выборку
	row := db.QueryRow("SELECT id, fio, login, birthday, email FROM clients WHERE id = :id", sql.Named("id", id))
	err := row.Scan(&client.ID, &client.FIO, &client.Login, &client.Birthday, &client.Email)

	return client, err
}
