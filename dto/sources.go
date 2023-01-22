package dto

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/domain/models"
)

func (c DtoClient) ListSources(ctx context.Context, limit int32) ([]models.SourceDto, error) {
	var res []models.SourceDto

	items, err := c.db.ListSources(ctx, limit)
	if err != nil {
		return res, err
	}

	for _, item := range items {
		res = append(res, c.ConvertToSource(item))
	}

	return res, nil
}

func (c DtoClient) ListSourcesBySource(ctx context.Context, sourceName string) ([]models.SourceDto, error) {
	var res []models.SourceDto

	items, err := c.db.ListSourcesBySource(ctx, strings.ToLower(sourceName))
	if err != nil {
		return res, err
	}

	for _, item := range items {
		res = append(res, c.ConvertToSource(item))
	}

	return res, nil
}

func (c DtoClient) GetSourceById(ctx context.Context, id uuid.UUID) (models.SourceDto, error) {
	var res models.SourceDto

	item, err := c.db.GetSourceByID(ctx, id)
	if err != nil {
		return res, err
	}

	return c.ConvertToSource(item), nil
}

func (c DtoClient) GetSourceByNameAndSource(ctx context.Context, name, source string) (models.SourceDto, error) {
	var res models.SourceDto

	item, err := c.db.GetSourceByNameAndSource(ctx, database.GetSourceByNameAndSourceParams{
		Name: name,
		Source: source,
	})
	if err != nil {
		return res, err
	}

	return c.ConvertToSource(item), nil
}

func (c DtoClient) ConvertToSource(i database.Source) models.SourceDto {
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
