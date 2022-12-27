// var test = document.getElementById('test');
// console.log(test.id);
// test.innerHTML = "Привет друг";

// var newOrder = document.getElementById('newOrder');
// console.log(newOrder.method);

function newUserForm (el){
    var login = el.login.value;
    var password = el.password.value;
    var passwoerdCon = el.passwordCon.value;
    console.log(login, password, passwoerdCon);
    if (login == "" || login.lenght < 3){

        document.getElementById("loginErr").innerHTML = "Не корректный логин";
       // window.location = 'https://www.google.com';
        return false;
    }

    return false;
}


function createOrder (el){
el.innerHTML = "hihihi";

}