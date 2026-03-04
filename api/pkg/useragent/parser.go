package useragent

import "strings"

type UserAgent struct {
	Browser        string
	BrowserVersion string
	OS             string
	OSVersion      string
	DeviceType     string
}

func Parse(ua string) UserAgent {
	result := UserAgent{
		DeviceType: "desktop", // Default
	}

	if ua == "" {
		return result
	}

	// Detect OS
	if strings.Contains(ua, "Windows") {
		result.OS = "Windows"
	} else if strings.Contains(ua, "Macintosh") || strings.Contains(ua, "Mac OS X") {
		result.OS = "macOS"
	} else if strings.Contains(ua, "Android") {
		result.OS = "Android"
		result.DeviceType = "mobile"
	} else if strings.Contains(ua, "iPhone") || strings.Contains(ua, "iPad") || strings.Contains(ua, "iPod") {
		result.OS = "iOS"
		result.DeviceType = "mobile"
		if strings.Contains(ua, "iPad") {
			result.DeviceType = "tablet"
		}
	} else if strings.Contains(ua, "Linux") {
		result.OS = "Linux"
	}

	// Detect Browser
	if strings.Contains(ua, "Firefox") {
		result.Browser = "Firefox"
	} else if strings.Contains(ua, "Chrome") { // Chrome usually also contains Safari, so check first
		result.Browser = "Chrome"
	} else if strings.Contains(ua, "Safari") {
		result.Browser = "Safari"
	} else if strings.Contains(ua, "Edge") {
		result.Browser = "Edge"
	}

	return result
}
