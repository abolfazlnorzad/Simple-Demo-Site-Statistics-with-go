package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error in Upgrade Connction, err", err)
			return
		}
		hub := newHub(conn)
		hub.run()

	})

	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
