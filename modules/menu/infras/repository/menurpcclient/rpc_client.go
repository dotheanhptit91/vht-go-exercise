package menurpcclient

type FoodRPCClient struct {
	foodServiceURI string
}

func NewFoodRPCClient(foodServiceURI string) *FoodRPCClient {
	return &FoodRPCClient{foodServiceURI: foodServiceURI}
}

type RestaurantRPCClient struct {
	restaurantServiceURI string
}

func NewRestaurantRPCClient(restaurantServiceURI string) *RestaurantRPCClient {
	return &RestaurantRPCClient{restaurantServiceURI: restaurantServiceURI}
}

