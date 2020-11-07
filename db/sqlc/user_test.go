package db

import (
	"context"
	"database/sql"
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

func TestListUsers(t *testing.T) {

	for i := 0; i < 10 ; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams {
		Limit: 5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)

	require.NoError(t, err)

	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

}

func TestDeleteUser(t *testing.T) {

	user := createRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user.ID)

	require.NoError(t, err)

	notGottenUser, err := testQueries.GetUser(context.Background(), user.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, notGottenUser)
}