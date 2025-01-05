package inventory

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/ref-it/gisela/server/common"
)

func PlacesAddHandler(w http.ResponseWriter, r *http.Request) {
	var place Place
	err := json.NewDecoder(r.Body).Decode(&place)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close(context.Background())

	var placeID int
	err = connection.QueryRow(context.Background(), `INSERT INTO places(
			place_title,
			place_description
		) VALUES(
			$1, $2
		) RETURNING place_id`,
		place.Title,
		place.Description).Scan(&placeID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	common.WriteJsonData(w, []byte(strconv.Itoa(placeID)))
}

func PlacesIndexHandler(w http.ResponseWriter, r *http.Request) {
	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer connection.Close(context.Background())

	var places []Place
	rows, err := connection.Query(context.Background(), "SELECT * FROM places")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = pgxscan.ScanAll(&places, rows)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sort.Slice(places, func(i, j int) bool {
		return places[i].Title < places[j].Title
	})

	data, err := json.Marshal(places)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	common.WriteJsonData(w, data)
}

func PlacesDataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	placeID := vars["id"]

	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer connection.Close(context.Background())

	var place Place
	err = connection.QueryRow(context.Background(), "SELECT place_id, place_title, place_description FROM places WHERE place_id = $1", placeID).Scan(&place.PlaceID, &place.Title, &place.Description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//return
		panic(err)
	}

	data, err := json.Marshal(place)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	common.WriteJsonData(w, data)
}

func PlacesEditHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	placeID := vars["id"]

	var place Place
	err := json.NewDecoder(r.Body).Decode(&place)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close(context.Background())

	_, err = connection.Exec(context.Background(), "UPDATE places SET place_title = $2, place_description = $3 WHERE place_id = $1", placeID, place.Title, place.Description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func PlacesUsedHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	placeID := vars["id"]

	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close(context.Background())

	var itemIDs []int
	rows, err := connection.Query(context.Background(), "SELECT item_id FROM items WHERE item_place = $1", placeID)
	if err != nil {
		panic(err)
	}

	err = pgxscan.ScanAll(&itemIDs, rows)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(itemIDs)
	if err != nil {
		panic(err)
	}

	common.WriteJsonData(w, data)
}
