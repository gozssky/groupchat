package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func parseRequestPage(c *gin.Context) (pageIndex, pageSize int, err error) {
	var page struct {
		Index int `json:"pageIndex"`
		Size  int `json:"pageSize"`
	}
	if err = c.ShouldBindJSON(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	return page.Index, page.Size, nil
}

func convertPageToRange(totalSize int, pageIndex int, pageSize int) (start, end int) {
	if pageIndex >= 0 {
		start = pageIndex * pageSize
		end = start + pageSize
	} else {
		start = totalSize + pageIndex*pageSize
		end = start + pageSize
	}
	if start < 0 {
		start = 0
	}
	if end > totalSize {
		end = totalSize
	}
	if start > end {
		return 0, 0
	}
	return start, end
}
