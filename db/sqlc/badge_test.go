package db

import (
	"context"
	"testing"

	"github.com/gaku101/my-portfolio/util"
	"github.com/stretchr/testify/require"
)

func createRandomBadge(t *testing.T) Badge {
	name := util.RandomString(6)
	badge, err := testQueries.CreateBadge(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, badge)

	require.Equal(t, name, badge.Name)

	require.NotZero(t, badge.ID)
	return badge
}

func TestCreateBadge(t *testing.T) {
	createRandomBadge(t)
}
