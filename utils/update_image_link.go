package utils

import "strings"

func GetCloudwalkLink(s3Link string) string {
	return strings.Replace(s3Link, "https://ressy2.s3.amazonaws.com", "http://images.ressyapp.com", 1)
}
