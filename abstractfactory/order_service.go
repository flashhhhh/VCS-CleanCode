package main

type OrderService struct {
	
}

func (o *OrderService) makeDatabase() IDatabase {
	return &Database{
		host: "localhost",
		port: 5432,
		user: "postgres",
		password: "12345678",
		database: "order_db",
	}
}

func (o *OrderService) makeCache() ICache {
	return &Cache{
		host: "localhost",
		port: 6379,
	}
}