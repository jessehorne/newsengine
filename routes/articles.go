package routes

import (
	"encoding/json"
	"fmt"
	"github.com/jessehorne/newsengine/db"
	"github.com/jessehorne/newsengine/types"
	"github.com/jessehorne/newsengine/util"
	"net/http"
)

type article struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
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

	sqlStr := fmt.Sprintf(`
		INSERT INTO articles (title, body, created_at, created_by, updated_at, updated_by)
		VALUES ('%v', '%v', CURRENT_TIMESTAMP, %v, CURRENT_TIMESTAMP, %v);
	
	`, a.Title, a.Body, 0, 0)
	fmt.Println(sqlStr)
	if _, err := db.DB.Exec(sqlStr); err != nil {
		fmt.Println(err)
		util.APIResponseError(w, http.StatusBadRequest, "couldn't store article")
		return
	}

	util.APIResponse(w, 200, types.APIResponse{"msg": "OK"})
}
