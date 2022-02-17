package repository

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"kingsmith.com.cn/ks-horizon/infrastructure/persistence"
)

type regionRepository struct {
}

var _ regionRepositoryInterface = &regionRepository{}

type regionRepositoryInterface interface {
	Get(id uint) (user *persistence.XJUserRegion, err error)
	GetAll(mods qm.QueryMod) (userList []*persistence.XJUserRegion, err error)
	Upsert(user *persistence.XJUserRegion, updateColumns boil.Columns, insertColumns boil.Columns) (err error)
	Update(user *persistence.XJUserRegion, updateColumns boil.Columns) (id int64, err error)
	UpdateBy(cols persistence.M, mods ...qm.QueryMod) (rowsAffected int64, err error)
	Delete(user *persistence.XJUserRegion) (id int64, err error)
	DeleteBy(mods ...qm.QueryMod) (rowsAffected int64, err error)
	Count(mods ...qm.QueryMod) (count int64, err error)
}

func (r *regionRepository) Get(id uint) (user *persistence.XJUserRegion, err error) {
	user, err = persistence.FindXJUserRegionG(id)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *regionRepository) GetAll(mods qm.QueryMod) (userList []*persistence.XJUserRegion, err error) {
	userList, err = persistence.XJUserRegions(mods).AllG()
	if err != nil {
		return nil, err
	}
	return userList, err
}

func (r *regionRepository) Upsert(user *persistence.XJUserRegion, updateColumns boil.Columns, insertColumns boil.Columns) (err error) {
	return user.UpsertG(updateColumns, insertColumns)
}

func (r *regionRepository) Update(user *persistence.XJUserRegion, updateColumns boil.Columns) (id int64, err error) {
	id, err = user.UpdateG(updateColumns)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *regionRepository) UpdateBy(cols persistence.M, mods ...qm.QueryMod) (rowsAffected int64, err error) {
	rowsAffected, err = persistence.XJUserRegions(mods...).UpdateAllG(cols)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (r *regionRepository) Delete(user *persistence.XJUserRegion) (id int64, err error) {
	id, err = user.DeleteG()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *regionRepository) DeleteBy(mods ...qm.QueryMod) (rowsAffected int64, err error) {
	rowsAffected, err = persistence.XJUserRegions(mods...).DeleteAllG()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (r *regionRepository) Count(mods ...qm.QueryMod) (count int64, err error) {
	count, err = persistence.XJUserRegions(mods...).DeleteAllG()
	if err != nil {
		return 0, err
	}
	return count, nil
}
