package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kaihendry/upwithlib/foo"
)

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	// HOW DO I KNOW bar CAME FROM "github.com/kaihendry/upwithlib/foo" ?
	fmt.Fprintln(w, bar.HelloFromLibrary())
}
