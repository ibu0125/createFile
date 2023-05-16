package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", fileCreate)
	http.ListenAndServe(":8080", nil)

}

func fileCreate(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 30; i++ {
		fileName := fmt.Sprintf("baka%d.txt", i)
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()
	}
}
