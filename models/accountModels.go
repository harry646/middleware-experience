package models

type (
	ResHttpGeneral struct {
		Code       int
		ResGeneral ResGeneral
	}
	ResGeneral struct {
		RC    string      `json:"rc"`
		MSG   string      `json:"msg"`
		Data  interface{} `json:"data,omitempty"`
		Error interface{} `json:"errors,omitempty"`
	}

	ReqLoginAccount struct {
		Username string `example:"john.doe@bankina.co.id" json:"username" validate:"required"`
		Password string `example:"CLIENT" json:"password" validate:"required"`
	}

	ReqForgotPassword struct {
		Username string `example:"johndoe@bankina.co.id" json:"username" validate:"required"`
	}

	ReqUpdatePassword struct {
		Username    string `example:"33f8d3e76464f185a50745a67d35c86cb794bac07df69c902" json:"username" validate:"required"`
		NewPassword string `example:"33f8d3e76464f185a50745a67d35c86cb794bac07df69c902" json:"new_password" validate:"required"`
	}

	ReqCaptchaToken struct {
		Token string `example:"33f8d3e76464f185a50745a67d35c86cb794bac07df69c902" json:"token" validate:"required"`
	}

	ReqToken struct {
		ClientCode string `json:"client_code" example:"abcdefghijklmnopqrstuvwxyz" validate:"required"`
		ClientKey  string `json:"client_key" example:"abcdefghijklmnopqrstuvwxyz" validate:"required"`
		Timestamp  int64  `json:"timestamp" example:"1234567890" validate:"required"`
		UID        string `json:"uid"`
		Device     string `json:"device,omitempty"`
		Signature  string `json:"signature,omitempty"`
		ExtraInfo1 string `json:"extra_info1,omitempty"`
		ExtraInfo2 string `json:"extra_info2,omitempty"`
		NIK        string `json:"nik,omitempty"`
		Mobile     string `json:"mobile,omitempty"`
		Email      string `json:"email,omitempty"`
	}

	ReqValidatePage struct {
		Token       string `json:"access_token" example:"abcdefghijklmnopqrstuvwxyz" validate:"required"`
		ProjectCode string `json:"project_code" example:"ONBRD" validate:"required"`
		State       int32  `json:"state" example:"1"`
		Timestamp   int64  `json:"timestamp" example:"1234567890" validate:"required"`
	}

	ReqInitUserState struct {
		Token     string `json:"access_token" example:"abcdefghijklmnopqrstuvwxyz" validate:"required"`
		UserId    int64  `json:"user_id" example:"1" validate:"required"`
		Timestamp int64  `json:"timestamp" example:"1234567890" validate:"required"`
	}

	ReqGetUserState struct {
		Token       string `json:"access_token" example:"abcdefghijklmnopqrstuvwxyz" validate:"required"`
		UserId      int64  `json:"user_id" example:"1" validate:"required"`
		ProjectCode string `json:"project_code" example:"ABC" validate:"required"`
		Timestamp   int64  `json:"timestamp" example:"1234567890" validate:"required"`
	}

	AccessToken struct {
		Token string `json:"access_token" validate:"required"`
	}

	ReqFirstUpdatePassword struct {
		Username    string `example:"33f8d3e76464f185a50745a67d35c86cb794bac07df69c902" json:"username" validate:"required"`
		OldPassword string `example:"33f8d3e76464f185a50745a67d35c86cb794bac07df69c902" json:"old_password" validate:"required"`
		NewPassword string `example:"33f8d3e76464f185a50745a67d35c86cb794bac07df69c902" json:"new_password" validate:"required"`
	}

	ProfileByUsername struct {
		Profile_Username string `example:"john.doe@bankina.co.id" json:"username" validate:"required"`
	}
)
