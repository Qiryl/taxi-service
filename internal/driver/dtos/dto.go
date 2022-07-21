package dtos

import "github.com/Qiryl/taxi-service/internal/driver/domain"

type DriverDTO struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	TaxiType string `json:"taxi-type"`
}

func (dto *DriverDTO) ToModel() *domain.Driver {
	return &domain.Driver{
		Name:     dto.Name,
		Phone:    dto.Phone,
		Email:    dto.Email,
		Password: dto.Password,
		TaxiType: dto.TaxiType,
	}
}

func toDriverDto(model *domain.Driver) *DriverDTO {
	return &DriverDTO{
		Name:     model.Name,
		Phone:    model.Phone,
		Email:    model.Email,
		Password: model.Password,
		TaxiType: model.TaxiType,
	}
}

type LoginDTO struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (dto *LoginDTO) ToModel() *domain.Login {
	return &domain.Login{
		Phone:    dto.Phone,
		Password: dto.Password,
	}
}

func toLoginDto(model *domain.Login) *LoginDTO {
	return &LoginDTO{
		Phone:    model.Phone,
		Password: model.Password,
	}
}
