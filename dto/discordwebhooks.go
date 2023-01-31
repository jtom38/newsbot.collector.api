package dto

import (
	"context"

	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/domain/models"
)

func (c *DtoClient) ListDiscordWebHooks(ctx context.Context, total int32) ([]models.DiscordWebHooksDto, error) {
	var res []models.DiscordWebHooksDto

	items, err := c.db.ListDiscordWebhooks(ctx, total)
	if err != nil {
		return res, nil
	}

	for _, item := range items {
		res = append(res, c.ConvertDiscordWebhook(item))
	}

	return res, nil
}

func (c *DtoClient) GetDiscordWebhook(ctx context.Context, id uuid.UUID) (models.DiscordWebHooksDto, error) {
	var res models.DiscordWebHooksDto

	item, err := c.db.GetDiscordWebHooksByID(ctx, id)
	if err != nil {
		return res, err
	}

	return c.ConvertDiscordWebhook(item), nil
}

func (c *DtoClient) GetDiscordWebHookByServerAndChannel(ctx context.Context, server, channel string) ([]models.DiscordWebHooksDto, error) {
	var res []models.DiscordWebHooksDto

	items, err := c.db.GetDiscordWebHooksByServerAndChannel(ctx, database.GetDiscordWebHooksByServerAndChannelParams{
		Server:  server,
		Channel: channel,
	})
	if err != nil {
		return res, err
	}

	for _, item := range items {
		res = append(res, c.ConvertDiscordWebhook(item))
	}

	return res, nil
}

func (c *DtoClient) ConvertDiscordWebhook(i database.Discordwebhook) models.DiscordWebHooksDto {
	return models.DiscordWebHooksDto{
		ID:      i.ID,
		Url:     i.Url,
		Server:  i.Server,
		Channel: i.Channel,
		Enabled: i.Enabled,
	}
}
