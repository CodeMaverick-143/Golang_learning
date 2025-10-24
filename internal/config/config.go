package config
import (
	"flag"
	"log"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPserver struct{
	Address string `yaml:"address" env-required:"true"`

}

// env-default:"production"

type Config struct{
	Env string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HTTPServer HTTPserver `yaml:"http_server" env:"HTTP_SERVER" env-required:"true"`
}


func MustLoad() *Config{

	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == ""{
		flags := flag.String("config", "config/local.yaml", "config file path")
		flag.Parse()
		configPath = *flags

		if configPath == ""{
			log.Fatal("config path not provided")
		}
	}

	if _,err := os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("config file does not exist: %s", configPath)
	}


	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil{
		log.Fatalf("error reading config file: %s", err.Error())
	}

	return &cfg

 
}