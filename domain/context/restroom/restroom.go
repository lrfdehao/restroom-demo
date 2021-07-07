package restroom

import (
	"demo/domain/context"
	"demo/domain/context/vo"
)

var _ context.IDomain = new(Restroom)

const (
	male   Gender = 1
	female Gender = 2
)

type (
	Gender int

	IRestroom interface {
		// 新增厕所
		New(ir IRestroom) Restroom
		// 移除厕所
		Remove(toiletId int) error

		// 修改厕所性别
		ChangeGender(restroomId int, g Gender) error
	}

	IToiletRepo interface {
	}

	Restroom struct {
		RestroomId       vo.GenId
		Floor            int
		Address          string
		Gender           Gender
		Maxsize          int
		IsActive         bool
		LastModifiedTime int64
		CreatedTime      int64
	}
)
