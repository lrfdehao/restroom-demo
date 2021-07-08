package gorm

import (
	"demo/domain/context/toilet"
	"demo/domain/context/vo"
)

var _ toilet.IToiletRepo = new(toiletRepo)

type toiletRepo struct {
}

type ToiletModel struct {
	BaseModel
	ToiletId   string
	RestroomId string
	Status     int32
	Kind       int32
}

func InitToiletRepo() *toiletRepo {
	return &toiletRepo{}
}

func (t toiletRepo) SaveToilet(toilet toilet.Toilet) toilet.Toilet {
	tm := &ToiletModel{}
	tm.toiletToModel(toilet)
	if DB.Create(tm).Error != nil {
		return toilet
	}
	return toilet
}

func (t toiletRepo) GetToiletByStatus(status toilet.Status) []*toilet.Toilet {
	panic("implement me")
}

func (t toiletRepo) GetToiletByStatusAndRestroomId(status toilet.Status, restroomId int) []*toilet.Toilet {
	panic("implement me")
}

func (t toiletRepo) UpdateToiletKind(kind toilet.Kind) error {
	panic("implement me")
}

func (t toiletRepo) UpdateToiletStatus(s toilet.Status) error {
	panic("implement me")
}

func (t toiletRepo) RemoveToilet(id int) error {
	panic("implement me")
}

func (t toiletRepo) GetToiletById(toiletId string) (toilet.Toilet, error) {
	tm := ToiletModel{}
	d := DB.Where("toilet_id = ?", toiletId).First(&tm)

	return tm.modelToToilet(), d.Error
}

func (tm *ToiletModel) toiletToModel(t toilet.Toilet) {
	tm.ToiletId = t.ToiletId.Value
	tm.RestroomId = t.RestroomId.Value
	tm.Status = int32(t.Status)
	tm.Kind = int32(t.Kind)
}

func (tm *ToiletModel) modelToToilet() toilet.Toilet {
	return toilet.Toilet{
		ToiletId:         vo.GenId{Value: tm.ToiletId},
		RestroomId:       vo.GenId{Value: tm.RestroomId},
		Status:           toilet.GetStatus(tm.Status),
		Kind:             toilet.GetKind(tm.Kind),
		LastModifiedTime: tm.UpdatedAt.UnixNano(),
		CreatedTime:      tm.CreatedAt.UnixNano(),
	}
}
