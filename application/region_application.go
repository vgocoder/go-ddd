package application

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"kingsmith.com.cn/ks-horizon/domain/repository"
	"kingsmith.com.cn/ks-horizon/domain/service"
	"kingsmith.com.cn/ks-horizon/infrastructure/persistence"
)

type regionApplication struct {
	repositoryFacades *repository.Facades
	serviceFacades    *service.Facades
}

func newRegionApplication() *regionApplication {
	return &regionApplication{
		repositoryFacades: repository.NewFacades(),
		serviceFacades:    service.NewFacades(),
	}
}

var _ regionApplicationInterface = &regionApplication{}

type regionApplicationInterface interface {
	Get(id uint) (user *persistence.XJUserRegion, err error)
	GetAll(mods qm.QueryMod) (userList []*persistence.XJUserRegion, err error)
	Upsert(user *persistence.XJUserRegion, updateColumns boil.Columns, insertColumns boil.Columns) (err error)
	Update(user *persistence.XJUserRegion, updateColumns boil.Columns) (id int64, err error)
	UpdateBy(cols persistence.M, mods ...qm.QueryMod) (rowsAffected int64, err error)
	Delete(user *persistence.XJUserRegion) (id int64, err error)
	DeleteBy(mods ...qm.QueryMod) (rowsAffected int64, err error)
	Count(mods ...qm.QueryMod) (count int64, err error)
	UpsertRegion(xjId uint) (err error)
}

func (ra *regionApplication) Get(id uint) (user *persistence.XJUserRegion, err error) {
	return ra.repositoryFacades.RegionRepository.Get(id)
}

func (ra *regionApplication) GetAll(mods qm.QueryMod) (userList []*persistence.XJUserRegion, err error) {
	return ra.repositoryFacades.RegionRepository.GetAll(mods)
}

func (ra *regionApplication) Upsert(user *persistence.XJUserRegion, updateColumns boil.Columns, insertColumns boil.Columns) (err error) {
	return ra.repositoryFacades.RegionRepository.Upsert(user, updateColumns, insertColumns)
}

func (ra *regionApplication) Update(user *persistence.XJUserRegion, updateColumns boil.Columns) (id int64, err error) {
	return ra.repositoryFacades.RegionRepository.Update(user, updateColumns)
}

func (ra *regionApplication) UpdateBy(cols persistence.M, mods ...qm.QueryMod) (rowsAffected int64, err error) {
	return ra.repositoryFacades.RegionRepository.UpdateBy(cols, mods...)
}

func (ra *regionApplication) Delete(user *persistence.XJUserRegion) (id int64, err error) {
	return ra.repositoryFacades.RegionRepository.Delete(user)
}

func (ra *regionApplication) DeleteBy(mods ...qm.QueryMod) (rowsAffected int64, err error) {
	return ra.repositoryFacades.RegionRepository.DeleteBy(mods...)
}

func (ra *regionApplication) Count(mods ...qm.QueryMod) (count int64, err error) {
	return ra.repositoryFacades.RegionRepository.DeleteBy(mods...)
}

func (ra *regionApplication) UpsertRegion(xjId uint) (err error) {
	return ra.serviceFacades.UserService.UpsertRegion(xjId)
}
