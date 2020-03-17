package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "[ENV: ",os.Getenv("environment"),"] Cześć, 🚢 => ", name)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server started ENV: ",os.Getenv("environment"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
