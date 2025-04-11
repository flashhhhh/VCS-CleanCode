package main

import "fmt"

func main() {
	// Service Factory
	// userService := GetServicesFactory("user")
	orderService := GetServicesFactory("order")

	// Database Factory
	// userServiceDatabase := userService.makeDatabase()
	// orderServiceDatabase := orderService.makeDatabase()

	// Cache Factory
	// userServiceCache := userService.makeCache()
	orderServiceCache := orderService.makeCache()

	// userServiceDatbase.SetUpConnection()
	// resp, err := userServiceDatabase.ExecuteQuery("SELECT * FROM users")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(resp)

	// orderServiceDatabase.SetUpConnection()
	// resp, err := orderServiceDatabase.ExecuteQuery("SELECT * FROM orders")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(resp)


	orderServiceCache.SetUpConnection()
	if orderServiceCache == nil {
		panic("Cache is nil")
	}

	resp, err := orderServiceCache.SetQuery("name", "John Doe")
	if err != nil {
		panic(err)
	}

	resp, err = orderServiceCache.GetQuery("name")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}