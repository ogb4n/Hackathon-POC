function blur(){
    document.getElementById("blurrr").style.display = "block";
    document.getElementById("filter").style.filter = "blur(4px)";
    document.getElementById("search").style.filter = "blur(4px)";
    document.getElementById("post").style.filter = "blur(4px)";
}

function unblur(){
    document.getElementById("blurrr").style.display = "none";
    document.getElementById("filter").style.filter = "blur(0px)";
    document.getElementById("search").style.filter = "blur(0px)";
    document.getElementById("post").style.filter = "blur(0px)";
}

function loginButton() {
    var passValid =  /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[^a-zA-Z0-9])(?!.*\s).{8,15}$/;
    var password = document.getElementById("loginPass").value;
    if (passValid.test(password)){
        document.getElementById("loginBtn").disabled = false;
    }else {
        document.getElementById("loginBtn").disabled = true;
    }
    
}