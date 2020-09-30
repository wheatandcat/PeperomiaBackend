package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wheatandcat/PeperomiaBackend/domain"
	graph "github.com/wheatandcat/PeperomiaBackend/graph"
	"github.com/wheatandcat/PeperomiaBackend/graph/generated"
	"github.com/wheatandcat/PeperomiaBackend/handler"
	"github.com/wheatandcat/PeperomiaBackend/middleware"
	"github.com/wheatandcat/PeperomiaBackend/repository"
)

// Defining the Graphql handler
func graphqlHandler(h *handler.Handler) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	gh := gqlHandler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Handler: h,
	}}))

	return func(c *gin.Context) {
		gh.ServeHTTP(c.Writer, c.Request)
	}
}

func ginContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), domain.GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func main() {

	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DNS"),
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if hint.Context != nil {
				if req, ok := hint.Context.Value(sentry.RequestContextKey).(*http.Request); ok {
					fmt.Println(req)
				}
			}
			fmt.Println(event)
			return event
		},
		Debug:            false,
		AttachStacktrace: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	r := gin.Default()

	r.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

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
		app.Use(func(ctx *gin.Context) {
			ctx.Set("role", handler.RoleApp)
			ctx.Next()
		})

		app.Use(m.FirebaseAuthMiddleWare)

		h, err := handler.NewHandler(ctx, f)
		if err != nil {
			panic(err)
		}

		app.POST("/CreateUser", h.CreateUser)
		app.POST("/CreateItem", h.CreateItem)
		app.POST("/UpdateItem", h.UpdateItem)
		app.POST("/UpdateItemPublic", h.UpdateItemPublic)
		app.POST("/UpdateItemPrivate", h.UpdateItemPrivate)
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
		app.Use(func(ctx *gin.Context) {
			ctx.Set("role", handler.RoleAdmin)
			ctx.Next()
		})

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
		app.Use(func(ctx *gin.Context) {
			ctx.Set("role", handler.RoleCron)
			ctx.Next()
		})

		h, err := handler.NewHandler(ctx, f)
		if err != nil {
			panic(err)
		}

		cr.GET("/SendCalendarPushNotifications", h.SendCalendarPushNotifications)
	}

	gqr := r.Group("/graphql")
	{
		gqr.Use(ginContextToContextMiddleware())
		app.Use(func(ctx *gin.Context) {
			ctx.Set("role", handler.RoleGraphql)
			ctx.Next()
		})

		h, err := handler.NewHandler(ctx, f)
		if err != nil {
			panic(err)
		}
		gqr.POST("", graphqlHandler(h))
	}

	gqrApp := r.Group("/app/graphql")
	{
		gqrApp.Use(ginContextToContextMiddleware())
		gqrApp.Use(m.FirebaseAuthMiddleWare)
		app.Use(func(ctx *gin.Context) {
			ctx.Set("role", handler.RoleAppGraphql)
			ctx.Next()
		})

		h, err := handler.NewHandler(ctx, f)
		if err != nil {
			panic(err)
		}
		gqrApp.POST("", graphqlHandler(h))
	}

	r.Run()
}
