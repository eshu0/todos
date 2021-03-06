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

type ProjectsController struct {
	ProjectsHandler *handlers.ProjectsHandler // Storage handler
	Server *RSServer.RServer
}

func NewProjectsController(handler *handlers.ProjectsHandler, Server *RSServer.RServer) *ProjectsController {
	ds := ProjectsController{}
	ds.ProjectsHandler = handler
	ds.Server = Server
	return &ds
}

func (controller *ProjectsController) HandleRequest(request Request.ServerRequest) per.IQueryResult {  //.Project {
	data := request.Payload.(models.Project)
	
	if request.Request.Method == "POST" {
		controller.Server.LogDebug("HandleRequest", "Calling to insert a new Project")
		result := controller.ProjectsHandler.Create(data)
		return result

	} else if request.Request.Method == "PUT" { 
	
		controller.Server.LogDebug("HandleRequest", "Calling to update Project")
		result := controller.ProjectsHandler.Update(data)
		return result

	} else if request.Request.Method == "DELETE" { 
	
		controller.Server.LogDebug("HandleRequest", "Calling to update the (DELETE) Project")
		result := controller.ProjectsHandler.Update(data)
		return result

	} else {
		
		Id := controller.Server.RequestHelper.GetRequestId(request.Request,"Id")
		if Id != nil {
			controller.Server.LogDebugf("HandleRequest", "Id was not nil and have the following to lookup %d", *Id)
			result := controller.ProjectsHandler.FindById(int64(*Id))
			return result
		} else {
			controller.Server.LogError("HandleRequest", "Id was nil")
		}
	}
	
	controller.Server.LogError("HandleRequest", "Failed returning empty SQLLiteResult")
	return SQLL.NewEmptyFailedSQLLiteQueryResult()
}

func (controller *ProjectsController) HandleReadAllRequest(request Request.ServerRequest) per.IQueryResult { 
	controller.Server.LogDebug("HandleRequest", "Calling to read all Project")
	result := controller.ProjectsHandler.ReadAll()
	return result
}


