package service

import (
	"log"
	"micro/app/locale"
	"micro/model"
	"micro/repo"

	common "github.com/krishnarajvr/micro-common"
)

type ITenantService interface {
	List(page common.Pagination) (model.TenantDtos, error)
	Get(id int) (*model.TenantDto, error)
	Add(tenant *model.TenantForm) (*model.TenantDto, error)
	Update(form *model.TenantForm, id int) (*model.TenantDto, error)
	Patch(form *model.TenantPatchForm, id int) (*model.TenantDto, error)
	Delete(id int) (*model.TenantDto, error)
}

type TenantService struct {
	TenantRepo repo.ITenantRepo
	Lang       *locale.Locale
}

func NewTenantService(c *ServiceConfig) ITenantService {
	return &TenantService{
		TenantRepo: c.TenantRepo,
		Lang:       c.Lang,
	}
}

func (s *TenantService) List(page common.Pagination) (model.TenantDtos, error) {

	tenants, err := s.TenantRepo.List(page)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(s.Lang.Get("hi", "GetAll"))

	tenantsDto := tenants.ToDto()

	return tenantsDto, err
}

func (s *TenantService) Get(id int) (*model.TenantDto, error) {

	tenant, err := s.TenantRepo.Get(id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(tenant)

	tenantDto := tenant.ToDto()

	log.Println(tenantDto)

	return tenantDto, nil
}

func (s *TenantService) Add(form *model.TenantForm) (*model.TenantDto, error) {

	tenant, err := s.TenantRepo.Add(form)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(tenant)

	tenantDto := tenant.ToDto()

	log.Println(tenantDto)

	return tenantDto, nil
}

func (s *TenantService) Update(form *model.TenantForm, id int) (*model.TenantDto, error) {

	tenant, err := s.TenantRepo.Update(form, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(tenant)

	tenantDto := tenant.ToDto()

	log.Println(tenantDto)

	return tenantDto, nil
}

func (s *TenantService) Patch(form *model.TenantPatchForm, id int) (*model.TenantDto, error) {

	tenant, err := s.TenantRepo.Patch(form, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(tenant)

	tenantDto := tenant.ToDto()

	log.Println(tenantDto)

	return tenantDto, nil
}

func (s *TenantService) Delete(id int) (*model.TenantDto, error) {

	tenant, err := s.TenantRepo.Delete(id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(tenant)

	tenantDto := tenant.ToDto()

	log.Println(tenantDto)

	return tenantDto, nil
}
