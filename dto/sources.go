package dto

import (
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/domain/models"
)

func (c DtoClient) ConvertToSourceDto(i database.Source) models.SourceDto {
	var deleted bool
	if !i.Deleted.Valid {
		deleted = true
	}

	return models.SourceDto{
		ID:      i.ID,
		Site:    i.Site,
		Name:    i.Name,
		Source:  i.Source,
		Type:    i.Type,
		Value:   i.Value.String,
		Enabled: i.Enabled,
		Url:     i.Url,
		Tags:    c.SplitTags(i.Tags),
		Deleted: deleted,
	}
}
