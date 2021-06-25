package app

type Config struct {
	Db
}

type Db struct {
	Addr     string
	User     string
	Password string
	Database string
}

var Conf Config
