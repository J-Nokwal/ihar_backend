package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "net/http"
	"github.com/J-Nokwal/ihar_backend/pkg/routes"
)

func main() {
	r := gin.Default()

	routes.RegisterBookStoreRoutes(r)
	// r.Use(cors.New(cors.Config{
	// 	// AllowOrigins:     []string{"http://localhost:30090", "http://192.168.18.29:46590"},
	// 	AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
	// 	AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
	// 	AllowCredentials: true,
	// 	// AllowOriginFunc: func(origin string) bool {
	// 	// 	fmt.Println(origin)
	// 	// 	return true
	// 	// },
	// 	AllowAllOrigins: true,
	// }))

	// r.Use(cors.New(cors.Config{
	// 	// AllowOrigins:     []string{"http://localhost:30090", "http://192.168.18.29:46590"},
	// 	AllowMethods:           []string{"*"},
	// 	AllowHeaders:           []string{"*"},
	// 	ExposeHeaders:          []string{"*"},
	// 	MaxAge:                 time.Second * 5,
	// 	AllowWildcard:          true,
	// 	AllowBrowserExtensions: true,
	// 	AllowFiles:             true,
	// 	AllowWebSockets:        true,
	// 	AllowCredentials:       true,
	// 	// AllowOriginFunc: func(origin string) bool {
	// 	// 	fmt.Println(origin)
	// 	// 	return true
	// 	// },
	// 	AllowAllOrigins: true,
	// }))
	r.Use(cors.New(CORSConfig()))
	// r.Use(cors.Default())
	r.Run()

}
func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE", "OPTIONS")
	// corsConfig.AllowAllOrigins = true
	corsConfig.AllowOrigins = []string{"*"}
	return corsConfig

}
