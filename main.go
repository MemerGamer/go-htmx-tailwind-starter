package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func main() {
	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Serve index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("static/index.html"))
		tmpl.Execute(w, nil)
	})

	// Get counter
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		countStr := strconv.Itoa(counter) // Convert int to string safely
		mutex.Unlock()
		log.Printf("/get: Counter is %s\n", countStr)
		fmt.Fprintf(w, `<span class="text-blue-500 font-bold">%s</span>`, countStr)
	})

	// Increment counter
	http.HandleFunc("/increment", func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		counter++
		countStr := strconv.Itoa(counter) // Convert int to string safely
		mutex.Unlock()
		log.Printf("/increment: Counter incremented to %s\n", countStr)
		fmt.Fprintf(w, `<span class="text-green-500 font-bold">%s</span>`, countStr)
	})

	log.Println("Server started at http://localhost:4356")
	log.Fatal(http.ListenAndServe(":4356", nil))
}
