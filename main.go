package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	http.HandleFunc("/Products/", func(w http.ResponseWriter, r *http.Request) {

		data, err := query()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}

func query() (productsData, error) {
	repos := productsData{}

	db, err := sql.Open("sqlite3", "testdb.db")
	if err != nil {
		return productsData{}, err
	}
	defer db.Close()
	rows, err := db.Query("select Name,Description from Product")
	if err != nil {
		return productsData{}, err
	}
	for rows.Next() {
		d := productData{}
		err = rows.Scan(&d.Name, &d.Description)
		if err != nil {
			return productsData{}, err
		}
		repos.ProductsData = append(repos.ProductsData, d)
	}
	return repos, nil
}

type productsData struct {
	ProductsData []productData
}
type productData struct {
	Name        string `json:"name"`
	Description string `json:"Description"`
}
