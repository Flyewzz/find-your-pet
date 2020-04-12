package normalizer

import (
	"strings"

	"github.com/Kotyarich/find-your-pet/errs"
)

func SexNormalize(sex string) (string, error) {
	sex = strings.ToLower(sex)
	if sex != "f" && sex != "m" && sex != "n/a" {
		if sex == "" {
			// Unknown gender
			sex = "n/a"
		} else {
			return "", errs.IncorrectGender
		}
	}
	return sex, nil
}
