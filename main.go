package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	twitch "github.com/proletariatgames/go-twitch/twitch"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	http.HandleFunc("/", idx)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("pub"))))
	log.Printf("Listening on port %s\n\n", port)
	http.ListenAndServe(":"+port, nil)
}

func idx(w http.ResponseWriter, req *http.Request) {
	os.Setenv("GO_TWITCH_CLIENTID", "5hk0e3wz856a198lypq5bai57nf13u")

	client := twitch.NewClient(&http.Client{})

	opt := &twitch.ListOptions{
		Game: "programming",
	}

	sd, e := client.Streams.List(opt)
	check(e)
	//fmt.Println(sd.Streams)

	err := tpl.ExecuteTemplate(w, "index.html", sd.Total)
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
