package services

import "errors"

var (
	ErrMissingAuthorName  = errors.New("unable to find the author name")
	ErrMessingAuthorImage = errors.New("unable to find the author image")
	ErrMissingThumbnail   = errors.New("unable to find the thumbnail url")
	ErrMissingPublishDate = errors.New("unable to find the publish date")
	ErrMissingTags        = errors.New("unable to find the tags")
)

const DATETIME_FORMAT string = "1/2/2006 3:4 PM"
