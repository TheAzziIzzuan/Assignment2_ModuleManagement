package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

const dsn = "root:root@tcp(db:3306)/assignment2"

func DataMigrations() { //Using Gorm to Connect to database and check if it has been created
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {

		fmt.Println(err.Error())
		panic("Database Unable to connect")
	}
	db.AutoMigrate(&Modules{})
	db.AutoMigrate(&ModuleTutor{})
}

func main() { /*Connect to DB*/
	DataMigrations()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/module/create", CreateModules).Methods("POST")
	router.HandleFunc("/api/v1/modules/", GetAllModules).Methods("GET")
	router.HandleFunc("/api/v1/module/{modulecode}", GetModules).Methods("GET")
	router.HandleFunc("/api/v1/module/assign", AssignModuleToTutor).Methods("PUT")
	router.HandleFunc("/api/v1/module/change/{modulecode}", UpdateModule).Methods("PUT")
	router.HandleFunc("/api/v1/module/tutor/{modulecode}", GetAllTutorByModuleCode).Methods("GET")
	router.HandleFunc("/api/v1/module/delete/{modulecode}", DeleteModule).Methods("DELETE")

	http.ListenAndServe(":9141", router)
	fmt.Println("Listening at port 9141")
	log.Fatal(http.ListenAndServe(":9141", router))
}

type Modules struct {
	ModuleID          int    `gorm:"primaryKey"`
	ModuleCode        string `json:"modulecode"`
	ModuleName        string `json:"modulename"`
	Synopis           string `json:"synopis"`
	LearningObjective string `json:"learningobjective"`
	Deleted           gorm.DeletedAt
}

type ModuleTutor struct {
	TutorID    string `json:"tutorid"`
	ModuleID   int
	ModuleCode string `json:"modulecode"`
	Deleted    gorm.DeletedAt
}

//type Tutor struct {
//Input 3.3 Api Struct//

//Function to Create Modules
func CreateModules(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newmodule Modules
	var moduledb Modules

	if err == nil {
		json.Unmarshal(reqBody, &newmodule)

		if newmodule.ModuleCode == "" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("Please Enter module code"))
			return
		} else if newmodule.ModuleName == "" {
			w.WriteHeader(
				http.StatusUnprocessableEntity)
			w.Write([]byte(
				"Please Enter module name"))
			return
		} else if newmodule.Synopis == "" {
			w.WriteHeader(
				http.StatusUnprocessableEntity)
			w.Write([]byte(
				"Please enter module synopis"))
			return
		} else if newmodule.LearningObjective == "" {
			w.WriteHeader(
				http.StatusUnprocessableEntity)
			w.Write([]byte(
				"Please enter learning objective"))
			return
		}
	}

	err := db.Where("module_code = ?", newmodule.ModuleCode).First(&moduledb).Error
	if err == nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Module Code has been used")
		return
	}

	//If all Validation passed, It will tan create a table or insert into the table if already existed
	db.Create(&newmodule)
	json.NewEncoder(w).Encode(newmodule)
}

//Get Individual Module
func GetModules(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var module Modules
	err := db.Where("module_code = ?", params["modulecode"]).First(&module).Error
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(module)
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "not found")
	}
}

//Get All Modules Created
func GetAllModules(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var modules []Modules
	db.Find(&modules)
	json.NewEncoder(w).Encode(modules)
}

//Assign Module To Tutor
func AssignModuleToTutor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var assignmodule ModuleTutor

	if err == nil {
		json.Unmarshal(reqBody, &assignmodule)

		if assignmodule.TutorID == "" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("Please Enter module code"))
			return
		} else if assignmodule.ModuleCode == "" {
			w.WriteHeader(
				http.StatusUnprocessableEntity)
			w.Write([]byte("Please Enter Tutor ID"))
			return
		}
	}

	//If all Validation passed, It will than create a table or insert into the table if already existed
	db.Create(&assignmodule)
	json.NewEncoder(w).Encode(assignmodule)
}

//Get Tutor by Module Code
func GetAllTutorByModuleCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var modules []ModuleTutor
	err := db.Where("module_code = ?", params["modulecode"]).Find(&modules).Error
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(modules)
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "not found")
	}

}

func UpdateModule(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	var module Modules

	if err != nil {
		fmt.Printf("Module not found")
		return
	} else {
		json.NewDecoder(router.Body).Decode(&module)
		db.Model(&Modules{}).Where("module_code=?", params["modulecode"]).Updates(module)
		db.Model(&ModuleTutor{}).Where("module_code=?", params["modulecode"]).Updates(module)
	}
}

func DeleteModule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var modules Modules
	var moduletutor ModuleTutor
	err := db.Where("module_code = ?", params["modulecode"]).Find(&modules).Error
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		db.Where("module_code=?", params["modulecode"]).Delete(&modules)
		db.Where("module_code=?", params["modulecode"]).Delete(&moduletutor)
		json.NewEncoder(w).Encode(modules)
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "not found")
	}

}
