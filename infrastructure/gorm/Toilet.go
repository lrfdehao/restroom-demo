package gorm

import "demo/domain/context/toilet"

var _ toilet.IToiletRepo = new(ToiletModel)

type ToiletModel struct {
	BaseModel
	ToiletId   string
	RestroomId string
	Status     int32
	Kind       int32
}

func (t ToiletModel) SaveToilet(toilet toilet.Toilet) toilet.Toilet {
	panic("implement me")
}

func (t ToiletModel) GetToiletByStatus(status toilet.Status) []*toilet.Toilet {
	panic("implement me")
}

func (t ToiletModel) GetToiletByStatusAndRestroomId(status toilet.Status, restroomId int) []*toilet.Toilet {
	panic("implement me")
}

func (t ToiletModel) UpdateToiletKind(kind toilet.Kind) error {
	panic("implement me")
}

func (t ToiletModel) UpdateToiletStatus(s toilet.Status) error {
	panic("implement me")
}

func (t ToiletModel) RemoveToilet(id int) error {
	panic("implement me")
}

func (t ToiletModel) GetToiletById(toiletId int) (toilet.Toilet, error) {
	panic("implement me")
}
