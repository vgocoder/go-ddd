package application

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"kingsmith.com.cn/ks-horizon/domain/repository"
	"kingsmith.com.cn/ks-horizon/infrastructure/persistence"
)

type userApplication struct {
	repositoryFacades *repository.Facades
}

func newUserApplication() *userApplication {
	return &userApplication{
		repositoryFacades: repository.NewFacades(),
	}
}

var _ userApplicationInterface = &userApplication{}

type userApplicationInterface interface {
	Get(id uint) (user *persistence.XJUser, err error)
	GetOne(mods ...qm.QueryMod) (user *persistence.XJUser, err error)
	GetAll(mods ...qm.QueryMod) (userList []*persistence.XJUser, err error)
	Upsert(user *persistence.XJUser, updateColumns boil.Columns, insertColumns boil.Columns) (err error)
	Update(user *persistence.XJUser, updateColumns boil.Columns) (id int64, err error)
	UpdateBy(cols persistence.M, mods ...qm.QueryMod) (rowsAffected int64, err error)
	Delete(user *persistence.XJUser) (id int64, err error)
	DeleteBy(mods ...qm.QueryMod) (rowsAffected int64, err error)
	Count(mods ...qm.QueryMod) (count int64, err error)
}

func (ua *userApplication) Get(id uint) (user *persistence.XJUser, err error) {
	return ua.repositoryFacades.UserRepository.Get(id)
}

func (ua *userApplication) GetOne(mods ...qm.QueryMod) (user *persistence.XJUser, err error) {
	return ua.repositoryFacades.UserRepository.GetOne(mods...)
}

func (ua *userApplication) GetAll(mods ...qm.QueryMod) (userList []*persistence.XJUser, err error) {
	return ua.repositoryFacades.UserRepository.GetAll(mods...)
}

func (ua *userApplication) Upsert(user *persistence.XJUser, updateColumns boil.Columns, insertColumns boil.Columns) (err error) {
	return ua.repositoryFacades.UserRepository.Upsert(user, updateColumns, insertColumns)
}

func (ua *userApplication) Update(user *persistence.XJUser, updateColumns boil.Columns) (id int64, err error) {
	return ua.repositoryFacades.UserRepository.Update(user, updateColumns)
}

func (ua *userApplication) UpdateBy(cols persistence.M, mods ...qm.QueryMod) (rowsAffected int64, err error) {
	return ua.repositoryFacades.UserRepository.UpdateBy(cols, mods...)
}

func (ua *userApplication) Delete(user *persistence.XJUser) (id int64, err error) {
	return ua.repositoryFacades.UserRepository.Delete(user)
}

func (ua *userApplication) DeleteBy(mods ...qm.QueryMod) (rowsAffected int64, err error) {
	return ua.repositoryFacades.UserRepository.DeleteBy(mods...)
}

func (ua *userApplication) Count(mods ...qm.QueryMod) (count int64, err error) {
	return ua.repositoryFacades.UserRepository.Count(mods...)
}
