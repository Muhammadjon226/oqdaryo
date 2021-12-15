var slideIndex = 1;
showSlides(slideIndex);

// Next/previous controls
function plusSlides(n) {
  showSlides(slideIndex += n);
}

// Thumbnail image controls
function currentSlide(n) {
  showSlides(slideIndex = n);
}

function showSlides(n) {
  var i;
  var slides = document.getElementsByClassName("mySlides");
  var dots = document.getElementsByClassName("dot");
  if (n > slides.length) {slideIndex = 1}
  if (n < 1) {slideIndex = slides.length}
  for (i = 0; i < slides.length; i++) {
      slides[i].style.display = "none";
  }
  for (i = 0; i < dots.length; i++) {
      dots[i].className = dots[i].className.replace(" active", "");
  }
  slides[slideIndex-1].style.display = "block";
  dots[slideIndex-1].className += " active";
}

// Hamburger menu 
(() => {
  const hamburger = document.getElementById("hamburger");
  const menu = document.getElementById("overlay");
  let open = false;

  const change = () => {
    if (!open) {
      hamburger.classList.add("open");
      menu.classList.add("menu");
    } else {
      hamburger.classList.remove("open");
      menu.classList.remove("menu");
    }
    open = !open;
  };

  hamburger.addEventListener("click", change);
})();


// form to JSON format
let comments = [];
// const myJson = JSON.stringify(comments);
// get date, example: {id:123423423543, name: "Anvar", email: anvar12345@gmail.com, commnet: "Hello, everybody"}
const addComment = (a) => {
  a.preventDefault(); 
  let comment = {
    id: Date.now(),
    name: document.getElementById("sender_name").value,
    email: document.getElementById("sender_email").value,
    comment: document.getElementById("sender_comment").value
  };
  comments.push(comment); // comment obyektini comments massiviga joylash
  document.forms[0].reset();

  let getJson = document.querySelector('#msg pre');
  getJson.textContent = '\n' + JSON.stringify(comments, '\t', 2);

};




document.addEventListener('DOMContentLoaded', ()=>{
  document.getElementById('sendBtn').addEventListener('click', addComment);
});

