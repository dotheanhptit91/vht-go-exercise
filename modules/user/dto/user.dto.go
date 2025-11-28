package userdto

import (
	"errors"
	"strings"
	userdomain "vht-go/modules/user/domain"
)

type RegisterUserDTO struct {
	LoginUserDTO
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
}

func (dto *RegisterUserDTO) Validate() error {
	if err := dto.LoginUserDTO.Validate(); err != nil {
		return err
	}

	dto.LastName = strings.TrimSpace(dto.LastName)
	dto.FirstName = strings.TrimSpace(dto.FirstName)

	if dto.FirstName == "" {
		return errors.New(userdomain.ErrUserFirstNameRequired)
	}

	if dto.LastName == "" {
		return errors.New(userdomain.ErrUserLastNameRequired)
	}

	return nil
}

type LoginUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *LoginUserDTO) Validate() error {
	dto.Email = strings.TrimSpace(dto.Email)
	dto.Password = strings.TrimSpace(dto.Password)

	if dto.Email == "" {
		return errors.New(userdomain.ErrUserEmailRequired)
	}

	if dto.Password == "" {
		return errors.New(userdomain.ErrUserPasswordRequired)
	}

	return nil
}

type LoginResponseDTO struct {
	Token string `json:"token"`
}