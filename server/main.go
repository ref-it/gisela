package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/motemen/go-loghttp/global"

	"github.com/ref-it/gisela/server/auth"
	"github.com/ref-it/gisela/server/common"
	"github.com/ref-it/gisela/server/inventory"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Join internally call path.Clean to prevent directory traversal
	path := filepath.Join(h.staticPath, r.URL.Path)

	// check whether a file exists or is a directory at the given path
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		// file does not exist or path is a directory, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	}

	if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static file
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

// main is the start function for this program
func main() {
	// Get configuration data
	config := common.GetConfiguration()

	// Do database migrations
	/*databaseURL := "postgres://"
	databaseURL += config.Database.User
	databaseURL += ":"
	databaseURL += config.Database.Password
	databaseURL += "@"
	databaseURL += config.Database.Host
	databaseURL += ":"
	databaseURL += strconv.Itoa(config.Database.Port)
	databaseURL += "/"
	databaseURL += config.Database.Name
	//migrationsPath := "file://database/migrations"

	// Create new migrations instance
	//migrations, err := migrate.New(migrationsPath, databaseURL)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Perform migrations
	//if err := migrations.Up(); err != nil {
	//	log.Fatal(err)
	//}

	*/

	// ===== SAML =====
	samlSP := auth.InitSamlAuth()

	// ===== Routing =====

	r := mux.NewRouter()

	// SAML Routing
	samlRouter := r.PathPrefix("/saml/").Subrouter()

	samlRouter.HandleFunc("/metadata", samlSP.ServeMetadata)
	samlRouter.HandleFunc("/acs", samlSP.ServeACS)
	//samlRouter.HandleFunc("/slo", samlSP.End)

	// API Routing
	apiRouter := r.PathPrefix("/api/").Subrouter()
	apiRouter.Use(samlSP.RequireAccount)

	apiRouter.HandleFunc("/inventory/items/index", inventory.InventoryIndexHandler)
	apiRouter.HandleFunc("/inventory/items/add", inventory.InventoryAddHandler)
	apiRouter.HandleFunc("/inventory/items/data/{id}", inventory.InventoryDataHandler)
	apiRouter.HandleFunc("/inventory/items/edit/{id}", inventory.InventoryEditHandler)
	apiRouter.HandleFunc("/inventory/categories/index", inventory.CategoriesIndexHandler)
	apiRouter.HandleFunc("/inventory/categories/add", inventory.CategoriesAddHandler)
	apiRouter.HandleFunc("/inventory/categories/data/{id}", inventory.CategoriesDataHandler)
	apiRouter.HandleFunc("/inventory/categories/edit/{id}", inventory.CategoriesEditHandler)
	apiRouter.HandleFunc("/inventory/categories/used/{id}", inventory.CategoriesUsedHandler)
	//apiRouter.HandleFunc("/inventory/categories/delete/{id}", inventory.CategoriesDeleteHandler)
	apiRouter.HandleFunc("/inventory/places/index", inventory.PlacesIndexHandler)
	apiRouter.HandleFunc("/inventory/places/add", inventory.PlacesAddHandler)
	apiRouter.HandleFunc("/inventory/places/data/{id}", inventory.PlacesDataHandler)
	apiRouter.HandleFunc("/inventory/places/edit/{id}", inventory.PlacesEditHandler)
	apiRouter.HandleFunc("/inventory/places/used/{id}", inventory.PlacesUsedHandler)
	//apiRouter.HandleFunc("/inventory/places/delete/{id}", inventory.PlacesDeleteHandler)
	apiRouter.HandleFunc("/inventory/groups/index", inventory.GroupsIndexHandler)
	apiRouter.HandleFunc("/inventory/groups/add", inventory.GroupsAddHandler)
	apiRouter.HandleFunc("/inventory/groups/data/{id}", inventory.GroupsDataHandler)
	apiRouter.HandleFunc("/inventory/groups/edit/{id}", inventory.GroupsEditHandler)
	apiRouter.HandleFunc("/inventory/groups/used/{id}", inventory.GroupsUsedHandler)
	//apiRouter.HandleFunc("/inventory/groups/delete/{id}", inventory.GroupsDeleteHandler)
	apiRouter.HandleFunc("/inventory/export/csv", inventory.ExportCsvHandler)
	apiRouter.HandleFunc("/inventory/export/pdf", inventory.ExportPdfHandler)

	// Single Page Application Routing
	spaRouter := r.PathPrefix("/").Subrouter()
	spaRouter.Use(samlSP.RequireAccount)

	spa := spaHandler{staticPath: "../client/dist", indexPath: "index.html"}
	spaRouter.PathPrefix("/").Handler(spa)

	// Handlers
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{`*`})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	srv := &http.Server{
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk)(r),
		Addr:         config.Web.Address + ":" + strconv.Itoa(config.Web.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start webserver
	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		panic(err)
	}
}
