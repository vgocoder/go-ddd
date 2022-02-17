package repository

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"kingsmith.com.cn/ks-horizon/infrastructure/persistence"
)

type didRepository struct {
}

var _ didRepositoryInterface = &didRepository{}

type didRepositoryInterface interface {
	Get(id uint) (user *persistence.XJUserDid, err error)
	GetAll(mods ...qm.QueryMod) (userList []*persistence.XJUserDid, err error)
	Upsert(user *persistence.XJUserDid, updateColumns boil.Columns, insertColumns boil.Columns) (err error)
	Update(user *persistence.XJUserDid, updateColumns boil.Columns) (id int64, err error)
	UpdateBy(cols persistence.M, mods ...qm.QueryMod) (rowsAffected int64, err error)
	Delete(user *persistence.XJUserDid) (id int64, err error)
	DeleteBy(mods ...qm.QueryMod) (rowsAffected int64, err error)
	Count(mods ...qm.QueryMod) (count int64, err error)
}

func (d *didRepository) Get(id uint) (user *persistence.XJUserDid, err error) {
	user, err = persistence.FindXJUserDidG(id)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (d *didRepository) GetAll(mods ...qm.QueryMod) (userList []*persistence.XJUserDid, err error) {
	userList, err = persistence.XJUserDids(mods...).AllG()
	if err != nil {
		return nil, err
	}
	return userList, err
}

func (d *didRepository) Upsert(user *persistence.XJUserDid, updateColumns boil.Columns, insertColumns boil.Columns) (err error) {
	return user.UpsertG(updateColumns, insertColumns)
}

func (d *didRepository) Update(user *persistence.XJUserDid, updateColumns boil.Columns) (id int64, err error) {
	id, err = user.UpdateG(updateColumns)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *didRepository) UpdateBy(cols persistence.M, mods ...qm.QueryMod) (rowsAffected int64, err error) {
	rowsAffected, err = persistence.XJUserDids(mods...).UpdateAllG(cols)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (d *didRepository) Delete(user *persistence.XJUserDid) (id int64, err error) {
	id, err = user.DeleteG()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *didRepository) DeleteBy(mods ...qm.QueryMod) (rowsAffected int64, err error) {
	rowsAffected, err = persistence.XJUserDids(mods...).DeleteAllG()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (d *didRepository) Count(mods ...qm.QueryMod) (count int64, err error) {
	count, err = persistence.XJUserDids(mods...).CountG()
	if err != nil {
		return 0, err
	}
	return count, nil
}
