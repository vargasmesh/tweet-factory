package model

type User struct {
	ID       int64
	IsFamous bool

	tweetFactory TweetFactory
}

func (u *User) CreateTweet(content string) (Tweet, error) {
	tweet := u.tweetFactory.NewTweet(*u, content)

	err := tweet.Validate()
	if err != nil {
		return Tweet{}, err
	}

	err = tweet.Save()
	if err != nil {
		return Tweet{}, err
	}

	return *tweet, nil
}
