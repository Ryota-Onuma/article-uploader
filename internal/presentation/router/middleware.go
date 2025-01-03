package router

import (
	"context"
	"net/http"
	usecase "onion/internal/usecase/interfaces"

	"github.com/google/uuid"
)

func LoggingMiddleware(logger usecase.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// traceIDを生成し、コンテキストに設定
		ctx := context.WithValue(r.Context(), usecase.TraceIDKey{}, uuid.New().String())

		// 次のハンドラを呼び出す
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
