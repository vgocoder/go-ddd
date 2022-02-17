package application

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"kingsmith.com.cn/ks-horizon/domain/repository"
	"kingsmith.com.cn/ks-horizon/domain/service"
	"kingsmith.com.cn/ks-horizon/infrastructure/persistence"
)

type didApplication struct {
	repositoryFacades *repository.Facades
	serviceFacades    *service.Facades
}

func newDidApplication() *didApplication {
	return &didApplication{
		repositoryFacades: repository.NewFacades(),
	}
}

var _ didApplicationInterface = &didApplication{}

type didApplicationInterface interface {
	Get(id uint) (user *persistence.XJUserDid, err error)
	GetAll(mods ...qm.QueryMod) (userList []*persistence.XJUserDid, err error)
	Upsert(user *persistence.XJUserDid, updateColumns boil.Columns, insertColumns boil.Columns) (err error)
	Update(user *persistence.XJUserDid, updateColumns boil.Columns) (id int64, err error)
	UpdateBy(cols persistence.M, mods ...qm.QueryMod) (rowsAffected int64, err error)
	Delete(user *persistence.XJUserDid) (id int64, err error)
	DeleteBy(mods ...qm.QueryMod) (rowsAffected int64, err error)
	Count(mods ...qm.QueryMod) (count int64, err error)
}

func (da *didApplication) Get(id uint) (user *persistence.XJUserDid, err error) {
	return da.repositoryFacades.DidRepository.Get(id)
}

func (da *didApplication) GetAll(mods ...qm.QueryMod) (userList []*persistence.XJUserDid, err error) {
	return da.repositoryFacades.DidRepository.GetAll(mods...)
}

func (da *didApplication) Upsert(user *persistence.XJUserDid, updateColumns boil.Columns, insertColumns boil.Columns) (err error) {
	return da.repositoryFacades.DidRepository.Upsert(user, updateColumns, insertColumns)
}

func (da *didApplication) Update(user *persistence.XJUserDid, updateColumns boil.Columns) (id int64, err error) {
	return da.repositoryFacades.DidRepository.Update(user, updateColumns)
}

func (da *didApplication) UpdateBy(cols persistence.M, mods ...qm.QueryMod) (rowsAffected int64, err error) {
	return da.repositoryFacades.DidRepository.UpdateBy(cols, mods...)
}

func (da *didApplication) Delete(user *persistence.XJUserDid) (id int64, err error) {
	return da.repositoryFacades.DidRepository.Delete(user)
}

func (da *didApplication) DeleteBy(mods ...qm.QueryMod) (rowsAffected int64, err error) {
	return da.repositoryFacades.DidRepository.DeleteBy(mods...)
}

func (da *didApplication) Count(mods ...qm.QueryMod) (count int64, err error) {
	return da.repositoryFacades.DidRepository.DeleteBy(mods...)
}
