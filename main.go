package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	dumpPort := getEnvOrDefault("DUMPSTER_PORT", "3002")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Starting Request handler")
		defer r.Body.Close()
		bits, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		log.Println("URL:", r.URL.String())
		log.Println("BODY:")
		log.Println(string(bits))

		w.WriteHeader(http.StatusOK)
	})

	log.Println("Listening on port", dumpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", dumpPort), nil)
	if err != nil {
		log.Panic(err)
	}
}

func getEnvOrDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}

	return v
}
