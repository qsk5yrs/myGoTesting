package enum

// Constants of redis key
// Redis Key的格式为:
//   项目名:模块名:键名

const (
	REDIS_KEY_DEMO_ORDER_DETAIL = "GOMALL:DEMO:ORDER_DETAIL_%s"
)

const (
	REDIS_KEY_ACCESS_TOKEN        = "GOMALL:USER:ACCESS_TOKEN_%s"
	REDIS_KEY_REFRESH_TOKEN       = "GOMALL:USER:REFRESH_TOKEN_%s"
	REDIS_KEY_USER_SESSION        = "GOMALL:USER:SESSION_%d"
	REDIS_KEY_TOKEN_REFRESH_LOCK  = "GOMALL:USER:TOKEN_REFRESH_LOCK_%s"
	REDIS_KEY_PASSWORDRESET_TOKEN = "GOMALL:USER:PASSWORD_RESET_TOKEN_%s"
)
