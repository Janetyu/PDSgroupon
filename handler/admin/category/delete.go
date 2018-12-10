package category

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"PDSgroupon/model"
	. "PDSgroupon/handler"
	"PDSgroupon/pkg/errno"
)

// 如果删除主类别，则该类别的子类别也同时删除
func DeleteMain(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteMainCategory(uint64(categoryId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}


// 如果删除子类别，则只删除该子类别本身
func DeleteSub(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteSubCategory(uint64(categoryId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}