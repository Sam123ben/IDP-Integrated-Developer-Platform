document.addEventListener("DOMContentLoaded", function() {
    const settingsIcon = document.getElementById("settings-icon");
    const settingsDropdown = document.getElementById("settings-dropdown");

    // Toggle dropdown visibility when settings icon is clicked
    settingsIcon.addEventListener("click", function() {
        settingsDropdown.classList.toggle("show");
    });

    // Close dropdown when clicking outside
    document.addEventListener("click", function(event) {
        if (!settingsIcon.contains(event.target) && !settingsDropdown.contains(event.target)) {
            settingsDropdown.classList.remove("show");
        }
    });

    // Handling theme change
    const lightThemeOption = document.getElementById("light-theme-option");
    const darkThemeOption = document.getElementById("dark-theme-option");

    lightThemeOption.addEventListener("click", function() {
        document.body.classList.remove("dark-theme");
        settingsDropdown.classList.remove("show"); // Close dropdown after selecting
    });

    darkThemeOption.addEventListener("click", function() {
        document.body.classList.add("dark-theme");
        settingsDropdown.classList.remove("show"); // Close dropdown after selecting
    });
});