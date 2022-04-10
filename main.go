package main

import (
	"context"
	"fmt"
	"net/http"
	"log"
	"challenge-dev/translates"
	"challenge-dev/utils"
	"github.com/julienschmidt/httprouter"
)

func GetTranslate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, cancel := context.WithCancel(context.Background())
    defer cancel()
	var target = ps.ByName("targetLang")
	result, err := translates.TranslateText(target, "Hello, my name is Aloy. My experience in TopCoder is extremely delightful with feeling of competence that I can get at most.")
	if err != nil {
		fmt.Println(err)
	}
	utils.ResponseJSON(w, result, http.StatusOK)
}

func main() {
	router := httprouter.New()
	router.GET("/translate/:targetLang", GetTranslate)
	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}