package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, World!")
		if err != nil {
			fmt.Println("Error writing response:", err)
			return
		}
		fmt.Printf("Number of bytes written: %d\n", n)
	})

	fmt.Println("Server starting on :8080")
	_ = http.ListenAndServe(":8080", nil)
}
