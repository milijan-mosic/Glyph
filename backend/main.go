package main

import (
	"fmt"
	"glyph/database"
	"glyph/handlers/blog/article"
	"glyph/handlers/blog/comment"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
	"github.com/go-chi/jwtauth/v5"
	"gorm.io/gorm"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil) // replace with secret key

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func main() {
	dbConn := database.InitializeDatabase()

	port := ":5000"
	fmt.Printf("Starting server on %v\n", port)
	http.ListenAndServe(port, newRouter(dbConn))
}

func newRouter(dbConn *gorm.DB) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.AllowContentEncoding("gzip"))
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.Compress(5, "application/json"))
	router.Use(middleware.CleanPath)
	router.Use(httprate.LimitByIP(1000, 1*time.Minute))
	router.Use(middleware.Timeout(5 * time.Second))
	router.Use(middleware.Heartbeat("/health"))
	router.Use(middleware.Recoverer)

	urlPrefix := "/api/1.0"

	router.Route(urlPrefix+"/article", func(r chi.Router) {
		router.Get("/list", article.List(dbConn))
		router.Get("/get/{articleId}", article.GetByID(dbConn))
		router.Post("/create", article.Create(dbConn))
		router.Put("/update", article.Update(dbConn))
		router.Delete("/delete/{articleId}", article.DeleteByID(dbConn))
	})

	router.Route(urlPrefix+"/comment", func(r chi.Router) {
		r.Post("/create", comment.Create(dbConn))
		r.Get("/list/{articleId}", comment.ListByArticleID(dbConn))
		r.Get("/pending", comment.ListPending(dbConn))
		r.Put("/approve/{commentId}", comment.Approve((dbConn)))
		r.Delete("/delete/{commentId}", comment.DeleteByID(dbConn))
	})

	return router
}

// func router() http.Handler {
// 	r := chi.NewRouter()

// 	// Protected routes
// 	r.Group(func(r chi.Router) {
// 		// Seek, verify and validate JWT tokens
// 		r.Use(jwtauth.Verifier(tokenAuth))

// 		// Handle valid / invalid tokens. In this example, we use
// 		// the provided authenticator middleware, but you can write your
// 		// own very easily, look at the Authenticator method in jwtauth.go
// 		// and tweak it, its not scary.
// 		r.Use(jwtauth.Authenticator)

// 		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
// 			_, claims, _ := jwtauth.FromContext(r.Context())
// 			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
// 		})
// 	})

// 	// Public routes
// 	r.Group(func(r chi.Router) {
// 		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 			w.Write([]byte("welcome anonymous"))
// 		})
// 	})

// 	return r
// }
