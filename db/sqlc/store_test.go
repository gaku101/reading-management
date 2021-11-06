package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	fmt.Println(">> before:", user1.Points, user2.Points)

	n := 3
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromUserID: user1.ID,
				ToUserID:   user2.ID,
				Amount:     amount,
			})
			errs <- err
			results <- result
		}()
	}

	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, user1.ID, transfer.FromUserID)
		require.Equal(t, user2.ID, transfer.ToUserID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, user1.ID, fromEntry.UserID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, user2.ID, toEntry.UserID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check user
		fromUser := result.FromUser
		require.NotEmpty(t, fromUser)
		require.Equal(t, user1.ID, fromUser.ID)

		toUser := result.ToUser
		require.NotEmpty(t, toUser)
		require.Equal(t, user2.ID, toUser.ID)

		// check user's points
		fmt.Println(">> tx:", fromUser.Points, toUser.Points)

		diff1 := user1.Points - fromUser.Points
		diff2 := toUser.Points - user2.Points
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0) // 1 * amount, 2 * amount, 3 * amount, ..., n * amount

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}
	// check the final updated user's point
	updatedUser1, err := store.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)

	updatedAccount2, err := store.GetUser(context.Background(), user2.Username)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedUser1.Points, updatedAccount2.Points)

	require.Equal(t, user1.Points-int64(n)*amount, updatedUser1.Points)
	require.Equal(t, user2.Points+int64(n)*amount, updatedAccount2.Points)
}
