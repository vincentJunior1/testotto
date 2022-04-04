package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/copier"
	"github.com/vincentJunior1/test-kriya/httpEntity"
	"github.com/vincentJunior1/test-kriya/models"
	"github.com/vincentJunior1/test-kriya/utils"
)

func GetOmzet(c *gin.Context) {
	merchant := &[]models.Transaction{}
	pagination := utils.GeneratePaginationFromRequest(c)
	merchantID, _ := strconv.Atoi(c.Param("merchant_id"))
	query := c.Request.URL.Query()
	var startDate time.Time
	var endDate time.Time

	userLogin := &models.User{}

	if c.Value("UserLogin") != nil {
		userLogin = c.Value("UserLogin").(*models.User)
	}

	for i, val := range query {
		fmt.Println(i)
		if len(query[i][0]) > 0 {
			if i == "start_date" {
				tmp, _ := strconv.ParseInt(val[0], 10, 64)
				startDate = time.Unix(tmp, 0)
			} else if i == "end_date" {
				tmp, _ := strconv.ParseInt(val[0], 10, 64)
				endDate = time.Unix(tmp, 0)
			}
		}
	}

	err := models.GetMerchantOmzet(merchantID, startDate, endDate, userLogin.ID, merchant, &pagination)

	if err != nil {
		log.Fatal(err)
		response := &httpEntity.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Error get data",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
		return
	}
	res := []httpEntity.OmzetResponse{}
	for _, val := range *merchant {
		tmp := httpEntity.OmzetResponse{
			MerchantName:    val.MerchantName,
			Omzet:           val.Omzet,
			TransactionDate: val.CreatedAt.Unix(),
		}
		res = append(res, tmp)
	}
	response := &httpEntity.Response{
		Code:       http.StatusOK,
		Message:    "Success",
		Status:     true,
		Data:       res,
		Pagination: pagination,
	}
	c.JSON(http.StatusOK, response)
	c.Abort()
}
