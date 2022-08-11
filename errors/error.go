package errors

import "errors"

var (
	ERR_ADDRESS_IS_EMPTY      = errors.New("address list is empty")
	ERR_PROXY_PATH_IS_INVALID = errors.New("proxy path is invalid")
	ERR_NUM_OF_WORKER_LIMITED = errors.New("num of worker is limited to 3 worker")
	ERR_ADDRESS_LIST_LIMITED  = errors.New("address list limited to 3 address")
)
