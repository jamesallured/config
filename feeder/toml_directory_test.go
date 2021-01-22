package feeder_test

import (
	"testing"

	"github.com/golobby/config/feeder"
	"github.com/stretchr/testify/assert"
)

func Test_TomlDirectory_Feed_Not_Existing_Dir_It_Should_Return_Error(t *testing.T) {
	toml := feeder.TomlDirectory{Path: "/404"}

	_, err := toml.Feed()

	assert.Error(t, err)
}

func Test_TomlDirectory_Feed_Invalid_File_It_Should_Return_Error(t *testing.T) {
	toml := feeder.TomlDirectory{Path: "test/invalid-toml"}

	_, err := toml.Feed()

	assert.Error(t, err)
}

func Test_TomlDirectory_Feed_Sample(t *testing.T) {
	toml := feeder.TomlDirectory{Path: "test/toml"}

	m, err := toml.Feed()

	assert.NoError(t, err)
	assert.Equal(t, m["app"].(map[string]interface{})["name"], "MyAppUsingConfig")
	assert.Equal(t, m["app"].(map[string]interface{})["version"], 3.14)
}
