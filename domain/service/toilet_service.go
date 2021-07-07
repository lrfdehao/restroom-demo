package service

import "demo/domain/context/toilet"

var _ toilet.IToilet = new(ToiletImpl)

type ToiletImpl struct {
	toiletRepo toilet.IToiletRepo
}

func NewToiletImpl(t toilet.IToiletRepo) *ToiletImpl {
	return &ToiletImpl{toiletRepo: t}
}

func (t ToiletImpl) New(toilet toilet.Toilet) toilet.Toilet {
	return t.toiletRepo.SaveToilet(toilet)
}

func (t ToiletImpl) Remove(toiletId int) error {
	return t.toiletRepo.RemoveToilet(toiletId)
}

func (t ToiletImpl) GetUselessToilet(restroomId int) []*toilet.Toilet {
	return t.toiletRepo.GetToiletByStatusAndRestroomId(toilet.SUseless, restroomId)
}

func (t ToiletImpl) GetToiletDetail(toiletId int) (toilet.Toilet, error) {
	return t.toiletRepo.GetToiletById(toiletId)
}

func (t ToiletImpl) ChangeStatus(s toilet.Status) error {
	return t.toiletRepo.UpdateToiletStatus(s)
}
