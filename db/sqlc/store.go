package db

import (
	"context"
	"database/sql"
	"fmt"
)

//Store provides all functions to execute db queries and transaction
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	DeletePostTx(ctx context.Context, arg DeletePostTxParams) error
	DeleteUserTx(ctx context.Context, arg DeleteUserTxParams) error
	LoginPointTx(ctx context.Context, arg LoginPointTxParams) (LoginPointTxResult, error)
}

//SQLStore provides all functions to execute SQL queries and transaction
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromUserID int64 `json:"from_user_id"`
	ToUserID   int64 `json:"to_user_id"`
	Amount     int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer  Transfer        `json:"transfer"`
	FromUser  UpdatePointsRow `json:"from_user"`
	ToUser    UpdatePointsRow `json:"to_user"`
	FromEntry Entry           `json:"from_entry"`
	ToEntry   Entry           `json:"to_entry"`
}

// TransferTx performs a points transfer from one user to the other.
// It creates the transfer, add entries, and update users' points within a database transaction
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromUserID: arg.FromUserID,
			ToUserID:   arg.ToUserID,
			Amount:     arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			UserID: arg.FromUserID,
			Amount: -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			UserID: arg.ToUserID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromUserID < arg.ToUserID {
			result.FromUser, result.ToUser, err = addPoints(ctx, q, arg.FromUserID, -arg.Amount, arg.ToUserID, arg.Amount)
		} else {
			result.ToUser, result.FromUser, err = addPoints(ctx, q, arg.ToUserID, arg.Amount, arg.FromUserID, -arg.Amount)
		}

		return err
	})

	return result, err
}

func addPoints(
	ctx context.Context,
	q *Queries,
	userID1 int64,
	amount1 int64,
	userID2 int64,
	amount2 int64,
) (user1 UpdatePointsRow, user2 UpdatePointsRow, err error) {
	user1, err = q.UpdatePoints(ctx, UpdatePointsParams{
		ID:     userID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	user2, err = q.UpdatePoints(ctx, UpdatePointsParams{
		ID:     userID2,
		Amount: amount2,
	})
	return
}

type LoginPointTxParams struct {
	UserID int64 `json:"user_id"`
	Amount int64 `json:"amount"`
}

type LoginPointTxResult struct {
	User  UpdatePointsRow `json:"to_user"`
	Entry Entry           `json:"to_entry"`
}

func (store *SQLStore) LoginPointTx(ctx context.Context, arg LoginPointTxParams) (LoginPointTxResult, error) {
	var result LoginPointTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Entry, err = q.CreateEntry(ctx, CreateEntryParams{
			UserID: arg.UserID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}
		result.User, err = q.UpdatePoints(ctx, UpdatePointsParams{
			ID:     arg.UserID,
			Amount: arg.Amount,
		})

		return err
	})

	return result, err
}

type DeletePostTxParams struct {
	ID int64 `json:"id"`
}

func (store *SQLStore) DeletePostTx(ctx context.Context, arg DeletePostTxParams) error {

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		err = q.DeleteComments(ctx, arg.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("post_id = %v on Comments not set", arg.ID)
			} else {
				return err
			}
		}
		err = q.DeletePostFavorites(ctx, arg.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("post_id = %v on PostFavorites not set", arg.ID)
			} else {
				return err
			}
		}
		err = q.DeletePostCategory(ctx, arg.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("post_id = %v on PostCategory not set", arg.ID)
			} else {
				return err
			}
		}
		err = q.DeleteNotes(ctx, arg.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("post_id = %v on PostCategory not set", arg.ID)
			} else {
				return err
			}
		}
		err = q.DeletePost(ctx, arg.ID)
		if err != nil {
			return err
		}

		return err
	})

	return err
}

type DeleteUserTxParams struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

func (store *SQLStore) DeleteUserTx(ctx context.Context, arg DeleteUserTxParams) error {

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		err = q.DeleteFollows(ctx, arg.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("id = %v on follow not set", arg.ID)
			} else {
				return err
			}
		}
		posts, err := q.ListMyAllPosts(ctx, arg.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("Username = %v on posts not set", arg.Username)
			} else {
				return err
			}
		}
		for i := range posts {
			post := posts[i]
			err = q.DeleteComments(ctx, post.ID)
			if err != nil {
				if err == sql.ErrNoRows {
					fmt.Printf("post_id = %v on Comments not set", post.ID)
				} else {
					return err
				}
			}
			err = q.DeletePostFavorites(ctx, post.ID)
			if err != nil {
				if err == sql.ErrNoRows {
					fmt.Printf("post_id = %v on PostFavorites not set", post.ID)
				} else {
					return err
				}
			}
			err = q.DeletePostCategory(ctx, post.ID)
			if err != nil {
				if err == sql.ErrNoRows {
					fmt.Printf("post_id = %v on PostCategory not set", post.ID)
				} else {
					return err
				}
			}
			err = q.DeleteNotes(ctx, post.ID)
			if err != nil {
				if err == sql.ErrNoRows {
					fmt.Printf("post_id = %v on PostCategory not set", arg.ID)
				} else {
					return err
				}
			}
			err = q.DeletePost(ctx, post.ID)
			if err != nil {
				if err == sql.ErrNoRows {
					fmt.Printf("post_id = %v on Posts not set", post.ID)
				} else {
					return err
				}
			}
		}
		err = q.DeleteMyFavoritePosts(ctx, arg.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("id = %v on post_favorites not set", arg.ID)
			} else {
				return err
			}
		}
		err = q.DeleteUser(ctx, arg.Username)
		if err != nil {
			return err
		}

		return err
	})

	return err
}
