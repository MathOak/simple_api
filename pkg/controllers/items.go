package controllers

import (
	"encoding/json"
	"net/http"
	"simple_api/pkg/models"
	"strconv"

	"gorm.io/gorm"
)

func GetItems(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var items []models.Item
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit
	db.Limit(limit).Offset(offset).Find(&items)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
