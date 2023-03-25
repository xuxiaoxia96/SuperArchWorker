package conf

import (
	"os"
)

var Cfg Config

type Config struct {
	Basic struct {

	}
	RabbitMQ struct {
		Host string
		Port int64
		Username string
		Password string
		Exchange string
	}
	Redis struct {
		Host string
		Port int64
		Password string
		DB int64
	}
	PostgresDB struct {
		Host string
		Port int64
		Username string
		Password string
		DBNAME string
		SCHEMA string
	}

	WorkerConf map[string]map[string]string
}

func InitConfig()  {
	if len(os.Getenv("DEBUG")) > 0{
		// rabbitMQ
		Cfg.RabbitMQ.Host = "localhost"
		Cfg.RabbitMQ.Port = 5672
		Cfg.RabbitMQ.Username = "admin"
		Cfg.RabbitMQ.Password = "admin"
		Cfg.RabbitMQ.Exchange = "super_ex"

		// Redis
		Cfg.Redis.Host = "localhost"
		Cfg.Redis.Port = 6379
		Cfg.Redis.Password = "testadmin123"
		Cfg.Redis.DB = 0

		// Postgres
		Cfg.PostgresDB.Host = "localhost"
		Cfg.PostgresDB.Port = 5432
		Cfg.PostgresDB.Username = "admin"
		Cfg.PostgresDB.Password = "admin"
		Cfg.PostgresDB.DBNAME = "admin"
		Cfg.PostgresDB.SCHEMA = "superarch"

		// Worker Config
		imageScoreConfig := map[string]string{}
		imageScoreConfig["api_url"] = "http://127.0.0.1:5001"

		txt2ImageConfig := map[string]string{}
		txt2ImageConfig["api_url"] = "http://82.156.47.200:8888"

		image2ImageConfig := map[string]string{}
		image2ImageConfig["api_url"] = "http://82.156.47.200:8888"

		image2txtConfig := map[string]string{}
		image2txtConfig["api_url"] = "http://127.0.0.1:5001"

		imagevqaConfig := map[string]string{}
		imagevqaConfig["api_url"] = "http://127.0.0.1:5001"

		workerConf := map[string]map[string]string{}
		workerConf["ai.painter.imagescore"] = imageScoreConfig
		workerConf["ai.painter.txt2image"] = txt2ImageConfig
		workerConf["ai.painter.image2image"] = image2ImageConfig
		workerConf["ai.painter.image2txt"] = image2txtConfig
		workerConf["ai.painter.imagevqa"] = imagevqaConfig

		Cfg.WorkerConf = workerConf
	}else{
		// rabbitMQ
		Cfg.RabbitMQ.Host = "localhost"
		Cfg.RabbitMQ.Port = 5672
		Cfg.RabbitMQ.Username = "admin"
		Cfg.RabbitMQ.Password = "testadmin123ashore"
		Cfg.RabbitMQ.Exchange = "super_ex"

		// Redis
		Cfg.Redis.Host = "localhost"
		Cfg.Redis.Port = 6379
		Cfg.Redis.Password = "testadmin123ashore"
		Cfg.Redis.DB = 0

		// Postgres
		Cfg.PostgresDB.Host = "localhost"
		Cfg.PostgresDB.Port = 5432
		Cfg.PostgresDB.Username = "admin"
		Cfg.PostgresDB.Password = "testadmin123ashore"
		Cfg.PostgresDB.DBNAME = "admin"
		Cfg.PostgresDB.SCHEMA = "superarch"

		// Worker Config
		imageScoreConfig := map[string]string{}
		imageScoreConfig["api_url"] = "http://82.156.47.200:8810"

		txt2ImageConfig := map[string]string{}
		txt2ImageConfig["api_url"] = "http://82.156.47.200:8808"

		image2ImageConfig := map[string]string{}
		image2ImageConfig["api_url"] = "http://82.156.47.200:8808"

		image2txtConfig := map[string]string{}
		image2txtConfig["api_url"] = "http://82.156.47.200:8809"

		imagevqaConfig := map[string]string{}
		imagevqaConfig["api_url"] = "http://82.156.47.200:8809"

		workerConf := map[string]map[string]string{}
		workerConf["ai.painter.imagescore"] = imageScoreConfig
		workerConf["ai.painter.txt2image"] = txt2ImageConfig
		workerConf["ai.painter.image2image"] = image2ImageConfig
		workerConf["ai.painter.image2txt"] = image2txtConfig
		workerConf["ai.painter.imagevqa"] = imagevqaConfig

		Cfg.WorkerConf = workerConf
	}

}