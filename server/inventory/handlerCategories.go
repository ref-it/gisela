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

func CategoriesAddHandler(w http.ResponseWriter, r *http.Request) {
	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		panic(err)
	}

	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close(context.Background())

	var categoryID int
	err = connection.QueryRow(context.Background(), `INSERT INTO categories(
			category_title,
			category_description
		) VALUES(
			$1, $2
		) RETURNING category_id`,
		category.Title,
		category.Description).Scan(&categoryID)
	if err != nil {
		panic(err)
	}

	common.WriteJsonData(w, []byte(strconv.Itoa(categoryID)))
}

func CategoriesIndexHandler(w http.ResponseWriter, r *http.Request) {
	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer connection.Close(context.Background())

	var categories []Category
	rows, err := connection.Query(context.Background(), "SELECT * FROM categories")
	if err != nil {
		panic(err)
	}

	err = pgxscan.ScanAll(&categories, rows)
	if err != nil {
		panic(err)
	}

	sort.Slice(categories, func(i, j int) bool {
		return categories[i].Title < categories[j].Title
	})

	data, err := json.Marshal(categories)
	if err != nil {
		panic(err)
	}

	common.WriteJsonData(w, data)
}

func CategoriesDataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]

	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer connection.Close(context.Background())

	var category Category
	err = connection.QueryRow(context.Background(), "SELECT category_id, category_title, category_description FROM categories WHERE category_id = $1", categoryID).Scan(&category.CategoryID, &category.Title, &category.Description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//return
		panic(err)
	}

	data, err := json.Marshal(category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	common.WriteJsonData(w, data)
}

func CategoriesEditHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]

	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
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

	_, err = connection.Exec(context.Background(), "UPDATE categories SET category_title = $2, category_description = $3 WHERE category_id = $1", categoryID, category.Title, category.Description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CategoriesUsedHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]

	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close(context.Background())

	var itemIDs []int
	rows, err := connection.Query(context.Background(), "SELECT item_id FROM items WHERE item_category = $1", categoryID)
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
