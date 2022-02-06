package utils

import (
	"strconv"

	"github.com/vincentJunior1/test-kriya/httpEntity"

	"github.com/gin-gonic/gin"
)

//GeneratePaginationFromRequest ..
func GeneratePaginationFromRequest(c *gin.Context) httpEntity.Pagination {
	// Initializing default
	//	var mode string
	limit := 5
	page := 1
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		}
	}
	return httpEntity.Pagination{
		Limit: limit,
		Page:  page,
	}
}
