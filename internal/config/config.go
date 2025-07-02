// package config

// import (
// 	"flag"
// 	"log"
// 	"os"
// 	"github.com/ilyakaznacheev/cleanenv"
// )

// // "go get -u github.com/ilyakaznacheev/cleanenv" -> run this command in your terminal then you will be able to do like an example

// // ---------this is only an example---------
// // type ConfigDatabase struct {
// //     Port     string `yaml:"port" env:"PORT" env-default:"5432"`
// //     Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
// //     Name     string `yaml:"name" env:"NAME" env-default:"postgres"`
// //     User     string `yaml:"user" env:"USER" env-default:"user"`
// //     Password string `yaml:"password" env:"PASSWORD"`
// // }

// type HTTPServer struct {
// 	Addr string
// }

// // env-default:"production"
// type Config struct {
// 	// this is a struct tags -> `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
// 	Env        string `yaml:"env" env:"ENV" env-required:"true"`
// 	StoragePth string `yaml:"storage_path" env-required:"true"`

// 	// Embedding Struct
// 	HTTPServer `yaml:"http_server"`
// }

// func MustLoad() *Config{
// 	var configPath string
// 	configPath = os.Getenv("CONFIG_PATH")

// 	if configPath == ""{
// 		flags := flag.String("config", "", "path to the configuration file")
// 		flag.Parse()

// 		configPath = *flags

// 		if configPath == ""{
// 			log.Fatal("Config path is not set")
// 		}
// 	}

// 	if _, err := os.Stat(configPath); os.IsNotExist(err){
// 		log.Fatal("config file does not exist: %s", configPath)
// 	}

// 	var cfg Config

// 	err := cleanenv.ReadConfig(configPath, &cfg)
// 	if err != nil{
// 		log.Fatal("can not read config file: %s", err.Error())
// 	}

// 	return &cfg
// }

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() *Config{
	//Local .env file
	if err := godotenv.Load();
	err != nil{
		log.Fatal("No .env file found")
	}

	return &Config{
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "Shashank @2004"),
		DBName: getEnv("DB_NAME", "user_data"),

	}
}

	func getEnv(key, defaultValue string)string{
		value := os.Getenv(key)
		if(value == ""){
			return defaultValue
		}
		return value
	}
		

