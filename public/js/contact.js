function submitData(event){
    event.preventDefault();
    let name = document.getElementById("input-name").value;
    let email = document.getElementById("input-email").value;
    let phone = document.getElementById("input-phone").value;
    let subject = document.getElementById("input-subject").value;
    let message = document.getElementById("input-message").value;
    
    // console.log(`${name}\n${email}\n${phone}\n${subject}\n${message}`);
    
    if(name == ""){
        return alert("Name input fields must be not empty");
    } else if(email == ""){
        return alert("Email input fields must be not empty");
    } else if(phone == ""){
        return alert("Phone input fields must be not empty");
    } else if(subject == ""){
        return alert("Subject input fields must be not empty");
    } else if(message == ""){
        return alert("Message input fields must be not empty");
    }
    
    let a = document.createElement("a");
    a.href = `mailto:${email}?Subject=${subject}&body= Hello my name is ${name}, ${message}. Please contact my phone number at ${phone}`;
    a.target = "_blank";
    a.click();

    let submitDataObject = {
        name, email, phone, subject, message,
    }
    console.log(submitDataObject);
}


// Logic Input Color
let inputBorderColor1 = document.querySelectorAll(".border-color");
let inputBorderColor2 = document.querySelector(".form-container form");
let input = document.querySelectorAll("input");
let btnSubmit = document.querySelector(".btn-submit");

inputBorderColor1.forEach(function(ibc1) {
    ibc1.addEventListener('click', function(event) {
        
        inputBorderColor1.forEach(function(ibc1) {
            if(ibc1.classList.contains("violet-border")){
                ibc1.classList.remove("violet-border");
            }
        });    
        event.target.classList.add("violet-border");
    });

    btnSubmit.addEventListener('mouseover', function() {
        inputBorderColor1.forEach(function(ibc1) {
            ibc1.classList.remove("violet-border");
        });  
    });  
});

inputBorderColor2.addEventListener('mouseenter', function(){
    inputBorderColor1.forEach(function(ibc1){
        ibc1.classList.add("black-border");
    });
    
    inputBorderColor2.addEventListener('mouseleave', function(){
        inputBorderColor1.forEach(function(ibc1){
            ibc1.classList.remove("black-border");
            ibc1.classList.remove("violet-border");
        });
    });
});
// End logic input color

// Logic slide navbar
const slide = document.querySelector('.slide');
const menuToggle = document.querySelector('.menu-toggle input');
const nav = document.querySelector('.navbar ul');

menuToggle.addEventListener('click', function(){
    nav.classList.toggle('slide');
});
// End logic end navbar



  