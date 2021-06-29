package cip

import (
	"github.com/wizedkyle/sumologic-go-sdk/v2"
)

func UserAgent() string {
	return "sumologic-go-sdk/" + Version()
}

func Version() string {
	return version.Number
}