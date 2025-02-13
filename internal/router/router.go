package router

import (
	"context"
	"easy-cooking/internal/handler"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var r *gin.Engine

func InitRouter(handler *handler.Handler) {
	r = gin.Default()
	//r.Use(middleware.CORSMiddleware())
	//r.Static("/uploads/images", "./uploads/images")

	//recipe
	RecipeRouter(r, handler)
}

func StartRouter(port string) {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("Server is running on port %s", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
