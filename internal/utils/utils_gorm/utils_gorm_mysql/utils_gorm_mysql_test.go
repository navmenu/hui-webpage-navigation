package utils_gorm_mysql

import (
	"context"
	"hui-webpage-navigation/internal/utils/utils_gorm"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	m.Run()
	os.Exit(0)
}

func TestExample(t *testing.T) {
	var err error
	require.NoError(t, err)

	var b = true
	if !assert.True(t, b) {
		return
	}
}

func TestNewGormDB(t *testing.T) {
	db, err := NewGormDB("root:123456@tcp(127.0.0.1:3306)/golang_demo_20230515?charset=utf8mb4&parseTime=True&loc=Local", true)
	require.NoError(t, err)
	err = utils_gorm.Ping(context.Background(), db)
	require.NoError(t, err)
	err = utils_gorm.Close(db)
	require.NoError(t, err)
}
