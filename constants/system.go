package constants

var AccessToken = ""

const (
	CODE_SUCCESS     = "200"
	CODE_SUCCESS_MSG = "Success"

	CODE_CREATED     = "201"
	CODE_CREATED_MSG = "Created"

	CODE_REJECT     = "400"
	CODE_REJECT_MSG = "Bad Request"

	CODE_NOT_FOUND_REJECT     = "404"
	CODE_NOT_FOUND_REJECT_MSG = "Not Found"

	CODE_INTERNAL_SERVER_ERR     = "500"
	CODE_INTERNAL_SERVER_ERR_MSG = "Internal Server Error"

	CODE_SIGNATURE_ERR     = "505"
	CODE_SIGNATURE_ERR_MSG = "Signature Its Not Match !!"

	CODE_PENDING     = "02"
	CODE_PENDING_MSG = "Pending"

	//level log
	LEVEL_LOG_INFO    = 0
	LEVEL_LOG_WARNING = 1
	LEVEL_LOG_ERROR   = 2
	LEVEL_LOG_FATAL   = 3

	// send http
	HttpMethodGet               = "GET"
	HttpMethodPost              = "POST"
	HttpMethodPostWithoutHeader = "POST_WITHOUT_HEADER"
	HttpMethodPut               = "PUT"
	HttpMethodDelete            = "DELETE"

	MambuSendOnUs  = 1
	MambuSendOffUs = 2
	MambuSendFee   = 3

	// Error ID
	ERROR_ID_SUCCESS             = "0000"
	ERROR_ID_FAILED              = "0001"
	ERROR_ID_BAD_REQUEST         = "0400"
	ERROR_ID_UNAUTHORIZED        = "0404"
	ERROR_ID_INIT_JEAGER         = "4001"
	ERROR_ID_GET_RAW_DATA        = "4002"
	ERROR_ID_DECODE_REQ_BODY     = "4003"
	ERROR_ID_VALIDATE_REQ_FAILED = "4004"
	ERROR_ID_DECRYPT_BODY_FAILED = "4005"
	ERROR_ID_INVALID_SIGNATURE   = "4006"
	ERROR_ID_SERVICE_RETURN_400  = "4007"
	ERROR_ID_SERVICE_RETURN_401  = "4008"
	ERROR_ID_INVALID_TOKEN       = "4009"

	// Error Text
	ERROR_TEXT_SUCCESS      = "SUCCESS"
	ERROR_TEXT_FAILED       = "FAILED"
	ERROR_TEXT_BAD_REQUEST  = "BAD REQUEST"
	ERROR_TEXT_UNAUTHORIZED = "UNAUTHORIZED"

	//Secret Key For Signature
	SECRET_KEY_SIGNATURE = "d7a37b7c980f149b5dda8b1de3b7aa07"
)
