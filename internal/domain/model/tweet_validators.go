package model

import (
	"errors"
	"fmt"
)

type tweetValidator func(Tweet) error

var ErrCharactersExceeded = errors.New("max characters exceeded")

func TweetLengthValidator(maxCharacters int) tweetValidator {
	return func(t Tweet) error {
		if len(t.Content) > maxCharacters {
			return fmt.Errorf(
				"%w: got: %d max: %d",
				ErrCharactersExceeded, len(t.Content), maxCharacters,
			)
		}

		return nil
	}
}
