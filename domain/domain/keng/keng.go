package keng

import (
	"demo/domain/domain"
	"demo/domain/domain/vo"
)

var _ domain.IDomain = new(Keng)

var (
	ErrKengNotFound = domain.NewError(
		"keng_not_found", "找不到坑")
)

const (
	KStand Kind = 1
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

	IKeng interface {
		// 新增坑
		New(keng Keng) Keng
		// 移除坑
		Remove(KengId int) error
		// 获取未被占用的坑
		GetUselessKeng(toiletId int) []*Keng
		// 查看坑的详情
		GetKengDetail(kengId int) (Keng, error)
		// 修改坑状态
		ChangeStatus(s Status) error
	}

	IKengRepo interface {
		// 新增坑
		SaveKeng(keng Keng) Keng
		// 根据状态获取坑
		GetKengByStatus(status Status) []*Keng
		// 根据状态和厕所获取坑
		GetKengByStatusAndToiletId(status Status, toiletId int) []*Keng
		// 修改坑种类
		UpdateKengKind(kind Kind) error
		// 修改坑状态
		UpdateKengStatus(s Status) error
		// 删除坑
		RemoveKeng(id int) error
		// 根据Id获取坑
		GetKengById(kengId int) (Keng, error)
	}

	Keng struct {
		KengId           vo.Id
		toiletId         vo.GenId
		Status           Status
		Kind             Kind
		LastModifiedTime int64
		CreatedTime      int64
	}
)
