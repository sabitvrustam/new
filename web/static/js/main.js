// var test = document.getElementById('test');
// console.log(test.id);
// test.innerHTML = "Привет друг";

// var newOrder = document.getElementById('newOrder');
// console.log(newOrder.method);

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


    console.log(UserFirstName, UserLastName, UserMidlName, PhoneNombe, TypeEquipment, Brand, Model, SerialN, MasterId);
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
var Order = {User:{
    FirstName: UserFirstName,
}

};

    return false;
};

// function sendRequest(method, url, body = null) {
//     return new Promise((resolve, reject) => {
//         const xhr = new XMLHttpRequest;
//         xhr.open(method, url);
//         xhr.responseType = 'json';
//         xhr.setRequestHeader('Content-type', 'application/json');
//         xhr.onload = () => {
//             if (xhr.status >= 400) {
//                 reject(xhr.response);
//             } else {
//                 resolve(xhr.response);
//             }
//         }
//         xhr.onerror = () => {
//             reject(xhr.response);
//         }
//         xhr.send(JSON.stringify(body));
//     });
// };

// const requestURL = 'https://jsonplaceholder.typicode.com/users';

// sendRequest('post', requestURL, body = {
//     "name": "Rustam",
//     "username": "Hello"


// })
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