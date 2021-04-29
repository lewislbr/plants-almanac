package info

import (
	"users/user"
)

type (
	repository interface {
		GetUserInfo(string) (user.Info, error)
	}

	service struct {
		repo repository
	}
)

func NewService(repo repository) *service {
	return &service{repo}
}

func (s *service) UserInfo(userID string) (user.Info, error) {
	if userID == "" {
		return user.Info{}, user.ErrMissingData
	}

	userInfo, err := s.repo.GetUserInfo(userID)
	if err != nil {
		return user.Info{}, user.ErrNotFound
	}

	return userInfo, nil
}
