package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

// JWT token string = header string . message string. signature string
// signature string = sha256(header+message,secret)
func GenerateJWTToken(header string, payload map[string]string, secret string) (string, error) {
	//Create a hash using a crypto algorithm and a secret key
	h := hmac.New(sha256.New, []byte(secret))

	//Encode the header with base64
	EncodedHeader := base64.StdEncoding.EncodeToString([]byte(header))

	//Convert payload map to json string and then encode it to base64
	payloadString, err := json.Marshal(payload)
	if err != nil {
		return "", errors.New("payload marshal error")
	}
	EncodedPayload := base64.StdEncoding.EncodeToString(payloadString)

	//Now add Encoded header and encoded payload
	message := EncodedHeader + "." + EncodedPayload

	unsignedstr := header + string(payloadString)
	//write this unsignedstr to SHA256 algorithm to hash it
	h.Write([]byte(unsignedstr))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	jwt_token := message + "." + signature

	return jwt_token, nil
}

func ValidateToken(jwt_token string, secret string) (bool, error) {
	token_splits := strings.Split(jwt_token, ".")

	if len(token_splits) != 3 {
		return false, errors.New("invalid token length")
	}

	header, err := base64.StdEncoding.DecodeString(token_splits[0])
	if err != nil {
		return false, errors.New("invalid header in token")
	}

	payload, err := base64.StdEncoding.DecodeString(token_splits[1])
	if err != nil {
		return false, errors.New("invalid payload in token")
	}

	// Now create a signature string using the decoded header and payload
	h := hmac.New(sha256.New, []byte(secret))
	unsignedstr := string(header) + string(payload)
	h.Write([]byte(unsignedstr))
	generated_signedstr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	//compare generated_signedstr with the given signature string
	if generated_signedstr == token_splits[2] {
		return true, nil
	}
	return false, errors.New("invalid token signature")
}
