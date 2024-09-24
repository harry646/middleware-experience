package models

type (
	HeaderHostItem struct {
		Path  string
		Value string
	}

	ReqEncrypted struct {
		Data string `json:"data" example:"abcdefghijklmnopqrstuvwxyz" validate:"required"`
	}

	ResData struct {
		ResponseCode    string      `json:"response_code" example:"S0004"`
		ResponseMessage string      `json:"response_message" example:"Success"`
		Meta            ResMetaItem `json:"meta,omitempty"`
		ResponseData    interface{} `json:"response_data,omitempty"`
	}

	ResMetaItem struct {
		DebugParam string `json:"debug_param,omitempty" example:"RegistrationRequest.Oprmode Error:Field validation for Oprmode failed on the oneof tag"`
	}
)
