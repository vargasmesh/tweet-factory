package application

import "fmt"

type User struct {
	ID       int64
	IsFamous bool
}

type TweetTableGateway interface {
	CreateTweet(user User, text string) error
}

type UserService struct {
	tweetTableGateway TweetTableGateway
}

func (u *UserService) CreateTweet(user User, tweet string) error {
	charactersLimit := 140
	if user.IsFamous {
		charactersLimit = 280
	}

	if len(tweet) > charactersLimit {
		return fmt.Errorf(
			"max characters exceeded. Received: %d. Max: %d",
			len(tweet), charactersLimit,
		)
	}

	return u.tweetTableGateway.CreateTweet(user, tweet)
}
