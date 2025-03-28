package services

import (
	"bank-api/models"
	"bank-api/repositories"
	"bank-api/utils"
	"errors"
)

func Login(username, password string) (map[string]interface{}, error) {
	customer, err := repositories.FindCustomerByUsername(username)

	if err != nil || customer.Password != password {
		return map[string]interface{}{}, errors.New("invalid username or password")
	}

	token, err := utils.GenerateToken(customer.Username)
	if err != nil {
		return map[string]interface{}{}, err
	}

	err = repositories.SaveSession(models.Session{
		Username:  customer.Username,
		Token:     token["token"].(string),
		ExpiredAt: token["expired_at"].(string),
	})
	if err != nil {
		return map[string]interface{}{}, err
	}

	return token, nil
}

func Logout(token string) (string, error) {
	onlyToken := utils.RemoveBearerPrefix(token)

	session, err := repositories.GetSessionByToken(onlyToken)
	if err != nil || session.Token == "" {
		return "", errors.New("invalid token")
	}

	validateResult, err := utils.ValidateToken(token)
	if err != nil {
		return "", err
	}

	err = repositories.RemoveSession(validateResult.Token)
	if err != nil {
		return "", err
	}

	return validateResult.Token, nil
}
