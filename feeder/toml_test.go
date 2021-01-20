package feeder_test

import (
	"testing"

	"github.com/golobby/config/feeder"
	"github.com/stretchr/testify/assert"
)

func Test_Toml_Feed_Not_Existing_File_It_Should_Return_Error(t *testing.T) {
	toml := feeder.Toml{Path: "/404.toml"}

	_, err := toml.Feed()

	assert.Error(t, err)
}

func Test_Toml_Feed_Invalid_File_It_Should_Return_Error(t *testing.T) {
	toml := feeder.Toml{Path: "test/invalid.toml"}

	_, err := toml.Feed()

	assert.Error(t, err)
}

func Test_Toml_Feed_Sample(t *testing.T) {
	toml := feeder.Toml{Path: "test/config.toml"}

	m, err := toml.Feed()

	assert.NoError(t, err)
	assert.Equal(t, "MyAppUsingConfig", m["name"])
	assert.Equal(t, 3.14, m["version"])
}
