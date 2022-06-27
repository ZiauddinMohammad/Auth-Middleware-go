package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func ParseToken(r *http.Request) (string, error) {
	var token string
	data, err := ioutil.ReadAll(r.Body)
	token = string(data)
	//err := json.NewDecoder(body.Body).Decode(&token)
	if err != nil {
		log.Println("im here")
		return "", errors.New("error while parsing token")
	}
	return token, nil
}

func ParsePayloadFromToken(token string) (map[string]string, error) {
	payload_byte, _ := base64.StdEncoding.DecodeString(strings.Split(token, ".")[1])
	var payload_map map[string]string
	err := json.Unmarshal(payload_byte, &payload_map)
	if err != nil {
		return nil, errors.New("error while parsing payload token")
	}
	return payload_map, nil
}
