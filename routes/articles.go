package routes

import (
	"encoding/json"
	"github.com/jessehorne/newsengine/util"
	"net/http"
)

type article struct {
	Title string `validate:"required"`
	Body  string `validate:"required"`
}

func GetArticleByID(w http.ResponseWriter, r *http.Request) {

}

func UpdateArticleByID(w http.ResponseWriter, r *http.Request) {

}

func DeleteArticleByID(w http.ResponseWriter, r *http.Request) {

}

func GetArticles(w http.ResponseWriter, r *http.Request) {

}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var a article
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		util.APIResponseError(w, http.StatusBadRequest, "couldn't decode article json")
		return
	}

	if err := util.Validate(a); err != nil {
		util.APIResponseError(w, http.StatusBadRequest, "invalid data")
		return
	}
}
