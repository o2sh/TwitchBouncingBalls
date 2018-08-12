package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	twitch "github.com/proletariatgames/go-twitch/twitch"
)

//Docker fix
const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
)

//Output data send to the client
type Output struct {
	Total   int
	Streams []twitch.StreamNS
}

var tpl *template.Template
var client *twitch.Client

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	os.Setenv("GO_TWITCH_CLIENTID", "5hk0e3wz856a198lypq5bai57nf13u")
	client = twitch.NewClient(&http.Client{})
	http.HandleFunc("/", idx)
	http.HandleFunc("/submit", sub)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("pub"))))
	log.Printf("Listening on port %s\n\n", port)
	http.ListenAndServe(":"+port, nil)
}

func idx(w http.ResponseWriter, req *http.Request) {
	var err error

	out := getStreams("programming")

	err = tpl.ExecuteTemplate(w, "index.html", out)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func sub(w http.ResponseWriter, req *http.Request) {
	var game string
	var err error

	if req.Method == MethodPost {
		game = req.FormValue("game")
	}
	out := getStreams(game)
	err = tpl.ExecuteTemplate(w, "index.html", out)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getStreams(game string) Output {
	total := 20

	opt := &twitch.ListOptions{
		Game:  game,
		Limit: total,
	}

	sd, err := client.Streams.List(opt)
	check(err)

	out := Output{total, sd.Streams}
	return out
}
