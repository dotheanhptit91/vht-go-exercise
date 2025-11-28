package shared

const (
	KeyRequester = "requester"
)

const (
	RoleUser    = "user"
	RoleAdmin   = "admin"
	RoleShipper = "shipper"

	KeyGormComp        = "gorm"
	KeyLocalPubSubComp = "localPubSub"
	KeyNatsPubSubComp  = "natsPubSub"
	KeyRedisComp       = "redis"
	KeyGrpcServerComp  = "grpcServer"

	EvtRestaurantLiked   = "RestaurantLiked"
	EvtRestaurantUnliked = "RestaurantUnliked"
)