package dto

import (
	"context"

	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/domain/models"
)

func (c DtoClient) ListDiscordWebhookQueue(ctx context.Context, limit int32) {

}

func (c DtoClient) ListDiscordWebhookQueueDetails(ctx context.Context, limit int32) ([]models.DiscordQueueDetailsDto, error) {
	var res []models.DiscordQueueDetailsDto

	items, err := c.db.ListDiscordQueueItems(ctx, limit)
	if err != nil {
		return res, err
	}

	for _, item := range items {
		article, err := c.GetArticleDetails(ctx, item.ID)
		if err != nil {
			return res, err
		}

		res = append(res, models.DiscordQueueDetailsDto{
			ID:      item.ID,
			Article: article,
		})
	}

	return res, nil
}

func (c DtoClient) ConvertToDiscordQueueDto(i database.Discordqueue) models.DiscordQueueDto {
	return models.DiscordQueueDto{
		ID:        i.ID,
		Articleid: i.Articleid,
	}
}
