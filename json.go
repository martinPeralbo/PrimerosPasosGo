package main

import (
	"encoding/json"
	"net/http"
)

type Curso struct {
	Title        string
	NumeroVideos int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		curso := Cursos{{"Curso de go", 60},
			{"Curso de go", 60},
			{"Curso de go", 60},
			{"Curso de go", 60},
		}
		json.NewEncoder(w).Encode(curso)
	})

	http.ListenAndServe(":8000", nil)
}
