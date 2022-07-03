package model

import (
	"fmt"
)

type TweetRepository interface {
	Save(Tweet) (int64, error)
}

type Tweet struct {
	ID      int64
	Content string
	User    User

	repository TweetRepository
}

func NewTweet(user User, content string, repository TweetRepository) *Tweet {
	return &Tweet{Content: content, User: user, repository: repository}
}

func (t Tweet) Validate(maxCharacters int) error {
	if len(t.Content) > maxCharacters {
		return fmt.Errorf(
			"max characters exceeded. Received: %d. Max: %d",
			len(t.Content), maxCharacters,
		)
	}

	return nil
}

func (t *Tweet) Save() error {
	id, err := t.repository.Save(*t)
	if err != nil {
		return err
	}

	t.ID = id

	return nil
}
