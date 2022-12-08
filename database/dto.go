package database

import (
	"strings"

	"github.com/google/uuid"
)

type SourceDto struct {
	ID      uuid.UUID `json:"id"`
	Site    string    `json:"site"`
	Name    string    `json:"name"`
	Source  string    `json:"source"`
	Type    string    `json:"type"`
	Value   string    `json:"value"`
	Enabled bool      `json:"enabled"`
	Url     string    `json:"url"`
	Tags    []string  `json:"tags"`
	Deleted bool      `json:"deleted"`
}

func ConvertToSourceDto(i Source) SourceDto {
	var deleted bool
	if !i.Deleted.Valid {
		deleted = true
	}

	return SourceDto{
		ID:      i.ID,
		Site:    i.Site,
		Name:    i.Name,
		Source:  i.Source,
		Type:    i.Type,
		Value:   i.Value.String,
		Enabled: i.Enabled,
		Url:     i.Url,
		Tags:    splitTags(i.Tags),
		Deleted: deleted,
	}
}

func splitTags(t string) []string {
	items := strings.Split(t, ", ")
	return items
}
