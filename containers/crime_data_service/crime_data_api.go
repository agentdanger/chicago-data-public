package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	//"strings"
)


type crime_response []crime_data

type crime_data struct {
	Id                   string
	Case_number          *string
	Date                 *string
	Block                *string
	Iucr                 *string
	Primary_type         *string
	Description          *string
	Location_description *string
	Arrest               bool
	Domestic             bool
	Beat                 *string
	District             *string
	Ward                 *string
	Community_area       *string
	Fbi_code             string
	X_coordinate         *string
	Y_coordinate         *string
	Year                 string
	Updated_on           string
	Latitude             *string
	Longitude            *string
}

func main() {
	host := "10.100.138.27" // this is the host of the PSQL database on AWS.
	port := 5432
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname   := "fullstack_api"
	app_token := os.Getenv("API_TOKEN")

	crime_table := ` 
CREATE TABLE IF NOT EXISTS t_crime_data (
id serial PRIMARY KEY, 
case_number VARCHAR,  
date DATE,  
block VARCHAR, 
iucr VARCHAR, 
primary_type VARCHAR, 
description VARCHAR, 
location_description VARCHAR, 
arrest BOOLEAN, 
domestic BOOLEAN, 
beat VARCHAR, 
district VARCHAR, 
ward NUMERIC, 
community_area VARCHAR, 
fbi_code VARCHAR, 
x_coordinate NUMERIC, 
y_coordinate NUMERIC, 
year NUMERIC, 
updated_on DATE, 
latitude NUMERIC, 
longitude NUMERIC )
;
`

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	_, err = db.Exec(crime_table)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	max_offset_qry := ` 
SELECT id 
FROM t_crime_data 
ORDER BY id DESC 
LIMIT 1
;
`
	var max_offset int

	db.QueryRow(max_offset_qry).Scan(&max_offset)

	offset := strconv.Itoa(max_offset)

	url1 := fmt.Sprintf("https://data.cityofchicago.org/resource/crimes.json?$limit=100000&$where=id>%s&$order=id&$$app_token=%s",
		offset, app_token)

	var url string = url1

	response, err := http.Get(url)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	crime_json := crime_response{}
	json.Unmarshal([]byte(body), &crime_json)

	//fmt.Printf("%s", crime_json)  // for debugging

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	for _, data := range crime_json {
		// fmt.Printf("%s", data)  // for debugging
		sqlStatement := `
INSERT INTO t_crime_data (
id, 
case_number,  
date,  
block, 
iucr, 
primary_type, 
description, 
location_description, 
arrest, 
domestic, 
beat, 
district, 
ward, 
community_area, 
fbi_code, 
x_coordinate, 
y_coordinate, 
year, 
updated_on, 
latitude, 
longitude) 
VALUES (
CAST($1 AS NUMERIC), 
CAST($2 AS VARCHAR), 
CAST($3 AS DATE), 
CAST($4 AS VARCHAR), 
CAST($5 AS VARCHAR), 
CAST($6 AS VARCHAR), 
CAST($7 AS VARCHAR), 
CAST($8 AS VARCHAR), 
CAST($9 AS BOOLEAN), 
CAST($10 AS BOOLEAN), 
CAST($11 AS VARCHAR), 
CAST($12 AS VARCHAR), 
CAST($13 AS NUMERIC),  
CAST($14 AS VARCHAR), 
CAST($15 AS VARCHAR), 
CAST($16 AS NUMERIC), 
CAST($17 AS NUMERIC), 
CAST($18 AS NUMERIC), 
CAST($19 AS DATE), 
CAST($20 AS NUMERIC), 
CAST($21 AS NUMERIC) 
)
`
		_, err = db.Exec(sqlStatement,
			data.Id,
			data.Case_number,
			data.Date,
			data.Block,
			data.Iucr,
			data.Primary_type,
			data.Description,
			data.Location_description,
			data.Arrest,
			data.Domestic,
			data.Beat,
			data.District,
			data.Ward,
			data.Community_area,
			data.Fbi_code,
			data.X_coordinate,
			data.Y_coordinate,
			data.Year,
			data.Updated_on,
			data.Latitude,
			data.Longitude)

		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
			}
		}(db)
	}
}