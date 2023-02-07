async function sendRequest(method, url, body) {
    let response = await fetch(url, {
        method: method,
        headers: { 'Content-Type': 'application/json;charset=utf-8' },
        body: JSON.stringify(body)
    });
    data = await response.json();
    return data;
};
var id_user;
var id_device;

document.addEventListener("DOMContentLoaded", function () {
    document.getElementById("orderStatus").style.display = "none";
});

function readIdOrder(el) {
    var idOrder = el.id.value;
    let order = {};
    order["id"] = idOrder;
    readOrder(order);
    return false;
};

function readOrder(order) {
    document.getElementById("orderStatus").style.display = "block";
    sendRequest('get', '/api/order/' + order.id)
        .then(data => {
            order["user"] = data.user;
            order["device"] = data.device;
            order["statuse"] = data.status;
            order["master"] = data.master;
            id_user = order.user.id_user;
            id_device = order.device.id_device;
            readMasters(order);
        });
};

function readMasters(order) {
    sendRequest("get", "/api/masters")
        .then(data => {
            order["masters"] = data;
            readStatus(order);
        });
};

function readStatus(order) {
    sendRequest("get", "/api/status")
        .then(data => {
            order["statuses"] = data;
            readOrderParts(order);
        });
};

function readOrderParts(order) {
    sendRequest('get', '/api/orderparts/' + order.id)
        .then(data => {
            order["order_parts"] = data;
            readOrderWorks(order);
        });
};

function readOrderWorks(order) {
    sendRequest('get', '/api/orderworks/' + order.id)
        .then(data => {
            order["order_works"] = data;
            innerHtml(order);
        });
};

function innerHtml(order) {
    document.getElementById("idOrder").value = order.id;
    document.getElementById("UserFirstName").value = order.user.last_name;
    document.getElementById("UserLastName").value = order.user.first_name;
    document.getElementById("UserMidlName").value = order.user.midl_name;
    document.getElementById("PhoneNomber").value = order.user.phone;
    document.getElementById("TypeEquipment").value = order.device.type_equipment;
    document.getElementById("Brand").value = order.device.brand;
    document.getElementById("Model").value = order.device.model;
    document.getElementById("SerialN").value = order.device.sn;

    html = "";
    let index;
    for (index = 0; index < order.masters.length; ++index) {
        var master = (order.masters[index]);
        if (order.master.id === master.id) {
            html += (`<option selected value="${master.id}">${master.first_name}</option>`);
        };
        html += (`<option value="${master.id}">${master.first_name}</option>`);
    }
    document.getElementById("masters").innerHTML = html;

    html = "";
    for (index = 0; index < order.statuses.length; ++index) {
        var status = (order.statuses[index]);
        if (order.statuse.id_status === status.id_status) {
            html += (`<option selected value="${status.id_status}">${status.status_order}</option>`);
        };
        html += (`<option value="${status.id_status}">${status.status_order}</option>`);
    }
    document.getElementById("status").innerHTML = html;



    if (order.order_parts !== null) {
        let html = "";
        let index;
        for (index = 0; index < order.order_parts.length; ++index) {
            var part = (order.order_parts[index]);
            html += (`<tr>
                    <td><label>${index + 1}</label></td>
                    <td><label>${part.part_name}</label></td>
                    <td><label>${part.part_price}</label></td>
                </tr>  `);
        };
        document.getElementById("parts").innerHTML = html;
    } else {
        document.getElementById("parts").innerHTML = "";
    };

    if (order.order_works !== null) {
        let html = "";
        let index;
        for (index = 0; index < order.order_works.length; ++index) {
            var work = (order.order_works[index]);
            html += (`<tr>
                <td><label>${index + 1}</label></td>
                <td><label>${work.work_name}</label></td>
                <td><label>${work.work_price}</label></td>
            </tr>  `);
        };
        document.getElementById("works").innerHTML = html;
    } else {
        document.getElementById("works").innerHTML = "";
    };
};

function changeOrder(el){
    let id = el.idOrder.value;
    let UserFirstName = el.UserFirstName.value;
    let UserLastName = el.UserLastName.value;
    let UserMidlName = el.UserMidlName.value;
    let PhoneNombe = el.PhoneNomber.value;
    let TypeEquipment = el.TypeEquipment.value;
    let Brand = el.Brand.value;
    let Model = el.Model.value;
    let SerialN = el.SerialN.value;
    let idMaster = el.masters.value;
    let idStatus = el.status.value;

    let user = {
        first_name: UserFirstName,
        last_name: UserLastName,
        midl_name: UserMidlName,
        phone: PhoneNombe
    };

    let device = {
        type_equipment: TypeEquipment,
        brand: Brand,
        model: Model,
        sn: SerialN
    };

    let order = {
        id_user: id_user,
        id_device: id_device,
        id_master: parseInt(idMaster),
        id_status: parseInt(idStatus)
    };

    sendRequest('PUT', "/api/users/"+id_user, body = user)
        .then(data => { console.log(data)
            sendRequest('PUT', "/api/device/"+id_device, body = device)
            .then(data => { console.log(data)
                sendRequest('PUT', "/api/order/"+id, body = order)
                .then(data => { console.log(order)
                    
                });
            });
        });
    return false;
};
