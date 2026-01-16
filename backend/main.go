package main

import (
	"fmt"
	"glyph/database"
	"glyph/handlers/blog/article"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	urlPrefix := "/api/1.0"

	articlePrefix := urlPrefix + "/article"
	router.Get(articlePrefix+"/list", article.List(dbConn))
	router.Get(articlePrefix+"/get/{articleId}", article.Get(dbConn))
	router.Post(articlePrefix+"/create", article.Create(dbConn))
	router.Put(articlePrefix+"/update", article.Update(dbConn))
	router.Delete(articlePrefix+"/delete/{articleId}", article.Delete(dbConn))

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
