package service

const LoginCustomerAccessKeyPrefix = "LOGIN_CUSTOMER_ACCESS_KEY_"

type RegisterStatus int32

const (
	RegisterStatusFailSysErr = -1
	RegisterStatusSuccess    = 0
	RegisterStatusFailHasUse = 1
)

type LoginStatus int32

const (
	LoginStatusFailSysErr  = -1
	LoginStatusSuccess     = 0
	LoginStatusFailNotReg  = 1
	LoginStatusFailDisable = 2
)

type GetInfoStatus int32

const (
	GetInfoStatusFailSysErr   = -1
	GetInfoStatusSuccess      = 0
	GetInfoStatusFailNotLogin = 1
)
