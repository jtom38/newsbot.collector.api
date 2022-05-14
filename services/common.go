package services

import "errors"

var (
	ErrMissingAuthorName  = errors.New("unable to find the post author name")
	ErrMessingAuthorImage = errors.New("unable to find the post author image")
	ErrMissingThumbnail   = errors.New("unable to find the post thumbnail url")
	ErrMissingPublishDate = errors.New("unable to find the post publish date")
	ErrMissingTags        = errors.New("unable to find the post tags")
	ErrMissingDescription = errors.New("unable to find the post description")
)

const DATETIME_FORMAT string = "1/2/2006 3:4 PM"
