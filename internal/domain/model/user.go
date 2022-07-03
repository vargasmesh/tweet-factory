package model

type User struct {
	ID       int64
	IsFamous bool

	tweetRepository TweetRepository
}

func (u *User) CreateTweet(content string) (Tweet, error) {
	charactersLimit := 140
	if u.IsFamous {
		charactersLimit = 280
	}

	tweet := NewTweet(*u, content, u.tweetRepository)

	err := tweet.Validate(charactersLimit)
	if err != nil {
		return Tweet{}, err
	}

	err = tweet.Save()
	if err != nil {
		return Tweet{}, err
	}

	return *tweet, nil
}
