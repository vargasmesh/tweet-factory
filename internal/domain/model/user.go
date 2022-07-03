package model

import "fmt"

type User struct {
	ID       int64
	IsFamous bool

	tweetRepository TweetRepository
}

type TweetRepository interface {
	Create(User, string) (Tweet, error)
}

type Tweet struct {
	ID      int64
	Content string
}

func (u *User) CreateTweet(tweet string) (Tweet, error) {
	charactersLimit := 140
	if u.IsFamous {
		charactersLimit = 280
	}

	if len(tweet) > charactersLimit {
		return Tweet{}, fmt.Errorf(
			"max characters exceeded. Received: %d. Max: %d",
			len(tweet), charactersLimit,
		)
	}

	return u.tweetRepository.Create(*u, tweet)
}
