package util

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"

	"cooky-go/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	pageSize, err := com.StrTo(c.Query("pageSize")).Int()
	if err != nil {
		pageSize = setting.PageSize
	}
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
