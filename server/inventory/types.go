package inventory

import "database/sql"

type InventoryItem struct {
	ItemID           int     `db:"item_id" json:"id"`
	Title            string  `db:"item_title" json:"title"`
	Description      string  `db:"item_description" json:"description"`
	Condition        string  `db:"item_condition" json:"condition"`
	Category         int     `db:"item_category" json:"category"`
	Place            int     `db:"item_place" json:"place"`
	Group            int     `db:"item_group" json:"group"`
	Number           int     `db:"item_number" json:"number"`
	Image            string  `db:"item_image" json:"image"`
	ImagePreview     string  `db:"item_image_preview" json:"imagePreview"`
	DateOfReceipt    string  `db:"date_of_receipt" json:"dateOfReceipt"`
	DateOfRetirement string  `db:"date_of_retirement" json:"dateOfRetirement"`
	DateOfAccounting string  `db:"date_of_accounting" json:"dateOfAccounting"`
	AcquisitionCost  float64 `db:"acquisition_cost" json:"acquisitionCost"`
	Project          []int   `db:"project" json:"project"`
	Receipt          []int   `db:"receipt" json:"receipt"`
	Supplier         string  `db:"supplier" json:"supplier"`
	GuaranteeUntil   string  `db:"guarantee_until" json:"guaranteeUntil"`
	Borrowable       bool    `db:"borrowable" json:"borrowable"`
	ParentItem       int     `db:"parent_item" json:"parentItem"`
}

type InventoryItemDB struct {
	ItemID           int             `db:"item_id" json:"id"`
	Title            sql.NullString  `db:"item_title" json:"title"`
	Description      sql.NullString  `db:"item_description" json:"description"`
	Condition        sql.NullString  `db:"item_condition" json:"condition"`
	Category         sql.NullInt64   `db:"item_category" json:"category"`
	Place            sql.NullInt64   `db:"item_place" json:"place"`
	Group            sql.NullInt64   `db:"item_group" json:"group"`
	Number           sql.NullInt64   `db:"item_number" json:"number"`
	Image            sql.NullString  `db:"item_image" json:"image"`
	ImagePreview     sql.NullString  `db:"item_image_preview" json:"imagePreview"`
	DateOfReceipt    sql.NullString  `db:"date_of_receipt" json:"dateOfReceipt"`
	DateOfRetirement sql.NullString  `db:"date_of_retirement" json:"dateOfRetirement"`
	DateOfAccounting sql.NullString  `db:"date_of_accounting" json:"dateOfAccounting"`
	AcquisitionCost  sql.NullFloat64 `db:"acquisition_cost" json:"acquisitionCost"`
	Project          []int           `db:"project" json:"project"`
	Receipt          []int           `db:"receipt" json:"receipt"`
	Supplier         sql.NullString  `db:"supplier" json:"supplier"`
	GuaranteeUntil   sql.NullString  `db:"guarantee_until" json:"guaranteeUntil"`
	Borrowable       bool            `db:"borrowable" json:"borrowable"`
	ParentItem       sql.NullInt64   `db:"parent_item" json:"parentItem"`
}

type InventoryItemShort struct {
	ItemID     int    `db:"item_id" json:"id"`
	Title      string `db:"item_title" json:"title"`
	Category   int    `db:"item_category" json:"category"`
	Place      int    `db:"item_place" json:"place"`
	Group      int    `db:"item_group" json:"group"`
	Number     int    `db:"item_number" json:"number"`
	Borrowable bool   `db:"borrowable" json:"borrowable"`
}

type InventoryItemShortDB struct {
	ItemID     int            `db:"item_id" json:"id"`
	Title      sql.NullString `db:"item_title" json:"title"`
	Category   sql.NullInt64  `db:"item_category" json:"category"`
	Place      sql.NullInt64  `db:"item_place" json:"place"`
	Group      sql.NullInt64  `db:"item_group" json:"group"`
	Number     sql.NullInt64  `db:"item_number" json:"number"`
	Borrowable bool           `db:"borrowable" json:"borrowable"`
}

type InventoryItemsIndexResponse struct {
	Items      []InventoryItemShort `json:"items"`
	ItemsTotal int                  `json:"itemsTotal"`
}

type Category struct {
	CategoryID  int    `db:"category_id" json:"id"`
	Title       string `db:"category_title" json:"title"`
	Description string `db:"category_description" json:"description"`
}

type Place struct {
	PlaceID     int    `db:"place_id" json:"id"`
	Title       string `db:"place_title" json:"title"`
	Description string `db:"place_description" json:"description"`
}

type Group struct {
	GroupID     int      `db:"group_id" json:"id"`
	Title       string   `db:"group_title" json:"title"`
	Description string   `db:"group_description" json:"description"`
	EMail       string   `db:"group_email" json:"email"`
	Roles       []string `db:"group_roles" json:"roles"`
}

type ExportRequest struct {
	FilterDateOfReceipt    bool `json:"filterDateOfReceipt"`
	FilterDateOfRetirement bool `json:"filterDateOfRetirement"`
	FilterAcquisitionCost  bool `json:"filterAcquisitionCost"`
	FilterCategories       bool `json:"filterCategories"`
	FilterGroups           bool `json:"filterGroups"`
	FilterPlaces           bool `json:"filterPlaces"`
	FilterRetired          bool `json:"filterRetired"`
	FilterBorrowable       bool `json:"filterBorrowable"`

	RelationDateOfReceipt    string `json:"relationDateOfReceipt"`
	RelationDateOfRetirement string `json:"relationDateOfRetirement"`
	RelationAcquisitionCost  string `json:"relationAcqusitionCost"`

	DateOfReceipt    string  `json:"dateOfReceipt"`
	DateOfRetirement string  `json:"dateOfRetirement"`
	AcquisitionCost  float64 `json:"acquisitionCost"`
	Categories       []int   `json:"categories"`
	Groups           []int   `json:"groups"`
	Places           []int   `json:"places"`
	Retired          bool    `json:"retired"`
	Borrowable       bool    `json:"borrowable"`
}

type ExportMaps struct {
	Categories map[int]string
	Groups     map[int]string
	Places     map[int]string
}

type ExportData struct {
	ExportRequest  ExportRequest
	InventoryItems []InventoryItem
	Maps           ExportMaps
}
