package pguhandlers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	SQLL "github.com/eshu0/persist/pkg/sqllite"	
	per "github.com/eshu0/persist/pkg/interfaces"

	// these are from the package
	data "github.com/esh0/todos/pkg/Models"
)

//
// Built from:
// main - Todos.Db
/*
 CREATE TABLE Tasks (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, displayname TEXT NOT NULL, archived INTEGER DEFAULT (0) NOT NULL, completed INTEGER DEFAULT (0) NOT NULL)
 */
//

// Table fields

// Tasks
const tasksTName = "Tasks"

// Primay Key: id
const tasksIdCName = "id"


// displayname
const tasksDisplaynameCName = "displayname"

// archived
const tasksArchivedCName = "archived"

// completed
const tasksCompletedCName = "completed"



// HANDLER

type TasksHandler struct {
	per.IStorageHandler
	Parent *SQLL.SQLLiteDatastore
	Executor *SQLL.SQLLightQueryExecutor
}

func NewTasksHandler(datastore *SQLL.SQLLiteDatastore) *TasksHandler {
	ds := TasksHandler{}
	ds.Parent = datastore
	ds.Executor = SQLL.NewSQLLightQueryExecutor(datastore)
	return &ds
}

// Start IStorage Handler 
func (handler *TasksHandler) GetPersistantStorage() per.IPersistantStorage {
	return handler.Parent
}

func (handler *TasksHandler) SetPersistantStorage(persistant per.IPersistantStorage) { // per.IStorageHandler {
	res := persistant.(*SQLL.SQLLiteDatastore)
	handler.Parent = res
	//return handler
}

// This function creates the database table for Task 
func (handler *TasksHandler) CreateStructures() per.IQueryResult {
	handler.Parent.LogDebug("CreateStructures","Executing Query")
	return handler.Executor.ExecuteQuery(`CREATE TABLE IF NOT EXISTS Tasks (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, displayname TEXT NOT NULL, archived INTEGER DEFAULT (0) NOT NULL, completed INTEGER DEFAULT (0) NOT NULL)`)
}

// End Istorage 

// This function Task removes all data for the table
func (handler *TasksHandler) Wipe() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteQuery("DELETE FROM " + tasksTName))
}

// This adds Task to the database 
func (handler *TasksHandler) Create(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.Task)
	return handler.ConvertResult(handler.Executor.ExecuteInsertQuery("INSERT INTO " + tasksTName + " ( "+ "["+tasksDisplaynameCName+"]" +  ",["+tasksArchivedCName+"]" + ",["+tasksCompletedCName+"]" +" ) VALUES (?,?,?)", data.Displayname,data.Archived,data.Completed))
}

func (handler *TasksHandler) Update(Data per.IDataItem) SQLL.SQLLiteQueryResult  {
	data := Data.(data.Task)
	return handler.ConvertResult(handler.Executor.ExecuteQuery("UPDATE " + tasksTName + " SET "+ "["+tasksDisplaynameCName+"] = ? " +  ",["+tasksArchivedCName+"] = ? " + ",["+tasksCompletedCName+"] = ? " +"  WHERE [" + tasksIdCName + "] = ?",data.Displayname,data.Archived,data.Completed,data.Id))
}

func (handler *TasksHandler) ConvertResult(data per.IQueryResult) SQLL.SQLLiteQueryResult {
	// this needs to be implemented
	return  SQLL.ResultToSQLLiteQueryResult(data)
}








func (handler *TasksHandler) FindById(SearchData int64)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+tasksIdCName+"]," + "["+tasksDisplaynameCName+"]" +  ",["+tasksArchivedCName+"]" + ",["+tasksCompletedCName+"]" +"  FROM " + tasksTName + " WHERE " + tasksIdCName + " = ?",handler.ParseRows,SearchData))
}


func (handler *TasksHandler) FindByDisplayname(SearchData string)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+tasksIdCName+"]," + "["+tasksDisplaynameCName+"]" +  ",["+tasksArchivedCName+"]" + ",["+tasksCompletedCName+"]" +"  FROM " + tasksTName + " WHERE " + tasksDisplaynameCName + " = ?",handler.ParseRows,SearchData))
}


func (handler *TasksHandler) FindByArchived(SearchData int64)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+tasksIdCName+"]," + "["+tasksDisplaynameCName+"]" +  ",["+tasksArchivedCName+"]" + ",["+tasksCompletedCName+"]" +"  FROM " + tasksTName + " WHERE " + tasksArchivedCName + " = ?",handler.ParseRows,SearchData))
}


func (handler *TasksHandler) FindByCompleted(SearchData int64)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+tasksIdCName+"]," + "["+tasksDisplaynameCName+"]" +  ",["+tasksArchivedCName+"]" + ",["+tasksCompletedCName+"]" +"  FROM " + tasksTName + " WHERE " + tasksCompletedCName + " = ?",handler.ParseRows,SearchData))
}




func (handler *TasksHandler) ReadAll()  SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+tasksIdCName+"]," + "["+tasksDisplaynameCName+"]" +  ",["+tasksArchivedCName+"]" + ",["+tasksCompletedCName+"]" +"  FROM " + tasksTName, handler.ParseRows))
}

func (handler *TasksHandler) ParseRows(rows *sql.Rows) per.IQueryResult {
	
	var Id *int64
	
	var Displayname *string
	
	var Archived *int64
	
	var Completed *int64
	
	results := []per.IDataItem{} //Task{}

	for rows.Next() {
		err := rows.Scan(&Id,&Displayname,&Archived,&Completed)
		//fmt.Println("READ: id: " + string(id) + "- Displayname:"+  displayname + "- Description:" + description)
		if err != nil {
			handler.Parent.LogErrorEf("ParseRows","Row Scan errr: %s ",err)
		} else {
			res := data.Task{}
			
				if Id != nil {
					res.Id = *Id
					handler.Parent.LogDebugf("ParseRows","Set '%v' for Id",*Id)
				} else {
					handler.Parent.LogDebugf("ParseRows","{.Name}} was NULL")
				}
			
				if Displayname != nil {
					res.Displayname = *Displayname
					handler.Parent.LogDebugf("ParseRows","Set '%v' for Displayname",*Displayname)
				} else {
					handler.Parent.LogDebugf("ParseRows","{.Name}} was NULL")
				}
			
				if Archived != nil {
					res.Archived = *Archived
					handler.Parent.LogDebugf("ParseRows","Set '%v' for Archived",*Archived)
				} else {
					handler.Parent.LogDebugf("ParseRows","{.Name}} was NULL")
				}
			
				if Completed != nil {
					res.Completed = *Completed
					handler.Parent.LogDebugf("ParseRows","Set '%v' for Completed",*Completed)
				} else {
					handler.Parent.LogDebugf("ParseRows","{.Name}} was NULL")
				}
			

			results = append(results, res)
		}

	}
	return SQLL.NewDataQueryResult(true,results)
}
