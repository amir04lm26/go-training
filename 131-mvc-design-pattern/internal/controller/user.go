package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"mvc-design-pattern/internal/model"
	"mvc-design-pattern/internal/view"
	"mvc-design-pattern/pkg/util"

	"github.com/julienschmidt/httprouter"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// Methods have to be capitalized to be exported, eg, GetUser and not getUser
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := &model.User{
		Name:   "Amir",
		Gender: "male",
		Age:    27,
		Id:     p.ByName("id"),
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		util.BadRequestErr(w, err)
	}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := &model.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		util.BadRequestErr(w, err)
	}

	user.Id = "007"
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		util.InternalServerErr(w, err)
	}
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: only write code to delete user
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user\n")
}

func (uc *UserController) GetHome(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	view.Tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
