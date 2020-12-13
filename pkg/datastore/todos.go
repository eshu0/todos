package datastore

import (
	SQLL "github.com/eshu0/persist/pkg/sqllite"

	// these are from the package
	"github.com/eshu0/todos/pkg/handlers"
)

type TodosDatastore struct {
	Datastore *SQLL.SQLLiteDatastore

	//
	//ProjectHasJobsHandler *ProjectHasJobsHandler
	//
	//ProjectsHandler *ProjectsHandler
	//
	//JobsHandler *JobsHandler
	//
	//TasksHandler *TasksHandler
	//
	//JobHasTasksHandler *JobHasTasksHandler
	//
}

func CreateDataStorage(filename string) *TodosDatastore {
	res := TodosDatastore{}

	ds := SQLL.CreateOpenSQLLiteDatastore(filename)
	
	// tests the example
	ds.SetStorageHander("Generic",SQLL.NewSQLLiteTableHandler(ds)) 
	
	ds.SetStorageHander("ProjectHasJobs",hndlr.NewProjectHasJobsHandler(ds))
	
	ds.SetStorageHander("Projects",hndlr.NewProjectsHandler(ds))
	
	ds.SetStorageHander("Jobs",hndlr.NewJobsHandler(ds))
	
	ds.SetStorageHander("Tasks",hndlr.NewTasksHandler(ds))
	
	ds.SetStorageHander("JobHasTasks",hndlr.NewJobHasTasksHandler(ds))
	


	ds.CreateStructures()

	res.Datastore = ds
	
	return &res
}


func (fds *TodosDatastore) GetProjectHasJobsHandler() *hndlr.ProjectHasJobsHandler {

	data, ok := fds.Datastore.GetStorageHandler("ProjectHasJobs")
	if ok {
	  res := data.(*hndlr.ProjectHasJobsHandler)
	  return res
	}
	return nil
}

func (fds *TodosDatastore) GetProjectsHandler() *hndlr.ProjectsHandler {

	data, ok := fds.Datastore.GetStorageHandler("Projects")
	if ok {
	  res := data.(*hndlr.ProjectsHandler)
	  return res
	}
	return nil
}

func (fds *TodosDatastore) GetJobsHandler() *hndlr.JobsHandler {

	data, ok := fds.Datastore.GetStorageHandler("Jobs")
	if ok {
	  res := data.(*hndlr.JobsHandler)
	  return res
	}
	return nil
}

func (fds *TodosDatastore) GetTasksHandler() *hndlr.TasksHandler {

	data, ok := fds.Datastore.GetStorageHandler("Tasks")
	if ok {
	  res := data.(*hndlr.TasksHandler)
	  return res
	}
	return nil
}

func (fds *TodosDatastore) GetJobHasTasksHandler() *hndlr.JobHasTasksHandler {

	data, ok := fds.Datastore.GetStorageHandler("JobHasTasks")
	if ok {
	  res := data.(*hndlr.JobHasTasksHandler)
	  return res
	}
	return nil
}



