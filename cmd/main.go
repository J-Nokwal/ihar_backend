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
	r.Use(cors.New(CORSConfig()))
	// r.Use(cors.Default())
	r.Run()

}
func CORSConfig() cors.Config {
	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{"*"}
	// corsConfig.AllowCredentials = true
	// corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	// corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	// // corsConfig.AllowAllOrigins = true
	// corsConfig.AllowOriginFunc(func (s string )bool {return true})
	// return corsConfig
	return cors.Config{
		// AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	fmt.Println(origin)
		// 	return true
		// },
		AllowWildcard:   true,
		AllowAllOrigins: true,
	}
}
