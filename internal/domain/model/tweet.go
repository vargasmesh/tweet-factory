package model

type TweetRepository interface {
	Save(Tweet) (int64, error)
}

type Tweet struct {
	ID      int64
	Content string
	User    User

	repository TweetRepository
	validators []tweetValidator
}

func (t Tweet) Validate() error {
	for _, validator := range t.validators {
		if err := validator(t); err != nil {
			return err
		}
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
