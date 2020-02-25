
package main

import (
	"fmt"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func main() {
	port := getEnv("PORT", "8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¿NO ES ESTO GENIAL...?")
	})

	http.ListenAndServe(":"+port, nil)
}