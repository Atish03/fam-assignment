package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"example.com/application/util"
)

type video struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishedAt string `json:"published_at"`
	Thumbnail   string `json:"thumbnail"`
	URL         string `json:"url"`
}

type Resp struct {
	Videos      []video     `json:"videos"`
	CurrentPage int         `json:"currentPage"`
	TotalPages  int         `json:"totalPages"`
	SortedIn    string      `json:"sorted_in"`
	Filter      util.Filter `json:"filter"`
}

// Handler to get videos with pagination
func GetVideos(w http.ResponseWriter, r *http.Request) {
	resp := Resp{}
	filter := util.Filter{}

	filter.Title = r.URL.Query().Get("title")

	i, err := strconv.ParseInt(r.URL.Query().Get("start"), 10, 64)
    if err != nil {
		log.Println("invalid start time:", err)
        http.Error(w, "not a valid start time", http.StatusBadRequest)
		return
    }
    filter.StartDate = time.UnixMilli(i)

	i, err = strconv.ParseInt(r.URL.Query().Get("end"), 10, 64)
    if err != nil {
		log.Println("invalid end time", err)
        http.Error(w, "not a valid end time", http.StatusBadRequest)
		return
    }
    filter.EndDate = time.UnixMilli(i)	

	totalPages, err := util.GetTotalPages(filter)
	if err != nil {
		log.Println("Error getting total pages:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	resp.TotalPages = totalPages

	if resp.TotalPages == 0 {
		resp.TotalPages = 1
	}

	sortOptions := map[string]string{
		"published_at_desc": "published_at DESC",
		"title_asc":         "title ASC",
	}
	
	sortOrder := sortOptions["published_at_desc"]
	if val, ok := sortOptions[r.URL.Query().Get("sort")]; ok {
		sortOrder = val
	}

	page := 1
	if pageParam := r.URL.Query().Get("page"); pageParam != "" {
		var err error
		page, err = strconv.Atoi(pageParam)
		if err != nil || page <= 0 || page > resp.TotalPages {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
	}

	rows, err := util.GetVideoData(sortOrder, filter, page)
	if err != nil {
		log.Println("cannot get video data")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	videos := []video{}
	for rows.Next() {
		var video video
		if err := rows.Scan(&video.ID, &video.Title, &video.Description, &video.PublishedAt, &video.Thumbnail); err != nil {
			http.Error(w, "Error reading video data", http.StatusInternalServerError)
			log.Println("Error reading row:", err)
			return
		}
		video.URL = fmt.Sprintf("https://www.youtube.com/watch?v=%s", video.ID)
		videos = append(videos, video)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating through videos", http.StatusInternalServerError)
		log.Println("Error iterating rows:", err)
		return
	}

	resp.Videos = videos
	resp.CurrentPage = page
	resp.SortedIn = sortOrder
	resp.Filter = filter

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		log.Println("Error encoding response:", err)
		return
	}
}