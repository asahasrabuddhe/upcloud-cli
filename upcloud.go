package main

const API_VERSION = "1.2.8"
const API_URL = "https://api.cloud.com/%s/%s"

type Error struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
