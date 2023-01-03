package repository

import (
	"testing"

	"example.com/mod/domain"

	"github.com/stretchr/testify/require"
)

func TestSyncMapRepository(t *testing.T) {
	dbRepo := NewSyncMapTodoRepository()

	id := 1
	testData := domain.Todo{
		ID:        id,
		Text:      "test",
		Completed: false,
	}

	t.Run("AllGet Todo Test", func(t *testing.T) {
		r1, _ := dbRepo.AllGet()
		require.Empty(t, r1)
	})

	t.Run("Store Todo Test", func(t *testing.T) {
		dbRepo.Store(testData)
		r2, _ := dbRepo.AllGet()
		require.Equal(t, r2[0], testData)
	})

	t.Run("Status Update Test", func(t *testing.T) {
		dbRepo.StatusUpdate(id)
		r3, _ := dbRepo.AllGet()
		require.Equal(t, r3[0].Completed, true)
	})

	t.Run("Delete Todo Test", func(t *testing.T) {
		dbRepo.Delete(id)
		r4, _ := dbRepo.AllGet()
		require.Empty(t, r4)
	})
}
