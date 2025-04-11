package main

type IServicesFactory interface {
	makeDatabase() IDatabase
	makeCache() ICache
	// makeObjectStorage() IObjectStorage
}

func GetServicesFactory(serviceType string) IServicesFactory {
	switch serviceType {
		case "user":
			return &UserService{}
		case "order":
			return &OrderService{}
		// case "product":
		// 	return &ProductService{}
		default:
			return nil
	}
}