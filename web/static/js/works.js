async function sendRequest(method, url, body) {
    let response = await fetch(url, {
        method: method,
        headers: { 'Content-Type': 'application/json;charset=utf-8' },
        body: JSON.stringify(body)
    });
    data = await response.json();
    return data;
};
document.addEventListener("DOMContentLoaded", function () {
    readWorks();
});
function readWorks() {
    sendRequest("get", "/api/works")
        .then(data => {
            let html = "";
            let index;
            for (index = 0; index < data.length; ++index) {
                var work = (data[index]);
                html += (`<tr>
            <td><label>${index + 1}</label></td>
            <td><label>${work.work_name}</label></td>
            <td><label>${work.work_price}</label></td>
            <td><button class="btn btn-danger btn-sm" onclick="deleteWork(${work.id})">Удалить</button></td>
            <td><button class="btn btn-primary btn-sm" data-bs-toggle="modal" data-bs-target="#exampleModal" onclick="changeWork(${work.id}, '${work.work_name}', ${work.work_price})">Изменить</button></td>
        </tr>  `);
            }
            document.getElementById("works").innerHTML = html;

        });
}


function deleteWork(id) {
    sendRequest("delete", "/api/works/" + id).then(data => {
        console.log(data)
        readWorks();
    });
};

function changeWork(id, name, price) {
    console.log(id, name, price);
    document.getElementById("exampleModalLabel").innerHTML = "изменить запчасть";
    let html = "";
    html += (`
    <tr><td><input type="text" required="required" value="${id}" maxlength="10"
    name="id" id="id"></td></tr>
        <tr><td><input type="text" required="required" value="${name}" maxlength="10"
        name="workName" id="workName"></td></tr>
        <tr><rd><input type="text" required="required" value="${price}" maxlength="10"
        name="workPrice" id="workPrice"></td></tr>
      `);
    document.getElementById("changework").innerHTML = html;
};
function saveCangeWork(el) {
    const id = el.id.value;
    const name = el.workName.value;
    const price = el.workPrice.value;
    let work = {
        work_name: name,
        work_price: price
    };
    sendRequest('PUT', "/api/works/" + id, body = work)
        .then(data => { console.log(data) 
            readWorks();
        
        });


    return false;
};
function NewWork(el){
    const name = el.NewWorkName.value;
    const price = el.NewWorkPrice.value;
    let work = {
        work_name: name,
        work_price: price
    };
    sendRequest('POST', "/api/works", body = work)
        .then(data => { console.log(data) 
            readWorks();
        });
    return false;
};