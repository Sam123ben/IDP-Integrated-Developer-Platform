document.addEventListener("DOMContentLoaded", function() {
    const settingsIcon = document.getElementById("settings-icon");
    const settingsDropdown = document.getElementById("settings-dropdown");
    let hideTimeout;

    // Function to set a cookie
    function setCookie(name, value, days) {
        const date = new Date();
        date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
        document.cookie = `${name}=${value}; expires=${date.toUTCString()}; path=/`;
    }

    // Function to get a cookie by name
    function getCookie(name) {
        const nameEQ = name + "=";
        const ca = document.cookie.split(';');
        for(let i = 0; i < ca.length; i++) {
            let c = ca[i];
            while (c.charAt(0) == ' ') c = c.substring(1, c.length);
            if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length, c.length);
        }
        return null;
    }

    // Apply the stored theme on load
    const storedTheme = getCookie("theme");
    if (storedTheme === "dark") {
        document.body.classList.add("dark-theme");
    } else {
        document.body.classList.remove("dark-theme");
    }

    // Show dropdown when settings icon is clicked
    settingsIcon.addEventListener("click", function() {
        clearTimeout(hideTimeout); // Clear any existing timeout to hide dropdown
        settingsDropdown.classList.add("show");
    });

    // Show dropdown when hovering over settings icon
    settingsIcon.addEventListener("mouseenter", function() {
        clearTimeout(hideTimeout); // Clear any hiding timeouts if user hovers back
        settingsDropdown.classList.add("show");
    });

    // Show dropdown when hovering over dropdown itself
    settingsDropdown.addEventListener("mouseenter", function() {
        clearTimeout(hideTimeout); // Clear any hiding timeouts
        settingsDropdown.classList.add("show");
    });

    // Hide dropdown after a delay when mouse leaves dropdown
    settingsDropdown.addEventListener("mouseleave", function() {
        hideTimeout = setTimeout(function() {
            settingsDropdown.classList.remove("show");
        }, 2000); // Hide after 2 seconds
    });

    // Hide dropdown when clicking outside the dropdown or icon
    document.addEventListener("click", function(event) {
        if (!settingsIcon.contains(event.target) && !settingsDropdown.contains(event.target)) {
            settingsDropdown.classList.remove("show");
        }
    });

    // Handling theme change
    const lightThemeOption = document.getElementById("light-theme-option");
    const darkThemeOption = document.getElementById("dark-theme-option");

    if (lightThemeOption && darkThemeOption) {
        lightThemeOption.addEventListener("click", function() {
            document.body.classList.remove("dark-theme");
            settingsDropdown.classList.remove("show"); // Close dropdown after selecting
            setCookie("theme", "light", 7); // Store theme preference for 7 days
        });

        darkThemeOption.addEventListener("click", function() {
            document.body.classList.add("dark-theme");
            settingsDropdown.classList.remove("show"); // Close dropdown after selecting
            setCookie("theme", "dark", 7); // Store theme preference for 7 days
        });
    } else {
        console.error("Theme options not found in the DOM");
    }
});