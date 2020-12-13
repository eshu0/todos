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

type JobsController struct {
	JobsHandler *handlers.JobsHandler // Storage handler
	Server *RSServer.RServer
}

func NewJobsController(handler *handlers.JobsHandler, Server *RSServer.RServer) *JobsController {
	ds := JobsController{}
	ds.JobsHandler = handler
	ds.Server = Server
	return &ds
}

func (controller *JobsController) HandleRequest(request Request.ServerRequest) per.IQueryResult {  //.Job {
	data := request.Payload.(models.Job)
	
	if request.Request.Method == "POST" {
		controller.Server.LogDebug("HandleRequest", "Calling to insert a new Job")
		result := controller.JobsHandler.Create(data)
		return result

	} else if request.Request.Method == "PUT" { 
	
		controller.Server.LogDebug("HandleRequest", "Calling to update Job")
		result := controller.JobsHandler.Update(data)
		return result

	} else if request.Request.Method == "DELETE" { 
	
		controller.Server.LogDebug("HandleRequest", "Calling to update the (DELETE) Job")
		result := controller.JobsHandler.Update(data)
		return result

	} else {
		
		Id := controller.Server.RequestHelper.GetRequestId(request.Request,"Id")
		if Id != nil {
			controller.Server.LogDebugf("HandleRequest", "Id was not nil and have the following to lookup %d", *Id)
			result := controller.JobsHandler.FindById(int64(*Id))
			return result
		} else {
			controller.Server.LogError("HandleRequest", "Id was nil")
		}
	}
	
	controller.Server.LogError("HandleRequest", "Failed returning empty SQLLiteResult")
	return SQLL.NewEmptyFailedSQLLiteQueryResult()
}

func (controller *JobsController) HandleReadAllRequest(request Request.ServerRequest) per.IQueryResult { 
	controller.Server.LogDebug("HandleRequest", "Calling to read all Job")
	result := controller.JobsHandler.ReadAll()
	return result
}


