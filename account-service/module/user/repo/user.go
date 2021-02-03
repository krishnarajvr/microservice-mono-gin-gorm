package repo

import (
	"micro/module/user/model"

	"gorm.io/gorm"

	common "github.com/krishnarajvr/micro-common"
)

type IUserRepo interface {
	List(page common.Pagination) (model.Users, *common.PageResult, error)
	Get(id int) (*model.User, error)
	Add(form *model.UserForm) (*model.User, error)
	Update(form *model.UserForm, id int) (*model.User, error)
	Patch(form *model.UserPatchForm, id int) (*model.User, error)
	Delete(id int) (*model.User, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{
		DB: db,
	}
}

func (r UserRepo) List(page common.Pagination) (model.Users, *common.PageResult, error) {

	users := make([]*model.User, 0)

	var totalCount int64
	if err := r.DB.Table("users").Count(&totalCount).Error; err != nil {
		return nil, nil, err
	}

	if err := r.DB.Scopes(common.Paginate(page)).Find(&users).Error; err != nil {
		return nil, nil, err
	}

	pageResult := common.PageInfo(page, totalCount)

	if len(users) == 0 {
		return nil, nil, nil
	}

	return users, &pageResult, nil
}

func (r UserRepo) Get(id int) (*model.User, error) {
	user := new(model.User)
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
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

func (r UserRepo) Add(form *model.UserForm) (*model.User, error) {
	user, err := form.ToModel()
	//Todo - Get from token
	user.TenantId = 1

	if err != nil {
		return nil, err
	}

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
