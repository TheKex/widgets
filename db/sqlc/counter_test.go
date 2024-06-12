package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"widgets/util"
)

func createRandomCounter(t *testing.T) Counter {
	widget := createRandomWidget(t)

	arg := CreateCounterParams{
		WidgetID:  widget.ID,
		Order:     sql.NullInt32{Int32: int32(util.RandomInt(0, 100)), Valid: true},
		Value:     sql.NullInt32{Int32: int32(util.RandomInt(0, 100)), Valid: true},
		Increment: sql.NullInt32{Int32: int32(util.RandomInt(0, 100)), Valid: true},
	}

	counter, err := testQueries.CreateCounter(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, counter)

	require.Equal(t, arg.WidgetID, counter.WidgetID)
	require.Equal(t, arg.Order, counter.Order)
	require.Equal(t, arg.Value, counter.Value)
	require.Equal(t, arg.Increment, counter.Increment)

	require.NotZero(t, counter.ID)

	return counter
}

func TestCreateCounter(t *testing.T) {
	createRandomCounter(t)
}

func TestGetCounter(t *testing.T) {
	counter1 := createRandomCounter(t)

	counter2, err := testQueries.GetCounter(context.Background(), counter1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, counter2)

	require.Equal(t, counter1.ID, counter2.ID)
	require.Equal(t, counter1.WidgetID, counter2.WidgetID)
	require.Equal(t, counter1.Order, counter2.Order)
	require.Equal(t, counter1.Value, counter2.Value)
	require.Equal(t, counter1.Increment, counter2.Increment)
}

func TestUpdateCounter(t *testing.T) {
	counter1 := createRandomCounter(t)

	arg := UpdateCounterParams{
		ID:        counter1.ID,
		Order:     sql.NullInt32{Int32: int32(util.RandomInt(0, 10)), Valid: true},
		Value:     sql.NullInt32{Int32: int32(util.RandomInt(0, 10)), Valid: true},
		Increment: sql.NullInt32{Int32: int32(util.RandomInt(0, 10)), Valid: true},
	}

	counter2, err := testQueries.UpdateCounter(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, counter2)

	require.Equal(t, counter1.ID, counter2.ID)
	require.Equal(t, counter1.WidgetID, counter2.WidgetID)
	require.Equal(t, arg.Order, counter2.Order)
	require.Equal(t, arg.Value, counter2.Value)
	require.Equal(t, arg.Increment, counter2.Increment)
}

func TestDeleteCounter(t *testing.T) {
	counter1 := createRandomCounter(t)

	err := testQueries.DeleteCounter(context.Background(), counter1.ID)
	require.NoError(t, err)

	counter2, err := testQueries.GetCounter(context.Background(), counter1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, counter2)
}

func TestListCounters(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomCounter(t)
	}

	arg := ListCountersParams{
		Limit:  5,
		Offset: 5,
	}

	counters, err := testQueries.ListCounters(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, counters, 5)

	for _, counter := range counters {
		require.NotEmpty(t, counter)
	}
}

func TestUpdateCounterValue(t *testing.T) {
	counter1 := createRandomCounter(t)

	arg := UpdateCounterValueParams{
		ID:    counter1.ID,
		Value: sql.NullInt32{Int32: int32(util.RandomInt(-10, 10)), Valid: true},
	}

	counter2, err := testQueries.UpdateCounterValue(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, counter2)

	require.Equal(t, counter1.ID, counter2.ID)
	require.Equal(t, arg.Value, counter2.Value)
}
