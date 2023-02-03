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
    readOrders();
});
function readOrders(){
    sendRequest("get", "/api/order/100/0")
    .then(data => {
        let html = "";
        let index;
        for (index = 0; index < data.length; ++index) {
            var order = (data[index]);
            let collor = "";
            if (order.status.id_status === 8){
collor = ("table-success");
            };
            if (order.status.id_status === 2){
                collor = ("table-info"); 
            };
            if (order.status.id_status === 1){
                collor = ("table-primary"); 
            };
            if (order.status.id_status === 5){
                collor = ("table-warning"); 
            };

            html += (`<tr class="${collor}">
        <td style="text-align: center;"><label>${order.id_order}</label></td>
        <td style="text-align: center;"><label>${order.user.last_name} ${order.user.first_name} ${order.user.midl_name}</label></td>
        <td style="text-align: center;"><label>${order.device.type_equipment}</label></td>
        <td style="text-align: center;"><label>${order.device.model}</label></td>
        <td style="text-align: center;"><label>${order.master.last_name}</label></td>
        <td style="text-align: center;"><label>${order.status.status_order}</label></td>
   `);
        }
        document.getElementById("orders").innerHTML = html;

    });

};