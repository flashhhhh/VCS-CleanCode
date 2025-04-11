package main

type UserService struct {

}

func (u *UserService) makeDatabase() IDatabase {
	return &Database{
		host: "localhost",
		port: 5432,
		user: "postgres",
		password: "12345678",
		database: "user_db",
	}
}

func (u *UserService) makeCache() ICache {
	return &Cache{
		host: "localhost",
		port: 6379,
	}
}