package api

import (
	"demo/domain/common"
	"demo/domain/context/toilet"
	"demo/domain/context/vo"
	"demo/domain/service"
	"demo/infrastructure/gorm"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"io/ioutil"
	"time"
)

type toiletReq struct {
	RestroomId string        `json:"restroom_id"`
	Status     toilet.Status `json:"status"`
	Kind       toilet.Kind   `json:"kind"`
}

func CreateToilet(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		SendFailResponse(c, common.ErrToiletNotFound)
		return
	}
	var req toiletReq
	_ = json.Unmarshal(body, &req)

	id := vo.GenId{Value: uuid.New()}
	errG := id.Valid()
	if errG != nil {
		SendFailResponse(c, errG, nil)
		return
	}

	t := NewImpl().New(toilet.Toilet{
		ToiletId:         id,
		RestroomId:       vo.GenId{Value: req.RestroomId},
		Status:           req.Status,
		Kind:             req.Kind,
		LastModifiedTime: time.Nanosecond.Nanoseconds(),
		CreatedTime:      time.Nanosecond.Nanoseconds(),
	})

	SendOkResponse(c, t.ToiletId)
}

func GetToilet(c *gin.Context) {
	toiletIdReq := c.Query("toilet_id")

	detail, err := NewImpl().GetToiletDetail(toiletIdReq)
	if err != nil {
		SendFailResponse(c, common.ErrToiletNotFound)
		return
	}

	SendOkResponse(c, detail)
}

func NewImpl() *service.ToiletImpl {
	return service.NewToiletImpl(gorm.InitToiletRepo())
}
