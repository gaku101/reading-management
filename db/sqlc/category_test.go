package db

import (
	"context"
	"testing"

	"github.com/gaku101/my-portfolio/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	name := util.RandomString(6)
	category, err := testQueries.CreateCategory(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, name, category.Name)

	require.NotZero(t, category.ID)
	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomBadge(t)
}

func GetCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
}

func TestListCategories(t *testing.T) {
	categories, err := testQueries.ListCategories(context.Background())
	require.NoError(t, err)
	for _, category := range categories {
		require.NotEmpty(t, category)
	}
}
