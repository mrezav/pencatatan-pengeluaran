package application

import (
	"your/path/project/domain_myspending/controller/restapi"
	"your/path/project/domain_myspending/gateway/prod"
	"your/path/project/domain_myspending/usecase/getallexpense"
	"your/path/project/domain_myspending/usecase/runexpensecreate"
	"your/path/project/shared/driver"
	"your/path/project/shared/infrastructure/config"
	"your/path/project/shared/infrastructure/logger"
	"your/path/project/shared/infrastructure/server"
	"your/path/project/shared/infrastructure/util"
)

type myapp struct {
	httpHandler *server.GinHTTPHandler
	controller  driver.Controller
}

func (c myapp) RunApplication() {
	c.controller.RegisterRouter()
	c.httpHandler.RunApplication()
}

func NewMyapp() func() driver.RegistryContract {
	return func() driver.RegistryContract {

		cfg := config.ReadConfig()

		appID := util.GenerateID(4)

		appData := driver.NewApplicationData("myapp", appID)

		log := logger.NewSimpleJSONLogger(appData)

		httpHandler := server.NewGinHTTPHandlerDefault(log, appData, cfg)

		datasource := prod.NewGateway(log, appData, cfg)

		return &myapp{
			httpHandler: &httpHandler,
			controller: &restapi.Controller{
				Log:                    log,
				Config:                 cfg,
				Router:                 httpHandler.Router,
				RunExpenseCreateInport: runexpensecreate.NewUsecase(datasource),
				GetAllExpenseInport:    getallexpense.NewUsecase(datasource),
			},
		}

	}
}
