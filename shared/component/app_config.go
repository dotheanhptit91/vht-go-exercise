package sharedcomponent

import (
	"flag"

	sctx "github.com/viettranx/service-context"
)

const (
	AppConfigID                  = "AppConfig"
	DefaultJWTExpIn              = 60 * 60 * 24 * 7
	DefaultCategoryServiceURI    = "http://localhost:3600/v1/rpc/categories"
	DefaultRestaurantServiceURI  = "http://localhost:3600/v1/rpc/restaurants"
	DefaultFoodServiceURI        = "http://localhost:3600/v1/rpc/foods"
)

type AppConfig struct {
	// id                 string
	jwtSecretKey         string
	jwtExpIn             int
	categoryServiceURI   string
	restaurantServiceURI string
	foodServiceURI       string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

func (a *AppConfig) ID() string {
	return AppConfigID
}

func (a *AppConfig) InitFlags() {
	flag.StringVar(&a.jwtSecretKey, "jwt-secret-key", "", "JWT secret key")
	flag.IntVar(&a.jwtExpIn, "jwt-exp-in", DefaultJWTExpIn, "JWT expiration in seconds")
	flag.StringVar(&a.categoryServiceURI, "category-service-uri", DefaultCategoryServiceURI, "Category service URI")
	flag.StringVar(&a.restaurantServiceURI, "restaurant-service-uri", DefaultRestaurantServiceURI, "Restaurant service URI")
	flag.StringVar(&a.foodServiceURI, "food-service-uri", DefaultFoodServiceURI, "Food service URI")
}

func (a *AppConfig) Activate(_ sctx.ServiceContext) error {
	return nil
}

func (a *AppConfig) Stop() error {
	return nil
}

func (a *AppConfig) JwtSecretKey() string {
	return a.jwtSecretKey
}

func (a *AppConfig) JwtExpIn() int {
	return a.jwtExpIn
}

func (a *AppConfig) CategoryServiceURI() string {
	return a.categoryServiceURI
}

func (a *AppConfig) RestaurantServiceURI() string {
	return a.restaurantServiceURI
}

func (a *AppConfig) FoodServiceURI() string {
	return a.foodServiceURI
}

type IAppConfig interface {
	JwtSecretKey() string
	JwtExpIn() int
	CategoryServiceURI() string
	RestaurantServiceURI() string
	FoodServiceURI() string
}