package goviper

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestJSONFile(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	// read cfg based file type..
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "go-viper-cfg-loader-json", config.GetString("app.name"))
	assert.Equal(t, "Azmi Farhan", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))

	jsonCfg := fmt.Sprintf("App Name \t: %s\nVersion \t: %s\nFile Type \t: %s", config.GetString("app.name"), config.GetString("app.version"), config.GetString("general.file_type"))
	fmt.Println(jsonCfg)
}

func TestYAMLFile(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.yaml")
	// config.SetConfigName("config")
	// config.SetConfigType("yaml")
	config.AddConfigPath(".")

	// read cfg based file type..
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "go-viper-cfg-loader-yaml", config.GetString("app.name"))
	assert.Equal(t, "Azmi Farhan", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))

	yamlCfg := fmt.Sprintf("App Name \t: %s\nVersion \t: %s\nFile Type \t: %s", config.GetString("app.name"), config.GetString("app.version"), config.GetString("general.this_file_type"))
	fmt.Println(yamlCfg)
}

func TestENVFile(t *testing.T) {

	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	// read config based file type
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "go-viper-cfg-loader-env", config.GetString("APP_NAME"))
	assert.Equal(t, "Azmi Farhan", config.GetString("APP_AUTHOR"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))

	envCfg := fmt.Sprintf("App Name \t: %s\nVersion \t: %s\nFile Type \t: %s", config.GetString("APP_NAME"), config.GetString("APP_VERSION"), config.GetString("APP_FILE_TYPE"))
	fmt.Println(envCfg)
}

func TestENV(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	// read config based file type
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "go-viper-cfg-loader-env", config.GetString("APP_NAME"))
	assert.Equal(t, "Azmi Farhan", config.GetString("APP_AUTHOR"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))

	assert.Equal(t, "Hello", config.GetString("FROM_ENV"))
}
