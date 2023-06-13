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
    readInginer();
});

function readInginer() {
    sendRequest("get", "/api/masters")
        .then(data => {
            let html = "";
            let index;
            for (index = 0; index < data.length; ++index) {
                var inginer = (data[index]);
                html += (`<tr>
            <td><label>${index + 1}</label></td>
            <td><label>${inginer.first_name}</label></td>
            <td><label>${inginer.last_name}</label></td>
            <td><label>${inginer.midl_name}</label></td>
            <td><label>${inginer.phone}</label></td>
            <td><button class="btn btn-danger btn-sm" onclick="deleteIngener(${inginer.id})">Удалить</button></td>
           <td><button class="btn btn-primary btn-sm" data-bs-toggle="modal" data-bs-target="#exampleModal" onclick="changeIngener(${inginer.id}, '${inginer.first_name}',
        '${inginer.last_name}', '${inginer.midl_name}', '${inginer.phone}')">Изменить</button></td>
        </tr>  `);
            }
            document.getElementById("inginers").innerHTML = html;

        });
}

function newIngenerForm(el) {
    var FirstName = el.IngenerFirstName.value;
    var LastName = el.IngenerLastName.value;
    var MidName = el.IngenerMidlName.value;
    var Phone = el.PhoneNomber.value;

    document.getElementById("IngenerFirstNameErr").innerHTML = "";
    document.getElementById("IngenerLastNameErr").innerHTML = "";
    document.getElementById("IngenerMidlNameErr").innerHTML = "";
    document.getElementById("PhoneNomberErr").innerHTML = "";


    if (FirstName == "" || FirstName <= 3) {
        document.getElementById("IngenerFirstNameErr").innerHTML = "не коректная фамилия мастера";
        return false;
    }
    if (LastName == "" || LastName <= 3) {
        document.getElementById("IngenerLastNameErr").innerHTML = "не коректное имя мастера";
        return false;
    }
    if (MidName == "" || MidName <= 3) {
        document.getElementById("IngenerMidlNameErr").innerHTML = "не коректное отчество";
        return false;
    }
    if (Phone == "" || Phone <= 11) {
        document.getElementById("PhoneNomberErr").innerHTML = "не коректный номер телефона";
        return false;
    }

    let NewIngener = {
        first_name: FirstName,
        last_name:  LastName,
        midl_name: MidName,
        phone: Phone
    };

    sendRequest('POST', "/api/masters", body = NewIngener)
    .then(data => { console.log(data);
    });
return true;
   
};

function deleteIngener(id) {
    sendRequest("delete", "/api/masters/" + id).then(data => {
        console.log(data)
        readParts();
    });
    readInginer()
};

function changeIngener(id, first_name, last_name, mid_name, phone) {
    console.log(id, first_name, last_name, mid_name, phone);
    document.getElementById("exampleModalLabel").innerHTML = "изменить данные мастера";
    let html = "";
    html += (`
        <tr><td><input type="text" required="required" value="${id}" maxlength="10"
        name="id" id="id"></td></tr>
        <tr><td><input type="text" required="required" value="${first_name}" maxlength="12"
        name="first_name" id="first_name"></td></tr>
        <tr><td><input type="text" required="required" value="${last_name}" maxlength="12"
        name="last_name" id="last_name"></td></tr>
        <tr><td><input type="text" required="required" value="${mid_name}" maxlength="12"
        name="mid_name" id="mid_name"></td></tr>
        <tr><rd><input type="text" required="required" value="${phone}" maxlength="12"
        name="phone" id="phone"></td></tr>
      `);
    document.getElementById("changeIngener").innerHTML = html;
};
function saveChangeIngener(el) {
    const id = el.id.value;
    const firstName = el.first_name.value;
    const lastName = el.last_name.value;
    const midName = el.mid_name.value;
    const phone = el.phone.value;


    let ingener = {
        first_name: firstName,
        last_name: lastName,
        midl_name: midName,
        phone: phone
    };

    sendRequest('PUT', "/api/masters/" + id, body = ingener)
        .then(data => { console.log(data) 
            readInginer();
        
        });


    return false;
};

