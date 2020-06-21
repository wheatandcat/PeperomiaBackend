package main

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wheatandcat/PeperomiaBackend/backend/handler"
	"github.com/wheatandcat/PeperomiaBackend/backend/middleware"
	"github.com/wheatandcat/PeperomiaBackend/backend/repository"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "https://app.peperomia.info"},
		AllowMethods: []string{
			http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions,
		},
		AllowHeaders: []string{
			"Authorization",
			"Content-Type",
			"Cache-Control",
		},
		AllowCredentials: true,
		MaxAge:           86400,
	}))

	ctx := context.Background()
	f, err := repository.FirebaseApp(ctx)

	if err != nil {
		panic(err)
	}

	m := middleware.NewMiddleware(f)

	app := r.Group("")
	{
		app.Use(m.FirebaseAuthMiddleWare)

		h, err := handler.NewHandler(ctx, f)
		if err != nil {
			panic(err)
		}

		app.POST("/CreateUser", h.CreateUser)
		app.POST("/CreateItem", h.CreateItem)
		app.POST("/UpdateItem", h.UpdateItem)
		app.POST("/DeleteItem", h.DeleteItem)
		app.POST("/CreateItemDetail", h.CreateItemDetail)
		app.POST("/UpdateItemDetail", h.UpdateItemDetail)
		app.POST("/DeleteItemDetail", h.DeleteItemDetail)
		app.POST("/CreateCalendar", h.CreateCalendar)
		app.POST("/UpdateCalendar", h.UpdateCalendar)
		app.POST("/DeleteCalendar", h.DeleteCalendar)

		app.POST("/SyncItems", h.SyncItems)
		app.POST("/LoginWithAmazon", h.LoginWithAmazon)
		app.POST("/CreatePushToken", h.CreatePushToken)

	}

	ad := r.Group("/admin")
	{
		ad.Use(m.FirebaseAuthMiddleWare)
		ad.Use(m.AdminMiddleWare)
		h, err := handler.NewHandler(ctx, f)
		if err != nil {
			panic(err)
		}
		// Push通知のテスト
		ad.POST("/SentPushNotifications", h.SentPushNotifications)
	}

	am := r.Group("/amazon")
	{
		am.Use(m.AmazonAuthMiddleWare)
		h, err := handler.NewHandler(ctx, f)
		if err != nil {
			panic(err)
		}

		am.POST("/RegisterItem", h.AmazonRegisterItem)
	}

	cr := r.Group("/cron")
	{
		h, err := handler.NewHandler(ctx, f)
		if err != nil {
			panic(err)
		}

		cr.GET("/SendCalendarPushNotifications", h.SendCalendarPushNotifications)
	}

	r.Run()
}
