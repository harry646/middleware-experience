package helpers

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"middleware-experience/constants"
	"middleware-experience/models"
	"middleware-experience/utils"
	"strings"
	"unicode"
)

func GenerateSignatureLogin(request_body models.ReqLoginAccount, header_timestamps string, header_httpMethod string) (res_signature string) {
	secretKey := constants.SECRET_KEY_SIGNATURE

	reqEndpoint := "/login"

	minifiedBody, errMini := minifyFunc(request_body)
	if errMini != nil {
		utils.LogData("Failed to Minify", "Failed to Minify", constants.LEVEL_LOG_ERROR, errMini.Error())
	}

	stringReq := header_httpMethod + strings.ToLower(reqEndpoint) + minifiedBody + request_body.Username + header_timestamps

	signature := generateSignature(secretKey, stringReq)

	return signature
}

func minifyFunc(body interface{}) (string, error) {
	// Convert the struct to JSON
	jsonData, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return "", fmt.Errorf("error marshaling to JSON: %v", err)
	}

	// Convert the JSON bytes to a string (similar to JSON.stringify in JavaScript)
	jsonString := string(jsonData)
	return jsonString, nil
}

// RemoveSpacesAndNewlines removes tabs, spaces, and newlines from a string
func RemoveSpacesAndNewlines(input string) string {
	var builder strings.Builder
	for _, r := range input {
		if !unicode.IsSpace(r) || (r != '\t' && r != '\n' && r != '\r' && r != ' ') {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// Generate the HMAC signature using SHA-512
func generateSignature(secretKey string, stringToSign string) string {
	h := hmac.New(sha512.New, []byte(secretKey))
	h.Write([]byte(stringToSign))

	signature := h.Sum(nil)

	verify := hex.EncodeToString(signature)
	return verify
}
