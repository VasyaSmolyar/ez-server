package bcrypt

import "golang.org/x/crypto/bcrypt"

func Generate(str string) (string, error) {
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword([]byte(str), bcrypt.MinCost)
	return string(hashedPasswordBytes), err
}

func Check(str, hashedStr string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedStr), []byte(str))
	return err == nil
}
