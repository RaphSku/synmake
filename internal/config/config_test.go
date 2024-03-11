package config

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

func isEmptyStruct(s interface{}) bool {
	typ := reflect.TypeOf(s)
	if typ.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		zeroValue := reflect.Zero(field.Type)
		fieldValue := reflect.ValueOf(s).Field(i)
		if !reflect.DeepEqual(fieldValue.Interface(), zeroValue.Interface()) {
			return false
		}
	}

	return true
}

func TestConfigManagerInit(t *testing.T) {
	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	expFilePath := "./config.yaml"
	configManager := NewConfigManager(logger, expFilePath)
	assert.Equal(t, expFilePath, configManager.filePath)
	assert.Equal(t, true, isEmptyStruct(configManager.config))
}

func TestDefaultConfigAsYAML(t *testing.T) {
	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	configString := GetConfigAsYamlString(logger)
	var config Config
	err = yaml.Unmarshal([]byte(configString), &config)
	assert.NoError(t, err)
	assert.Equal(t, false, isEmptyStruct(config))
}
