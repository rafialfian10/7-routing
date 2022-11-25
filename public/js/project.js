// Logic input color
let inputBorderColor1 = document.querySelectorAll(".border-color");
let inputBorderColor2 = document.querySelector(".form-container form");
let btnSubmit = document.querySelector(".btn-submit");
let label = document.querySelector("label.border-color");
let p = document.querySelector(".img-upload p")
let imgUpload = document.querySelector(".img-upload label img")

inputBorderColor1.forEach(function(ibc1) {
    ibc1.addEventListener('click', function(event) {
        
        inputBorderColor1.forEach(function(ibc1) {
            if(ibc1.classList.contains("violet-border")){
                ibc1.classList.remove("violet-border");
            }
        });    
        event.target.classList.add("violet-border");
        p.classList.remove("violet-border");
        imgUpload.classList.remove("violet-border");  
    });
 
    btnSubmit.addEventListener('mouseover', function() {
        inputBorderColor1.forEach(function(ibc1) {
            ibc1.classList.remove("violet-border");
        }); 
    });  
});

inputBorderColor2.addEventListener('mouseenter', function(){
    label.style.border = "1px solid black"
    inputBorderColor1.forEach(function(ibc1){
        ibc1.classList.add("black-border");
    });
    
    inputBorderColor2.addEventListener('mouseleave', function(){
        label.style.border = "2px solid white"
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





