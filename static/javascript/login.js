function loginButton() {
    var passValid =  /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[^a-zA-Z0-9])(?!.*\s).{8,15}$/;
    var password = document.getElementById("loginPass").value;
    if (passValid.test(password)){
        document.getElementById("loginBtn").style.display = "block";
    }else {
        document.getElementById("loginBtn").style.display = "none";
    }
    
}

var  updateLogin = () => document.getElementById("login").style.display = document.getElementById("login").style.display == "block" ? "none" : "block"
var closeLogin = () => document.getElementById("login").style.display = "none"

function hideLogin() {
    document.getElementById("login").style.display = "none"
}