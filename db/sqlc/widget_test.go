package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"widgets/util"
)

func createRandomWidget(t *testing.T) Widget {
	user := createRandomUser(t)

	arg := CreateWidgetParams{
		Name:   util.RandomString(10),
		UserID: user.ID,
	}

	widget, err := testQueries.CreateWidget(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, widget)

	require.NotZero(t, widget.ID)
	require.NotZero(t, widget.CreatedAt)

	return widget
}

func TestCreateWidget(t *testing.T) {
	createRandomWidget(t)
}

func TestGetWidget(t *testing.T) {
	widget1 := createRandomWidget(t)
	widget2, err := testQueries.GetWidget(context.Background(), widget1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, widget2)

	require.Equal(t, widget1.ID, widget2.ID)
	require.Equal(t, widget1.Name, widget2.Name)
	require.Equal(t, widget1.UserID, widget2.UserID)
	require.WithinDuration(t, widget1.CreatedAt, widget2.CreatedAt, time.Second)
}

func TestUpdateWidget(t *testing.T) {
	widget1 := createRandomWidget(t)

	arg := UpdateWidgetParams{
		ID:   widget1.ID,
		Name: util.RandomString(10),
	}

	widget2, err := testQueries.UpdateWidget(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, widget2)

	require.Equal(t, widget1.ID, widget2.ID)
	require.Equal(t, arg.Name, widget2.Name)
	require.Equal(t, widget1.UserID, widget2.UserID)
	require.WithinDuration(t, widget1.CreatedAt, widget2.CreatedAt, time.Second)
}

func TestDeleteWidget(t *testing.T) {
	widget1 := createRandomWidget(t)

	err := testQueries.DeleteUser(context.Background(), widget1.ID)
	require.NoError(t, err)

	widget2, err := testQueries.GetUser(context.Background(), widget1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, widget2)
}

func TestListWidget(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomWidget(t)
	}

	arg := ListWidgetsParams{
		Limit:  5,
		Offset: 5,
	}

	widgets, err := testQueries.ListWidgets(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, widgets, 5)

	for _, widget := range widgets {
		require.NotEmpty(t, widget)
	}
}
