package dept

import (
	models "cooky-go/models/sys"
	"cooky-go/pkg/e"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitDeptRouter(r *gin.Engine) {
	dept := r.Group("/dept")

	dept.GET("", SelectDept)
	dept.POST("", AddDept)
	dept.PUT("", EditDept)
	dept.DELETE("/:deptId", DeleteDept)
}

func SelectDept(ctx *gin.Context) {
	depts := models.SelectAllDept()
	result := make(map[string]interface{})
	result["data"] = depts
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"data": result,
	})
}

func AddDept(ctx *gin.Context) {
	var dept models.Dept
	if err := ctx.ShouldBindJSON(&dept); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "新增部门失败",
		})
		return
	}
	models.AddDept(dept)

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "新增部门成功",
	})
}

func EditDept(ctx *gin.Context) {
	var dept models.Dept
	if err := ctx.ShouldBindJSON(&dept); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "修改部门失败",
		})
		return
	}
	models.EditDept(dept)

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "修改部门成功",
	})
}

func DeleteDept(ctx *gin.Context) {
	deptId := com.StrTo(ctx.Param("deptId")).MustInt()
	models.DeleteDept(deptId)
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "删除部门成功",
	})
}
