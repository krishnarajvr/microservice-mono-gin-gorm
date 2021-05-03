package service

import (
	"log"
	"regexp"

	common "github.com/krishnarajvr/micro-common"

	"github.com/krishnarajvr/micro-common/locale"
	"micro/module/tenant/model"
	"micro/module/tenant/repo"
	userModel "micro/module/user/model"
	userRepo "micro/module/user/repo"
	"micro/util/password"
)

type ITenantService interface {
	List(page common.Pagination) (model.TenantDtos, error)
	Get(id int) (*model.TenantDto, error)
	Add(tenant *model.TenantForm) (*model.TenantDto, error)
	Register(tenant *model.TenantRegisterForm) (*model.TenantDto, error)
	Update(form *model.TenantForm, id int) (*model.TenantDto, error)
	Patch(form *model.TenantPatchForm, id int) (*model.TenantDto, error)
	Delete(id int) (*model.TenantDto, error)
}

type ServiceConfig struct {
	TenantRepo repo.ITenantRepo
	UserRepo   userRepo.IUserRepo
	Lang       *locale.Locale
}

type Service struct {
	TenantRepo repo.ITenantRepo
	UserRepo   userRepo.IUserRepo
	Lang       *locale.Locale
}

func NewService(c ServiceConfig) ITenantService {
	return &Service{
		TenantRepo: c.TenantRepo,
		UserRepo:   c.UserRepo,
		Lang:       c.Lang,
	}
}

func (s *Service) List(page common.Pagination) (model.TenantDtos, error) {

	tenants, err := s.TenantRepo.List(page)

	if err != nil {
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
		return nil, err
	}

	tenantDto := tenant.ToDto()

	return tenantDto, nil
}

func (s *Service) Add(form *model.TenantForm) (*model.TenantDto, error) {

	tenant, err := s.TenantRepo.Add(form)

	if err != nil {
		return nil, err
	}

	tenantDto := tenant.ToDto()

	return tenantDto, nil
}

//Register a tenant which will create a root user as well
func (s *Service) Register(form *model.TenantRegisterForm) (*model.TenantDto, error) {

	tenantForm := form.ToTenantForm()

	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	code := reg.ReplaceAllString(tenantForm.Name, "")
	tenantForm.Code = code

	tenant, err := s.TenantRepo.Add(tenantForm)
	if err != nil {
		return nil, err
	}

	userForm := userModel.UserForm{
		Name:      form.Name,
		Email:     form.Email,
		FirstName: form.FirstName,
		LastName:  form.LastName,
	}
	userModel, err := userForm.ToModel()
	userModel.TenantId = tenant.ID
	userModel.Password = password.Encrypt(form.Password)

	_, errUser := s.UserRepo.Add(userModel)

	if errUser != nil {
		//Todo - Rollback Added Tenant
		log.Println(errUser)
		return nil, errUser
	}

	tenantDto := tenant.ToDto()

	return tenantDto, nil
}

func (s *Service) Update(form *model.TenantForm, id int) (*model.TenantDto, error) {

	tenant, err := s.TenantRepo.Update(form, id)

	if err != nil {
		return nil, err
	}

	tenantDto := tenant.ToDto()

	return tenantDto, nil
}

func (s *Service) Patch(form *model.TenantPatchForm, id int) (*model.TenantDto, error) {

	tenant, err := s.TenantRepo.Patch(form, id)

	if err != nil {
		return nil, err
	}

	tenantDto := tenant.ToDto()

	return tenantDto, nil
}

func (s *Service) Delete(id int) (*model.TenantDto, error) {

	tenant, err := s.TenantRepo.Delete(id)

	if err != nil {
		return nil, err
	}

	tenantDto := tenant.ToDto()

	return tenantDto, nil
}
