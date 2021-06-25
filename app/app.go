package app

import (
	"errors"
	"my-test-app/handlers"
	"my-test-app/service/datastore"
	"my-test-app/service/validators"
	"my-test-app/usecases"

	"github.com/go-playground/validator"

	"github.com/spf13/viper"

	"github.com/go-pg/pg/v10"

	"github.com/labstack/echo/v4"
)

const Name = "my-test-app"

var (
	db *pg.DB
)

func InitApp(config string) {
	loadConfig(config)
	opt := &pg.Options{
		Addr:     Conf.Addr,
		User:     Conf.User,
		Password: Conf.Password,
		Database: Conf.Database,
	}
	db = pg.Connect(opt)
}

func loadConfig(config string) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(config)
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}

func Start() {

	server := echo.New()
	server.Validator = validators.NewCustomValidator(validator.New())
	rep := datastore.NewUserRepository(db)

	getUsers := usecases.NewGetUsersUsecase(rep)
	getUser := usecases.NewGetUserUsecase(rep)
	postUser := usecases.NewPostUserUsecase(rep)
	putUser := usecases.NewPutUserUsecase(rep)
	deleteUser := usecases.NewDeleteUserUsecase(rep)

	server.GET("/users", handlers.GetUsers(getUsers))
	server.GET("/users/:id", handlers.GetUser(getUser))
	server.POST("/users", handlers.PostUser(postUser))
	server.PUT("/users/:id", handlers.PutUser(putUser))
	server.DELETE("/users/:id", handlers.DeleteUser(deleteUser))

	server.Logger.Fatal(server.Start(":8080"))
}

func DB() (*pg.DB, error) {
	if db != nil {
		return db, nil
	}
	return nil, errors.New("no db")
}
