package models

import "testing"

func NewReview(stars int, comment string) *CreateReviewCMD {
	return &CreateReviewCMD{
		Stars:   stars,
		Comment: comment,
	}
}

func Test_createReviewValidate(t *testing.T) {
	r := NewReview(4, "The quick brown fox jumps over the lazy dog")

	err := r.validate()
	if err != nil {
		t.Error("the validation did not pass")
		t.Fail()

	}
}
