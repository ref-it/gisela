package common

import (
	"encoding/json"
	"os"
	"strconv"
)

// GetConfiguration reads the configuration file and provides the data as Configuration struct
func GetConfiguration() Configuration {
	configFile, _ := os.ReadFile("config/config.json")
	config := Configuration{}
	err := json.Unmarshal([]byte(configFile), &config)
	if err != nil {
		panic(err)
	}
	return config
}

// GetWebserverConfiguration extracts the webserver configuration
func GetWebserverConfiguration() WebserverConfiguration {
	config := GetConfiguration()
	return config.Web
}

// GetWebserverConfiguration extracts the webserver configuration
func GetDatabaseConfiguration() DatabaseConfiguration {
	config := GetConfiguration()
	return config.Database
}

func GetSmtpConfiguration() SmtpConfiguration {
	config := GetConfiguration()
	return config.Email.SMTP
}

func GetDatabaseConnectionString() string {
	dbConfig := GetDatabaseConfiguration()
	var connectionString string
	if dbConfig.Password == "" {
		connectionString = "postgres://" + dbConfig.Name + "?host=" + dbConfig.Host
	} else {
		connectionString = "postgres://" + dbConfig.User + ":" + dbConfig.Password + "@" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + "/" + dbConfig.Name
	}
	return connectionString
}
