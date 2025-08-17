package db

import (
	"database/sql"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// ManPage represents the manPage table structure


type ManPage struct {
    DropDownMenu []string `json:"dropDownMenu"`
}
func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")
	if err != nil {
		return err
	}

    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)

    return CreateTables()
}

// CreateTables ensures required tables exist
func CreateTables() error {
    createManPageTable := `
    CREATE TABLE IF NOT EXISTS manPage (
		menu TEXT
    )`

    _, err := DB.Exec(createManPageTable)
    if err != nil {
        return err
    }

    return nil
}



func GetAllManPages() ([]ManPage, error) {
    rows, err := DB.Query("SELECT menu FROM manPage")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var pages []ManPage
    for rows.Next() {
        var rawJSON string
        var mp ManPage

        // scan raw JSON from DB
        if err := rows.Scan(&rawJSON); err != nil {
            return nil, err
        }

        // unmarshal directly into []string
        if err := json.Unmarshal([]byte(rawJSON), &mp); err != nil {
            return nil, err
        }

        pages = append(pages, mp)
    }
    return pages, nil
}


// GetManPageByID retrieves a specific manPage by ID
// func GetManPageByID(id int) (*ManPage, error) {
//     var mp ManPage
//     err := DB.QueryRow("SELECT id, name FROM manPage WHERE id = ?", id).Scan(&mp.ID, &mp.Name)
//     if err != nil {
//         return nil, err
//     }
//     return &mp, nil
// }

// CreateManPage inserts a new manPage record
// func CreateManPage(name string) (*ManPage, error) {
// 	result, err := DB.Exec("INSERT INTO manPage (name) VALUES (?)", name)
// 	if err != nil {
// 		return nil, err
// 	}

// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &ManPage{ID: int(id), Name: name}, nil
// }

// // UpdateManPage updates an existing manPage record
// func UpdateManPage(id int, name string) error {
// 	_, err := DB.Exec("UPDATE manPage SET name = ? WHERE id = ?", name, id)
// 	return err
// }

// // DeleteManPage deletes a manPage record by ID
// func DeleteManPage(id int) error {
// 	_, err := DB.Exec("DELETE FROM manPage WHERE id = ?", id)
// 	return err
// }

