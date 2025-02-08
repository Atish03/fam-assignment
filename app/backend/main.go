package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"example.com/application/handler"
	"example.com/application/util"
)

func main() {
	util.InitDB()
	defer util.DB.Close()

	exportDir := "./frontend"

	mux := http.NewServeMux()
	mux.HandleFunc("/api/videos", handler.GetVideos)

	fs := http.FileServer(http.Dir(exportDir))
	mux.Handle("/", fs)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
