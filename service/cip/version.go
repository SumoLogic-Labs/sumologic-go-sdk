package cip

// Number
const number = "Dev"

func UserAgent() string {
	return "sumologic-go-sdk/" + Version()
}

func Version() string {
	return number
}
