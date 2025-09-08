package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// ManPage represents the manPage table structure


type ManPage struct {
    DropDownMenu []string `json:"dropDownMenu"`
}

type RecentImages struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Summary string `json:"summary" db:"summary"`
	Link    string `json:"link" db:"link"`
}
type Episode struct {
    ID            int    `json:"id"`
    SeriesID      int    `json:"series_id"`
    EpisodeNumber int    `json:"episode_number"`
    ReleaseDate   string `json:"release_date"`
    Title         string `json:"title"`
    SeasonNumber  int    `json:"season_number"`
}

func InitDB() error {
    var err error
    DB, err = sql.Open("sqlite3", "./api.db")
    if err != nil {
        return err
    }

    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)

    if err := CreateTables(); err != nil {
        return err
    }
    if err := CreateImageTable(); err != nil {
        return err
    }


    return nil
}


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
func CreateImageTable() error {
    createImageTable := `
    CREATE TABLE IF NOT EXISTS RecentUploads (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        summary TEXT,
        link TEXT
    )`
    _, err := DB.Exec(createImageTable)
    return err
}

func GetRecentUploadsFromDB() ([]RecentImages, error) {
    rows, err := DB.Query("SELECT id, name, summary, link FROM RecentUploads  LIMIT 15 OFFSET 0")
    if err != nil {
        return nil, fmt.Errorf("failed to query database: %w", err)
    }
    defer rows.Close()

    fmt.Println(rows,"rows")
    var images []RecentImages
    for rows.Next() {
        var img RecentImages
        
        if err := rows.Scan(&img.ID, &img.Name, &img.Summary, &img.Link); err != nil {
            return nil, fmt.Errorf("failed to scan row: %w", err)
        }
        
        images = append(images, img)
    }
    
    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("error during row iteration: %w", err)
    }
    fmt.Println(images,"images")
    
    return images, nil
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

     
        if err := rows.Scan(&rawJSON); err != nil {
            return nil, err
        }

        if err := json.Unmarshal([]byte(rawJSON), &mp); err != nil {
            return nil, err
        }

        pages = append(pages, mp)
    }
    return pages, nil
}

func getEpisodes() ([]Episode, error) {
    rows, err := DB.Query(`
    SELECT e.id, e.series_id, e.episode_number, e.release_date, s.title, s.season_number
    FROM Episodes e
    JOIN Series s ON s.id = e.series_id
    WHERE e.release_date BETWEEN ? AND ?
    ORDER BY e.release_date ASC
`, "2025-09-01", "2025-09-30")

if err != nil {
    log.Fatal(err)
}
defer rows.Close()

var episodes []Episode
for rows.Next() {
    var ep Episode
    if err := rows.Scan(&ep.ID, &ep.SeriesID, &ep.EpisodeNumber, &ep.ReleaseDate, &ep.Title, &ep.SeasonNumber); err != nil {
        log.Fatal(err)
    }
    episodes = append(episodes, ep)
}
return episodes, nil
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

