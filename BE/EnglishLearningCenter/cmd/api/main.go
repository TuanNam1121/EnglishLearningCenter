package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Nhóm 2 LuckyWind Chiến hết mình nhé !!!")
	})

	http.ListenAndServe(":8080", nil)
}
