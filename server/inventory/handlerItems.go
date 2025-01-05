package inventory

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/ref-it/gisela/server/common"
)

func stringListToInt(list string) []int {
	var intList []int
	if list != "" {
		splitted := strings.Split(list, ",")
		for i := 0; i < len(splitted); i++ {
			singleInt, err := strconv.Atoi(splitted[i])
			if err != nil {
				panic(err)
			}
			intList = append(intList, singleInt)
		}
	} else {
		intList = []int{}
	}
	return intList
}

func processItemFormValues(r *http.Request) InventoryItemDB {
	var item InventoryItemDB

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		panic(err)
	}

	if r.Form.Get("title") != "undefined" && r.Form.Get("title") != "" {
		item.Title.String = r.Form.Get("title")
		item.Title.Valid = true
	}
	if r.Form.Get("description") != "undefined" && r.Form.Get("description") != "" {
		item.Description.String = r.Form.Get("description")
		item.Description.Valid = true
	}
	if r.Form.Get("condition") != "undefined" && r.Form.Get("condition") != "" {
		item.Condition.String = r.Form.Get("condition")
		item.Condition.Valid = true
	}
	if r.Form.Get("category") != "undefined" && r.Form.Get("category") != "" {
		item.Category.Int64, err = strconv.ParseInt(r.Form.Get("category"), 0, 64)
		if err != nil {
			panic(err)
		} else {
			item.Category.Valid = true
		}
	}
	if r.Form.Get("place") != "undefined" && r.Form.Get("place") != "" {
		item.Place.Int64, err = strconv.ParseInt(r.Form.Get("place"), 0, 64)
		if err != nil {
			panic(err)
		} else {
			item.Place.Valid = true
		}
	}
	if r.Form.Get("group") != "undefined" && r.Form.Get("group") != "" {
		item.Group.Int64, err = strconv.ParseInt(r.Form.Get("group"), 0, 64)
		if err != nil {
			panic(err)
		} else {
			item.Group.Valid = true
		}
	}
	if r.Form.Get("number") != "undefined" && r.Form.Get("number") != "" {
		item.Number.Int64, err = strconv.ParseInt(r.Form.Get("number"), 0, 64)
		if err != nil {
			panic(err)
		} else {
			item.Number.Valid = true
		}
	}
	if r.Form.Get("image") != "undefined" && r.Form.Get("image") != "" {
		item.Image.String = r.Form.Get("image")
		item.Image.Valid = true
	}
	if r.Form.Get("dateOfReceipt") != "undefined" && r.Form.Get("dateOfReceipt") != "" {
		item.DateOfReceipt.String = r.Form.Get("dateOfReceipt")
		item.DateOfReceipt.Valid = true
	}
	if r.Form.Get("dateOfRetirement") != "undefined" && r.Form.Get("dateOfRetirement") != "" {
		item.DateOfRetirement.String = r.Form.Get("dateOfRetirement")
		item.DateOfRetirement.Valid = true
	}
	if r.Form.Get("dateOfAccounting") != "undefined" && r.Form.Get("dateOfAccounting") != "" {
		item.DateOfAccounting.String = r.Form.Get("dateOfAccounting")
		item.DateOfAccounting.Valid = true
	}
	if r.Form.Get("acquisitionCost") != "undefined" && r.Form.Get("acquisitionCost") != "" {
		item.AcquisitionCost.Float64, err = strconv.ParseFloat(r.Form.Get("acquisitionCost"), 64)
		if err != nil {
			panic(err)
		} else {
			item.AcquisitionCost.Valid = true
		}
	}
	if r.Form.Get("project") != "undefined" && r.Form.Get("project") != "" {
		item.Project = stringListToInt(r.Form.Get("project"))
	} else {
		item.Project = []int{}
	}
	if r.Form.Get("receipt") != "undefined" && r.Form.Get("receipt") != "" {
		item.Receipt = stringListToInt(r.Form.Get("receipt"))
	} else {
		item.Receipt = []int{}
	}
	if r.Form.Get("supplier") != "undefined" && r.Form.Get("supplier") != "" {
		item.Supplier.String = r.Form.Get("supplier")
		item.Supplier.Valid = true
	}
	if r.Form.Get("guaranteeUntil") != "undefined" && r.Form.Get("guaranteeUntil") != "" {
		item.GuaranteeUntil.String = r.Form.Get("guaranteeUntil")
		item.GuaranteeUntil.Valid = true
	}
	if r.Form.Get("borrowable") != "undefined" && r.Form.Get("borrowable") != "" {
		item.Borrowable, err = strconv.ParseBool(r.Form.Get("borrowable"))
		if err != nil {
			panic(err)
		}
	}
	if r.Form.Get("parentItem") != "undefined" && r.Form.Get("parentItem") != "" {
		item.ParentItem.Int64, err = strconv.ParseInt(r.Form.Get("parentItem"), 0, 64)
		if err != nil {
			panic(err)
		} else {
			item.ParentItem.Valid = true
		}
	}

	// Resize base64 encoded image for the preview image

	if item.Image.Valid && strings.Contains(item.Image.String, "data:image/png") {
		item.ImagePreview.String = common.ResizeImage(item.Image.String, []int{250, 250})
		item.ImagePreview.Valid = true

		item.Image.String = common.ResizeImage(item.Image.String, []int{1000, 1000})
	}

	return item
}

