package keng

import "demo/domain/domain/keng"

var _ keng.IKeng = new(kengImpl)

type kengImpl struct {
	kengRepo keng.IKengRepo
}

func (k kengImpl) New(keng keng.Keng) keng.Keng {
	return k.kengRepo.SaveKeng(keng)
}

func (k kengImpl) Remove(KengId int) error {
	return k.kengRepo.RemoveKeng(KengId)
}

func (k kengImpl) GetUselessKeng(toiletId int) []*keng.Keng {
	return k.kengRepo.GetKengByStatusAndToiletId(keng.SUseless, toiletId)
}

func (k kengImpl) GetKengDetail(kengId int) (keng.Keng, error) {
	return k.kengRepo.GetKengById(kengId)
}

func (k kengImpl) ChangeStatus(s keng.Status) error {
	panic("implement me")
}
