package e

var MsgFlags = map[int]string{
	SUCCESS:               "SUCCESS",
	ERROR_BAD_REQUEST:     "ERROR_BAD_REQUEST",
	ERROR_AUTH:            "ERROR_AUTH",
	ERROR_PERMISSION:      "ERROR_PERMISSION",
	ERROR_NOT_FOUND:       "ERROR_NOT_FOUND",
	ERROR_INTERNAL_SERVER: "ERROR_INTERNAL_SERVER",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	return MsgFlags[code]
}
