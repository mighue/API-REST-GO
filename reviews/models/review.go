package models

import (
	"errors"
	"time"
)

const maxLengthInCommets = 400

//Review represent an anon review from some website
type Review struct {
	Id      int64
	Stars   int
	Comment string
	Date    time.Time //created at
}

type CreateReviewCMD struct {
	Stars   int    `json:"stars"`
	Comment string `json:"comment"`
}

func (cmd *CreateReviewCMD) validate() error {
	if cmd.Stars < 1 || cmd.Stars > 5 {
		return errors.New("stars must be between 1 and 5")
	}

	if len(cmd.Comment) > maxLengthInCommets {
		return errors.New("comment must be less than 400 characters long")
	}
	return nil
}
