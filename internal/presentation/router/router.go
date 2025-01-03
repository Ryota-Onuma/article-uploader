package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	usecase "onion/internal/usecase/interfaces"
	"sort"
	"text/template"
)

const (
	indexFilePath = "internal/presentation/template/index.html"
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

func (r *Router) AddTopHandler(uc usecase.FetchArticlesUsecase) {
	const path = "GET /"
	r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		articles, err := uc.Run(req.Context())
		if err != nil {
			r.logger.Error(req.Context(), "Failed to fetch articles", fmt.Sprintf("%+v", err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// createdAtの降順にソート
		sort.Slice(articles, func(i, j int) bool {
			return articles[i].CreatedAt.After(articles[j].CreatedAt)
		})

		articlesData := make([]Article, 0, len(articles))
		for _, article := range articles {
			articlesData = append(articlesData, convertArticleDomainModelToArticleDataModel(article))
		}

		data := struct {
			Articles []Article
		}{
			Articles: articlesData,
		}

		t := template.New("index.html")
		t, err = t.ParseFiles(indexFilePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := t.ExecuteTemplate(w, "index.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
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
		_, err := uc.Run(req.Context(), req.FormValue("title"), req.FormValue("body"))
		if err != nil {
			r.logger.Error(req.Context(), "Failed to create article", fmt.Sprintf("%+v", err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, req, "/", http.StatusMovedPermanently)
	})
	r.logger.Info(context.Background(), fmt.Sprintf("Added handler for %s", path))
}
