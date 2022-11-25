// Logic slide navbar
const slide = document.querySelector('.slide');
const menuToggle = document.querySelector('.menu-toggle input');
const nav = document.querySelector('.navbar ul');

menuToggle.addEventListener('click', function(){
    nav.classList.toggle('slide');
});

// End logic end navbar