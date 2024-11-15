package main

import (
	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"

	"github.com/deej-tsn/blog-go/internal/helper"
	"github.com/deej-tsn/blog-go/internal/routes"
	"github.com/deej-tsn/blog-go/web/components"
)

func main() {

	//db := sql.ConnectToDatabase()
	e := echo.New()

	//CONTROLLERS

	controller := routes.NewFileReader()

	// MIDDLEWARE

	// logs all http requests
	e.Use(middleware.Logger())

	// static files in public folder
	e.Static("static", "./web/public")

	// resets server if error
	e.Use(middleware.Recover())

	// CORS - cross origin resource sharing
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// allow any origin - DEVELOPMENT ONLY
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	index := components.Index("Posts", components.LoadPostList())

	// ROUTES
	e.GET("/", func(c echo.Context) error {
		return helper.Render(c, 100, index)
	})

	e.GET("posts/:slug", controller.GetPost)
	e.GET("posts", controller.GetAllPosts)

	e.Logger.Fatal(e.Start(":1323"))
}
