package main

import (
	"flag"

	RESTCommands "github.com/eshu0/RESTServer/pkg/commands"
	RSConfig "github.com/eshu0/RESTServer/pkg/config"
	RSServer "github.com/eshu0/RESTServer/pkg/server"

  	// these are from the package 
	"github.com/eshu0/todos/pkg/datastore"
	"github.com/eshu0/todos/pkg/models"
	"github.com/eshu0/todos/pkg/REST/controllers"		
)

func main() {

	dbname := flag.String("db", "./todos.db", "Database defaults to ./todos.db")
	flag.Parse()

	// create a new server
	conf := RSConfig.NewRServerConfig()

	// Create a new REST Server
	server := RSServer.NewRServer(conf)

	// load this first
	server.ConfigFilePath = "./config.json"

	ok := server.LoadConfig()

	if !ok {
		server.LogErrorf("Main", "Error : %s","Failed to load configuration server not started")
		return
	}
	
	server.LogInfof("Main", "Error : %s","Failed to load configuration server not started")
	
	// add the defaults here
	RESTCommands.AddDefaults(server)
	RESTCommands.SetDefaultFunctionalMap(server)

	fds := datastore.CreateDataStorage(*dbname)

	

	ProjectHasJobsHandler := fds.GetProjectHasJobsHandler()
	ProjectHasJobsController := controllers.NewProjectHasJobsController(ProjectHasJobsHandler, server)
	server.Register("ProjectHasJobsController",ProjectHasJobsController)

	

	ProjectsHandler := fds.GetProjectsHandler()
	ProjectsController := controllers.NewProjectsController(ProjectsHandler, server)
	server.Register("ProjectsController",ProjectsController)

	

	JobsHandler := fds.GetJobsHandler()
	JobsController := controllers.NewJobsController(JobsHandler, server)
	server.Register("JobsController",JobsController)

	

	TasksHandler := fds.GetTasksHandler()
	TasksController := controllers.NewTasksController(TasksHandler, server)
	server.Register("TasksController",TasksController)

	

	JobHasTasksHandler := fds.GetJobHasTasksHandler()
	JobHasTasksController := controllers.NewJobHasTasksController(JobHasTasksHandler, server)
	server.Register("JobHasTasksController",JobHasTasksController)

	

	

	newProjectHasJob  := models.ProjectHasJob{}
	
	server.AddJSONFunctionHandler("/ProjectHasJob/{Id}/","HandleRequest","GET","ProjectHasJobsController",newProjectHasJob)
	server.AddJSONFunctionHandler("/ProjectHasJob/","HandleRequest","","ProjectHasJobsController",newProjectHasJob)
	server.AddJSONFunctionHandler("/ProjectHasJobs/","HandleReadAllRequest","GET","ProjectHasJobsController",newProjectHasJob)

	

	newProject  := models.Project{}
	
	server.AddJSONFunctionHandler("/Project/{Id}/","HandleRequest","GET","ProjectsController",newProject)
	server.AddJSONFunctionHandler("/Project/","HandleRequest","","ProjectsController",newProject)
	server.AddJSONFunctionHandler("/Projects/","HandleReadAllRequest","GET","ProjectsController",newProject)

	

	newJob  := models.Job{}
	
	server.AddJSONFunctionHandler("/Job/{Id}/","HandleRequest","GET","JobsController",newJob)
	server.AddJSONFunctionHandler("/Job/","HandleRequest","","JobsController",newJob)
	server.AddJSONFunctionHandler("/Jobs/","HandleReadAllRequest","GET","JobsController",newJob)

	

	newTask  := models.Task{}
	
	server.AddJSONFunctionHandler("/Task/{Id}/","HandleRequest","GET","TasksController",newTask)
	server.AddJSONFunctionHandler("/Task/","HandleRequest","","TasksController",newTask)
	server.AddJSONFunctionHandler("/Tasks/","HandleReadAllRequest","GET","TasksController",newTask)

	

	newJobHasTask  := models.JobHasTask{}
	
	server.AddJSONFunctionHandler("/JobHasTask/{Id}/","HandleRequest","GET","JobHasTasksController",newJobHasTask)
	server.AddJSONFunctionHandler("/JobHasTask/","HandleRequest","","JobHasTasksController",newJobHasTask)
	server.AddJSONFunctionHandler("/JobHasTasks/","HandleReadAllRequest","GET","JobHasTasksController",newJobHasTask)

	


	// start Listen Server, this build the mapping and creates Handler/
	// also fires the "http listen and server method"
	server.ListenAndServe()

}


