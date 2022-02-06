var form = document.getElementById("assignmodule");

form.addEventListener("submit", assignTutor); //listen to the form where submit button is clicked

getTutor(); //call the getTutor Function
getModule(); //call the getModule Function

function getmodassign() { //map all module details from html to dropdown
    var dropdown = document.getElementById("modulecode");
    var modulename = document.getElementById("modulename");
    var moduleid = document.getElementById("moduleid");
    var synopsis = document.getElementById("synopsis");
    var learningobjective = document.getElementById("learningobjective");

    var selection = JSON.parse(dropdown.value);

    chosenModuleCode = selection.modulecode;
    modulename.value = selection.modulename;
    moduleid.value = selection.moduleid;
    synopsis.value = selection.synopsis;
    learningobjective.value = selection.learningobjective;
}

function tutorfunc() { //map all tutor details from html to dropdown
    var dropdown = document.getElementById("name");
    var tutor_id = document.getElementById("tutor_id");
    var email = document.getElementById("email");
    var descriptions = document.getElementById("descriptions");

    var selection = JSON.parse(dropdown.value);

    console.log(selection);

    chosenTutorName = selection.name;
    tutor_id.value = selection.tutor_id;
    email.value = selection.email;
    descriptions.value = selection.descriptions;
}

function getModule() { //get module using api
    let dropdown = document.getElementById("modulecode");
    dropdown.length = 0;

    let defaultOption = document.createElement("option");
    defaultOption.text = "Module Code"; //set drop down to Module Code

    dropdown.add(defaultOption);
    dropdown.selectedIndex = 0;

    const url = "http://10.31.11.12:9141/api/v1/modules/"; //Get api of modules

    fetch(url)
        .then(function(response) {
            if (response.status !== 200) {
                console.warn(
                    "Looks like there was a problem. Status Code: " + response.status
                );
                return;
            }

            // Examine the text in the response
            response.json().then(function(data) {
                let option;

                for (let i = 0; i < data.length; i++) { //for loop to stringify the data of api
                    option = document.createElement("option"); //setting the option of dropdown list
                    option.text = data[i].modulecode; //sets each of the text box using modulecode
                    option.value = JSON.stringify(data[i]);

                    dropdown.add(option); //add to dropdown
                }
            });
        })
        .catch(function(err) {
            console.error("Fetch Error -", err);
        });
}


function getTutor() { //get Tutor using api
    let dropdown = document.getElementById("name");
    dropdown.length = 0;

    let defaultOption = document.createElement("option");
    defaultOption.text = "Tutor Name";

    dropdown.add(defaultOption);
    dropdown.selectedIndex = 0;

    const url = "http://10.31.11.12:9181/api/v1/GetAllTutor";

    fetch(url)
        .then(function(response) {
            if (response.status !== 200) {
                console.warn(
                    "Looks like there was a problem. Status Code: " + response.status
                );
                return;
            }

            // Examine the text in the response
            response.json().then(function(data) {
                let option;

                for (let i = 0; i < data.length; i++) { //for loop to stringify the data of api
                    option = document.createElement("option"); //setting the option of dropdown list
                    option.text = data[i].name; //sets each of the text box using modulecode
                    option.value = JSON.stringify(data[i]);

                    dropdown.add(option); //add to dropdown
                }
            });
        })
        .catch(function(err) {
            console.error("Fetch Error -", err);
        });
}

function assignTutor(e) { //get the value of moduletutor table in the backend and pass in the update function
    e.preventDefault();

    var tutor_id = document.getElementById("tutor_id").value;
    var email = document.getElementById("email").value;
    var descriptions = document.getElementById("descriptions").value;
    var moduleid = document.getElementById("moduleid").value;
    var modulecode = JSON.parse(document.getElementById("modulecode").value).modulecode;
    var modulename = document.getElementById("modulename").value;
    var synopsis = document.getElementById("synopsis").value;
    var learningobjective = document.getElementById("learningobjective").value;

    fetch("http://10.31.11.12:9141/api/v1/module/assign", {
        method: "PUT",
        body: JSON.stringify({
            tutorid: parseInt(tutor_id),
            email: email,
            descriptions: descriptions,
            moduleid: parseInt(moduleid),
            modulecode: modulecode,
            modulename: modulename,
            synopsis: synopsis,
            learningobjective: learningobjective,
        }),
        headers: {
            "Content-Type": "application/json; charset=UTF-8",
        },
    }).then(function(response) {
        if (response.status == 428) {
            alert("Please fill in module details");
            return;
        } else if (response.status == 422) {
            alert("Module Existed!");
            return;
        } else {
            alert("Module Assigned!");
            return;
        }
    });
}