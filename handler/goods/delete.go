package goods

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
)

func Delete(c *gin.Context) {
	var filepath string

	gId, _ := strconv.Atoi(c.Param("id"))
	goods, err := model.GetGoodsModelById(uint64(gId))
	if err != nil {
		SendResponse(c, errno.ErrGoodsNotFount, nil)
		return
	}

	filepath = goods.GoodsPhoto

	if err := os.Remove(filepath); err != nil {
		log.Errorf(err, "del shoplogo occured error is :")
	}

	if err := model.DeleteGoodsModel(uint64(gId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}

func DeleteByMain(c *gin.Context) {
	var filepath string

	merid, err := strconv.Atoi(c.DefaultQuery("merid", "0"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	mid, err := strconv.Atoi(c.DefaultQuery("mid", "0"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	infos,_ ,err := model.ListGoodsModelForMerchantAndMain(merid,mid)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	for _, goods := range infos {
		filepath = goods.GoodsPhoto

		if err := os.Remove(filepath); err != nil {
			log.Errorf(err, "del shoplogo occured error is :")
		}
	}

	if err := model.DeleteGoodsModelByMain(merid,mid); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}


func DeleteBySub(c *gin.Context) {
	var filepath string

	merid, err := strconv.Atoi(c.DefaultQuery("merid", "0"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	sid, err := strconv.Atoi(c.DefaultQuery("sid", "0"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	infos,_ ,err := model.ListGoodsModelForMerchantAndSubAll(merid,sid)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	for _, goods := range infos {
		filepath = goods.GoodsPhoto

		if err := os.Remove(filepath); err != nil {
			log.Errorf(err, "del shoplogo occured error is :")
		}
	}

	if err := model.DeleteGoodsModelBySub(merid,sid); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}