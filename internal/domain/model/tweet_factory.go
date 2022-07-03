package model

type TweetFactory struct {
	lazyRepository func() TweetRepository
}

func (t *TweetFactory) NewTweet(user User, content string) *Tweet {
	charactersLimit := 140
	if user.IsFamous {
		charactersLimit = 280
	}

	return &Tweet{
		repository: t.lazyRepository(),
		Content:    content,
		User:       user,
		validators: []tweetValidator{
			TweetLengthValidator(charactersLimit),
		},
	}
}
