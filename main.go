package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	twitch "github.com/proletariatgames/go-twitch/twitch"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	os.Setenv("GO_TWITCH_CLIENTID", "5hk0e3wz856a198lypq5bai57nf13u")

	client := twitch.NewClient(&http.Client{})

	opt := &twitch.ListOptions{
		Game: "programming",
	}

	sd, err := client.Streams.List(opt)
	check(err)
	fmt.Println(sd)
	/*port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	http.HandleFunc("/", idx)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("pub"))))
	log.Printf("Listening on port %s\n\n", port)
	http.ListenAndServe(":"+port, nil)*/
}

/*func idx(w http.ResponseWriter, req *http.Request) {
	client := NewClient("5hk0e3wz856a198lypq5bai57nf13u")

	users, e := client.GetUsersByLogin("lirik")
	check(e)

	fmt.Println(users[0].Login)

	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}*/

func check(e error) {
	if e != nil {
		panic(e)
	}
}
