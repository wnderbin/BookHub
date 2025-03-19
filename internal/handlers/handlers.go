package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"handlers/mysql"

	"github.com/gorilla/mux"
)

func http_error(w http.ResponseWriter, err error, mes string, code int) {
	log.Println(err)
	http.Error(w, mes, code)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "not_found.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	data := []string{
		"Not Found",
	}

	err = tmpl.ExecuteTemplate(w, "not_found.html", data[0])
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}
}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "main_page.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	data, err := mysql.GetAllCategoriesBooks()
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "main_page.html", map[string]interface{}{"categories": data})
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}
}

func RegisterNewUserHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "register.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	data := []string{
		"Register", // title
	}

	err = tmpl.ExecuteTemplate(w, "register.html", data[0])
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}
}

type UserForPostform struct {
	Title string
	Name  string
	Pass  string
}

func PostformRegisterHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	data := UserForPostform{
		Title: "Postform",
		Name:  username,
		Pass:  password,
	}

	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "postform_register.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	err = mysql.CreateUsersTable()
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	err = mysql.InsertUser(username, password)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "postform_register.html", data)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}
}

func AddNewBook(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "add_book.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	data := []string{
		"Add Book", // title
	}

	err = tmpl.ExecuteTemplate(w, "add_book.html", data[0])
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}
}

type BookForPostform struct {
	Title  string
	Status string

	Name        string
	Description string
	Url         string
	Author      string
	Category    string

	Username string
	Password string
}

func PostformAddBookHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name") // book
	description := r.FormValue("description")
	url := r.FormValue("url")
	author := r.FormValue("author")
	category := r.FormValue("category")

	username := r.FormValue("username") // user
	password := r.FormValue("password")

	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "postform_add_book.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	users, err := mysql.GetAllUsers()
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	var existence_status bool = false

	for _, u := range users {
		if u.Name == username && u.Password == password {
			existence_status = true
			break
		}
	}

	if existence_status {
		data := BookForPostform{
			Title:       "Postform",
			Status:      "Your book has been successfully added to the database :)",
			Name:        name,
			Description: description,
			Url:         url,
			Author:      author,
			Category:    category,
			Username:    username,
			Password:    password,
		}

		err = mysql.CreateBooksTable()
		if err != nil {
			http_error(w, err, "Internal Server Error", 500)
			return
		}

		err = mysql.InsertBook(name, description, url, author, category)
		if err != nil {
			http_error(w, err, "Internal Server Error", 500)
			return
		}

		err = tmpl.ExecuteTemplate(w, "postform_add_book.html", data)
		if err != nil {
			http_error(w, err, "Internal Server Error", 500)
			return
		}
	} else {
		data := BookForPostform{
			Title:       "Postform",
			Status:      "Your book has not been added to the database because you have not registered yet...",
			Name:        name,
			Description: description,
			Url:         url,
			Author:      author,
			Category:    category,
			Username:    username,
			Password:    password,
		}

		err = tmpl.ExecuteTemplate(w, "postform_add_book.html", data)
		if err != nil {
			http_error(w, err, "Internal Server Error", 500)
			return
		}
	}
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "books.html"),
	}

	data, err := mysql.GetAllBooks()
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "books.html", map[string]interface{}{"books": data})
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "users.html"),
	}

	data, err := mysql.GetAllUsers()
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "users.html", map[string]interface{}{"users": data})
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}
}

func UpdateBookFormHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "update_book.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	data := []string{
		"Update Book",
	}

	err = tmpl.ExecuteTemplate(w, "update_book.html", data[0])
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}
}

func PostformUpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id")) // book
	name := r.FormValue("name")
	description := r.FormValue("description")
	url := r.FormValue("url")
	author := r.FormValue("author")
	category := r.FormValue("category")

	username := r.FormValue("username") // user
	password := r.FormValue("password")

	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "postform_update_book.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	users, err := mysql.GetAllUsers()
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	var existence_status_user bool = false

	for _, u := range users {
		if u.Name == username && u.Password == password {
			existence_status_user = true
			break
		}
	}

	all_id, err := mysql.GetAllBooksId()
	if err != nil {
		log.Println(err)
		return
	}

	var existence_status_id bool = false

	for _, i := range all_id {
		if i.Id == id {
			existence_status_id = true
			break
		}
	}

	if existence_status_user && existence_status_id {
		data := BookForPostform{
			Title:  "Postform",
			Status: "Your book has been updated",

			Name:        name,
			Description: description,
			Url:         url,
			Author:      author,
			Category:    category,

			Username: username,
			Password: password,
		}

		err = mysql.UpdateBook(id, name, description, url, author, category)
		if err != nil {
			http_error(w, err, "Internal Server Error", 500)
			return
		}

		err = tmpl.ExecuteTemplate(w, "postform_update_book.html", data)
		if err != nil {
			http_error(w, err, "Internal Server Error", 500)
			return
		}
	} else {
		data := BookForPostform{
			Title:  "Postform",
			Status: "Your book has not been updated because you have not registered yet or you specified a non-existent id",

			Name:        name,
			Description: description,
			Url:         url,
			Author:      author,
			Category:    category,

			Username: username,
			Password: password,
		}

		err = tmpl.ExecuteTemplate(w, "postform_update_book.html", data)
		if err != nil {
			http_error(w, err, "Internal Server Error", 500)
			return
		}
	}
}

func UpdateUserFormHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "update_user.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	data := []string{
		"Update User",
	}

	err = tmpl.ExecuteTemplate(w, "update_user.html", data[0])
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}
}

