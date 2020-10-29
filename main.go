package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	threadDelivery "github.com/kostikans/bdProject/restapi/thread/delivery"
	threadRepository "github.com/kostikans/bdProject/restapi/thread/repository"
	threadUsecase "github.com/kostikans/bdProject/restapi/thread/usecase"

	forumDelivery "github.com/kostikans/bdProject/restapi/forum/delivery"

	forumRepository "github.com/kostikans/bdProject/restapi/forum/repository"
	forumUsecase "github.com/kostikans/bdProject/restapi/forum/usecase"

	"github.com/kostikans/bdProject/pkg/logger"
	"github.com/kostikans/bdProject/restapi/user/delivery"

	userUsecase "github.com/kostikans/bdProject/restapi/user/usecase"

	userRepository "github.com/kostikans/bdProject/restapi/user/repository"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/kostikans/bdProject/configs"
	_ "github.com/lib/pq"
)

func initDB() *sqlx.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		configs.BdConfig.User,
		configs.BdConfig.Password,
		configs.BdConfig.DBName)

	fmt.Println(connStr)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	staticDir := "../static/"
	router.
		PathPrefix("/static").
		Handler(http.StripPrefix("/static", http.FileServer(http.Dir(staticDir))))

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	return router
}

func initRelativePath() string {
	_, fileName, _, _ := runtime.Caller(0)
	return filepath.ToSlash(filepath.Dir(filepath.Dir(fileName))) + "/"
}

func main() {

	configs.Init()
	db := initDB()
	r := NewRouter()

	configs.PrefixPath = initRelativePath()
	logOutput, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}
	defer logOutput.Close()

	log := logger.NewLogger(logOutput)

	uRepo := userRepository.NewUserRepository(db)
	fRepo := forumRepository.NewForumRepository(db)
	tRepo := threadRepository.NewThreadRepository(db)

	uUse := userUsecase.NewUserUseCase(uRepo)
	fUse := forumUsecase.NewForumUseCase(fRepo)
	tUse := threadUsecase.NewThreadUseCase(tRepo)

	delivery.NewUserHandler(r, uUse, log)
	forumDelivery.NewForumHandler(r, fUse, log)
	threadDelivery.NewThreadHandler(r, tUse, log)

	log.Info("Server started at port", configs.Port)
	http.ListenAndServe(configs.Port, r)
}
