// The converter package lives between the database calls and the API calls.
// This way if any new methods like RPC calls are added later, the API does not need to be reworked as much
package dto

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/domain/models"
)

type DtoClient struct {
	db *database.Queries
}

func NewDtoClient(db *database.Queries) *DtoClient {
	return &DtoClient{
		db: db,
	}
}

func (c *DtoClient) ListArticles(ctx context.Context, limit, page int) ([]models.ArticleDto, error) {
	var res []models.ArticleDto

	a, err := c.db.ListArticles(ctx, database.ListArticlesParams{
		Limit: int32(limit),
		Offset: int32(limit * page),
	})
	if err != nil {
		return res, err
	}

	for _, article := range a {
		res = append(res, c.convertArticle(article))
	}
	return res, nil
}

func (c *DtoClient) ListArticlesByPage(ctx context.Context, page, limit int32 ) ([]models.ArticleDto, error) {
	var res []models.ArticleDto

	a, err := c.db.ListArticlesByPage(ctx, database.ListArticlesByPageParams{
		Limit: limit,
		Offset: page * limit,
	})
	if err != nil {
		return res, err
	}

	for _, article := range a {
		res = append(res, c.convertArticle(article))
	}

	return res, nil
}

func (c *DtoClient) GetArticle(ctx context.Context, ID uuid.UUID) (models.ArticleDto, error) {
	a, err := c.db.GetArticleByID(ctx, ID)
	if err != nil {
		return models.ArticleDto{}, err
	}

	return c.convertArticle(a), nil
}

func (c *DtoClient) GetArticleDetails(ctx context.Context, ID uuid.UUID) (models.ArticleDetailsDto, error) {
	a, err := c.db.GetArticleByID(ctx, ID)
	if err != nil {
		return models.ArticleDetailsDto{}, err
	}

	s, err := c.db.GetSourceByID(ctx, a.Sourceid)
	if err != nil {
		return models.ArticleDetailsDto{}, err
	}

	res := c.convertArticleDetails(a, s)

	return res, nil
}

func (c *DtoClient) ListNewArticlesBySourceId(ctx context.Context, SourceID uuid.UUID, limit, page int) ([]models.ArticleDto, error) {
	var res []models.ArticleDto
	a, err := c.db.ListNewArticlesBySourceId(ctx, database.ListNewArticlesBySourceIdParams{
		Sourceid: SourceID,
		Limit: int32(limit),
		Offset: int32(limit * page),
	})
	if err != nil {
		return res, err
	}

	for _, article := range a {
		res = append(res, c.convertArticle(article))
	}

	return res, nil
}

func (c *DtoClient) convertArticle(i database.Article) models.ArticleDto {
	return models.ArticleDto{
		ID:          i.ID,
		Source:      i.Sourceid,
		Tags:        c.SplitTags(i.Tags),
		Title:       i.Title,
		Url:         i.Url,
		Pubdate:     i.Pubdate,
		Video:       i.Video.String,
		Videoheight: i.Videoheight,
		Videowidth:  i.Videoheight,
		Thumbnail:   i.Thumbnail,
		Description: i.Description,
		Authorname:  i.Authorname.String,
		Authorimage: i.Authorimage.String,
	}
}

func (c *DtoClient) convertArticleDetails(i database.Article, s database.Source) models.ArticleDetailsDto {
	return models.ArticleDetailsDto{
		ID:          i.ID,
		Source:      c.ConvertToSource(s),
		Tags:        c.SplitTags(i.Tags),
		Title:       i.Title,
		Url:         i.Url,
		Pubdate:     i.Pubdate,
		Video:       i.Video.String,
		Videoheight: i.Videoheight,
		Videowidth:  i.Videoheight,
		Thumbnail:   i.Thumbnail,
		Description: i.Description,
		Authorname:  i.Authorname.String,
		Authorimage: i.Authorimage.String,
	}
}

func (c DtoClient) SplitTags(t string) []string {
	return strings.Split(t, ", ")
}
