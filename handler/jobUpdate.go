package handler

import (
	"net/http"

	"github.com/robfig/cron/v3"
)

// HandleUpdate update a exsit job
func HandleUpdate(c *cron.Cron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
