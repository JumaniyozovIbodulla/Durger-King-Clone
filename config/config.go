package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	BotToken         string
	Admins           uint64
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	ServiceName      string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("no .env file: ", err)
	}

	config := Config{}

	config.BotToken = cast.ToString(getORReturnDefaultValue("BOT_TOKEN", ""))
	config.Admins = cast.ToUint64(getORReturnDefaultValue("ADMINS", ""))
	config.PostgresHost = cast.ToString(getORReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getORReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getORReturnDefaultValue("POSTGRES_USER", ""))
	config.PostgresPassword = cast.ToString(getORReturnDefaultValue("POSTGRES_PASSWORD", ""))
	config.PostgresDatabase = cast.ToString(getORReturnDefaultValue("POSTGRES_DB", ""))
	config.ServiceName = cast.ToString(getORReturnDefaultValue("SERVICE_NAME", ""))
	return config
}

func getORReturnDefaultValue(key string, defaultValue interface{}) interface{} {

	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}
	return defaultValue
}