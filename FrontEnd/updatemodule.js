updateModule(); //call updatemodule function
getModule(); //call getmodule function
deleteModule(); //call deletemodul function

function updateModule() { //update using update module api
    var form = document.getElementById("update");

    form.addEventListener("click", function(e) {
        e.preventDefault();

        var modulecode = JSON.parse(
            document.getElementById("modulecode").value
        ).modulecode;
        var modulename = document.getElementById("modulename").value;
        var synopsis = document.getElementById("synopsis").value;
        var learningobjective = document.getElementById("learningobjective").value;

        fetch("http://localhost:9141/api/v1/module/change/" + chosenModuleCode, {
            method: "PUT",
            body: JSON.stringify({
                modulecode: modulecode,
                modulename: modulename,
                synopsis: synopsis,
                learningobjective: learningobjective,
            }),
            headers: {
                "Content-Type": "application/json; charset=UTF-8",
            },
        }).then(function(response) {
            if (response.status != 202) {
                alert("Error updating!");
                return;
            } else {

                alert("Update Success!");
                return;
            }
        });
    });
}

function getmodfunc() { //upon fetching, sets the dropdown to the api of getmodules
    var dropdown = document.getElementById("modulecode");
    var modulename = document.getElementById("modulename");
    var synopsis = document.getElementById("synopsis");
    var learningobjective = document.getElementById("learningobjective");

    var selection = JSON.parse(dropdown.value);

    console.log(selection);

    chosenModuleCode = selection.modulecode;
    modulename.value = selection.modulename;
    synopsis.value = selection.synopsis;
    learningobjective.value = selection.learningobjective;
}

function getModule() { //uses getallmodules api 
    let dropdown = document.getElementById("modulecode");
    dropdown.length = 0;

    let defaultOption = document.createElement("option");
    defaultOption.text = "Module Code";

    dropdown.add(defaultOption);
    dropdown.selectedIndex = 0;

    const url = "http://localhost:9141/api/v1/modules/";

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

                for (let i = 0; i < data.length; i++) {
                    option = document.createElement("option");
                    option.text = data[i].modulecode;
                    option.value = JSON.stringify(data[i]);

                    dropdown.add(option);
                }
            });
        })
        .catch(function(err) {
            console.error("Fetch Error -", err);
        });
}

function deleteModule() {
    var form = document.getElementById("delete");

    form.addEventListener("click", function(e) {
        e.preventDefault();

        var modulecode = JSON.parse(
            document.getElementById("modulecode").value
        ).modulecode;
        var modulename = document.getElementById("modulename").value;
        var synopsis = document.getElementById("synopsis").value;
        var learningobjective = document.getElementById("learningobjective").value;

        fetch("http://localhost:9141/api/v1/module/delete/" + chosenModuleCode, {
            method: "DELETE",
            body: JSON.stringify({
                modulecode: modulecode,
                modulename: modulename,
                synopsis: synopsis,
                learningobjective: learningobjective,
            }),
            headers: {
                "Content-Type": "application/json; charset=UTF-8",
            },
        }).then(function(response) {
            if (response.status != 202) {
                alert("Error Deleting!");
                return;
            } else {
                alert("Delete Success!");
                fetch(`http://localhost:9101/api/v1/class?ModuleCode=${chosenModuleCode}&key=2c78afaf-97da-4816-bbee-9ad239abb296`, {
                    method: "DELETE",
                    body: JSON.stringify({
                        modulecode: modulecode,
                        modulename: modulename,
                        synopsis: synopsis,
                        learningobjective: learningobjective,
                    }),
                    headers: {
                        "Content-Type": "application/json; charset=UTF-8",
                    },
                })
            }
        })
    })
}