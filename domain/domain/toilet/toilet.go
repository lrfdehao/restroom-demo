package toilet

import (
	"demo/domain/domain"
	"demo/domain/domain/vo"
)

var _ domain.IDomain = new(Toilet)

var (
	ErrToiletNotFound = domain.NewError("toilet_not_found", "找不到厕所")
)

const (
	male   Gender = 1
	female Gender = 2
)

type (
	Gender int

	IToilet interface {
		// 新增厕所
		New(toilet IToilet) Toilet
		// 移除厕所
		Remove(toiletId int) error

		// 修改厕所性别
		ChangeGender(toilet int, g Gender) error
	}

	IToiletRepo interface {
	}

	Toilet struct {
		ToiletId         vo.GenId
		Floor            int
		Address          string
		Gender           Gender
		Maxsize          int
		IsActive         bool
		LastModifiedTime int64
		CreatedTime      int64
	}
)
