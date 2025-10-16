package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertUpdateDelete(t *testing.T) {
	// prepare
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)

	defer func() {
		_ = db.Close()
	}()

	newClient := Client{
		FIO:      "TEST",
		Login:    "TEST",
		Birthday: "TEST",
		Email:    "TEST",
	}

	// insert
	id, err := insertClient(db, newClient)

	require.NoError(t, err)
	require.NotEmpty(t, id)
	newClient.ID = int(id)

	got, err := selectClient(db, id)
	require.NoError(t, err)
	require.Equal(t, newClient, got)

	// update
	newLogin := "TEST_NEW"
	err = updateClientLogin(db, newLogin, id)
	require.NoError(t, err)

	got, err = selectClient(db, id)
	require.NoError(t, err)
	require.Equal(t, newLogin, got.Login)

	// delete
	err = deleteClient(db, id)
	require.NoError(t, err)

	got, err = selectClient(db, id)
	require.Equal(t, sql.ErrNoRows, err)
	require.Empty(t, got)
}
