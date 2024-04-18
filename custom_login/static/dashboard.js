// ACTIVE BUTTON AND CONTENT REMAIN AFTER RELOAD
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

// SIDEBAR NAVIGATION TOGGLE
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

// Function to handle logout
function handleLogout() {
  // Clear sidebar state and active button ID from local storage
  localStorage.removeItem("sidebarState");
  localStorage.removeItem("activeButtonId");
  
  // Close the sidebar
  closeNav();
  
  // Reset active class and display content for Button 1 (Random Quotes)
  const randomQuotesButton = document.getElementById("random-quotes-btn");
  if (randomQuotesButton) {
    // Add active class to Button 1
    randomQuotesButton.classList.add("active");
    // Call the function to display content for Button 1
    showRandomQuotes();
  }
}

// ZOOM AUTHOR PROFILE
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

// HIDE AND REVEAL THE QUOTE CONTENT
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
  // Get the active button ID from local storage
  const activeButtonId = localStorage.getItem("activeButtonId");

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
      // Add cases for other buttons if needed
    }
  }

  // Add event listeners to all buttons
  buttons.forEach(button => {
    button.addEventListener("click", handleButtonClick);
  });

  // If there's an active button ID stored, set the corresponding button as active
  if (activeButtonId) {
    const activeButton = document.getElementById(activeButtonId);
    if (activeButton) {
      activeButton.classList.add("active");
      
      // Trigger the click event of the active button to ensure proper styling and content display
      activeButton.click();
    }
  } else {
    // If no active button ID is stored (e.g., when the user logs in for the first time), 
    // default to the Random Quotes button
    const randomQuotesButton = document.getElementById("random-quotes-btn");
    if (randomQuotesButton) {
      randomQuotesButton.classList.add("active");
      randomQuotesButton.click(); // Trigger the click event to ensure proper styling and content display
    }
  }
});

