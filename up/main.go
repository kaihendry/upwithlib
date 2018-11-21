package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alecthomas/template"
	"github.com/kaihendry/upwithlib"
)

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	// HOW DO I KNOW bar CAME FROM "github.com/kaihendry/upwithlib/foo" ?
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, bar.HelloFromLibrary())
}
