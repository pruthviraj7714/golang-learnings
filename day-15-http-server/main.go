package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name != "" {
		fmt.Fprintf(w, "Hello, %s", name)
	} else {
		fmt.Fprintln(w, "Hello, Guest")
	}
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name: "John Doe",
		Age:  30,
	}

	// Set the content type header to application/json
	// This tells the client that the response is in JSON format
	w.Header().Set("Content-Type", "application/json")

	// Encode the user struct to JSON and write it to the response writer
	json.NewEncoder(w).Encode(user)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "Hello, JSON",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

type RequestObject struct {
	URL    string
	Method string
	Header map[string][]string
	Body   any
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go server 🚀")
	})

	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/user", userHandler)

	http.HandleFunc("/request-object", func(w http.ResponseWriter, r *http.Request) {

		reqObject := RequestObject{
			URL:    r.URL.String(),
			Method: r.Method,
			Header: r.Header,
			Body:   r.Body,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reqObject)
		// Method, URL, Header and Body are the attributes of the request object
		// fmt.Fprintln(w, "Method : ", r.Method)
		// fmt.Fprintln(w, "URL : ", r.URL)
		// fmt.Fprintln(w, "Header : ", r.Header)
		// fmt.Fprintln(w, "Body : ", r.Body)
	})

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
