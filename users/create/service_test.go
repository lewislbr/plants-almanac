package create

import (
	"testing"
	u "users/user"

	"github.com/stretchr/testify/require"
)

type mockRepo struct {
	Users []u.User
}

func (m mockRepo) InsertOne(new u.User) (interface{}, error) {
	m.Users = append(m.Users, new)

	return new, nil
}

func (m mockRepo) FindOne(email string) (u.User, error) {
	for _, u := range m.Users {
		if email == u.Email {
			return u, nil
		}
	}

	return u.User{}, u.ErrNotFound
}

func TestCreate(t *testing.T) {
	t.Run("should error when there are missing fields", func(t *testing.T) {
		t.Parallel()

		repo := mockRepo{
			Users: []u.User{},
		}
		createService := NewCreateService(repo)
		newUser := u.User{
			Email:    "test@test.com",
			Password: "1234",
		}
		err := createService.Create(newUser)

		require.EqualError(t, err, u.ErrMissingData.Error())
	})

	t.Run("should error when the user already exists", func(t *testing.T) {
		t.Parallel()

		repo := mockRepo{
			Users: []u.User{
				{
					Name:     "test",
					Email:    "test@test.com",
					Password: "1234",
				},
			},
		}
		createService := NewCreateService(repo)
		newUser := u.User{
			Name:     "test",
			Email:    "test@test.com",
			Password: "1234",
		}
		err := createService.Create(newUser)

		require.EqualError(t, err, u.ErrUserExists.Error())
	})

	t.Run("should create a user with no error", func(t *testing.T) {
		t.Parallel()

		repo := mockRepo{
			Users: []u.User{},
		}
		createService := NewCreateService(repo)
		newUser := u.User{
			Name:     "test",
			Email:    "test@test.com",
			Password: "1234",
		}
		err := createService.Create(newUser)

		require.NoError(t, err)
	})
}
