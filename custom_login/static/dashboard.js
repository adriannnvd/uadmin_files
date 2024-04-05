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


  
  
  