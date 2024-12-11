package enum

import "time"

const (
	UserBlockStateNormal  = 0
	UserBlockStateBlocked = 1
)

const (
	AccessTokenDuration            = 2 * time.Hour
	RefreshTokenDuration           = 24 * time.Hour * 10
	OldRefreshTokenHoldingDuration = 6 * time.Hour    // 刷新token时，老的RefreshToken会保留一段时间（用于发现RefreshToken被窃取）
	PasswordTokenDuration          = 15 * time.Minute // 重置密码的验证token有效期
)

const (
	AddressIsNotUserDefault = 0
	AddressIsUserDefault    = 1 // 用户收货地址状态--默认地址
)
