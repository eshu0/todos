package controllers

import (
	// external
	Request "github.com/eshu0/RESTServer/pkg/request"
	RSServer "github.com/eshu0/RESTServer/pkg/server"
	per "github.com/eshu0/persist/pkg/interfaces"
	"github.com/eshu0/persist/pkg/sqllite"

	// these are from the package
	"github.com/eshu0/todos/pkg/handlers"
	"github.com/eshu0/todos/pkg/models"
)

// Controller

type ProjectHasJobsController struct {
	ProjectHasJobsHandler *hndlr.ProjectHasJobsHandler // Storage handler
	Server *RSServer.RServer
}

func NewProjectHasJobsController(handler *hndlr.ProjectHasJobsHandler, Server *RSServer.RServer) *ProjectHasJobsController {
	ds := ProjectHasJobsController{}
	ds.ProjectHasJobsHandler = handler
	ds.Server = Server
	return &ds
}

func (controller *ProjectHasJobsController) HandleRequest(request Request.ServerRequest) per.IQueryResult {  //.ProjectHasJob {
	data := request.Payload.(models.ProjectHasJob)
	
	if request.Request.Method == "POST" {
		controller.Server.LogDebug("HandleRequest", "Calling to insert a new ProjectHasJob")
		result := controller.ProjectHasJobsHandler.Create(data)
		return result

	} else if request.Request.Method == "PUT" { 
	
		controller.Server.LogDebug("HandleRequest", "Calling to update ProjectHasJob")
		result := controller.ProjectHasJobsHandler.Update(data)
		return result

	} else if request.Request.Method == "DELETE" { 
	
		controller.Server.LogDebug("HandleRequest", "Calling to update the (DELETE) ProjectHasJob")
		result := controller.ProjectHasJobsHandler.Update(data)
		return result

	} else {
		
		Id := controller.Server.RequestHelper.GetRequestId(request.Request,"Id")
		if Id != nil {
			controller.Server.LogDebugf("HandleRequest", "Id was not nil and have the following to lookup %d", *Id)
			result := controller.ProjectHasJobsHandler.FindById(int64(*Id))
			return result
		} else {
			controller.Server.LogError("HandleRequest", "Id was nil")
		}
	}
	
	controller.Server.LogError("HandleRequest", "Failed returning empty SQLLiteResult")
	return SQLL.NewEmptyFailedSQLLiteQueryResult()
}

func (controller *ProjectHasJobsController) HandleReadAllRequest(request Request.ServerRequest) per.IQueryResult { 
	controller.Server.LogDebug("HandleRequest", "Calling to read all ProjectHasJob")
	result := controller.ProjectHasJobsHandler.ReadAll()
	return result
}


