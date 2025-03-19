package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*
Database structure.
	Table-1: BOOKS:
		| id INT AUTO_INCREMENT PRIMARY KEY | name varchar(20) | Description varchar(100)| URL varchar(50)| Author varchar(20) | Category varchar(20) |
		|                 1                 |    Java Book     |   ...                   | ...            |  ...               |                      |
	Table-2: USERS:
		| id INT AUTO_INCREMENT PRIMARY KEY | Name varchar(20) | Password varchar(20) |
		|               1                   | ...              | ...                  |
*/

func CreateBooksTable() error { // ++
	db, err := sql.Open("mysql", "books:pass@tcp(127.0.0.1:3306)/web_books")
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	result, err := db.Exec("CREATE TABLE IF NOT EXISTS books (id int auto_increment primary key, name varchar(20) not null, description varchar(50) not null, url varchar(50) not null, author varchar(20) not null, category varchar(20) not null);")
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())
	return nil
}

func CreateUsersTable() error { // ++
	db, err := sql.Open("mysql", "books:pass@tcp(127.0.0.1:3306)/web_books")
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	result, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id int auto_increment primary key, name varchar(20) not null, password varchar(20) not null);")
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())
	return nil
}

func InsertBook(name string, description string, url string, author string, category string) error { // ++
	db, err := sql.Open("mysql", "books:pass@tcp(127.0.0.1:3306)/web_books")
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	result, err := db.Exec("insert into web_books.books (name, description, url, author, category) values (?, ?, ?, ?, ?)", name, description, url, author, category)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())
	return nil
}

func InsertUser(name string, password string) error { // ++
	db, err := sql.Open("mysql", "books:pass@tcp(127.0.0.1:3306)/web_books")
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	result, err := db.Exec("insert into web_books.users (name, password) values (?, ?)", name, password)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())
	return nil
}

type Book struct {
	Id          int
	Name        string
	Description string
	Url         string
	Author      string
	Category    string
}

func GetAllBooks() ([]Book, error) { // ++
	db, err := sql.Open("mysql", "books:pass@tcp(127.0.0.1:3306)/web_books")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select * from web_books.books")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		b := Book{}
		err = rows.Scan(&b.Id, &b.Name, &b.Description, &b.Url, &b.Author, &b.Category)
		if err != nil {
			log.Println(err)
			continue
		}
		books = append(books, b)
	}

	return books, nil
}

func GetCategoryBooks(category string) ([]Book, error) { /////////////// !!!!!!!!!!!! // --
	db, err := sql.Open("mysql", "books:pass@tcp(127.0.0.1:3306)/web_books")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select * from web_books.books where category like ?", category)
	// SELECT * FROM web_books.Books WHERE category LIKE ?;
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		b := Book{}
		err = rows.Scan(&b.Id, &b.Name, &b.Description, &b.Url, &b.Author, &b.Category)
		if err != nil {
			log.Println(err)
			continue
		}
		books = append(books, b)
	}
	return books, nil
}

type User struct {
	Id       int
	Name     string
	Password string
}

func GetAllUsers() ([]User, error) { // ++
	db, err := sql.Open("mysql", "books:pass@tcp(127.0.0.1:3306)/web_books")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select * from web_books.users")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.Id, &u.Name, &u.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, u)
	}

	return users, nil
}

func UpdateBook(id int, name string, description string, url string, author string, category string) error { // update by html-form // --
	db, err := sql.Open("mysql", "books:pass@tcp(127.0.0.1:3306)/web_books")
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	// id int auto_increment primary key, name varchar(20) not null, description varchar(50) not null,
	// url varchar(50) not null, author varchar(50) nor null, category varchar(20) not null)

	result, err := db.Exec("update web_books.books set name = ? where id = ?", name, id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Update name results:")
	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())

	result, err = db.Exec("update web_books.books set description = ? where id = ?", description, id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Update description results:")
	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())

	result, err = db.Exec("update web_books.books set url = ? where id = ?", url, id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Update url results:")
	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())

	result, err = db.Exec("update web_books.books set author = ? where id = ?", author, id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Update author results:")
	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())

	result, err = db.Exec("update web_books.books set category = ? where id = ?", category, id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Update category results:")
	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())

	return nil
}

func UpdateUser(id int, name string, password string) error { // --
	db, err := sql.Open("mysql", "books:pass@tcp(127.0.0.1:3306)/web_books")
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	result, err := db.Exec("update web_books.users set name = ? where id = ?", name, id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Update name results:")
	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())

	result, err = db.Exec("update web_books.users set password = ? where id = ?", password, id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Update password results:")
	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())

	return nil
}

func DeleteBook(id int) error { // --
	db, err := sql.Open("mysql", "books:pass@tcp(127.0.0.1:3306)/web_books")
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	result, err := db.Exec("delete from web_books.books where id = ?", id)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())
	return nil
}

func DeleteUser(id int) error { // --
	db, err := sql.Open("mysql", "books:pass@tcp(127.0.0.1:3306)/web_books")
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	result, err := db.Exec("delete from web_books.users where id = ?", id)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())
	return nil
}

type DbId struct {
	Id int
}

func GetAllBooksId() ([]DbId, error) {
	db, err := sql.Open("mysql", "books:pass@tcp(127.0.0.1:3306)/web_books")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select id from web_books.books")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	id_slice := []DbId{}
	for rows.Next() {
		i := DbId{}
		err = rows.Scan(&i.Id)
		if err != nil {
			log.Println(err)
			continue
		}
		id_slice = append(id_slice, i)
	}
	return id_slice, nil
}

func GetAllUsersId() ([]DbId, error) {
	db, err := sql.Open("mysql", "books:pass@tcp(127.0.0.1:3306)/web_books")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select id from web_books.users")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	id_slice := []DbId{}
	for rows.Next() {
		i := DbId{}
		err = rows.Scan(&i.Id)
		if err != nil {
			log.Println(err)
			continue
		}
		id_slice = append(id_slice, i)
	}
	return id_slice, nil
}

type CategoryBook struct {
	Category string
}

func GetAllCategoriesBooks() ([]CategoryBook, error) {
	db, err := sql.Open("mysql", "books:pass@tcp(127.0.0.1:3306)/web_books")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select category from web_books.books")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	category_slice := []CategoryBook{}
	for rows.Next() {
		c := CategoryBook{}
		err = rows.Scan(&c.Category)
		if err != nil {
			log.Println(err)
			continue
		}

		category_slice = append(category_slice, c)
	}

	// remove duplicates

	keys := make(map[CategoryBook]bool)
	list := []CategoryBook{}

	for _, entry := range category_slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list, nil
}
