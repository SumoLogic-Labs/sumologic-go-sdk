package cip

// Number
const number = "Dev"

func UserAgent() string {
	return "sumologic-go-sdk/" + version()
}

func version() string {
	return number
}
