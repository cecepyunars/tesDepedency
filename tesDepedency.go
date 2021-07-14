package tesDepedency

func PrintTesting (gender string) string{
	if gender == "" {
		return ""
	}

	if gender == "male" {
		gender = "laki-laki"
	} else if gender == "female" {
		gender = "perempuan"
	}

	return ""
}