package utils

func HashPassword(password string) (string, error) {
	// TODO: handle hash password
	return "hashed" + password, nil
}

func VerifyPassword(password string, hashPassword string) (bool, error) {
	// TODO: handle verify password
	return true, nil
}
