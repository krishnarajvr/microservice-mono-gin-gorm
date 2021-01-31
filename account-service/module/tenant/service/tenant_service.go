package service

import (
	"log"

	common "github.com/krishnarajvr/micro-common"

	"micro/app/locale"
	"micro/module/tenant/model"
	"micro/module/tenant/repo"
)

type ITenantService interface {
	List(page common.Pagination) (model.TenantDtos, error)
	Get(id int) (*model.TenantDto, error)
	Add(tenant *model.TenantForm) (*model.TenantDto, error)
	Update(form *model.TenantForm, id int) (*model.TenantDto, error)
	Patch(form *model.TenantPatchForm, id int) (*model.TenantDto, error)
	Delete(id int) (*model.TenantDto, error)
}

type ServiceConfig struct {
	TenantRepo repo.ITenantRepo
	Lang       *locale.Locale
}

type Service struct {
	TenantRepo repo.ITenantRepo
	Lang       *locale.Locale
}

func NewService(c ServiceConfig) ITenantService {
	return &Service{
		TenantRepo: c.TenantRepo,
		Lang:       c.Lang,
	}
}

func (s *Service) List(page common.Pagination) (model.TenantDtos, error) {

	tenants, err := s.TenantRepo.List(page)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(s.Lang.Get("hi", "GetAll"))

	m := map[string]interface{}{
		"Name": "John Doe",
		"Age":  66,
	}

	s.Lang.SetLang("en-US")
	log.Println(s.Lang.Get("intro", m))

	tenantsDto := tenants.ToDto()

	return tenantsDto, err
}

func (s *Service) Get(id int) (*model.TenantDto, error) {

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

func (s *Service) Add(form *model.TenantForm) (*model.TenantDto, error) {

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

func (s *Service) Update(form *model.TenantForm, id int) (*model.TenantDto, error) {

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

func (s *Service) Patch(form *model.TenantPatchForm, id int) (*model.TenantDto, error) {

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

func (s *Service) Delete(id int) (*model.TenantDto, error) {

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
