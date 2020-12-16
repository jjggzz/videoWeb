package service

const LoginCustomerAccessKeyPrefix = "LOGIN_CUSTOMER_ACCESS_KEY_"

type RegisterStatus int32

const (
	RegisterStatus_FAIL_SYS_ERR = -1
	RegisterStatus_SUCCESS      = 0
	RegisterStatus_FAIL_HAS_USE = 1
)

type LoginStatus int32

const (
	LoginStatus_FAIL_SYS_ERR = -1
	LoginStatus_SUCCESS      = 0
	LoginStatus_FAIL_NOT_REG = 1
	LoginStatus_FAIL_DISABLE = 2
)
