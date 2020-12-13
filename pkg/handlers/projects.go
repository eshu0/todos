package handlers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	SQLL "github.com/eshu0/persist/pkg/sqllite"	
	per "github.com/eshu0/persist/pkg/interfaces"

	// these are from the package
	"github.com/eshu0/todos/pkg/models"
)

//
// Built from:
// main - Todos.Db
/*
 CREATE TABLE Projects (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, displayname TEXT NOT NULL, description TEXT, archived INTEGER NOT NULL DEFAULT (0), completed INTEGER DEFAULT (0) NOT NULL)
 */
//

// Table fields

// Projects
const projectsTName = "Projects"

// Primay Key: id
const projectsIdCName = "id"


// displayname
const projectsDisplaynameCName = "displayname"

// description
const projectsDescriptionCName = "description"

// archived
const projectsArchivedCName = "archived"

// completed
const projectsCompletedCName = "completed"



// HANDLER

type ProjectsHandler struct {
	per.IStorageHandler
	Parent *SQLL.SQLLiteDatastore
	Executor *SQLL.SQLLightQueryExecutor
}

func NewProjectsHandler(datastore *SQLL.SQLLiteDatastore) *ProjectsHandler {
	ds := ProjectsHandler{}
	ds.Parent = datastore
	ds.Executor = SQLL.NewSQLLightQueryExecutor(datastore)
	return &ds
}

// Start IStorage Handler 
func (handler *ProjectsHandler) GetPersistantStorage() per.IPersistantStorage {
	return handler.Parent
}

func (handler *ProjectsHandler) SetPersistantStorage(persistant per.IPersistantStorage) { // per.IStorageHandler {
	res := persistant.(*SQLL.SQLLiteDatastore)
	handler.Parent = res
	//return handler
}

// This function creates the database table for Project 
func (handler *ProjectsHandler) CreateStructures() per.IQueryResult {
	handler.Parent.LogDebug("CreateStructures","Executing Query")
	return handler.Executor.ExecuteQuery(`CREATE TABLE IF NOT EXISTS Projects (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, displayname TEXT NOT NULL, description TEXT, archived INTEGER NOT NULL DEFAULT (0), completed INTEGER DEFAULT (0) NOT NULL)`)
}

// End Istorage 

// This function Project removes all data for the table
func (handler *ProjectsHandler) Wipe() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteQuery("DELETE FROM " + projectsTName))
}

// This adds Project to the database 
func (handler *ProjectsHandler) Create(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(models.Project)
	return handler.ConvertResult(handler.Executor.ExecuteInsertQuery("INSERT INTO " + projectsTName + " ( "+ "["+projectsDisplaynameCName+"]" +  ",["+projectsDescriptionCName+"]" + ",["+projectsArchivedCName+"]" + ",["+projectsCompletedCName+"]" +" ) VALUES (?,?,?,?)", data.Displayname,data.Description,data.Archived,data.Completed))
}

func (handler *ProjectsHandler) Update(Data per.IDataItem) SQLL.SQLLiteQueryResult  {
	data := Data.(models.Project)
	return handler.ConvertResult(handler.Executor.ExecuteQuery("UPDATE " + projectsTName + " SET "+ "["+projectsDisplaynameCName+"] = ? " +  ",["+projectsDescriptionCName+"] = ? " + ",["+projectsArchivedCName+"] = ? " + ",["+projectsCompletedCName+"] = ? " +"  WHERE [" + projectsIdCName + "] = ?",data.Displayname,data.Description,data.Archived,data.Completed,data.Id))
}

func (handler *ProjectsHandler) ConvertResult(data per.IQueryResult) SQLL.SQLLiteQueryResult {
	// this needs to be implemented
	return  SQLL.ResultToSQLLiteQueryResult(data)
}








func (handler *ProjectsHandler) FindById(SearchData int64)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+projectsIdCName+"]," + "["+projectsDisplaynameCName+"]" +  ",["+projectsDescriptionCName+"]" + ",["+projectsArchivedCName+"]" + ",["+projectsCompletedCName+"]" +"  FROM " + projectsTName + " WHERE " + projectsIdCName + " = ?",handler.ParseRows,SearchData))
}


func (handler *ProjectsHandler) FindByDisplayname(SearchData string)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+projectsIdCName+"]," + "["+projectsDisplaynameCName+"]" +  ",["+projectsDescriptionCName+"]" + ",["+projectsArchivedCName+"]" + ",["+projectsCompletedCName+"]" +"  FROM " + projectsTName + " WHERE " + projectsDisplaynameCName + " = ?",handler.ParseRows,SearchData))
}


func (handler *ProjectsHandler) FindByDescription(SearchData string)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+projectsIdCName+"]," + "["+projectsDisplaynameCName+"]" +  ",["+projectsDescriptionCName+"]" + ",["+projectsArchivedCName+"]" + ",["+projectsCompletedCName+"]" +"  FROM " + projectsTName + " WHERE " + projectsDescriptionCName + " = ?",handler.ParseRows,SearchData))
}


func (handler *ProjectsHandler) FindByArchived(SearchData int64)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+projectsIdCName+"]," + "["+projectsDisplaynameCName+"]" +  ",["+projectsDescriptionCName+"]" + ",["+projectsArchivedCName+"]" + ",["+projectsCompletedCName+"]" +"  FROM " + projectsTName + " WHERE " + projectsArchivedCName + " = ?",handler.ParseRows,SearchData))
}


func (handler *ProjectsHandler) FindByCompleted(SearchData int64)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+projectsIdCName+"]," + "["+projectsDisplaynameCName+"]" +  ",["+projectsDescriptionCName+"]" + ",["+projectsArchivedCName+"]" + ",["+projectsCompletedCName+"]" +"  FROM " + projectsTName + " WHERE " + projectsCompletedCName + " = ?",handler.ParseRows,SearchData))
}




func (handler *ProjectsHandler) ReadAll()  SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+projectsIdCName+"]," + "["+projectsDisplaynameCName+"]" +  ",["+projectsDescriptionCName+"]" + ",["+projectsArchivedCName+"]" + ",["+projectsCompletedCName+"]" +"  FROM " + projectsTName, handler.ParseRows))
}

func (handler *ProjectsHandler) ParseRows(rows *sql.Rows) per.IQueryResult {
	
	var Id *int64
	
	var Displayname *string
	
	var Description *string
	
	var Archived *int64
	
	var Completed *int64
	
	results := []per.IDataItem{} //Project{}

	for rows.Next() {
		err := rows.Scan(&Id,&Displayname,&Description,&Archived,&Completed)
		//fmt.Println("READ: id: " + string(id) + "- Displayname:"+  displayname + "- Description:" + description)
		if err != nil {
			handler.Parent.LogErrorEf("ParseRows","Row Scan errr: %s ",err)
		} else {
			res := models.Project{}
			
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
			
				if Description != nil {
					res.Description = *Description
					handler.Parent.LogDebugf("ParseRows","Set '%v' for Description",*Description)
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
