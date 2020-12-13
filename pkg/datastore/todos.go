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
	
	ds.SetStorageHander("ProjectHasJobs",handlers.NewProjectHasJobsHandler(ds))
	
	ds.SetStorageHander("Projects",handlers.NewProjectsHandler(ds))
	
	ds.SetStorageHander("Jobs",handlers.NewJobsHandler(ds))
	
	ds.SetStorageHander("Tasks",handlers.NewTasksHandler(ds))
	
	ds.SetStorageHander("JobHasTasks",handlers.NewJobHasTasksHandler(ds))
	


	ds.CreateStructures()

	res.Datastore = ds
	
	return &res
}


func (fds *TodosDatastore) GetProjectHasJobsHandler() *handlers.ProjectHasJobsHandler {

	data, ok := fds.Datastore.GetStorageHandler("ProjectHasJobs")
	if ok {
	  res := data.(*handlers.ProjectHasJobsHandler)
	  return res
	}
	return nil
}

func (fds *TodosDatastore) GetProjectsHandler() *handlers.ProjectsHandler {

	data, ok := fds.Datastore.GetStorageHandler("Projects")
	if ok {
	  res := data.(*handlers.ProjectsHandler)
	  return res
	}
	return nil
}

func (fds *TodosDatastore) GetJobsHandler() *handlers.JobsHandler {

	data, ok := fds.Datastore.GetStorageHandler("Jobs")
	if ok {
	  res := data.(*handlers.JobsHandler)
	  return res
	}
	return nil
}

func (fds *TodosDatastore) GetTasksHandler() *handlers.TasksHandler {

	data, ok := fds.Datastore.GetStorageHandler("Tasks")
	if ok {
	  res := data.(*handlers.TasksHandler)
	  return res
	}
	return nil
}

func (fds *TodosDatastore) GetJobHasTasksHandler() *handlers.JobHasTasksHandler {

	data, ok := fds.Datastore.GetStorageHandler("JobHasTasks")
	if ok {
	  res := data.(*handlers.JobHasTasksHandler)
	  return res
	}
	return nil
}



