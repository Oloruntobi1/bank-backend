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

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	gottenUser, err := testQueries.GetUser(context.Background(), user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, gottenUser)

	require.Equal(t, user.Name, gottenUser.Name)
	require.Equal(t, user.Email, gottenUser.Email)
	require.Equal(t, user.Password, gottenUser.Password)
	

}