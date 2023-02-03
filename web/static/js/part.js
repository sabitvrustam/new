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
    readParts();
});
function readParts() {
    sendRequest("get", "/api/parts")
        .then(data => {
            let html = "";
            let index;
            for (index = 0; index < data.length; ++index) {
                var part = (data[index]);
                html += (`<tr>
            <td><label>${index + 1}</label></td>
            <td><label>${part.parts_name}</label></td>
            <td><label>${part.parts_price}</label></td>
            <td><button class="btn btn-danger btn-sm" onclick="deletePart(${part.id})">Удалить</button></td>
            <td><button class="btn btn-primary btn-sm" data-bs-toggle="modal" data-bs-target="#exampleModal" onclick="changePart(${part.id}, '${part.parts_name}', ${part.parts_price})">Изменить</button></td>
        </tr>  `);
            }
            document.getElementById("parts").innerHTML = html;

        });
}


function deletePart(id) {
    sendRequest("delete", "/api/parts/" + id).then(data => {
        console.log(data)
        readParts();
    });
};

function changePart(id, name, price) {
    console.log(id, name, price);
    document.getElementById("exampleModalLabel").innerHTML = "изменить запчасть";
    let html = "";
    html += (`
    <tr><td><input type="text" required="required" value="${id}" maxlength="10"
    name="id" id="id"></td></tr>
        <tr><td><input type="text" required="required" value="${name}" maxlength="10"
        name="partsName" id="partsName"></td></tr>
        <tr><rd><input type="text" required="required" value="${price}" maxlength="10"
        name="partsPrice" id="partsPrice"></td></tr>
      `);
    document.getElementById("changePart").innerHTML = html;
};
function saveCangePart(el) {
    const id = el.id.value;
    const name = el.partsName.value;
    const price = el.partsPrice.value;
    let part = {
        parts_name: name,
        parts_price: price
    };
    sendRequest('PUT', "/api/parts/" + id, body = part)
        .then(data => { console.log(data) 
            readParts();
        
        });


    return false;
};
function NewPart(el){
    const name = el.NewPartsName.value;
    const price = el.NewPartsPrice.value;
    let part = {
        parts_name: name,
        parts_price: price
    };
    sendRequest('POST', "/api/parts", body = part)
        .then(data => { console.log(data) 
            readParts();
        });
    return false;
};