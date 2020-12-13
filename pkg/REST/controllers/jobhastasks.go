package pgucontrollers

import (
	// external
	Request "github.com/eshu0/RESTServer/pkg/request"
	RSServer "github.com/eshu0/RESTServer/pkg/server"
	per "github.com/eshu0/persist/pkg/interfaces"
	"github.com/eshu0/persist/pkg/sqllite"

	// these are from the package
	hndlr "github.com/eshu0/todos/pkg/handlers"
	models "github.com/eshu0/todos/pkg/models"
)

// Controller

type JobHasTasksController struct {
	JobHasTasksHandler *hndlr.JobHasTasksHandler // Storage handler
	Server *RSServer.RServer
}

func NewJobHasTasksController(handler *hndlr.JobHasTasksHandler, Server *RSServer.RServer) *JobHasTasksController {
	ds := JobHasTasksController{}
	ds.JobHasTasksHandler = handler
	ds.Server = Server
	return &ds
}

func (controller *JobHasTasksController) HandleRequest(request Request.ServerRequest) per.IQueryResult {  //.JobHasTask {
	data := request.Payload.(models.JobHasTask)
	
	if request.Request.Method == "POST" {
		controller.Server.LogDebug("HandleRequest", "Calling to insert a new JobHasTask")
		result := controller.JobHasTasksHandler.Create(data)
		return result

	} else if request.Request.Method == "PUT" { 
	
		controller.Server.LogDebug("HandleRequest", "Calling to update JobHasTask")
		result := controller.JobHasTasksHandler.Update(data)
		return result

	} else if request.Request.Method == "DELETE" { 
	
		controller.Server.LogDebug("HandleRequest", "Calling to update the (DELETE) JobHasTask")
		result := controller.JobHasTasksHandler.Update(data)
		return result

	} else {
		
		Id := controller.Server.RequestHelper.GetRequestId(request.Request,"Id")
		if Id != nil {
			controller.Server.LogDebugf("HandleRequest", "Id was not nil and have the following to lookup %d", *Id)
			result := controller.JobHasTasksHandler.FindById(int64(*Id))
			return result
		} else {
			controller.Server.LogError("HandleRequest", "Id was nil")
		}
	}
	
	controller.Server.LogError("HandleRequest", "Failed returning empty SQLLiteResult")
	return SQLL.NewEmptyFailedSQLLiteQueryResult()
}

func (controller *JobHasTasksController) HandleReadAllRequest(request Request.ServerRequest) per.IQueryResult { 
	controller.Server.LogDebug("HandleRequest", "Calling to read all JobHasTask")
	result := controller.JobHasTasksHandler.ReadAll()
	return result
}


