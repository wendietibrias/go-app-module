package app_module

func SayName(name string, gender string) string {
	var formatedName string

	if gender == "male" || gender == "Male" {
		formatedName = "Mr." + name
	} else if gender == "female" || gender == "Female" {
		formatedName = "Mrs." + name
	}

	return "Hello " + formatedName
}
