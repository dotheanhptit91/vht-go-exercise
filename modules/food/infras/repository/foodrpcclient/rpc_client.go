package foodrpcclient

type CategoryRPCClient struct {
	categoryServiceURI string
}

func NewCategoryRPCClient(categoryServiceURI string) *CategoryRPCClient {
	return &CategoryRPCClient{categoryServiceURI: categoryServiceURI}
}

type RestaurantRPCClient struct {
	restaurantServiceURI string
}

func NewRestaurantRPCClient(restaurantServiceURI string) *RestaurantRPCClient {
	return &RestaurantRPCClient{restaurantServiceURI: restaurantServiceURI}
}

