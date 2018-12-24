package category

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/service"
)

func MainList(c *gin.Context) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "1"))

	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	infos, count, err := model.ListMainCategory(offset, limit)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "getlist from database has error: ")
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount:   count,
		CategoryList: infos,
	})
}

func SubList(c *gin.Context) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "1"))

	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	pid, err := strconv.Atoi(c.DefaultQuery("pid", ""))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	infos, count, err := model.ListSubCategory(offset, limit, pid)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "getlist from database has error: ")
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount:   count,
		CategoryList: infos,
	})
}

func MainListWithSubCount(c *gin.Context) {

	infos, count, err := service.ListMainCategoryWithSubCount()
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "getlist from database has error: ")
		return
	}

	SendResponse(c, nil, MainListWithSubCountResponse{
		TotalCount:   count,
		CategoryList: infos,
	})
}

func MainListAll(c *gin.Context) {

	infos, count, err := model.ListMainCategoryAll()
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "getlist from database has error: ")
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount:   count,
		CategoryList: infos,
	})
}

func SubListAllByMainId(c *gin.Context) {

	pid, err := strconv.Atoi(c.DefaultQuery("pid", ""))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	infos, err := model.ListSubCategoryAll(pid)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "getlist from database has error: ")
		return
	}

	SendResponse(c, nil, ListSortAllResponse{
		CategoryList: infos,
	})
}

func SubListAll(c *gin.Context) {
	infos, err := model.ListSubNoMainCategory()
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "getlist from database has error: ")
		return
	}

	SendResponse(c, nil, ListSortAllResponse{
		CategoryList: infos,
	})
}