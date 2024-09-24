package accounting_portal

import "time"

type (
	// GENERAL
	ResProcess struct {
		Response_Code    string      `json:"response_code"`
		Response_Message string      `json:"response_message"`
		Response_Data    interface{} `json:"response_data"`
	}

	// REQUEST TOKEN
	ReqRequestToken struct {
		ClientCode string `json:"client_code"`
		ClientKey  string `json:"client_key"`
		Timestamp  int64  `json:"timestamp"`
		UID        string `json:"uid"`
		Device     string `json:"device,omitempty"`
		Signature  string `json:"signature"`
		ExtraInfo1 string `json:"extra_info1,omitempty"`
		ExtraInfo2 string `json:"extra_info2,omitempty"`
		NIK        string `json:"nik,omitempty"`
		Mobile     string `json:"mobile,omitempty"`
		Email      string `json:"email,omitempty"`
	}
	ResRequestToken struct {
		AccessToken string    `json:"access_token"`
		Expired     time.Time `json:"expired"`
	}

	// VALIDATE PAGE
	ReqSaveLogState struct {
		Token       string `json:"access_token"`
		ProjectCode string `json:"project_code"`
		State       int32  `json:"state"`
		Timestamp   int64  `json:"timestamp"`
	}
	ResValidatePage struct {
		IsValid bool `json:"is_valid"`
	}
)
