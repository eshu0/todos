package pgucontrollers

import (
	// external
	Request "github.com/eshu0/RESTServer/pkg/request"
	per "github.com/eshu0/persist/pkg/interfaces"
	RSServer "github.com/eshu0/RESTServer/pkg/server"
	"github.com/eshu0/persist/pkg/sqllite"

	// these are from the package
	hndlr "eshu0/todos/github.com/pkg/handlers"
	models "eshu0/todos/github.com/pkg/models"
)

// Controller

type ProjectsController struct {
	ProjectsHandler *hndlr.ProjectsHandler // Storage handler
	Server *RSServer.RServer
}

func NewProjectsController(handler *hndlr.ProjectsHandler, Server *RSServer.RServer) *ProjectsController {
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

