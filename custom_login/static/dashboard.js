// //filled heart
// document.querySelector("#heart").addEventListener("click", function() {
//     if (this.classList.contains("filled-heart")) {
//       this.classList.remove("filled-heart");
//       this.classList.remove("fa-solid");
//     } else {
//       this.classList.add("filled-heart");
//       this.classList.add("fa-solid");
//     }
//   });

//   //filled star
// document.querySelector("#star").addEventListener("click", function() {
//   if (this.classList.contains("filled-star")) {
//     this.classList.remove("filled-star");
//     this.classList.remove("fa-solid");
//   } else {
//     this.classList.add("filled-star");
//     this.classList.add("fa-solid");
//   }
// });

// //filled heart
// document.querySelector("#bookmark").addEventListener("click", function() {
//   if (this.classList.contains("filled-bookmark")) {
//     this.classList.remove("filled-bookmark");
//     this.classList.remove("fa-solid");
//   } else {
//     this.classList.add("filled-bookmark");
//     this.classList.add("fa-solid");
//   }
// });


  // //sidebar navigation
  // function openNav() {
  //   document.getElementById("mySidenav").style.width = "320px";
  //   document.getElementById("main").style.marginLeft = "320px";
  //   document.body.style.backgroundColor = "rgba(0,0,0,0.4)";
  // }
  
  // function closeNav() {
  //   document.getElementById("mySidenav").style.width = "0";
  //   document.getElementById("main").style.marginLeft= "0";
  //   document.body.style.backgroundColor = "white";
  // }

  //CHANGING THE ACTIVE BUTTON FOR SIDEBAR BUTTONS
// var btnContainer = document.getElementById("quotes-type");

// var btns = btnContainer.getElementsByClassName("quotes-btn");

// for (var i = 0; i < btns.length; i++) {
//   btns[i].addEventListener("click", function() {
//     var current = document.getElementsByClassName("active");
//     current[0].className = current[0].className.replace(" active", "");

//     this.className += " active";
//   });
// }


// document.querySelector("#quotes-btn").addEventListener("click", function() {
//   var heartIcon = document.getElementById("heart");
//   if (heartIcon.classList.contains("fa-regular")) {
//       heartIcon.classList.remove("fa-regular");
//       heartIcon.classList.add("fa-solid");
//   } else {
//       heartIcon.classList.remove("fa-solid");
//       heartIcon.classList.add("fa-regular");
//   }
// });

// Zoom Authot Profile
const modal = document.getElementById('modal');
const modalImg = document.getElementById("modal-image");
const zoomableImg = document.querySelector('.zoomable-image');
const closeBtn = document.getElementsByClassName("close")[0];

zoomableImg.onclick = function(){
    modal.style.display = "block";
    modalImg.src = this.src;
}

closeBtn.onclick = function() {
    modal.style.display = "none";
}

window.onclick = function(event) {
    if (event.target == modal) {
        modal.style.display = "none";
    }
}



var div = document.getElementById('quotes-cover');

div.style.display = 'flex';
div.style.opacity = '1';

setTimeout(function() {
    div.style.opacity = '0';
    
    setTimeout(function() {
        div.style.display = 'none';
    }, 1000);
}, 2000);


function hideAndReveal() {
  var div = document.getElementById('quotes-cover');

        div.style.display = 'flex';
        div.style.opacity = '1';

        setTimeout(function() {
            div.style.opacity = '0';
            
            setTimeout(function() {
                div.style.display = 'none';
            }, 500);
        }, 2000);
}

document.addEventListener("DOMContentLoaded", function() {
  // Get all buttons in the sidebar
  const buttons = document.querySelectorAll(".quotes-btn");

  // Function to handle button click
  function handleButtonClick(event) {
    // Remove active class from all buttons
    buttons.forEach(btn => btn.classList.remove("active"));

    // Add active class to the clicked button
    event.target.classList.add("active");

    // Store the ID of the clicked button in local storage
    localStorage.setItem("activeButtonId", event.target.id);

    // Call the appropriate function to display content based on the clicked button
    // Assuming you have functions like showRandomQuotes(), showInspirationalQuotes(), etc.
    // Adjust this part based on your actual implementation
    switch (event.target.id) {
      case "random-quotes-btn":
        showRandomQuotes();
        break;
      case "inspirational-quotes-btn":
        showInspirationalQuotes();
        break;
      case "motivational-quotes-btn":
        showMotivationalQuotes();
        break;
      case "love-quotes-btn":
        showLoveQuotes();
        break;
      case "perseverance-quotes-btn":
        showPerseveranceQuotes();
        break;
      case "encouragement-quotes-btn":
        showEncouragementQuotes();
        break;
      default:
        // Handle default case if needed
        break;
    }
  }

  // Add event listeners to all buttons
  buttons.forEach(button => {
    button.addEventListener("click", handleButtonClick);
  });

  // Check if there's a stored active button ID and set it as active on page load
  const activeButtonId = localStorage.getItem("activeButtonId");
  if (activeButtonId) {
    const activeButton = document.getElementById(activeButtonId);
    if (activeButton) {
      activeButton.classList.add("active");
      
      // Trigger the display of content corresponding to the last active button
      // Assuming you have functions like showRandomQuotes(), showInspirationalQuotes(), etc.
      // Adjust this part based on your actual implementation
      switch (activeButtonId) {
        case "random-quotes-btn":
          showRandomQuotes();
          break;
        case "inspirational-quotes-btn":
          showInspirationalQuotes();
          break;
        case "motivational-quotes-btn":
          showMotivationalQuotes();
          break;
        case "love-quotes-btn":
          showLoveQuotes();
          break;
        case "perseverance-quotes-btn":
          showPerseveranceQuotes();
          break;
        case "encouragement-quotes-btn":
          showEncouragementQuotes();
          break;
        default:
          // Handle default case if needed
          break;
      }
    }
  }
});

// Function to open the sidebar
function openNav() {
  document.getElementById("mySidenav").style.width = "320px";
  document.getElementById("main").style.marginLeft = "320px";
  document.body.style.backgroundColor = "rgba(0,0,0,0.4)";
  // Store the state of the sidebar as open in local storage
  localStorage.setItem("sidebarState", "open");
}

// Function to close the sidebar
function closeNav() {
  document.getElementById("mySidenav").style.width = "0";
  document.getElementById("main").style.marginLeft= "0";
  document.body.style.backgroundColor = "white";
  // Store the state of the sidebar as closed in local storage
  localStorage.setItem("sidebarState", "closed");
}

// Check if the sidebar state is stored in local storage and set the sidebar accordingly
document.addEventListener("DOMContentLoaded", function() {
  const sidebarState = localStorage.getItem("sidebarState");
  if (sidebarState === "open") {
      openNav();
  } else {
      closeNav();
  }
});

// Function to toggle the sidebar state
function toggleNav() {
  const sidebarState = localStorage.getItem("sidebarState");
  if (sidebarState === "open") {
      closeNav();
  } else {
      openNav();
  }
}



