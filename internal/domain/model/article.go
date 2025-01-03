package model

import (
	"errors"

	"github.com/google/uuid"
)

type Article struct {
	ID    ArticleID
	Title ArticleTitle
	Body  ArticleBody
}

func NewArticle(id ArticleID, title ArticleTitle, body ArticleBody) (*Article, error) {
	return &Article{
		ID:    id,
		Title: title,
		Body:  body,
	}, nil
}

type ArticleID struct {
	value string
}

func NewArticleID(value string) (ArticleID, error) {
	// uuidかどうか
	id, err := uuid.Parse(value)
	if err != nil {
		return ArticleID{}, errors.New("invalid id")
	}

	return ArticleID{
		value: id.String(),
	}, nil
}

func (i ArticleID) Value() string {
	return i.value
}

type ArticleTitle struct {
	value string
}

func NewArticleTitle(value string) (ArticleTitle, error) {
	if value == "" {
		return ArticleTitle{}, errors.New("title is required")
	}

	return ArticleTitle{
		value: value,
	}, nil
}

func (t ArticleTitle) Value() string {
	return t.value
}

type ArticleBody struct {
	value string
}

func NewArticleBody(value string) (ArticleBody, error) {
	if value == "" {
		return ArticleBody{}, errors.New("content is required")
	}

	return ArticleBody{
		value: value,
	}, nil
}

func (b ArticleBody) Value() string {
	return b.value
}
