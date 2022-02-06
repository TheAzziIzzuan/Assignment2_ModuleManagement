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
| Frontend  | 9030 | http://10.31.11.12:9030 |
| Tutor API  | 9091 | http://10.31.11.12:9031 |
| Testing API  | 9092 | http://10.31.11.12:9042 |


# Tutor Microservice API Documentation
### [GET] /api/v1/tutor
Test API if working
```
Endpoint
http://10.31.11.12:9031/api/v1/tutor
Response 
Status: Tutor API is working
```

### [GET] /api/v1/tutor/profile/{TutorID}
Get tutor by TutorID
```
Endpoint
http://10.31.11.12:9031/api/v1/tutor/profile/{TutorID}
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
Tutor 
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
