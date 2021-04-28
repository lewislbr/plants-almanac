package info

import (
	"users/user"
)

type (
	Getter interface {
		GetUserInfo(string) (user.Info, error)
	}

	service struct {
		repo Getter
	}
)

func NewService(repo Getter) *service {
	return &service{repo}
}

func (s *service) UserInfo(id string) (user.Info, error) {
	if id == "" {
		return user.Info{}, user.ErrMissingData
	}

	userInfo, err := s.repo.GetUserInfo(id)
	if err != nil {
		return user.Info{}, user.ErrNotFound
	}

	return userInfo, nil
}
