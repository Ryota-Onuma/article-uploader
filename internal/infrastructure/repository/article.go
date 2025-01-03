package repository

import (
	"encoding/json"
	"fmt"
	"onion/internal/domain/model"
	"onion/internal/domain/repository"
	"os"
	"path/filepath"
)

const dataPath = "testdata/articles"

type Article struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var _ repository.ArticleRepository = (*ArticleRepository)(nil)

// DBを用意するのが面倒なので、ファイルに保存する

type ArticleRepository struct{}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{}
}

func (a *ArticleRepository) FetchArticles() ([]model.Article, error) {
	files, err := os.ReadDir(dataPath)
	if err != nil {
		return nil, err
	}

	var articles []model.Article
	for _, file := range files {
		article, err := loadArticleFromFile(filepath.Join(dataPath, file.Name()))
		if err != nil {
			return nil, err
		}
		articles = append(articles, *article)
	}
	return articles, nil
}

func (a *ArticleRepository) CreateArticle(article model.Article) error {
	fileName := fmt.Sprintf("%s/%s.json", dataPath, article.ID.Value())
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := json.Marshal(&Article{ID: article.ID.Value(), Title: article.Title.Value(), Body: article.Body.Value()})
	if err != nil {
		return err
	}

	if _, err := file.Write(data); err != nil {
		return err
	}

	return nil
}

func (a *ArticleRepository) FetchArticle(id model.ArticleID) (model.Article, error) {
	files, err := os.ReadDir(dataPath)
	if err != nil {
		return model.Article{}, err
	}

	for _, file := range files {
		if file.Name() != fmt.Sprintf("%s.json", id.Value()) {
			continue
		}

		article, err := loadArticleFromFile(filepath.Join(dataPath, file.Name()))
		if err != nil {
			return model.Article{}, err
		}
		return *article, nil
	}
	return model.Article{}, fmt.Errorf("article not found")
}

func loadArticleFromFile(fileName string) (*model.Article, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var article Article
	if err := json.Unmarshal(file, &article); err != nil {
		return nil, err
	}
	return article.ToDomainModel()
}

func (a *Article) ToDomainModel() (*model.Article, error) {
	id, err := model.NewArticleID(a.ID)
	if err != nil {
		return nil, err
	}

	title, err := model.NewArticleTitle(a.Title)
	if err != nil {
		return nil, err
	}

	body, err := model.NewArticleBody(a.Body)
	if err != nil {
		return nil, err
	}

	return model.NewArticle(id, title, body)
}
