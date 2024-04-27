package validators

import "regexp"

func IsEmailValid(email string) bool {
	rxEmail := regexp.MustCompile(
		`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`,
	)

	return len(email) > 3 || len(email) < 254 || rxEmail.MatchString(email)
}