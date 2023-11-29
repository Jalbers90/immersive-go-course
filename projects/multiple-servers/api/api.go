package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Image struct {
	ImageId int    `json:"image_id"`
	Url     string `json:"url"`
	Title   string `json:"title"`
	AltText string `json:"alt_text"`
}

var SQL_GET_IMAGES = "SELECT * FROM IMAGES"

func Run(port *string, dbURL string) {
	// Open a connection to the database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database!")
	http.HandleFunc("/api/images.json", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			w.Write([]byte("api server POST"))

		} else if r.Method == "GET" {
			// w.Write([]byte("api server GET"))
			sqlRes, err := db.Query(SQL_GET_IMAGES)
			if err != nil {
				http.Error(w, "Internal Server Error", 500)
				return
			}
			defer sqlRes.Close()
			images := []Image{}
			for sqlRes.Next() {
				var image Image
				err := sqlRes.Scan(&image.ImageId, &image.Title, &image.Url, &image.AltText)
				if err != nil {
					http.Error(w, "Internal Server Error", 500)
					return
				}
				images = append(images, image)
			}
			jsonData, err := json.Marshal(images)
			if err != nil {
				http.Error(w, "Internal Server Error", 500)
				return
			}
			// w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin
			// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			// w.Header().Set("Content-Type", "application/json")
			w.Header().Add("Content-Type", "text/json")
			w.Header().Add("Access-Control-Allow-Origin", "*")

			w.Write(jsonData)
		}
	})
	listenAddr := fmt.Sprintf(":%s", *port)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
