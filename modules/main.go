package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DataMigrations() { //Using Gorm to Connect to database and check if it has been created
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {

		fmt.Println(err.Error())
		panic("Database Unable to connect")
	}
	db.AutoMigrate(&Modules{})     //Creates the Module Table using Module Struct
	db.AutoMigrate(&ModuleTutor{}) //Creates the ModuleTutor table using ModuleTutor Struct
}

func main() { /*Connect to DB*/
	DataMigrations()
	router := mux.NewRouter()

	// This is to allow the headers, origins and methods all to access CORS resource sharing
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}) //CORS handler
	origins := handlers.AllowedOrigins([]string{"*"})                                                 //CORS handler
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})                      //CORS handler

	router.HandleFunc("/api/v1/module/create", CreateModules).Methods("POST")                              //Create Module
	router.HandleFunc("/api/v1/modules/", GetAllModules).Methods("GET")                                    //Get all Module
	router.HandleFunc("/api/v1/module/{modulecode}", GetModules).Methods("GET")                            //Show module information by module code
	router.HandleFunc("/api/v1/module/assign", AssignModuleToTutor).Methods("PUT")                         //Assign Module to tutor via put
	router.HandleFunc("/api/v1/module/change/{modulecode}", UpdateModule).Methods("PUT")                   //update module by modulecode
	router.HandleFunc("/api/v1/module/tutor/{modulecode}", GetAllTutorByModuleCode).Methods("GET")         //List all tutor that is teaching that specific module
	router.HandleFunc("/api/v1/module/alltutor/{tutor_id}", GetAllTutorModuleByTutorId).Methods("GET")     //List all of the specific tutor by tutor id
	router.HandleFunc("/api/v1/module/alltutorname/{name}", GetAllTutorModuleByTutorName).Methods("GET")   //List all of the specific tutor by tutor name
	router.HandleFunc("/api/v1/module/delete/{modulecode}", DeleteModule).Methods("DELETE")                //Delete module by modulecode
	router.HandleFunc("/api/v1/module/deleteassignedtutor/{email}", DeleteAssignedTutor).Methods("DELETE") // Delete assigned tutors in the moduletutor table

	fmt.Println("Listening at port 9141")
	log.Fatal(http.ListenAndServe(":9141", handlers.CORS(headers, origins, methods)(router))) //CORS handler

}

var db *gorm.DB
var err error

const dsn = "root:root@tcp(db:3306)/assignment2" //Database Connection string
//connection string for docker "root:root@tcp(db:3306)/assignment2"
//connection string for local "root:root@tcp(127.0.0.1:3306)/assignment2?charset=utf8mb4&parseTime=True&loc=Local"

//Module Structure
type Modules struct {
	ModuleID          int    `gorm:"primaryKey" json:"moduleid"`
	ModuleCode        string `json:"modulecode"`
	ModuleName        string `json:"modulename"`
	Synopsis          string `json:"synopsis"`
	LearningObjective string `json:"learningobjective"`
	Deleted           gorm.DeletedAt
}

