createModule();

function createModule() {
    var form = document.getElementById("createNewModule"); //post using module create api

    form.addEventListener("submit", function(e) {
        e.preventDefault();

        var modulecode = document.getElementById("modulecodes").value;
        var modulename = document.getElementById("modulename").value;
        var synopsis = document.getElementById("synopsis").value;
        var learningobjective = document.getElementById("learningobjective").value;

        fetch("http://localhost:9141/api/v1/module/create", {
            method: "POST",
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
            if (response.status == 428) {
                alert("Please fill in module details");
                return;
            } else if (response.status == 422) {
                alert("Module Existed!");
                return;
            } else {
                alert("Module Created!");
                return;
            }
        });
    });
}