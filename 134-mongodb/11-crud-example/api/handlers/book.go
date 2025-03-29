package handlers

import (
	"crud-example/internal/constant"
	"crud-example/internal/model"
	"crud-example/internal/tpl"
	apperr "crud-example/pkg/util/app_err"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BookController struct {
}

func NewBookController() *BookController {
	return &BookController{}
}

func (uc *BookController) GetBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bks, err := model.FetchAllBooks()

	if err != nil {
		apperr.HandleHttpMongoErr(w, r, err)
		return
	}

	tpl.Tpl.ExecuteTemplate(w, "books.gohtml", bks)
}

func (uc *BookController) GetBookDetails(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	isbn := p.ByName(constant.SlugISBN)

	if isbn == "" {
		apperr.HandleBadRequest(w, constant.ErrMissingISBN)
		return
	}

	bk, err := model.FetchBook(isbn)
	if err != nil {
		apperr.HandleHttpMongoErr(w, r, err)
		return
	}

	tpl.Tpl.ExecuteTemplate(w, "details.gohtml", bk)
}

func (uc *BookController) GetCreateBook(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	tpl.Tpl.ExecuteTemplate(w, "create.gohtml", nil)
}

func (uc *BookController) CreateBookProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	bk := model.Book{
		Isbn:   r.FormValue("isbn"),
		Title:  r.FormValue("title"),
		Author: r.FormValue("author"),
	}
	p := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		apperr.HandleBadRequest(w, constant.ErrMissingSomeFields)
		return
	}

	if price, err := strconv.ParseFloat(p, 64); err != nil {
		apperr.HandleNotAcceptable(w, constant.ErrInvalidPriceField)
		return
	} else {
		bk.Price = price
	}

	_, err := model.InsertBook(bk)
	if err != nil {
		apperr.HandleInternalServerError(w, apperr.ErrNoMessage)
		return
	}

	tpl.Tpl.ExecuteTemplate(w, "created.gohtml", bk)
}

func (uc *BookController) GetUpdateBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	isbn := p.ByName(constant.SlugISBN)

	if isbn == "" {
		apperr.HandleBadRequest(w, constant.ErrMissingISBN)
		return
	}

	bk, err := model.FetchBook(isbn)
	if err != nil {
		apperr.HandleHttpMongoErr(w, r, err)
		return
	}

	tpl.Tpl.ExecuteTemplate(w, "update.gohtml", bk)
}

func (uc *BookController) PutUpdatedBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	isbn := p.ByName(constant.SlugISBN)

	if isbn == "" {
		apperr.HandleBadRequest(w, constant.ErrMissingISBN)
		return
	}

	bk := model.Book{
		Isbn:   r.FormValue("isbn"),
		Title:  r.FormValue("title"),
		Author: r.FormValue("author"),
	}
	pr := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || pr == "" {
		apperr.HandleBadRequest(w, constant.ErrMissingSomeFields)
		return
	}

	if price, err := strconv.ParseFloat(pr, 64); err != nil {
		apperr.HandleNotAcceptable(w, constant.ErrInvalidPriceField)
		return
	} else {
		bk.Price = price
	}

	updatedBk, err := model.UpdateBook(isbn, bk)
	if err != nil {
		apperr.HandleHttpMongoErr(w, r, err)
		return
	}

	tpl.Tpl.ExecuteTemplate(w, "updated.gohtml", updatedBk)
}

func (uc *BookController) DeleteBookProcess(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	isbn := p.ByName(constant.SlugISBN)

	if isbn == "" {
		apperr.HandleBadRequest(w, constant.ErrMissingISBN)
		return
	}

	if err := model.DeleteBook(isbn); err != nil {
		apperr.HandleHttpMongoErr(w, r, err)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
