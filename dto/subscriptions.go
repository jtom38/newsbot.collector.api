package dto

import (
	"context"

	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/domain/models"
)

func (c *DtoClient) ListSubscriptions(ctx context.Context, limit int32) ([]models.SubscriptionDto, error) {
	var res []models.SubscriptionDto

	items, err := c.db.ListSubscriptions(ctx, limit)
	if err != nil {
		return res, err
	}

	for _, item := range items {
		res = append(res, c.ConvertSubscription(item))
	}

	return res, nil
}

func (c *DtoClient) ListSubscriptionDetails(ctx context.Context, limit int32) ([]models.SubscriptionDetailsDto, error) {
	var res []models.SubscriptionDetailsDto

	items, err := c.ListSubscriptions(ctx, limit)
	if err != nil {
		return res, err
	}

	for _, item := range items {
		dwh, err := c.GetDiscordWebhook(ctx, item.DiscordWebhookId)
		if err != nil {
			return res, err
		}

		source, err := c.GetSourceById(ctx, item.SourceId)
		if err != nil {
			return res, err
		}

		res = append(res, models.SubscriptionDetailsDto{
			ID:             item.ID,
			Source:         source,
			DiscordWebHook: dwh,
		})
	}

	return res, nil
}

func (c *DtoClient) ListSubscriptionsByDiscordWebhookId(ctx context.Context, id uuid.UUID) ([]models.SubscriptionDto, error) {
	var res []models.SubscriptionDto

	items, err := c.db.GetSubscriptionsByDiscordWebHookId(ctx, id)
	if err != nil {
		return res, err
	}

	for _, item := range items {
		res = append(res, c.ConvertSubscription(item))
	}

	return res, nil
}

func (c *DtoClient) ListSubscriptionsBySourceId(ctx context.Context, id uuid.UUID) ([]models.SubscriptionDto, error) {
	var res []models.SubscriptionDto

	items, err := c.db.GetSubscriptionsBySourceID(ctx, id)
	if err != nil {
		return res, err
	}

	for _, item := range items {
		res = append(res, c.ConvertSubscription(item))
	}

	return res, nil
}

func (c *DtoClient) ConvertSubscription(i database.Subscription) models.SubscriptionDto {
	return models.SubscriptionDto{
		ID:               i.ID,
		DiscordWebhookId: i.Discordwebhookid,
		SourceId:         i.Sourceid,
	}
}
