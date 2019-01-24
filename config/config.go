package config

import (
	"errors"
	"os"
)

var conf *Config

type Config struct {
	ListenOn     string
	WordListPath string
	DB           DBConfig
}

type DBConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
}

func InitConfig() (err error) {
	conf = &Config{
		ListenOn:     os.Getenv("NUMEROLOGYSERVER_HOST"),
		WordListPath: os.Getenv("NUMEROLOGYSERVER_WORDLISTPATH"),
		DB: DBConfig{
			User:     os.Getenv("NUMEROLOGYSERVERDB_USER"),
			Password: os.Getenv("NUMEROLOGYSERVERDB_PASSWORD"),
			DBName:   os.Getenv("NUMEROLOGYSERVERDB_DBNAME"),
			Host:     os.Getenv("NUMEROLOGYSERVERDB_HOST"),
			Port:     os.Getenv("NUMEROLOGYSERVERDB_PORT"),
		},
	}

	// No errors currently defined
	return nil
}

func InitDB() (err error) {
	if conf == nil {
		return errors.New("Trying to open database connection with no config set")
	}

	// TODO postgres DB connection using configured settings
	/* NYI...
	func InitDB(conf config.Config) {
		connectionString := fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			conf.DB.Host,
			conf.DB.Port,
			conf.DB.User,
			conf.DB.DBName,
			conf.DB.Password,
		)

		tmp, err := gorm.Open("postgres", connectionString)
		if err != nil {
			log.Fatalf("Failed to connect to the database with the following connection string: %q. The error was: %q", connectionString, err.Error())
		}
		db = tmp
		Migrate()
	}
	*/

	// All done :-)
	return nil
}

func GetConfig() Config {
	return *conf
}
