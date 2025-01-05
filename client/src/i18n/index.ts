import { createI18n } from 'vue-i18n'

const messages = {
    de: {
        nav: {
            categories: 'Kategorien',
            exportData: 'Daten exportieren',
            groups: 'Gremien',
            headingInventory: 'Inventarverwaltung',
            headingLending: 'Ausleihe',
            inventoryRegister: 'Inventarverzeichnis',
            lendingRequests: 'Ausleihe-Anfragen',
            logout: 'Abmelden',
            myLendings: 'Meine Ausleihen',
            overview: 'Übersicht',
            places: 'Orte',
            submitLendingRequest: 'Ausleihe-Antrag stellen'
        },
        dashboard: {
            inventoryRegisterHeading: 'Inventarverzeichnis',
            inventoryRegisterDescription: 'Sieh dir eine Auflistung des gesamten registrierten Inventars an.'
        },
        common: {
            aboutGISELA: 'Über GISELA',
            add: 'Hinzufügen',
            borrowable: 'Ausleihbar',
            cancel: 'Abbrechen',
            category: 'Kategorie',
            close: 'Schließen',
            delete: 'Löschen',
            description: 'Beschreibung',
            edit: 'Bearbeiten',
            emailAddress: 'E-Mail-Adresse',
            first: 'Erste',
            group: 'Gremium',
            inventoryNumber: 'Nummer',
            itemsPerPage: 'Elemente pro Seite',
            last: 'Letzte',
            name: 'Name',
            next: 'Nächste',
            place: 'Lagerort',
            previous: 'Vorherige',
            remove: 'Entfernen',
            requestRetirement: 'Ausmusterung beantragen',
            roles: 'Rollen',
            title: 'Titel',
            view: 'Ansehen'
        },
        inventory: {
            acqusitionCost: 'Anschaffungskosten',
            addItem: 'Gegenstand hinzufügen',
            addPicture: 'Bild hinzufügen',
            borrowable: 'Ausleihbar',
            notBorrowable: 'Nicht ausleihbar',
            borrowableDescription: 'Dieser Gegenstand kann gemäß Ausleiheordnung ausgeliehen werden.',
            condition: 'Zustand',
            dateOfAccounting: 'Buchungsdatum',
            dateOfReceipt: 'Zugangsdatum',
            dateOfRetirement: 'Abgangsdatum',
            guaranteeUntil: 'Garantie bis',
            inventoryIntro: 'Eine Auflistung des gesamten registrierten Inventars.',
            inventoryTitle: 'Inventar',
            number: 'Anzahl',
            partOf: 'Teil von',
            picture: 'Bild',
            project: 'Projekt',
            retiredItems: 'Ausgemusterte Gegenstände',
            severalSeperatedByComma: 'mehrere durch Komma getrennt',
            supplier: 'Lieferant'
        },
        categories: {
            addCategory: 'Kategorie hinzufügen',
            categoriesIntro: 'Eine Auflistung aller angelegten Kategorien, in die Inventar eingruppiert werden kann.',
            categoriesTitle: 'Kategorien',
        },
        export: {
            exportAsCSV: 'Als CSV exportieren',
            exportAsPDF: 'Als PDF exportieren'
        },
        lending: {
            cartIntro: 'Für deine ausgewählten Gegenstände kannst du hier die Ausleihe beantragen.',
            cartTitle: 'Ausleihe-Warenkorb'
        }
    },
    en: {
        nav: {
            categories: 'Categories',
            exportData: 'Export Data',
            groups: 'Groups',
            headingInventory: 'Inventory Management',
            headingLending: 'Lending',
            inventoryRegister: 'Inventory Register',
            lendingRequests: 'Lending Requests',
            logout: 'Log out',
            myLendings: 'My Lendings',
            overview: 'Overview',
            places: 'Places',
            submitLendingRequest: 'Submit Lending Request'
        },
        dashboard: {
            inventoryRegisterHeading: 'Inventory Register',
            inventoryRegisterDescription: 'View a list of the entire registered inventory.'
        },
        common: {
            aboutGISELA: 'About GISELA',
            add: 'Add',
            borrowable: 'Borrowable',
            cancel: 'Cancel',
            category: 'Category',
            close: 'Close',
            delete: 'Delete',
            description: 'Description',
            edit: 'Edit',
            emailAddress: 'Email Address',
            first: 'First',
            group: 'Group',
            inventoryNumber: 'Inventory Number',
            itemsPerPage: 'items per page',
            last: 'Last',
            name: 'Name',
            next: 'Next',
            place: 'Place',
            previous: 'Previous',
            remove: 'Remove',
            requestRetirement: 'Request Retirement',
            roles: 'Roles',
            title: 'Title',
            view: 'View'
        },
        inventory: {
            acqusitionCost: 'Acquisition Cost',
            addItem: 'Add Item',
            addPicture: 'Add Picture',
            borrowable: 'Borrowable',
            notBorrowable: 'Not borrowable',
            borrowableDescription: 'This item can be borrowed in accordance with the lending regulations.',
            condition: 'Condition',
            dateOfAccounting: 'Date of Accounting',
            dateOfReceipt: 'Date of Receipt',
            dateOfRetirement: 'Date of Retirement',
            guaranteeUntil: 'Guarantee Until',
            inventoryIntro: 'A list of the entire registered inventory.',
            inventoryTitle: 'Inventory',
            number: 'Number',
            partOf: 'Part of',
            picture: 'Picture',
            project: 'Project',
            retiredItems: 'Retired Items',
            severalSeperatedByComma: 'several separated by a comma',
            supplier: 'Supplier'
        },
        categories: {
            addCategory: 'Add Category',
            categoriesIntro: 'A list of all created categories into which inventory can be grouped.',
            categoriesTitle: 'Categories',
        },
        export: {
            exportAsCSV: 'Export as CSV',
            exportAsPDF: 'Export as PDF'
        },
        lending: {
            cartIntro: 'You can apply to borrow your selected items here.',
            cartTitle: 'Lending Cart'
        }
    }
}

export default messages
