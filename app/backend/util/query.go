package util

import (
	"database/sql"
	"fmt"
	"math"
	"time"
)

type Filter struct {
	Title     string    `json:"title"`
	StartDate time.Time `json:"start"`
	EndDate   time.Time `json:"end"`
}

type Cursor struct {
	Id          int       `json:"id"`
	PublishedAt	time.Time `json:"published_at"`
}

var LIMIT int = 8

func GetVideoData(sortOrder string, filter Filter, page int) (*sql.Rows, error) {
	query := fmt.Sprintf(`
	SELECT video_id, title, description, published_at, thumbnail_url 
	FROM videos
	WHERE 1=1 AND
	title ILIKE $1 AND
	published_at BETWEEN $2 AND $3
	ORDER BY %s
	LIMIT $4 OFFSET $5`, sortOrder)

	offset := (page - 1) * LIMIT
	
	rows, err := DB.Query(query, "%" + filter.Title + "%", filter.StartDate, filter.EndDate, LIMIT, offset)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func GetTotalPages(filter Filter) (int, error) {
	var totalRecords int

	query := fmt.Sprintf(`
	SELECT COUNT(*)
	FROM videos
	WHERE 1=1 AND
	title ILIKE $1 AND
	published_at BETWEEN $2 AND $3`)

	err := DB.QueryRow(query, "%" + filter.Title + "%", filter.StartDate, filter.EndDate).Scan(&totalRecords)
	if err != nil {
		return 0, err
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(LIMIT)))
	return totalPages, nil
}