package main

import (
	"onion/internal/domain/service/articles"
	"onion/internal/infrastructure/config"
	errs "onion/internal/infrastructure/error"
	"onion/internal/infrastructure/logger"
	"onion/internal/infrastructure/repository"
	"onion/internal/presentation/router"
	"onion/internal/usecase"
	usecasearticles "onion/internal/usecase/articles"
)

func main() {
	config := config.New()
	if err := config.Setup(); err != nil {
		panic(err)
	}

	var l *logger.Logger
	if config.Server.IsLocal() {
		l = logger.New(logger.NewDebugHandler()) // ローカルでは見やすいロガーを使う
	} else {
		l = logger.New(logger.NewJSONHandler())
	}

	rt := router.New(config.Server.Port, l)

	// DIする
	baseUsecase := usecase.NewBaseUsecase(l, errs.New())
	rt.AddFetchArticlesHandler(
		usecasearticles.NewFetchArticlesUsecase(baseUsecase, repository.NewArticleRepository()),
	)
	rt.AddCreateArticleHandler(
		usecasearticles.NewCreateArticleUsecase(baseUsecase, articles.NewCreateArticleService(), repository.NewArticleRepository()),
	)

	// サーバーを起動する
	if err := rt.Run(); err != nil {
		panic(err)
	}
}
