document.addEventListener("DOMContentLoaded", function() {
    const settingsContainer = document.getElementById("settings-container");
    const settingsDropdown = document.getElementById("settings-dropdown");
    const body = document.body;
    let hideTimeout;

    // Show the dropdown when hovering over the settings icon
    settingsContainer.addEventListener("mouseenter", function() {
        clearTimeout(hideTimeout); // Clear any existing timeout to hide dropdown
        settingsDropdown.classList.add("show");
    });

    // Show the dropdown when hovering over the dropdown itself
    settingsDropdown.addEventListener("mouseenter", function() {
        clearTimeout(hideTimeout); // Clear any hide timeout if hovering over dropdown
        settingsDropdown.classList.add("show");
    });

    // Hide the dropdown with a delay when mouse leaves the settings icon or dropdown
    function hideDropdownWithDelay() {
        hideTimeout = setTimeout(function() {
            settingsDropdown.classList.remove("show");
        }, 300); // Delay of 300ms before hiding
    }

    // Hide dropdown when mouse leaves settings icon
    settingsContainer.addEventListener("mouseleave", hideDropdownWithDelay);
    // Hide dropdown when mouse leaves dropdown itself
    settingsDropdown.addEventListener("mouseleave", hideDropdownWithDelay);

    // Close the dropdown if clicking outside the settings icon and dropdown
    document.addEventListener("click", function(event) {
        if (!settingsContainer.contains(event.target) && !settingsDropdown.contains(event.target)) {
            settingsDropdown.classList.remove("show");
        }
    });

    // Theme switching functionality
    const lightThemeOption = document.getElementById("light-theme-option");
    const darkThemeOption = document.getElementById("dark-theme-option");

    function setTheme(theme) {
        if (theme === "dark") {
            body.classList.add("dark-theme");
            setCookie("theme", "dark", 7); // Store theme preference in a cookie for 7 days
        } else {
            body.classList.remove("dark-theme");
            setCookie("theme", "light", 7); // Store theme preference in a cookie for 7 days
        }
    }

    // Event listeners for theme options
    lightThemeOption.addEventListener("click", function() {
        setTheme("light");
        settingsDropdown.classList.remove("show"); // Close dropdown after selecting
    });

    darkThemeOption.addEventListener("click", function() {
        setTheme("dark");
        settingsDropdown.classList.remove("show"); // Close dropdown after selecting
    });

    // Set a cookie
    function setCookie(name, value, days) {
        const date = new Date();
        date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
        document.cookie = `${name}=${value}; expires=${date.toUTCString()}; path=/`;
    }

    // Get a cookie by name
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
        body.classList.add("dark-theme");
    } else {
        body.classList.remove("dark-theme");
    }
});