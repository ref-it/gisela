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

func GroupsAddHandler(w http.ResponseWriter, r *http.Request) {
	var group Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close(context.Background())

	var groupID int
	err = connection.QueryRow(context.Background(), `INSERT INTO groups(
			group_title,
			group_description,
			group_email,
			group_roles
		) VALUES(
			$1, $2, $3, $4
		) RETURNING group_id`,
		group.Title,
		group.Description,
		group.EMail,
		group.Roles).Scan(&groupID)
	if err != nil {
		panic(err)
	}

	common.WriteJsonData(w, []byte(strconv.Itoa(groupID)))
}

func GroupsIndexHandler(w http.ResponseWriter, r *http.Request) {
	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close(context.Background())

	var groups []Group
	rows, err := connection.Query(context.Background(), "SELECT * FROM groups")
	if err != nil {
		panic(err)
	}

	err = pgxscan.ScanAll(&groups, rows)
	if err != nil {
		panic(err)
	}

	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Title < groups[j].Title
	})

	data, err := json.Marshal(groups)
	if err != nil {
		panic(err)
	}

	common.WriteJsonData(w, data)
}

func GroupsDataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["id"]

	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer connection.Close(context.Background())

	var group Group
	err = connection.QueryRow(context.Background(), "SELECT group_id, group_title, group_description, group_email, group_roles FROM groups WHERE group_id = $1", groupID).Scan(&group.GroupID, &group.Title, &group.Description, &group.EMail, &group.Roles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//return
		panic(err)
	}

	data, err := json.Marshal(group)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	common.WriteJsonData(w, data)
}

func GroupsEditHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["id"]

	var group Group
	err := json.NewDecoder(r.Body).Decode(&group)
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

	_, err = connection.Exec(context.Background(), "UPDATE groups SET group_title = $2, group_description = $3, group_email = $4, group_roles = $5 WHERE group_id = $1", groupID, group.Title, group.Description, group.EMail, group.Roles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GroupsUsedHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["id"]

	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close(context.Background())

	var itemIDs []int
	rows, err := connection.Query(context.Background(), "SELECT item_id FROM items WHERE item_group = $1", groupID)
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
