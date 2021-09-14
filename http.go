package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

type AppHandler struct {
	http.Handler
}

var rd *render.Render = render.New(render.Options{
	Extensions: []string{".html", ".tmpl"},
})

func (a *AppHandler) osSepHandler(w http.ResponseWriter, r *http.Request) {
	osSep := string(os.PathSeparator)
	rd.JSON(w, http.StatusOK, map[string]string{"data": osSep})
}

func (a *AppHandler) initHandler(w http.ResponseWriter, r *http.Request) {
	data, err := LoadInit()
	if err != nil {
		log.Println(err)
		rd.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	rd.JSON(w, http.StatusOK, map[string]*InitData{"data": data})
}

func (a *AppHandler) makeInitHandler(w http.ResponseWriter, r *http.Request) {
	data := InitData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		rd.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	err = MakeInit(&data)
	if err != nil {
		log.Println(err)
		rd.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	rd.JSON(w, http.StatusOK, map[string]bool{"data": true})
}

func (a *AppHandler) pathHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	imgs, err := Search(path)
	if err != nil {
		log.Println(err)
		rd.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	rd.JSON(w, http.StatusOK, map[string][]Item{"data": imgs})
}

func (a *AppHandler) loadHandler(w http.ResponseWriter, r *http.Request) {
	data, err := LoadData()
	if err != nil {
		log.Println(err)
		rd.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	rd.JSON(w, http.StatusOK, map[string][]Item{"data": *data})
}

func (a *AppHandler) autoSaveHandler(w http.ResponseWriter, r *http.Request) {
	data := []Item{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		rd.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	err = Autosave(&data)
	if err != nil {
		log.Println(err)
		rd.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	rd.JSON(w, http.StatusOK, map[string]bool{"data": true})
}

func (a *AppHandler) publishHandler(w http.ResponseWriter, r *http.Request) {
	data := []Item{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		rd.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	newData, err := Publish(data)
	if err != nil {
		log.Println(err)
		rd.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	rd.JSON(w, http.StatusOK, map[string][]Item{"data": newData})
}

func (a *AppHandler) logHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/log.html", http.StatusTemporaryRedirect)
}

func (a *AppHandler) logUpdateHandler(w http.ResponseWriter, r *http.Request) {
	data, err := LogData()
	if err != nil {
		log.Println(err)
		rd.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	rd.JSON(w, http.StatusOK, map[string][]Job{"data": data})
}

func MakeHandler() *AppHandler {
	r := mux.NewRouter()
	n := negroni.Classic()
	n.UseHandler(r)
	a := &AppHandler{
		Handler: n,
	}
	r.HandleFunc("/init", a.initHandler)
	r.HandleFunc("/ossep", a.osSepHandler)
	r.HandleFunc("/makeinit", a.makeInitHandler).Methods("POST")
	r.HandleFunc("/search", a.pathHandler).Methods("POST")
	r.HandleFunc("/load", a.loadHandler)
	r.HandleFunc("/autosave", a.autoSaveHandler).Methods("POST")
	r.HandleFunc("/publish", a.publishHandler).Methods("POST")
	r.HandleFunc("/log", a.logHandler)
	r.HandleFunc("/logupdate", a.logUpdateHandler)
	return a
}
