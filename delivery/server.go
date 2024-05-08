package delivery

import (
	"eniqloStore/config"
	"eniqloStore/delivery/controller"
	"eniqloStore/delivery/middleware"
	"eniqloStore/manager"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type appServer struct {
	usecaseManager manager.UsecaseManager
	router         *gin.RouterGroup
	routerDev      *gin.RouterGroup
	engine         *gin.Engine
	host           string
}

// func (a *appServer) Router() *gin.Engine {
// 	return a.engine
// }

func Server() *appServer {
	router := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(appConfig)
	repoManager := manager.NewRepositoryManager(infra)
	// repoManager.NotifRepo()
	usecaseManager := manager.NewUsecaseManager(repoManager)

	host := appConfig.Url
	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"Origin", "Date", "Content-Length", "Content-Type", "Content-Disposition", "Accept", "X-Requested-With", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Authorization", "token"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	return &appServer{
		usecaseManager: usecaseManager,
		engine:         router,
		router:         router.Group("", middleware.NewAuthTokenMiddleware(usecaseManager.TokenUsecase()).RequiredToken()),

		routerDev: router.Group("activation/"),
		host:      host,
	}
}

func (a *appServer) initControllers() {
	// buat daftarin controller ada disini
	// setiap controller, isinya harus ada isian dari usecaseManager
	controller.NewUserController(a.router, a.routerDev, a.usecaseManager.RegisterUsecase(), a.usecaseManager.LoginUsecase())

	controller.NewProductController(a.router, a.routerDev, a.usecaseManager.ProductUsecase())
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
