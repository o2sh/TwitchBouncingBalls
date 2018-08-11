package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	twitch "github.com/proletariatgames/go-twitch/twitch"
)

type Output struct {
	Total   int
	Streams []string
}

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
	var total int
	var streams []string
	os.Setenv("GO_TWITCH_CLIENTID", "5hk0e3wz856a198lypq5bai57nf13u")

	client := twitch.NewClient(&http.Client{})

	opt := &twitch.ListOptions{
		Game:  "programming",
		Limit: 20,
	}

	sd, e := client.Streams.List(opt)
	check(e)
	total = sd.Total
	for _, stream := range sd.Streams {
		streams = append(streams, stream.Channel.Status)
	}

	out := Output{total, streams}

	err := tpl.ExecuteTemplate(w, "index.html", out)
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
