package repo

import (
	"micro/module/user/model"

	"gorm.io/gorm"

	common "github.com/krishnarajvr/micro-common"
)

type IUserRepo interface {
	List(tenantId int, page common.Pagination, filters model.UserFilterList) (model.Users, *common.PageResult, error)
	Get(tenantId int, id int) (*model.User, error)
	Add(form *model.User) (*model.User, error)
	Update(form *model.UserForm, id int) (*model.User, error)
	Patch(form *model.UserPatchForm, id int) (*model.User, error)
	Delete(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{
		DB: db,
	}
}

func (r UserRepo) List(tenantId int, page common.Pagination, filters model.UserFilterList) (model.Users, *common.PageResult, error) {
	users := make([]*model.User, 0)
	var totalCount int64

	db := r.DB.Scopes(common.Paginate(page)).
		Find(&users).
		Where("tenant_id = ? ", tenantId)

	if len(filters.Name) > 0 {
		db = db.Where("name like ?", filters.Name)
	}

	if len(filters.Email) > 0 {
		db = db.Where("email like ?", filters.Email)
	}

	err := db.Find(&users).Count(&totalCount).Error

	if err != nil {
		return nil, nil, err
	}

	pageResult := common.PageInfo(page, totalCount)

	if len(users) == 0 {
		return nil, nil, nil
	}

	return users, &pageResult, nil
}

func (r UserRepo) Get(tenantId int, id int) (*model.User, error) {
	user := new(model.User)
	err := r.DB.
		Where("id = ?", id).
		Where("tenant_id = ? ", tenantId).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepo) Delete(id int) (*model.User, error) {
	user := new(model.User)
	if err := r.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepo) Add(user *model.User) (*model.User, error) {

	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepo) Update(form *model.UserForm, id int) (*model.User, error) {
	user, err := form.ToModel()

	if err != nil {
		return nil, err
	}

	if err := r.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepo) Patch(form *model.UserPatchForm, id int) (*model.User, error) {
	user, err := form.ToModel()

	if err != nil {
		return nil, err
	}

	if err := r.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepo) GetByEmail(email string) (*model.User, error) {
	user := new(model.User)
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
