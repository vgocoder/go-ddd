package repository

import (
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"kingsmith.com.cn/ks-horizon/infrastructure/persistence"
)

type userRepository struct {
}

var _ userRepositoryInterface = &userRepository{}

type userRepositoryInterface interface {
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

func (u *userRepository) Get(id uint) (user *persistence.XJUser, err error) {
	user, err = persistence.FindXJUserG(id)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (u *userRepository) GetOne(mods ...qm.QueryMod) (user *persistence.XJUser, err error) {
	user, err = persistence.XJUsers(mods...).OneG()
	if err != nil {
		return nil, err
	}
	return user, err
}

func (u *userRepository) GetAll(mods ...qm.QueryMod) (userList []*persistence.XJUser, err error) {
	userList, err = persistence.XJUsers(mods...).AllG()
	if err != nil {
		return nil, err
	}
	return userList, err
}

func (u *userRepository) Upsert(user *persistence.XJUser, updateColumns boil.Columns, insertColumns boil.Columns) (err error) {
	return user.UpsertG(updateColumns, insertColumns)
}

func (u *userRepository) Update(user *persistence.XJUser, updateColumns boil.Columns) (id int64, err error) {
	id, err = user.UpdateG(updateColumns)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *userRepository) UpdateBy(cols persistence.M, mods ...qm.QueryMod) (rowsAffected int64, err error) {
	rowsAffected, err = persistence.XJUsers(mods...).UpdateAllG(cols)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (u *userRepository) Delete(user *persistence.XJUser) (id int64, err error) {
	id, err = user.DeleteG()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *userRepository) DeleteBy(mods ...qm.QueryMod) (rowsAffected int64, err error) {
	rowsAffected, err = persistence.XJUsers(mods...).DeleteAllG()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (u *userRepository) Count(mods ...qm.QueryMod) (count int64, err error) {
	fmt.Println(count)
	count, err = persistence.XJUsers(mods...).CountG()
	fmt.Println(count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
