package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

/*
Page represents a wiki page
*/
type Page struct {
	Title string
	Body  []byte
}

/*
User for RESTful API
*/
type User struct {
	UserID    int    `json:"user_id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

/*
UserError for RESTful API
*/
type UserError struct {
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

const dataDirectory = "data/"
const templateDirectory = "tmpl/"

var users = map[int]User{
	1: User{1, "Ephraim", "Kunz"},
	2: User{2, "Megan", "Kunz"},
	3: User{3, "Charity", "Gifford"},
}

var templates = template.Must(template.ParseFiles(
	templateDirectory+"edit.html",
	templateDirectory+"view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func (p *Page) save() error {
	filename := dataDirectory + p.Title + ".txt"
	err := ioutil.WriteFile(filename, p.Body, 0600) // Read and write by current user
	return err
}

func loadPage(title string) (*Page, error) {
	filename := dataDirectory + title + ".txt"
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	return &Page{title, data}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view.html", page)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}

	renderTemplate(w, "edit.html", page)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{title, []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	pieces := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(pieces[len(pieces)-1])

	if err != nil {
		error, _ := json.Marshal(UserError{http.StatusNotFound, "Invalid id"})
		http.Error(w, string(error), http.StatusNotFound)
		return
	}
	user, found := users[id]
	if !found {
		error, _ := json.Marshal(UserError{http.StatusNotFound, "No user with id " + strconv.Itoa(id)})
		http.Error(w, string(error), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/data/", dataHandler)
	http.HandleFunc("/", rootHandler)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.ListenAndServe(":8080", nil)
}