type UpdateUserForPostform struct {
	Title  string
	Status string

	Id int

	Current_Username string
	Current_Password string

	New_Username string
	New_Password string
}

func PostformUpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	cur_username := r.FormValue("cur_username")
	cur_password := r.FormValue("cur_password")

	new_username := r.FormValue("new_username")
	new_password := r.FormValue("new_password")

	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "postform_update_user.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	users, err := mysql.GetAllUsers()
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	var existence_status_user bool = false
	for _, u := range users {
		if u.Name == cur_username && u.Password == cur_password {
			existence_status_user = true
			break
		}
	}

	all_id, err := mysql.GetAllUsersId()
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	var existence_status_id bool = false
	for _, i := range all_id {
		if i.Id == id {
			existence_status_id = true
			break
		}
	}

	if existence_status_user && existence_status_id {
		data := UpdateUserForPostform{
			Title:  "Postform",
			Status: "Your name and password have been successfully changed in the database",

			Id: id,

			Current_Username: cur_username,
			Current_Password: cur_password,

			New_Username: new_username,
			New_Password: new_password,
		}

		err = mysql.UpdateUser(id, new_username, new_password)
		if err != nil {
			http_error(w, err, "Internal Server Error", 500)
			return
		}

		err = tmpl.ExecuteTemplate(w, "postform_update_user.html", data)
		if err != nil {
			http_error(w, err, "Internal Server Error", 500)
			return
		}
	} else {
		data := UpdateUserForPostform{
			Title:  "Posform",
			Status: "Your data was not updated because you are not registered or you specified a non-existent id",

			Id: id,

			Current_Username: cur_username,
			Current_Password: cur_password,

			New_Username: new_username,
			New_Password: new_password,
		}

		err = tmpl.ExecuteTemplate(w, "postform_update_user.html", data)
		if err != nil {
			http_error(w, err, "Internal Server Error", 500)
			return
		}
	}
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "delete_book.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	data := []string{
		"Delete Book", // title
	}

	err = tmpl.ExecuteTemplate(w, "delete_book.html", data[0])
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}
}

type PostformDeleteData struct {
	Title  string
	Status string

	Id       int
	Username string
	Password string
}

func PostformDeleteBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	username := r.FormValue("username")
	password := r.FormValue("password")

	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "postform_delete_book.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	users, err := mysql.GetAllUsers()
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	all_id, err := mysql.GetAllUsersId()
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	var existence_status_user bool = false
	var existence_status_id bool = false

	for _, u := range users {
		if u.Name == username && u.Password == password {
			existence_status_user = true
			break
		}
	}

	for _, i := range all_id {
		if i.Id == id {
			existence_status_id = true
			break
		}
	}

	if existence_status_user && existence_status_id {
		data := PostformDeleteData{
			Title:  "Postform",
			Status: "The book was successfully deleted",

			Id:       id,
			Username: username,
			Password: password,
		}

		err = mysql.DeleteBook(id)
		if err != nil {
			http_error(w, err, "Internal Server Error", 500)
			return
		}

		err = tmpl.ExecuteTemplate(w, "postform_delete_book.html", data)
		if err != nil {
			http_error(w, err, "Internal Server Error", 500)
			return
		}
	} else {
		data := PostformDeleteData{
			Title:  "Postform",
			Status: "The book was not deleted because you are not a registered user, or there is no book with the id you specified.",

			Id:       id,
			Username: username,
			Password: password,
		}

		err = tmpl.ExecuteTemplate(w, "postform_delete_book.html", data)
		if err != nil {
			http_error(w, err, "Internal Server Error", 500)
			return
		}
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "delete_user.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	data := []string{
		"Delete User",
	}

	err = tmpl.ExecuteTemplate(w, "delete_user.html", data[0])
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}
}

func PostformDeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	username := r.FormValue("username")
	password := r.FormValue("password")

	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "postform_delete_user.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	users, err := mysql.GetAllUsers()
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	all_id, err := mysql.GetAllUsersId()
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	var existence_status_user bool = false
	var existence_status_id = false

	for _, u := range users {
		if u.Name == username && u.Password == password {
			existence_status_user = true
			break
		}
	}

	for _, i := range all_id {
		if i.Id == id {
			existence_status_id = true
			break
		}
	}

	if existence_status_user && existence_status_id {
		data := PostformDeleteData{
			Title:  "Postform",
			Status: "The user was successfully deleted",

			Id:       id,
			Username: username,
			Password: password,
		}

		err = mysql.DeleteUser(id)
		if err != nil {
			http_error(w, err, "Internal Server Error", 500)
			return
		}

		err = tmpl.ExecuteTemplate(w, "postform_delete_user.html", data)
		if err != nil {
			http_error(w, err, "Internal Server Error", 500)
			return
		}
	} else {
		data := PostformDeleteData{
			Title:  "Postform",
			Status: "The user was not deleted because you are not registered or you specified a non-existent id",

			Id:       id,
			Username: username,
			Password: password,
		}

		err = tmpl.ExecuteTemplate(w, "postform_delete_user.html", data)
		if err != nil {
			http_error(w, err, "Internal Server Error", 500)
			return
		}
	}
}

func CategoryBooksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]

	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "category_books.html"),
	}

	data, err := mysql.GetCategoryBooks(category)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "category_books.html", map[string]interface{}{"books": data})
	if err != nil {
		http_error(w, err, "Internal Server Error", 500)
		return
	}
}