//ModuleTutor Structure
type ModuleTutor struct {
	TutorID           int    `json:"tutor_id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Descriptions      string `json:"descriptions"`
	ModuleID          int    `json:"moduleid"`
	ModuleCode        string `json:"modulecode"`
	ModuleName        string `json:"modulename"`
	Synopsis          string `json:"synopsis"`
	LearningObjective string `json:"learningobjective"`
	Deleted           gorm.DeletedAt
}

//Function to Create new Modules
func CreateModules(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newmodule Modules //Create a newmodule variable using modules Structure
	var moduledb Modules  //Create a moduledb variable using modules structure

	if err == nil { //if there is not error continue
		json.Unmarshal(reqBody, &newmodule) //unmarshal newmodule

		if newmodule.ModuleCode == "" { //check if modulecode is empty
			w.WriteHeader(http.StatusPreconditionRequired) //set http header
			w.Write([]byte("Please Enter module code"))
			return
		} else if newmodule.ModuleName == "" { //check if modulename is empty
			w.WriteHeader(
				http.StatusPreconditionRequired) //set http header
			w.Write([]byte(
				"Please Enter module name"))
			return
		} else if newmodule.Synopsis == "" { //check if Synopis is empty
			w.WriteHeader(
				http.StatusPreconditionRequired) //set http header
			w.Write([]byte(
				"Please enter module synopsis"))
			return
		} else if newmodule.LearningObjective == "" { //check if LearningObjective is empty
			w.WriteHeader(
				http.StatusPreconditionRequired) //set http header
			w.Write([]byte(
				"Please enter learning objective"))
			return
		}
	}

	err := db.Where("module_code = ?", newmodule.ModuleCode).First(&moduledb).Error //Check if input Modulecode has already existed in Database
	if err == nil {
		w.WriteHeader(http.StatusUnprocessableEntity) //set http header
		fmt.Fprintf(w, "Module Code has been used")
		return
	}

	//If all Validation passed, It will than create a table or insert into the table if already existed
	db.Create(&newmodule)             //create base on newmodules input
	w.WriteHeader(http.StatusCreated) //set http header
	json.NewEncoder(w).Encode(newmodule)

}

//Get Individual Module
func GetModules(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)                                                         //mux parameter
	var module Modules                                                            //Create a module variable using modules Structure
	err := db.Where("module_code = ?", params["modulecode"]).First(&module).Error //Search using the parameter input
	if err == nil {                                                               //if no error
		w.WriteHeader(http.StatusCreated) //set http header
		json.NewEncoder(w).Encode(module) //encode module variable
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity) //set http header
		fmt.Fprintf(w, "not found")                   //return a print line
	}
}

//Get All Modules Created
func GetAllModules(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var modules []Modules              //list of modules
	db.Find(&modules)                  //Find the modules using modules database
	json.NewEncoder(w).Encode(modules) //encode modules
}

//Assign Module To Tutor
func AssignModuleToTutor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var assignmodule ModuleTutor //Create the variable assignmodule using ModuleTutor Struct

	if err == nil {
		json.Unmarshal(reqBody, &assignmodule) //unmarshal assignmodule variable
		fmt.Println(assignmodule)              //check assignmodule

		//if assignmodule.TutorID == ""{
		//	w.WriteHeader(http.StatusUnprocessableEntit)
		//	w.Write([]byte("Please Enter Tutor ID")
		//	return
		if assignmodule.ModuleCode == "" { //if module code is empty
			w.WriteHeader(
				http.StatusUnprocessableEntity) //set http header
			w.Write([]byte("Please Enter Module Code"))
			return
		}
	}

	//If all Validation passed, It will than create a table or insert into the table if already existed
	db.Create(&assignmodule)
	json.NewEncoder(w).Encode(assignmodule)
}

//Get all Tutor that is currently teaching the specific module
func GetAllTutorByModuleCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)                                                         //mux parameter
	var modules []ModuleTutor                                                     //create module variable list using moduletutor struct
	err := db.Where("module_code = ?", params["modulecode"]).Find(&modules).Error //find the input parameter in the database
	if err == nil {                                                               //if success
		w.WriteHeader(http.StatusCreated)  //set http header
		json.NewEncoder(w).Encode(modules) //encode modules variable
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity) //set http header
		fmt.Fprintf(w, "not found")                   //print error line
	}

}

//Get modules that taught by specific tutorid
func GetAllTutorModuleByTutorId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)                                                    //mux parameter
	var modules []ModuleTutor                                                //create modules variable list using moduletutor struct
	err := db.Where("tutor_id = ?", params["tutor_id"]).Find(&modules).Error //find the input parameter in the database
	if err == nil {                                                          //if success
		w.WriteHeader(http.StatusCreated)  //set http header
		json.NewEncoder(w).Encode(modules) //encode modules variable
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity) //set http header
		fmt.Fprintf(w, "not found")                   //print error line
	}
}

func GetAllTutorModuleByTutorName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)                                            //mux parameter
	var modules []ModuleTutor                                        //create modules variable list using moduletutor struct
	err := db.Where("name = ?", params["name"]).Find(&modules).Error //find the input parameter in the database
	if err == nil {                                                  //if success
		w.WriteHeader(http.StatusCreated)  //set http header
		json.NewEncoder(w).Encode(modules) //encode modules variable
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity) //set http header
		fmt.Fprintf(w, "not found")                   //print error line
	}
}

//Update Module information
func UpdateModule(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)  //mux parameter
	var module Modules          //create a module variable using Modules Struct
	var moduletutor ModuleTutor //create moduletutor struct using

	if err != nil { //if error
		fmt.Printf("Module not found")
		return
	} else { //if success
		json.NewDecoder(router.Body).Decode(&module)                                               //decode module variable
		moduletutor.Synopsis = module.Synopsis                                                     //set moduletutor synopsis as module synopsis
		moduletutor.LearningObjective = module.LearningObjective                                   //set moduletutor LearningObjective as module LearningObjective
		moduletutor.ModuleName = module.ModuleName                                                 //set moduletutor ModuleName as module ModuleName
		db.Model(&Modules{}).Where("module_code=?", params["modulecode"]).Updates(module)          //update using the parameter in modules Table
		db.Model(&ModuleTutor{}).Where("module_code=?", params["modulecode"]).Updates(moduletutor) //update using the parameter in moduletutor Table
		w.WriteHeader(http.StatusAccepted)                                                         //set http header
		fmt.Fprintf(w, "Update Successful")
	}
}

func DeleteModule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)                                                         //mux parameter
	var modules Modules                                                           //create a module variable using Modules Struct
	var moduletutor ModuleTutor                                                   //create a module variable using Modules Struct
	err := db.Where("module_code = ?", params["modulecode"]).Find(&modules).Error //Find the specific Module using modulecode
	if err == nil {                                                               //if success
		w.WriteHeader(http.StatusAccepted)                                   //set http header
		db.Where("module_code=?", params["modulecode"]).Delete(&modules)     //delete using the parameter in modules table
		db.Where("module_code=?", params["modulecode"]).Delete(&moduletutor) //delete using the parameter in modules table
		json.NewEncoder(w).Encode(modules)                                   //encode modules
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity) //set http header
		fmt.Fprintf(w, "not found")
	}

}

//Delete Assigned Tutorby email
func DeleteAssignedTutor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)                                                  //set http header
	var moduletutor ModuleTutor                                            //create a module variable using Modules Struct
	err := db.Where("email = ?", params["email"]).Find(&moduletutor).Error //find using the parameter email in the moduletutor table
	if err == nil {
		w.WriteHeader(http.StatusAccepted)                        //set http header
		db.Where("email=?", params["email"]).Delete(&moduletutor) //delete using the parameter
		json.NewEncoder(w).Encode(moduletutor)                    //encode moduletutor variable
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity) //set http header
		fmt.Fprintf(w, "not found")
	}

}
