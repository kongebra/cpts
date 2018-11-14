package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title string
}

type Post struct {
	Title string
	Summary string
	Body string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8088", r))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	t, err := template.ParseFiles("web/tmpl/layout.html", "web/tmpl/header.html", "web/tmpl/footer.html")

	if err != nil {
		fmt.Println(err)
	}

	data := Post{
		Title: "Lorem ipsum",
		Summary: "Sed sed finibus nisl. Mauris nec ornare ipsum. Donec ac semper ligula. Nullam et sem ut eros vehicula convallis. Morbi placerat porttitor sem. Sed ullamcorper felis eu lorem egestas tempor. Aenean at vehicula libero.",
		Body: "Maecenas sagittis lectus ut arcu malesuada interdum. Phasellus dolor dui, viverra in metus at, scelerisque semper nisi. Sed ultricies mi nisl, ac congue tellus finibus ut. Curabitur tempus, tortor ut fringilla dapibus, ipsum dolor luctus ante, id sodales diam enim sit amet arcu. Ut quam dolor, laoreet vel sem sit amet, lobortis tempor augue. Vivamus congue efficitur congue. Etiam tincidunt, quam nec eleifend dapibus, justo odio finibus purus, ac venenatis libero arcu nec mauris. Proin consectetur quis est feugiat facilisis.<br><br>Morbi suscipit sem sed tellus efficitur, eu fringilla ligula semper. Phasellus neque turpis, condimentum in tempor vitae, dignissim a risus. Mauris accumsan risus et tortor tristique, nec pretium orci interdum. Phasellus sagittis facilisis libero a consectetur. Fusce semper eleifend ipsum nec cursus. Pellentesque finibus sed tellus id rhoncus. Etiam vehicula quam ut quam pellentesque pulvinar.",
	}

	t.ExecuteTemplate(w, "layout", data)
}