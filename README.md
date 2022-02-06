<div id="top"></div>


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#Assignment">Assignment Objective</a>
      <ul>
        <li><a href="#AssignmentRequirements">Assignment Requirements</a></li>
        <li><a href="#AssignmentObjectives">Assignment Objectives</a></li>
        <li><a href="#DesignConsiderationsforthemicroservices">Design Considerations for the microservices</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

EduFi Module Management, Manages EduFi module system, allows admin to create, update, delete, view modules
and assign modules to tutors.

<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

* [GOLANG](https://go.dev/)
* [GORM](https://gorm.io/index.html)
* [MYSQL](https://www.mysql.com/)
* [JAVASCRIPT](https://www.javascript.com/)
* [NGIX](https://www.nginx.com/)
<p align="right">(<a href="#top">back to top</a>)</p>



<!-- Assignment Objective-->
## Assignment Objectives
* To demonstrate ability to implement microservices with REST APIs in containers
* To demonstrate ability to work with multiple teams in designing and implementing microservice architecture


## Design Considerations for the microservices
Module Microservice allows anyone that requires the usage of module to be able to get all information regarding the module and also assigned tutor.
To ensure that the front-end is capable of displaying any necessary data that may be linked to other microservices. The back-end microservice would be linked Tutor microservice, in order to assign available tutor to created module. As for the backend which allows other microservices to heavily rely on module microservices, multiple function is created to allowed the display of data required by certain microservices.

Gorm, is an object-relational mapping (ORM) library for dealing with relational databases. The database/sql package is used to build this gorm library. an example would be instead of using query when excuting a SQL line, instead using GORM it simplifies the execution and insertion of data in the table. Another example of using GORM is simplifying the database creation, if the table does not exist, GORM can also be used for the initial migration,creation of the database table is automatically created upon launching the API thus making database migration easier.





<img src="/Architecture.png" alt="Logo" width="1080" height="720">

For the Module Microservice, there are 2 different table used to create modules, and another is to assign the module and finally the Front End to navigate,
The rest API communicates with the used of HTTP GET POST PUT methods, such as creating the new Modules, it will issue a POST request and from there the information that is inputted will be send to the module table for storing and Front End to view all created modules, this also applies to assigned modules to tutor, delete modules, update modules,
all while adhering to the loosely coupled philosophy that Microservices is known for.

The Module Microservice consist of 

Module Microservice
* Create Module (POST)
* Update Module Details (PUT)
* Delete Module(Delete)
* Get all Module by calling the Module Microservice (GET)
* Assign Modules to Tutor (PUT)


Module FrontEnd
* Display Modules
* Create Modules
* Update Modules
* Search for modules
* Assign Modules

Database Tables
* Modules - Used to store Modules Information
* ModulesTutor - Used to store and update modules that are assigned to the tutors


### Prerequisites

GOLANG and MYSQL must be installed in order for the program to work

1. SQL information
  ```sh
  Username : Root
  Password : Root
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/TheAzziIzzuan/Assignment2_ModuleManagement.git
   ```
2. Install libraries
   ```sh
    go get -u github.com/go-sql-driver/mysql
    go get -u github.com/gorilla/mux
    go get -u github.com/gorilla/handlers
    go get -u gorm.io/gorm
   ```
3. Database is not required to execute as gorm will automatically create the tables and database.



**4. Docker Hub Links**

- [Module Microservice Container](https://hub.docker.com/repository/docker/azziizzuan/assignment2_backendmodulemanagementcontainer)
- [Module Frontend Container](https://hub.docker.com/repository/docker/azziizzuan/assignment2_frontendmodulemanagementcontainer)
- [Module Database](https://hub.docker.com/repository/docker/azziizzuan/assignment2_databasemodulemanagementcontainer)

   
    
<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

1. Start the Microservices
   ```sh
   cd ETIAssignment2-Docker\modules
   go run main.go
   ```

   
2. Start the front end by clicking on the file in the folder /FrontEnd
   ```sh
    In the Folder include this html files
    index.html 
    AssignModule.html 
    CreateModule.html 
    GetModule.html
    UpdateModule.html
   ```
<p align="right">(<a href="#top">back to top</a>)</p>

## Endpoints:
| Microservice  | Port | Endpoint URL |
| ------------- | ---- | ------------ |
| Module Frontend  | 9140 | http://10.31.11.12:9140 |
| Module API  | 9141 | http://10.31.11.12:9141 |

<p align="right">(<a href="#top">back to top</a>)</p>

## Module Data Structure
| Attribute | Data Type |
| --------- | ---- |
| ModuleID | int |
| ModuleCode | string |
| ModuleName | string |
| Synopsis | string |
| LearningObjective | string |
| Deleted | gorm.DeletedAt |

### Rating Data Structure
| Attribute | Data Type |
| --------- | ---- |
| TutorID | int |
| Name | string) |
| Email | string |
| Descriptions | string |
| ModuleID | int |
| ModuleCode | string(500) |
| ModuleName | string |
| Synopsis | string |
| LearningObjective | string |
| Deleted | gorm.DeletedAt |



<p align="right">(<a href="#top">back to top</a>)</p>

# Module Microservice API Documentation
### [GET] /api/v1/modules/
Get all Module

Endpoint
http://10.31.11.12:9141/api/v1/modules/
| Name  | Type | Required| Description |
| ------| ---- | ------- | ----------- |
|modulecode|string|notrequired|Module Code Eg. CM, ADB, PRG1|
|modulename|string|notrequired|An module name that matched with module code|
|synopis|string|notrequired|An brief summary of module|
|learningobjective|string|notrequired|An summary of learning objective|

Example Request
```
cURL:
curl  --request GET 'http://localhost:9141/api/v1/modules/' or 'http://10.31.11.12:9140/api/v1/modules/'
Response:
The response will be a status code 200 is successful, return with all module that have been created
```


### [POST] /api/v1/module/create
Create Modules

Endpoint
http://10.31.11.12:9141/api/v1/modules/
| Name  | Type | Required| Description |
| ------| ---- | ------- | ----------- |
|modulecode|string|required|Module Code Eg. CM, ADB, PRG1|
|modulename|string|required|An module name that matched with module code|
|synopis|string|required|An brief summary of module|
|learningobjective|string|required|An summary of learning objective|

Example Request
```
cURL:
curl --location --request POST 'http://localhost:9141/api/v1/module/create' \
--header 'Content-Type: application/json' \
--data '{
    "modulecode": "CM",
    "modulename": "COMPUTING MATH",
    "synopis" : "Learn MATH",
    "learningobjective" : "UNION INTERSEC"
}'

Response
The response will be a status code 200 is successful, or an 422 Unprocessable Entity if the same module has already existed
```


### [GET] /api/v1/module/{modulecode}
Get Individual Module

Endpoint
http://10.31.11.12:9141/api/v1/modules/
| Name  | Type | Required| Description |
| ------| ---- | ------- | ----------- |
|ModuleID|int|notrequired|a primary key with auto increment and generated by gorm|
|modulecode|string|notrequired|Module Code Eg. CM, ADB, PRG1|
|modulename|string|notrequired|An module name that matched with module code|
|synopis|string|notrequired|An brief summary of module|
|learningobjective|string|notrequired|An summary of learning objective|
|Deleted|gorm.DeletedAt|notrequired|gorm soft delete timing|

Example Request
```
cURL
curl --location --request GET 'http://localhost:9141/api/v1/module/CM

Response
The response will be a status code 200 is successful, or an error code 422 with a corresponding status message if unsuccessful.
```


### [PUT] api/v1/module/assign
This endpoint is used to assign module to tutor

Endpoint
http://10.31.11.12:9141/api/v1/modules/
| Name  | Type | Required| Description |
| ------| ---- | ------- | ----------- |
|tutorid|int|required|a primary key with auto increment and generated by gorm|
|Name|string|required|Tutor Name|
|email|string|required|Tutor Email|
|descriptions|string|required|Tutor Description|
|moduleid|int|required|a primary key with auto increment and generated by gorm|
|modulecode|string|required|module code "CM" "ABD"|
|modulename|string|required|module name|
|synopis|string|required|Module Synopis|
|learnignobjective|string|required|Module Learning Objective|
|Deleted|gorm.DeletedAt|notrequired|gorm soft delete timing|

Example Request
```
cURL
curl --location --request PUT 'http://localhost:9141/api/v1/module/assign' \
--header 'Content-Type: application/json' \
--data-raw '{
    "tutorid": 90,
    "Name": "JunZhi",
    "email": "JunZhi@gmail.com",
    "descriptions": "Test",
    "moduleid": 1,
    "modulecode": "CM",
    "modulename": "CM",
    "synopsis": "CM",
    "learningobjective": "CM",
    "Deleted": null
}'

Response
The response will be a status code 200 is successful.
```

### [PUT] /api/v1/module/change/{modulecode}
update module by module codes

Endpoint
http://10.31.11.12:9141/api/v1/module/change/{modulecode}
| Name  | Type | Required| Description |
| ------| ---- | ------- | ----------- |
|modulecode|string|notrequired|Module Code Eg. CM, ADB, PRG1|
|modulename|string|notrequired|An module name that matched with module code|
|synopis|string|notrequired|An brief summary of module|
|learningobjective|string|notrequired|An summary of learning objective|

Example Request
```
cURL
curl --location --request PUT 'http://localhost:9141/api/v1/module/change/CMs' \
--header 'Content-Type: application/json' \
--data-raw '{
    "modulecode": "CM",
    "modulename": "Computing Math",
    "synopis": "Study English",
    "learningobjective": "Union Intersection Cut"
}'

Response
The response will be a status code 200 is successful
```


### [GET] /api/v1/module/tutor/{modulecode}
List all of the specific tutor by modulecode

Endpoint
http://10.31.11.12:9141/api/v1/module/tutor/{modulecode}
| Name  | Type | Required| Description |
| ------| ---- | ------- | ----------- |
|tutorid|int|notrequired|a primary key with auto increment and generated by gorm|
|Name|string|notrequired|Tutor Name|
|email|string|notrequired|Tutor Email|
|descriptions|string|notrequired|Tutor Description|
|moduleid|int|notrequired|a primary key with auto increment and generated by gorm|
|modulecode|string|notrequired|module code "CM" "ABD"|
|modulename|string|notrequired|module name|
|synopis|string|notrequired|Module Synopis|
|learnignobjective|string|notrequired|Module Learning Objective|
|Deleted|gorm.DeletedAt|notrequired|gorm soft delete timing|

Example Request
```
cURL
curl --location --request GET 'http://localhost:9141/api/v1/module/tutor/{modulecode}'

Response
The response will be a status code 200 is successful
```


### [GET] /api/v1/module/alltutor/{tutor_id}
List all of the specific tutor by tutor id

Endpoint
http://10.31.11.12:9141/api/v1/module/alltutor/{tutor_id}
| Name  | Type | Required| Description |
| ------| ---- | ------- | ----------- |
|tutorid|int|notrequired|a primary key with auto increment and generated by gorm|
|Name|string|notrequired|Tutor Name|
|email|string|notrequired|Tutor Email|
|descriptions|string|notrequired|Tutor Description|
|moduleid|int|notrequired|a primary key with auto increment and generated by gorm|
|modulecode|string|notrequired|module code "CM" "ABD"|
|modulename|string|notrequired|module name|
|synopis|string|notrequired|Module Synopis|
|learnignobjective|string|notrequired|Module Learning Objective|
|Deleted|gorm.DeletedAt|notrequired|gorm soft delete timing|

Example Request
```
cURL
curl --location --request GET 'http://localhost:9141/api/v1/module/alltutor/{tutorid}'

Response
The response will be a status code 200 is successful
```


### [GET] /api/v1/module/alltutorname/{name}
List all of the specific tutor by tutor name

Endpoint
http://10.31.11.12:9141/api/v1/module/alltutorname/{name}
| Name  | Type | Required| Description |
| ------| ---- | ------- | ----------- |
|tutorid|int|notrequired|a primary key with auto increment and generated by gorm|
|Name|string|notrequired|Tutor Name|
|email|string|notrequired|Tutor Email|
|descriptions|string|notrequired|Tutor Description|
|moduleid|int|notrequired|a primary key with auto increment and generated by gorm|
|modulecode|string|notrequired|module code "CM" "ABD"|
|modulename|string|notrequired|module name|
|synopis|string|notrequired|Module Synopis|
|learnignobjective|string|notrequired|Module Learning Objective|
|Deleted|gorm.DeletedAt|notrequired|gorm soft delete timing|

Example Request
```
cURL
curl --location --request GET 'http://localhost:9141/api/v1/module/alltutorname/{name}'

Response
The response will be a status code 200 is successful
```


### [DELETE] /api/v1/module/delete/{modulecode}
Delete module by modulecode

Endpoint
http://10.31.11.12:9141/api/v1/module/delete/{modulecode}
| Name  | Type | Required| Description |
| ------| ---- | ------- | ----------- |
|modulecode|string|notrequired|Module Code Eg. CM, ADB, PRG1|
|modulename|string|notrequired|An module name that matched with module code|
|synopis|string|notrequired|An brief summary of module|
|learningobjective|string|notrequired|An summary of learning objective|

Example Request
```
cURL
curl --location --request DELETE 'http://localhost:9141/api/v1/module/delete/{modulecode}'

Response
The response will be a status code 200 is successful
```


### [DELETE] /api/v1/module/deleteassignedtutor/{email}
Delete assigned tutors in the moduletutor table

Endpoint
http://10.31.11.12:9141/api/v1/module/deleteassignedtutor/{email}
| Name  | Type | Required| Description |
| ------| ---- | ------- | ----------- |
|tutorid|int|notrequired|a primary key with auto increment and generated by gorm|
|Name|string|notrequired|Tutor Name|
|email|string|notrequired|Tutor Email|
|descriptions|string|notrequired|Tutor Description|
|moduleid|int|notrequired|a primary key with auto increment and generated by gorm|
|modulecode|string|notrequired|module code "CM" "ABD"|
|modulename|string|notrequired|module name|
|synopis|string|notrequired|Module Synopis|
|learnignobjective|string|notrequired|Module Learning Objective|
|Deleted|gorm.DeletedAt|notrequired|gorm soft delete timing|

Example Request
```
cURL
curl --location -g --request DELETE 'http://localhost:9141/api/v1/module/deleteassignedtutor/{email}'

Response
The response will be a status code 200 is successful
```


### [PUT] /api/v1/module/tutor/updateassignedtutor/{email}
Update assigned tutors in the moduletutor table

Endpoint
http://10.31.11.12:9141/api/v1/module/tutor/updateassignedtutor/{email}
| Name  | Type | Required| Description |
| ------| ---- | ------- | ----------- |
|tutorid|int|required|a primary key with auto increment and generated by gorm|
|Name|string|required|Tutor Name|
|email|string|required|Tutor Email|
|descriptions|string|required|Tutor Description|
|moduleid|int|required|a primary key with auto increment and generated by gorm|
|modulecode|string|required|module code "CM" "ABD"|
|modulename|string|required|module name|
|synopis|string|required|Module Synopis|
|learnignobjective|string|required|Module Learning Objective|
|Deleted|gorm.DeletedAt|notrequired|gorm soft delete timing|

Example Request
```
cURL
cURL
curl --location --request PUT 'http://localhost:9141/api/v1/module/tutor/updateassignedtutor/{email}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "tutorid": ,
    "Name": "",
    "email": "",
    "descriptions": "",
    "moduleid": ,
    "modulecode": "",
    "modulename": "",
    "synopsis": "",
    "learningobjective": "",
    "Deleted": null
}'

Response
The response will be a status code 200 is successful
```



<p align="right">(<a href="#top">back to top</a>)</p>


<!-- CONTACT -->
## Contact
School Email
```sh
S10189579@connect.np.edu.sg
```

<p align="right">(<a href="#top">back to top</a>)</p>




<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/github_username/repo_name.svg?style=for-the-badge
[contributors-url]: https://github.com/github_username/repo_name/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/github_username/repo_name.svg?style=for-the-badge
[forks-url]: https://github.com/github_username/repo_name/network/members
[stars-shield]: https://img.shields.io/github/stars/github_username/repo_name.svg?style=for-the-badge
[stars-url]: https://github.com/github_username/repo_name/stargazers
[issues-shield]: https://img.shields.io/github/issues/github_username/repo_name.svg?style=for-the-badge
[issues-url]: https://github.com/github_username/repo_name/issues
[license-shield]: https://img.shields.io/github/license/github_username/repo_name.svg?style=for-the-badge
[license-url]: https://github.com/github_username/repo_name/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: images/screenshot.png
