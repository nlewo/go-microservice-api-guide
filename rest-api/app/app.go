package app

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	app.Router.
		Methods("GET").
		Path("/healthz").
		HandlerFunc(app.healthzFunction)

	app.Router.
		Methods("GET").
		Path("/endpoint/{id}").
		HandlerFunc(app.getFunction)

	app.Router.
		Methods("POST").
		Path("/endpoint").
		HandlerFunc(app.postFunction)
}

func (app *App) healthzFunction(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (app *App) getFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbdata := &DbData{}
	err := app.Database.QueryRow("SELECT id, date, name FROM `test` WHERE id = ?", id).Scan(&dbdata.ID, &dbdata.Date, &dbdata.Name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dbdata); err != nil {
		panic(err)
	}
}

func (app *App) postFunction(w http.ResponseWriter, r *http.Request) {
	_, err := app.Database.Exec("INSERT INTO `test` (name) VALUES ('myname')")
	if err != nil {
		log.Println("Failed to insert in DB", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
