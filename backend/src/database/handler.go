package database

import (
	"context"
	"github.com/go-pg/pg"
	"net/http"
)

type databaseKey string

func contextWithDatabase(ctx context.Context, db *pg.DB) context.Context {
	return context.WithValue(ctx, databaseKey("database"), db)
}

func GetDatabaseFromContext(ctx context.Context) *pg.DB {
	return ctx.Value(databaseKey("database")).(*pg.DB)
}

func NewHandler(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := contextWithDatabase(req.Context(), db)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
