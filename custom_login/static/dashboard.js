//filled heart
document.querySelector("#heart").addEventListener("click", function() {
    if (this.classList.contains("filled-heart")) {
      this.classList.remove("filled-heart");
      this.classList.remove("fa-solid");
    } else {
      this.classList.add("filled-heart");
      this.classList.add("fa-solid");
    }
  });

  //filled star
document.querySelector("#star").addEventListener("click", function() {
  if (this.classList.contains("filled-star")) {
    this.classList.remove("filled-star");
    this.classList.remove("fa-solid");
  } else {
    this.classList.add("filled-star");
    this.classList.add("fa-solid");
  }
});

//filled heart
document.querySelector("#bookmark").addEventListener("click", function() {
  if (this.classList.contains("filled-bookmark")) {
    this.classList.remove("filled-bookmark");
    this.classList.remove("fa-solid");
  } else {
    this.classList.add("filled-bookmark");
    this.classList.add("fa-solid");
  }
});


  //sidebar navigation
  function openNav() {
    document.getElementById("mySidenav").style.width = "320px";
    document.getElementById("main").style.marginLeft = "320px";
    document.body.style.backgroundColor = "rgba(0,0,0,0.4)";
  }
  
  function closeNav() {
    document.getElementById("mySidenav").style.width = "0";
    document.getElementById("main").style.marginLeft= "0";
    document.body.style.backgroundColor = "white";
  }

  //CHANGING THE ACTIVE BUTTON FOR SIDEBAR BUTTONS
var btnContainer = document.getElementById("quotes-type");

var btns = btnContainer.getElementsByClassName("quotes-btn");

for (var i = 0; i < btns.length; i++) {
  btns[i].addEventListener("click", function() {
    var current = document.getElementsByClassName("active");
    current[0].className = current[0].className.replace(" active", "");

    this.className += " active";
  });
}



  
  
  