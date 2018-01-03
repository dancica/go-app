package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "       *       \n       *     *    \n        *   *    \n* Little Snowflakes! *\n     *      *\n       *       *\n        *\n       *")
}
