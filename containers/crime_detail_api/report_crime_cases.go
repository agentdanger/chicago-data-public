package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"
)

type Crime_report struct {
	Id 						*string `json:"id"`
	Case_number 			*string `json:"case_number"`
	Date 					*string `json:"date"`
	Primary_type 			*string `json:"primary_type"`
	Description 			*string `json:"description"`
	Location_description 	*string `json:"location_description"`
	Arrest 					*string `json:"arrest"`
	Domestic 				*string `json:"domestic"`
	Year 					*string `json:"year"`
	Latitude 				*string `json:"latitude"`
	Longitude		 		*string `json:"longitude"`
	Zip 					*string `json:"zip"`
}

/*const (
	host = "10.100.138.27" // this is the host of the PSQL database on AWS.
	port = 5432
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = "fullstack_api"
)*/

func OpenConnection() *sql.DB {
	host := "10.100.138.27" // this is the host of the PSQL database on AWS.
	port := 5432
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname   := "fullstack_api"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func GETHandler_all(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	rows, err := db.Query("SELECT * FROM v_crime_report_detail;")
	if err != nil {
		log.Fatal(err)
	}

	var Crimes []Crime_report

	for rows.Next() {
		var crime Crime_report
		rows.Scan(&crime.Id, 
			&crime.Case_number, 
			&crime.Date,
			&crime.Primary_type,
			&crime.Description,
			&crime.Location_description,
			&crime.Arrest,
			&crime.Domestic,
			&crime.Year,
			&crime.Latitude,
			&crime.Longitude,
			&crime.Zip)
		Crimes = append(Crimes, crime)
	}

	crimesBytes, _ := json.MarshalIndent(Crimes, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(crimesBytes)

	defer rows.Close()
	defer db.Close()
}

func main() {
	
	crime_report_full := `
	CREATE OR REPLACE VIEW v_crime_report_detail AS 
		SELECT id, case_number, date, primary_type, description, location_description, 
		arrest, domestic, year, latitude, longitude, zip 
		FROM t_crime_data_wzip 
		WHERE year > 2020;
		`
	db := OpenConnection()


	_, err := db.Exec(crime_report_full)
	if err != nil {
		log.Fatalf("could not create view: %v", err)
	}
	http.HandleFunc("/", GETHandler_all)
	log.Fatal(http.ListenAndServe(":8080", nil))
}