package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	usecase "onion/internal/usecase/interfaces"
)

func New(port int, logger usecase.Logger) *Router {
	return &Router{
		mux:    http.NewServeMux(),
		port:   port,
		logger: logger,
	}
}

type Router struct {
	mux    *http.ServeMux
	port   int
	logger usecase.Logger
}

func (r *Router) Run() error {
	routerWithMiddleware := LoggingMiddleware(r.logger, r.mux) // リクエスト処理の前後での処理
	r.logger.Info(context.Background(), fmt.Sprintf("Server is running on port %d", r.port))
	addr := fmt.Sprintf(":%d", r.port)
	if err := http.ListenAndServe(addr, routerWithMiddleware); err != nil {
		return err
	}
	return nil
}

func (r *Router) AddFetchArticlesHandler(uc usecase.FetchArticlesUsecase) {
	const path = "GET /articles"
	r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		articles, err := uc.Run(req.Context())
		if err != nil {
			r.logger.Error(req.Context(), "Failed to fetch articles", fmt.Sprintf("%+v", err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		articlesData := make([]Article, 0, len(articles))
		for _, article := range articles {
			articlesData = append(articlesData, convertArticleDomainModelToArticleDataModel(article))
		}

		data, err := json.Marshal(articlesData)
		if err != nil {
			r.logger.Error(req.Context(), "Failed to marshal articles", fmt.Sprintf("%+v", err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
	r.logger.Info(context.Background(), fmt.Sprintf("Added handler for %s", path))
}

func (r *Router) AddCreateArticleHandler(uc usecase.CreateArticleUsecase) {
	const path = "POST /article"
	r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		article, err := uc.Run(req.Context(), req.FormValue("title"), req.FormValue("body"))
		if err != nil {
			r.logger.Error(req.Context(), "Failed to create article", fmt.Sprintf("%+v", err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		articleData := convertArticleDomainModelToArticleDataModel(article)
		data, err := json.Marshal(articleData)
		if err != nil {
			r.logger.Error(req.Context(), "Failed to marshal article", fmt.Sprintf("%+v", err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
	r.logger.Info(context.Background(), fmt.Sprintf("Added handler for %s", path))
}
