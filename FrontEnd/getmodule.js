displayAllModuleAPI();
searchModule();

function searchModule() { //sets the row of each module information to the table
    // Declare variables
    var input, filter, table, tr, th, i, txtValue;
    input = document.getElementById("myInput");
    filter = input.value.toUpperCase();
    table = document.getElementById("tableBody");
    tr = table.getElementsByTagName("tr");

    // Loop through all table rows, and hide those who don't match the search query
    for (i = 0; i < tr.length; i++) {
        th = tr[i].getElementsByTagName("th")[0];
        if (th) {
            txtValue = th.textContent || th.innerText;
            if (txtValue.toUpperCase().indexOf(filter) > -1) {
                tr[i].style.display = "";
            } else {
                tr[i].style.display = "none";
            }
        }
    }
}

async function displayAllModuleAPI() { //display using get all modules part
    const response = await fetch("http://10.31.11.12:9141/api/v1/modules/");
    data = await response.json();

    const tableData = data
        .map(function(value) {
            return `<tr>
              <th scope="row">${value.modulecode}</th>
              <th scope="row">${value.modulename}</th>
              <th scope="row">${value.synopsis}</th>
              <th scope="row">${value.learningobjective}</th>
          </tr>`;
        })
        .join("");

    const tableBody = document.querySelector("#tableBody");
    tableBody.innerHTML = tableData;
}