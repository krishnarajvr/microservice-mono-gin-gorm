package repo

import (
	"micro/module/tenant/model"

	"gorm.io/gorm"

	common "github.com/krishnarajvr/micro-common"
)

type ITenantRepo interface {
	List(page common.Pagination) (model.Tenants, error)
	Get(id int) (*model.Tenant, error)
	Add(form *model.TenantForm) (*model.Tenant, error)
	Update(form *model.TenantForm, id int) (*model.Tenant, error)
	Patch(form *model.TenantPatchForm, id int) (*model.Tenant, error)
	Delete(id int) (*model.Tenant, error)
}

type TenantRepo struct {
	DB *gorm.DB
}

func NewTenantRepo(db *gorm.DB) TenantRepo {
	return TenantRepo{
		DB: db,
	}
}

func (r TenantRepo) List(page common.Pagination) (model.Tenants, error) {

	tenants := make([]*model.Tenant, 0)
	if err := r.DB.Scopes(common.Paginate(page)).Find(&tenants).Error; err != nil {
		return nil, err
	}

	if len(tenants) == 0 {
		return nil, nil
	}

	return tenants, nil
}

func (r TenantRepo) Get(id int) (*model.Tenant, error) {
	tenant := new(model.Tenant)
	if err := r.DB.Where("id = ?", id).First(&tenant).Error; err != nil {
		return nil, err
	}

	return tenant, nil
}

func (r TenantRepo) Delete(id int) (*model.Tenant, error) {
	tenant := new(model.Tenant)
	if err := r.DB.Where("id = ?", id).Delete(&tenant).Error; err != nil {
		return nil, err
	}

	return tenant, nil
}

func (r TenantRepo) Add(form *model.TenantForm) (*model.Tenant, error) {
	tenant, err := form.ToModel()

	if err != nil {
		return nil, err
	}

	if err := r.DB.Create(&tenant).Error; err != nil {
		return nil, err
	}

	return tenant, nil
}

func (r TenantRepo) Update(form *model.TenantForm, id int) (*model.Tenant, error) {
	tenant, err := form.ToModel()

	if err != nil {
		return nil, err
	}

	if err := r.DB.Where("id = ?", id).Updates(&tenant).Error; err != nil {
		return nil, err
	}

	return tenant, nil
}

func (r TenantRepo) Patch(form *model.TenantPatchForm, id int) (*model.Tenant, error) {
	tenant, err := form.ToModel()

	if err != nil {
		return nil, err
	}

	if err := r.DB.Where("id = ?", id).Updates(&tenant).Error; err != nil {
		return nil, err
	}

	return tenant, nil
}
