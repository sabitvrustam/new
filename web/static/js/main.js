function sendRequest(method, url, body = null) {
    return new Promise((resolve, reject) => {
        const xhr = new XMLHttpRequest;
        xhr.open(method, url);
        xhr.responseType = 'json';
        xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        xhr.onload = () => {
            if (xhr.status >= 400) {
                reject(xhr.response);
            } else {
                resolve(xhr.response);
            }
        }
        xhr.onerror = () => {
            reject(xhr.response);
        }
        xhr.send(JSON.stringify(body));
    });
};

function newUserForm(el) {
    var UserFirstName = el.UserFirstName.value;
    var UserLastName = el.UserLastName.value;
    var UserMidlName = el.UserMidlName.value;
    var PhoneNombe = el.PhoneNombe.value;
    var TypeEquipment = el.TypeEquipment.value;
    var Brand = el.Brand.value;
    var Model = el.Model.value;
    var SerialN = el.SerialN.value;
    var MasterId = el.MasterId.value;

    document.getElementById("UserFirstNameErr").innerHTML = "";
    document.getElementById("UserLastNameErr").innerHTML = "";
    document.getElementById("UserMidlNameErr").innerHTML = "";
    document.getElementById("PhoneNombeErr").innerHTML = "";
    document.getElementById("TypeEquipmentErr").innerHTML = "";
    document.getElementById("BrandErr").innerHTML = "";
    document.getElementById("ModelErr").innerHTML = "";
    document.getElementById("SerialNErr").innerHTML = "";

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
    var Order = {
        user: {
            first_name: UserFirstName,
            last_name: UserLastName,
            midl_name: UserMidlName,
            phone: PhoneNombe
        },
        device: {
            type_equipment: TypeEquipment,
            brand: Brand,
            model: Model,
            sn: SerialN
        },
        masters:{
            id: MasterId
        },
        status:{
            status_order: "1",
        }

    };
    const requestURL = '/createOrder';

sendRequest('POST', requestURL, body = Order)
    .then(data => console.log(data))
    .catch(err =>console.log(err));


    return false;
};




//     .then(data => console.log(data))
//     .catch(err =>console.log(err));


// const testJson = async(url) => {
//     const data = await fetch(url);
//     console.log(data);
//     if(!data.ok){
//         throw new error('ошибка по адресу ${url}, статус ошибки ${data}');
//     }
//     return await data.json();
// };

// testJson(requestURL).then((data) => console.log(data));
// function createOrder (el){
// el.innerHTML = "hihihi";
// };



// const senddata = async (url, data) => {
//     const dataread = await fetch(url, {
//         metod: 'post',
//         body: 'data',

//     });
//     return await dataread.json();

// }