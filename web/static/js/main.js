var context = {}


async function sendRequest(method, url, contextName, body) {
    let response = await fetch(url, {
        method: method,
        headers: { 'Content-Type': 'application/json;charset=utf-8' },
        body: JSON.stringify(body)
    });
    // data = await response.json();
    context[contextName] = await response.json()
    return context;
};



document.addEventListener("DOMContentLoaded", function () {

    document.getElementById("createDevice").style.display = "none";
    document.getElementById("createOrder").style.display = "none";
    document.getElementById("orderStatus").style.display = "none";
    sendRequest("get", "/api/masters", 'masters')
        .then(data => {
            html = "";
            let index;
            for (index = 0; index < context.masters.length; ++index) {
                var master = (context.masters[index]);
                html += (`<option value="${master.id}">${master.first_name}</option>`);
            }
            console.log(html);
            console.log(context);
            document.getElementById("masters").innerHTML = html;

        });
    sendRequest("get", "/api/status", 'status')
        .then(data => {
            html = "";
            let index;
            for (index = 0; index < context.status.length; ++index) {
                var status = (context.status[index]);
                html += (`<option value="${status.id_status}">${status.status_order}</option>`);
            }
            console.log(html);
            console.log(context);
            document.getElementById("status").innerHTML = html;

        });

});

function newUserForm(el, context) {
    const UserFirstName = el.UserFirstName.value;
    const UserLastName = el.UserLastName.value;
    const UserMidlName = el.UserMidlName.value;
    const PhoneNombe = el.PhoneNombe.value;
    document.getElementById("UserFirstNameErr").innerHTML = "";
    document.getElementById("UserLastNameErr").innerHTML = "";
    document.getElementById("UserMidlNameErr").innerHTML = "";
    document.getElementById("PhoneNombeErr").innerHTML = "";


    if (UserFirstName == "" || UserFirstName.length <= 3) {
        document.getElementById("UserFirstNameErr").innerHTML = "Не корректная фамилия";
        return false;
    }
    if (UserLastName == "" || UserLastName.length <= 3) {
        document.getElementById("UserLastNameErr").innerHTML = "Не корректное имя";
        return false;
    }
    if (UserMidlName == "" || UserMidlName.length <= 3) {
        document.getElementById("UserMidlNameErr").innerHTML = "Не корректное отчество";
        return false;
    }
    if (PhoneNombe == "" || PhoneNombe <= 6) {
        document.getElementById("PhoneNombeErr").innerHTML = "не коректный номер телефона";
        return false;
    }
    let user = {
        first_name: UserFirstName,
        last_name: UserLastName,
        midl_name: UserMidlName,
        phone: PhoneNombe
    };
    const requestURL = '/api/users';
    sendRequest('POST', requestURL, 'user', body = user)
        .then(data => { console.log(context) });

    document.getElementById("createUser").style.display = "none";
    document.getElementById("createDevice").style.display = "block";
    return false;
};

function newDeviceForm(el, context) {
    var TypeEquipment = el.TypeEquipment.value;
    var Brand = el.Brand.value;
    var Model = el.Model.value;
    var SerialN = el.SerialN.value;

    document.getElementById("TypeEquipmentErr").innerHTML = "";
    document.getElementById("BrandErr").innerHTML = "";
    document.getElementById("ModelErr").innerHTML = "";
    document.getElementById("SerialNErr").innerHTML = "";


    if (TypeEquipment == "" || TypeEquipment <= 3) {
        document.getElementById("TypeEquipmentErr").innerHTML = "не коректный тип аппарата";
        return false;
    }
    if (Brand == "" || Brand <= 3) {
        document.getElementById("BrandErr").innerHTML = "не коректная фирма";
        return false;
    }
    if (Model == "" || Model <= 3) {
        document.getElementById("ModelErr").innerHTML = "не коректная модель";
        return false;
    }
    if (SerialN == "" || SerialN <= 3) {
        document.getElementById("SerialNErr").innerHTML = "не коректная модель";
        return false;
    }
    let device = {
        type_equipment: TypeEquipment,
        brand: Brand,
        model: Model,
        sn: SerialN
    };
    const requestURL = '/api/device';

    sendRequest('POST', requestURL, 'device', body = device)
        .then(data => { console.log(data) });

    document.getElementById("createUser").style.display = "none";
    document.getElementById("createDevice").style.display = "none";
    document.getElementById("createOrder").style.display = "block";
    console.log(context.user.id_user);
    return false;
};
function newOrderForm(el, context) {
    var idMaster = el.masters.value;
    var idStatus = el.status.value;
    console.log(idStatus);

    let order = {
        id_user: context.user.id_user,
        id_device: context.device.id_device,
        id_master: parseInt(idMaster),
        id_status: parseInt(idStatus)
    };
    const requestURL = '/api/order';

    sendRequest('POST', requestURL, 'order', body = order)
        .then(data => {
            document.getElementById("idOrder").innerHTML = data.order.id_order;
            document.getElementById("UserFirstNameOrder").innerHTML = data.order.user.last_name;
            document.getElementById("UserLastNameOrder").innerHTML = data.order.user.first_name;
            document.getElementById("UserMidlNameOrder").innerHTML = data.order.user.midl_name;
            document.getElementById("UserPhoneNomberOrder").innerHTML = data.order.user.phone;
            document.getElementById("DeviceTypeEquipment").innerHTML = data.order.device.type_equipment;
            document.getElementById("DeviceBrand").innerHTML = data.order.device.brand;
            document.getElementById("DeviceModel").innerHTML = data.order.device.model;
            document.getElementById("DeviceSerialN").innerHTML = data.order.device.sn;
            document.getElementById("MasterFirstName").innerHTML = data.order.master.first_name;
            document.getElementById("StatusOrder").innerHTML = data.order.status.status_order;


        });


    document.getElementById("orderStatus").style.display = "block";

    return false;
};
