package inventory

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ref-it/gisela/server/common"
)

func transformInventoryItems(itemsDB []InventoryItemDB) []InventoryItem {
	var items []InventoryItem
	for i := 0; i < len(itemsDB); i++ {
		var item InventoryItem
		item.ItemID = itemsDB[i].ItemID
		item.Title = itemsDB[i].Title.String
		item.Description = itemsDB[i].Description.String
		item.Condition = itemsDB[i].Condition.String
		item.Category = int(itemsDB[i].Category.Int64)
		item.Place = int(itemsDB[i].Place.Int64)
		item.Group = int(itemsDB[i].Group.Int64)
		item.Number = int(itemsDB[i].Number.Int64)
		item.DateOfReceipt, _, _ = strings.Cut(itemsDB[i].DateOfReceipt.String, "T")
		item.DateOfRetirement, _, _ = strings.Cut(itemsDB[i].DateOfRetirement.String, "T")
		item.DateOfAccounting, _, _ = strings.Cut(itemsDB[i].DateOfAccounting.String, "T")
		item.AcquisitionCost = itemsDB[i].AcquisitionCost.Float64
		item.Project = itemsDB[i].Project
		item.Receipt = itemsDB[i].Receipt
		item.Supplier = itemsDB[i].Supplier.String
		item.Borrowable = itemsDB[i].Borrowable
		item.ParentItem = int(itemsDB[i].ParentItem.Int64)
		items = append(items, item)
	}
	return items
}

