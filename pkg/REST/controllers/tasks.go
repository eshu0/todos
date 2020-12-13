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

type TasksController struct {
	TasksHandler *handlers.TasksHandler // Storage handler
	Server *RSServer.RServer
}

func NewTasksController(handler *handlers.TasksHandler, Server *RSServer.RServer) *TasksController {
	ds := TasksController{}
	ds.TasksHandler = handler
	ds.Server = Server
	return &ds
}

func (controller *TasksController) HandleRequest(request Request.ServerRequest) per.IQueryResult {  //.Task {
	data := request.Payload.(models.Task)
	
	if request.Request.Method == "POST" {
		controller.Server.LogDebug("HandleRequest", "Calling to insert a new Task")
		result := controller.TasksHandler.Create(data)
		return result

	} else if request.Request.Method == "PUT" { 
	
		controller.Server.LogDebug("HandleRequest", "Calling to update Task")
		result := controller.TasksHandler.Update(data)
		return result

	} else if request.Request.Method == "DELETE" { 
	
		controller.Server.LogDebug("HandleRequest", "Calling to update the (DELETE) Task")
		result := controller.TasksHandler.Update(data)
		return result

	} else {
		
		Id := controller.Server.RequestHelper.GetRequestId(request.Request,"Id")
		if Id != nil {
			controller.Server.LogDebugf("HandleRequest", "Id was not nil and have the following to lookup %d", *Id)
			result := controller.TasksHandler.FindById(int64(*Id))
			return result
		} else {
			controller.Server.LogError("HandleRequest", "Id was nil")
		}
	}
	
	controller.Server.LogError("HandleRequest", "Failed returning empty SQLLiteResult")
	return SQLL.NewEmptyFailedSQLLiteQueryResult()
}

func (controller *TasksController) HandleReadAllRequest(request Request.ServerRequest) per.IQueryResult { 
	controller.Server.LogDebug("HandleRequest", "Calling to read all Task")
	result := controller.TasksHandler.ReadAll()
	return result
}


