package repositories

import (
	"bank-api/config"
	"bank-api/models"
	"bank-api/utils"
	"encoding/json"
	"os"
)

func SaveSession(session models.Session) error {
	sessions, err := GetSessions()
	if err != nil {
		return err
	}

	sessions = append(sessions, session)

	data, err := json.MarshalIndent(sessions, "", "	") // format sessions data to json and applies indent
	if err != nil {
		return err
	}

	return os.WriteFile(config.SessionDB, data, 0644) // write to session file
}

func GetSessions() ([]models.Session, error) {
	data, err := os.ReadFile(config.SessionDB) // read from session file
	if err != nil {
		return nil, err
	}

	var sessions []models.Session
	err = json.Unmarshal(data, &sessions) // parse json data to session struct
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func GetSessionByToken(token string) (models.Session, error) {
	sessions, err := GetSessions()
	if err != nil {
		return models.Session{}, err
	}

	for _, session := range sessions {
		if session.Token == token {
			return session, nil
		}
	}

	return models.Session{}, nil
}

func RemoveSession(token string) error {
	token = utils.RemoveBearerPrefix(token)
	sessions, err := GetSessions()
	if err != nil {
		return err
	}

	var newSessions []models.Session
	for _, session := range sessions {
		if session.Token != token {
			newSessions = append(newSessions, session)
		}
	}

	data, err := json.MarshalIndent(newSessions, "", "	")
	if err != nil {
		return err
	}

	return os.WriteFile(config.SessionDB, data, 0644)
}
