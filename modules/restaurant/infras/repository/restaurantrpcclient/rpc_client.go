package restaurantrpcclient

type CategoryRPCClient struct {
	categoryServiceURI string
}

func NewCategoryRPCClient(categoryServiceURI string) *CategoryRPCClient {
	return &CategoryRPCClient{categoryServiceURI: categoryServiceURI}
}