func getDataForExport(exportRequest ExportRequest) []InventoryItem {
	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close(context.Background())

	queryStringBase := `SELECT 
		item_id,
		item_title,
		item_description,
		item_condition,
		item_category,
		item_place,
		item_group,
		item_number,
		date_of_receipt,
		date_of_retirement,
		date_of_accounting,
		acquisition_cost,
		project,
		receipt,
		supplier,
		borrowable,
		parent_item
		FROM items`

	var itemsDB []InventoryItemDB
	rows, err := connection.Query(context.Background(), queryStringBase)
	if err != nil {
		panic(err)
	}
	err = pgxscan.ScanAll(&itemsDB, rows)
	if err != nil {
		panic(err)
	}

	var itemsDB1 []InventoryItemDB
	if exportRequest.FilterDateOfReceipt {
		for i := 0; i < len(itemsDB); i++ {
			switch exportRequest.RelationDateOfReceipt {
			case "=":
				if itemsDB[i].DateOfReceipt.String == exportRequest.DateOfReceipt {
					itemsDB1 = append(itemsDB1, itemsDB[i])
				}
			case "<":
				if itemsDB[i].DateOfReceipt.String < exportRequest.DateOfReceipt {
					itemsDB1 = append(itemsDB1, itemsDB[i])
				}
			case ">":
				if itemsDB[i].DateOfReceipt.String > exportRequest.DateOfReceipt {
					itemsDB1 = append(itemsDB1, itemsDB[i])
				}
			case "≤":
				if itemsDB[i].DateOfReceipt.String <= exportRequest.DateOfReceipt {
					itemsDB1 = append(itemsDB1, itemsDB[i])
				}
			case "≥":
				if itemsDB[i].DateOfReceipt.String >= exportRequest.DateOfReceipt {
					itemsDB1 = append(itemsDB1, itemsDB[i])
				}
			case "!=":
				if itemsDB[i].DateOfReceipt.String != exportRequest.DateOfReceipt {
					itemsDB1 = append(itemsDB1, itemsDB[i])
				}
			}
		}
	} else {
		itemsDB1 = itemsDB
	}

	var itemsDB2 []InventoryItemDB
	if exportRequest.FilterDateOfRetirement {
		for i := 0; i < len(itemsDB1); i++ {
			switch exportRequest.RelationDateOfRetirement {
			case "=":
				if itemsDB1[i].DateOfRetirement.String == exportRequest.DateOfRetirement {
					itemsDB2 = append(itemsDB2, itemsDB1[i])
				}
			case "<":
				if itemsDB1[i].DateOfRetirement.String < exportRequest.DateOfRetirement {
					itemsDB2 = append(itemsDB2, itemsDB1[i])
				}
			case ">":
				if itemsDB1[i].DateOfRetirement.String > exportRequest.DateOfRetirement {
					itemsDB2 = append(itemsDB2, itemsDB1[i])
				}
			case "≤":
				if itemsDB1[i].DateOfRetirement.String <= exportRequest.DateOfRetirement {
					itemsDB2 = append(itemsDB2, itemsDB1[i])
				}
			case "≥":
				if itemsDB1[i].DateOfRetirement.String >= exportRequest.DateOfRetirement {
					itemsDB2 = append(itemsDB2, itemsDB1[i])
				}
			case "!=":
				if itemsDB1[i].DateOfRetirement.String != exportRequest.DateOfRetirement {
					itemsDB2 = append(itemsDB2, itemsDB1[i])
				}
			}
		}
	} else {
		itemsDB2 = itemsDB1
	}

	var itemsDB3 []InventoryItemDB
	if exportRequest.FilterAcquisitionCost {
		for i := 0; i < len(itemsDB2); i++ {
			switch exportRequest.RelationAcquisitionCost {
			case "=":
				if itemsDB2[i].AcquisitionCost.Float64 == exportRequest.AcquisitionCost {
					itemsDB3 = append(itemsDB3, itemsDB2[i])
				}
			case "<":
				if itemsDB2[i].AcquisitionCost.Float64 < exportRequest.AcquisitionCost {
					itemsDB3 = append(itemsDB3, itemsDB2[i])
				}
			case ">":
				if itemsDB2[i].AcquisitionCost.Float64 > exportRequest.AcquisitionCost {
					itemsDB3 = append(itemsDB3, itemsDB2[i])
				}
			case "≤":
				if itemsDB2[i].AcquisitionCost.Float64 <= exportRequest.AcquisitionCost {
					itemsDB3 = append(itemsDB3, itemsDB2[i])
				}
			case "≥":
				if itemsDB2[i].AcquisitionCost.Float64 >= exportRequest.AcquisitionCost {
					itemsDB3 = append(itemsDB3, itemsDB2[i])
				}
			case "!=":
				if itemsDB2[i].AcquisitionCost.Float64 != exportRequest.AcquisitionCost {
					itemsDB3 = append(itemsDB3, itemsDB2[i])
				}
			}
		}
	} else {
		itemsDB3 = itemsDB2
	}

	var itemsDB4 []InventoryItemDB
	if exportRequest.FilterCategories {
		for i := 0; i < len(itemsDB3); i++ {
			for j := 0; j < len(exportRequest.Categories); j++ {
				if itemsDB3[i].Category.Int64 == int64(exportRequest.Categories[j]) {
					itemsDB4 = append(itemsDB4, itemsDB3[i])
				}
			}
		}
	} else {
		itemsDB4 = itemsDB3
	}

	var itemsDB5 []InventoryItemDB
	if exportRequest.FilterGroups {
		for i := 0; i < len(itemsDB4); i++ {
			for j := 0; j < len(exportRequest.Groups); j++ {
				if itemsDB4[i].Group.Int64 == int64(exportRequest.Groups[j]) {
					itemsDB5 = append(itemsDB5, itemsDB4[i])
				}
			}
		}
	} else {
		itemsDB5 = itemsDB4
	}

	var itemsDB6 []InventoryItemDB
	if exportRequest.FilterPlaces {
		for i := 0; i < len(itemsDB5); i++ {
			for j := 0; j < len(exportRequest.Places); j++ {
				if itemsDB5[i].Place.Int64 == int64(exportRequest.Places[j]) {
					itemsDB6 = append(itemsDB6, itemsDB5[i])
				}
			}
		}
	} else {
		itemsDB6 = itemsDB5
	}

	var itemsDB7 []InventoryItemDB
	if exportRequest.FilterRetired {
		for i := 0; i < len(itemsDB6); i++ {
			if itemsDB6[i].DateOfRetirement.Valid == exportRequest.Retired {
				itemsDB7 = append(itemsDB7, itemsDB6[i])
			}
		}
	} else {
		itemsDB7 = itemsDB6
	}

	var itemsDB8 []InventoryItemDB
	if exportRequest.FilterBorrowable {
		for i := 0; i < len(itemsDB7); i++ {
			if itemsDB6[i].Borrowable == exportRequest.Borrowable {
				itemsDB8 = append(itemsDB8, itemsDB7[i])
			}
		}
	} else {
		itemsDB8 = itemsDB7
	}

	items := transformInventoryItems(itemsDB8)

	return items
}

func getGroups() map[int]string {
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

	groupsMap := make(map[int]string)
	for i := 0; i < len(groups); i++ {
		groupsMap[groups[i].GroupID] = groups[i].Title
	}

	return groupsMap
}

func getCategories() map[int]string {
	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
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

	categoriesMap := make(map[int]string)
	for i := 0; i < len(categories); i++ {
		categoriesMap[categories[i].CategoryID] = categories[i].Title
	}

	return categoriesMap
}

func getPlaces() map[int]string {
	connStr := common.GetDatabaseConnectionString()
	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close(context.Background())

	var places []Place
	rows, err := connection.Query(context.Background(), "SELECT * FROM places")
	if err != nil {
		panic(err)
	}

	err = pgxscan.ScanAll(&places, rows)
	if err != nil {
		panic(err)
	}

	placesMap := make(map[int]string)
	for i := 0; i < len(places); i++ {
		placesMap[places[i].PlaceID] = places[i].Title
	}

	return placesMap
}