func InventoryAddHandler(w http.ResponseWriter, r *http.Request) {
	item := processItemFormValues(r)

	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close(context.Background())

	var itemID int
	err = connection.QueryRow(context.Background(), `INSERT INTO items(
			item_title,
			item_description,
			item_condition,
			item_category,
			item_place,
			item_group,
			item_number,
			item_image,
			item_image_preview,
			date_of_receipt,
			date_of_retirement,
			date_of_accounting,
			acquisition_cost,
			project,
			receipt,
			supplier,
			guarantee_until,
			borrowable,
			parent_item
		) VALUES(
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19
		) RETURNING item_id`,
		item.Title,
		item.Description,
		item.Condition,
		item.Category,
		item.Place,
		item.Group,
		item.Number,
		item.Image,
		item.ImagePreview,
		item.DateOfReceipt,
		item.DateOfRetirement,
		item.DateOfAccounting,
		item.AcquisitionCost,
		item.Project,
		item.Receipt,
		item.Supplier,
		item.GuaranteeUntil,
		item.Borrowable,
		item.ParentItem).Scan(&itemID)
	if err != nil {
		panic(err)
	}

	common.WriteJsonData(w, []byte(strconv.Itoa(itemID)))
}

func InventoryIndexHandler(w http.ResponseWriter, r *http.Request) {
	queryNoOfItems, err := strconv.Atoi(r.URL.Query().Get("display"))
	if err != nil {
		panic(err)
	}
	queryStart, err := strconv.Atoi(r.URL.Query().Get("start"))
	if err != nil {
		panic(err)
	}
	noOfItems := 10
	start := 0
	if queryNoOfItems > 0 {
		noOfItems = queryNoOfItems
	}
	if queryStart >= 0 {
		start = queryStart
	}
	end := start + noOfItems

	retired := false
	if r.URL.Query().Get("retired") != "" {
		retired, err = strconv.ParseBool(r.URL.Query().Get("retired"))
		if err != nil {
			panic(err)
		}
	}

	var borrowable bool
	borrowableSet := false
	if r.URL.Query().Get("borrowable") != "" {
		borrowable, err = strconv.ParseBool(r.URL.Query().Get("borrowable"))
		if err != nil {
			panic(err)
		}
		borrowableSet = true
	}

	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		panic(err)
	}

	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close(context.Background())

	var itemsDB []InventoryItemShortDB

	var rows pgx.Rows

	sql := sq.Select("item_id, item_title, item_category, item_place, item_group, item_number, borrowable").From("items")

	fmt.Println(r.Form.Get("title"))

	// look for more secure solution
	if len(r.Form.Get("title")) > 0 {
		sql = sql.Where("UPPER(item_title) LIKE UPPER('%" + r.Form.Get("title") + "%')")
	}

	// look for more secure solution
	//categories := strings.Split(r.Form.Get("categories"), ",")
	_, err = strconv.Atoi(r.Form.Get("categories"))
	if strings.Contains(r.Form.Get("categories"), ",") || err == nil {
		sql = sql.Where("item_category IN (" + r.Form.Get("categories") + ")")
		//sql = sql.Where(sq.Eq{"item_category": categories})
	}

	// look for more secure solution
	_, err = strconv.Atoi(r.Form.Get("groups"))
	if strings.Contains(r.Form.Get("groups"), ",") || err == nil {
		sql = sql.Where("item_group IN (" + r.Form.Get("groups") + ")")
		//sql = sql.Where(sq.Eq{"item_category": categories})
	}

	if retired {
		sql = sql.Where(sq.NotEq{"date_of_retirement": nil})
	} else {
		sql = sql.Where(sq.Eq{"date_of_retirement": nil})
	}

	if borrowableSet && borrowable {
		sql = sql.Where("borrowable = TRUE")
	} else if borrowableSet && !borrowable {
		sql = sql.Where("borrowable = FALSE")
	}

	query, args, err := sql.ToSql()
	if err != nil {
		panic(err)
	}

	fmt.Println(query)

	rows, err = connection.Query(context.Background(), query, args...)
	if err != nil {
		panic(err)
	}

	err = pgxscan.ScanAll(&itemsDB, rows)
	if err != nil {
		panic(err)
	}

	var items []InventoryItemShort
	for i := 0; i < len(itemsDB); i++ {
		var item InventoryItemShort
		item.ItemID = itemsDB[i].ItemID
		item.Title = itemsDB[i].Title.String
		item.Category = int(itemsDB[i].Category.Int64)
		item.Place = int(itemsDB[i].Place.Int64)
		item.Group = int(itemsDB[i].Group.Int64)
		item.Number = int(itemsDB[i].Number.Int64)
		item.Borrowable = itemsDB[i].Borrowable
		items = append(items, item)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].ItemID < items[j].ItemID
	})

	var itemsResponse InventoryItemsIndexResponse
	if len(items) >= end {
		itemsResponse.Items = items[start:end]
	} else {
		itemsResponse.Items = items[start:]
	}
	itemsResponse.ItemsTotal = len(items)

	data, err := json.Marshal(itemsResponse)
	if err != nil {
		panic(err)
	}

	common.WriteJsonData(w, data)
}

func InventoryDataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]

	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close(context.Background())

	var itemDB InventoryItemDB
	err = connection.QueryRow(context.Background(), "SELECT * FROM items WHERE item_id = $1", itemID).Scan(
		&itemDB.ItemID,
		&itemDB.Title,
		&itemDB.Description,
		&itemDB.Condition,
		&itemDB.Category,
		&itemDB.Place,
		&itemDB.Group,
		&itemDB.Number,
		&itemDB.Image,
		&itemDB.ImagePreview,
		&itemDB.DateOfReceipt,
		&itemDB.DateOfRetirement,
		&itemDB.DateOfAccounting,
		&itemDB.AcquisitionCost,
		&itemDB.Project,
		&itemDB.Receipt,
		&itemDB.Supplier,
		&itemDB.GuaranteeUntil,
		&itemDB.Borrowable,
		&itemDB.ParentItem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//return
		panic(err)
	}

	if itemDB.DateOfReceipt.Valid {
		itemDB.DateOfReceipt.String = itemDB.DateOfReceipt.String[0:10]
	}

	if itemDB.DateOfRetirement.Valid {
		itemDB.DateOfRetirement.String = itemDB.DateOfRetirement.String[0:10]
	}

	if itemDB.DateOfAccounting.Valid {
		itemDB.DateOfAccounting.String = itemDB.DateOfAccounting.String[0:10]
	}

	if itemDB.GuaranteeUntil.Valid {
		itemDB.GuaranteeUntil.String = itemDB.GuaranteeUntil.String[0:10]
	}

	item := InventoryItem{
		ItemID:           itemDB.ItemID,
		Title:            itemDB.Title.String,
		Description:      itemDB.Description.String,
		Condition:        itemDB.Condition.String,
		Category:         int(itemDB.Category.Int64),
		Place:            int(itemDB.Place.Int64),
		Group:            int(itemDB.Group.Int64),
		Number:           int(itemDB.Number.Int64),
		Image:            itemDB.Image.String,
		ImagePreview:     itemDB.ImagePreview.String,
		DateOfReceipt:    itemDB.DateOfReceipt.String,
		DateOfRetirement: itemDB.DateOfRetirement.String,
		DateOfAccounting: itemDB.DateOfAccounting.String,
		AcquisitionCost:  itemDB.AcquisitionCost.Float64,
		Project:          itemDB.Project,
		Receipt:          itemDB.Receipt,
		Supplier:         itemDB.Supplier.String,
		GuaranteeUntil:   itemDB.GuaranteeUntil.String,
		Borrowable:       itemDB.Borrowable,
		ParentItem:       int(itemDB.ParentItem.Int64),
	}

	data, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	common.WriteJsonData(w, data)
}

func InventoryEditHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]

	item := processItemFormValues(r)

	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close(context.Background())

	_, err = connection.Exec(context.Background(), `UPDATE items SET
		item_title = $2,
		item_description = $3,
		item_condition = $4,
		item_category = $5,
		item_place = $6,
		item_group = $7,
		item_number = $8,
		item_image = $9,
		item_image_preview = $10,
		date_of_receipt = $11,
		date_of_retirement = $12,
		date_of_accounting = $13,
		acquisition_cost = $14,
		project = $15,
		receipt = $16,
		supplier = $17,
		guarantee_until = $18,
		borrowable = $19,
		parent_item = $20
		WHERE item_id = $1`,
		itemID,
		item.Title,
		item.Description,
		item.Condition,
		item.Category,
		item.Place,
		item.Group,
		item.Number,
		item.Image,
		item.ImagePreview,
		item.DateOfReceipt,
		item.DateOfRetirement,
		item.DateOfAccounting,
		item.AcquisitionCost,
		item.Project,
		item.Receipt,
		item.Supplier,
		item.GuaranteeUntil,
		item.Borrowable,
		item.ParentItem)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
