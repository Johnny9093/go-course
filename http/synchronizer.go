package http

import (
	"fmt"
	"net/http"
)

func Synchronizer() {
	requestIds := make(map[string]chan string)
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("task_id")
		if id == "" {
			w.WriteHeader(400)
			return
		}

		_, ok := requestIds[id]
		if !ok {
			requestIds[id] = make(chan string)
		} else {
			w.WriteHeader(400)
			w.Write([]byte("Task already exists"))
			return
		}

		c := requestIds[id]
		response := <- c

		w.Write([]byte(fmt.Sprintf("Response: %s", response)))
	})
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("task_id")
		if id == "" {
			w.WriteHeader(400)
			return
		}

		_, ok := requestIds[id]
		if !ok {
			w.WriteHeader(400)
			w.Write([]byte("Task not found"))
			return
		}

		c:=requestIds[id]
		c <- "Message"
		w.Write([]byte("Ok"))
	})
	http.ListenAndServe(":8080", nil)
}
