package db

import (
	"context"
	"testing"

	"github.com/Oloruntobi1/Oloruntobi1/bank_backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {

	arg := CreateUserParams {

		Name : util.RandomOwner(),
		Email: util.RandomEmail(),
		Password: util.RandomPassword(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user

}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}