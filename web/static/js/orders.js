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
        readOrders(0);
    
});
function readOrders(page){
    sendRequest("get", "/api/count")
    .then(data => {
      
   let ind;
    let i = 0;
    let html = "";
    for (ind = 0; ind <= data; ind = ind +10){
       
      html += (`<button class="btn btn-primary btn-sm" onclick="readOrders(${i})">${i}</button>   `);
      document.getElementById("count").innerHTML = html;
i = i+1;
    }

    });
   
   

   page = page*10;
   
    sendRequest("get", "/api/order/10/"+page)
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
            if (order.status.id_status === 4){
                collor = ("table-danger"); 
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