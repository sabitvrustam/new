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

    document.getElementById("orderStatus").style.display = "none";

});
function OrderStatus(el) {
    var idOrder = el.id.value;
    document.getElementById("orderStatus").style.display = "block";
    sendRequest('get', '/api/order/' + idOrder)
        .then(data => {
            document.getElementById("idOrder").innerHTML = data.id_order;
            document.getElementById("UserFirstNameOrder").innerHTML = data.user.last_name;
            document.getElementById("UserLastNameOrder").innerHTML = data.user.first_name;
            document.getElementById("UserMidlNameOrder").innerHTML = data.user.midl_name;
            document.getElementById("UserPhoneNomberOrder").innerHTML = data.user.phone;
            document.getElementById("DeviceTypeEquipment").innerHTML = data.device.type_equipment;
            document.getElementById("DeviceBrand").innerHTML = data.device.brand;
            document.getElementById("DeviceModel").innerHTML = data.device.model;
            document.getElementById("DeviceSerialN").innerHTML = data.device.sn;
            document.getElementById("MasterFirstName").innerHTML = data.master.first_name;
            document.getElementById("StatusOrder").innerHTML = data.status.status_order;
        });
    sendRequest('get', '/api/orderparts/' + idOrder)
        .then(data => {
            let html = "";
            if (data !== null) {
                let index;
                for (index = 0; index < data.length; ++index) {
                    var part = (data[index]);
                    html += (`<tr>
                <td><label>${index + 1}</label></td>
                <td><label>${part.part_name}</label></td>
                <td><label>${part.part_price}</label></td>
            </tr>  `);
                };
            };
            document.getElementById("parts").innerHTML = html;
        });
        sendRequest('get', '/api/orderworks/' + idOrder)
        .then(data => {
            let html = "";
            if (data !== null) {
                let index;
                for (index = 0; index < data.length; ++index) {
                    var work = (data[index]);
                    html += (`<tr>
                <td><label>${index + 1}</label></td>
                <td><label>${work.work_name}</label></td>
                <td><label>${work.work_price}</label></td>
            </tr>  `);
                };
            };
            document.getElementById("works").innerHTML = html;
        });




    return false;
};
