package application

import (
	"tweet-ddd-factory/internal/domain/model"
)

type EventPublisher interface {
	PublishTweet(model.User, model.Tweet)
}

type Logger interface {
	Debug(string)
	Error(string, error)
}

type UserService struct {
	logger         Logger
	eventPublisher EventPublisher
}

func (u *UserService) CreateTweet(user model.User, content string) error {
	tweet, err := user.CreateTweet(content)
	if err != nil {
		u.logger.Error("could not create tweet", err)
		return err
	}

	u.eventPublisher.PublishTweet(user, tweet)

	return nil
}