func ExportCsvHandler(w http.ResponseWriter, r *http.Request) {
	var exportRequest ExportRequest
	err := json.NewDecoder(r.Body).Decode(&exportRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	items := getDataForExport(exportRequest)

	groupsMap := getGroups()
	categoriesMap := getCategories()
	placesMap := getPlaces()

	csv := "ID,Titel,Beschreibung,Zustand,Gremium,Kategorie,Lagerort,Zugang,Abgang,Anzahl,Teil von,Lieferant,Anschaffungskosten,Projekt,Abrechnung,Buchung,Ausleihbar\n"

	for i := 0; i < len(items); i++ {
		csv += strconv.Itoa(items[i].ItemID) + ", "
		csv += "\"" + items[i].Title + "\", "
		csv += "\"" + items[i].Description + "\", "
		csv += "\"" + items[i].Condition + "\", "
		csv += groupsMap[items[i].Group] + ", "
		csv += categoriesMap[items[i].Category] + ", "
		csv += placesMap[items[i].Place] + ", "
		csv += items[i].DateOfReceipt + ", "
		csv += items[i].DateOfRetirement + ", "
		csv += strconv.Itoa(items[i].Number) + ", "
		csv += strconv.Itoa(items[i].ParentItem) + ", "
		csv += "\"" + items[i].Supplier + "\", "
		csv += fmt.Sprintf("%.2f", items[i].AcquisitionCost) + ", "
		csv += "\"" + strings.Join(strings.Fields(fmt.Sprint(items[i].Project)), ", ") + "\", "
		csv += "\"" + strings.Join(strings.Fields(fmt.Sprint(items[i].Receipt)), ", ") + "\", "
		csv += items[i].DateOfAccounting + ", "
		if items[i].Borrowable {
			csv += "ja"
		} else {
			csv += "nein"
		}
		csv += "\n"
	}

	w.Write([]byte(csv))
}

func ExportPdfHandler(w http.ResponseWriter, r *http.Request) {
	var exportRequest ExportRequest
	err := json.NewDecoder(r.Body).Decode(&exportRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	items := getDataForExport(exportRequest)

	categoriesMap := getCategories()
	groupsMap := getGroups()
	placesMap := getPlaces()

	data := ExportData{
		ExportRequest:  exportRequest,
		InventoryItems: items,
		Maps: ExportMaps{
			Categories: categoriesMap,
			Groups:     groupsMap,
			Places:     placesMap}}
	pdfFileName, dirpath, err := createPdfFile(data)
	if err != nil {
		panic(err)
	}

	file, err := os.ReadFile(pdfFileName)
	if err != nil {
		panic(err)
	}
	w.Write(file)

	err = os.RemoveAll(dirpath)
	if err != nil {
		panic(err)
	}
}

func createPdfFile(data ExportData) (string, string, error) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	tplPath := path + "/templates/export-latex.tpl"
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		panic(err)
	}

	exportDirName := uuid.New().String()
	err = os.MkdirAll("tmp/"+exportDirName, 0755)
	if err != nil {
		panic(err)
	}
	filename := "gisela-export"
	dirpath := path + "/tmp/" + exportDirName

	texFileName := filepath.Join(dirpath, filename+".tex")
	pdfFileName := filepath.Join(dirpath, filename+".pdf")

	file, err := os.Create(texFileName)
	if err != nil {
		err2 := os.RemoveAll(dirpath)
		if err2 != nil {
			panic(err2)
		}
		panic(err)
	}

	err = tpl.Execute(file, data)
	if err != nil {
		log.Println(err)
		e := os.RemoveAll(dirpath)
		if e != nil {
			return "", "", err
		}
		return "", "", err
	}

	err = runLatex(dirpath, texFileName)
	if err != nil {
		e := os.RemoveAll(dirpath)
		if e != nil {
			return "", "", err
		}
		return "", "", err
	}

	file.Close()

	return pdfFileName, dirpath, nil
}

func runLatex(dirpath string, filename string) error {
	latexPath := common.GetConfiguration().Latex.Path
	args := []string{"-output-directory=" + dirpath, "-synctex=1", "-interaction=nonstopmode", "-halt-on-error", filename}
	numberOfRuns := 2

	for i := 0; i < numberOfRuns; i++ {
		cmd := exec.Command(latexPath, args...)
		err := cmd.Run()
		if err != nil {
			log.Printf("Command finished with error: %v", err)
			panic(err)
			//return err
		}
	}

	err := os.Remove(filename)
	if err != nil {
		panic(err)
	}
	return err
}
