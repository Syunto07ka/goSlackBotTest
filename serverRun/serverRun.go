package serverRun

import "net/http"

func Start() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/good_morning", goodMorning)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func goodMorning(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("good morning! うぇーい"))
}
