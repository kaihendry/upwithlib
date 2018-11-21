//go:generate -command asset go run asset.go
//go:generate asset index.html
package main

import (
	"log"
	"net/http"
	"os"

	"html/template"

	"github.com/kaihendry/upwithlib"
)

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func html(a asset) string {
	return a.Content
}

func hello(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("name").Parse(index)
	t.Execute(w, bar.HelloFromLibrary())
}
