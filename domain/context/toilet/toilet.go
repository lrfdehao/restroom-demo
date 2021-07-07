package toilet

import (
	"demo/domain/context"
	"demo/domain/context/vo"
)

var _ context.IDomain = new(Toilet)

const (
	KSquat Kind = 1
	KSit   Kind = 2
)

const (
	SUsed    Status = 1
	SUseless Status = 2
	SUnknown Status = 3
)

type (
	// 坑种类
	Kind   int
	Status int

	IToilet interface {
		// 新增坑
		New(toilet Toilet) Toilet
		// 移除坑
		Remove(toiletId int) error
		// 获取未被占用的坑
		GetUselessToilet(restroomId int) []*Toilet
		// 查看坑的详情
		GetToiletDetail(toiletId int) (Toilet, error)
		// 修改坑状态
		ChangeStatus(s Status) error
	}

	IToiletRepo interface {
		// 新增坑
		SaveToilet(toilet Toilet) Toilet
		// 根据状态获取坑
		GetToiletByStatus(status Status) []*Toilet
		// 根据状态和厕所获取坑
		GetToiletByStatusAndRestroomId(status Status, restroomId int) []*Toilet
		// 修改坑种类
		UpdateToiletKind(kind Kind) error
		// 修改坑状态
		UpdateToiletStatus(s Status) error
		// 删除坑
		RemoveToilet(id int) error
		// 根据Id获取坑
		GetToiletById(toiletId int) (Toilet, error)
	}

	Toilet struct {
		ToiletId         vo.GenId
		RestroomId       vo.GenId
		Status           Status
		Kind             Kind
		LastModifiedTime int64
		CreatedTime      int64
	}
)
