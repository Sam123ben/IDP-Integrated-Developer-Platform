document.addEventListener("DOMContentLoaded", function() {
    const settingsContainer = document.getElementById("settings-container");
    const settingsDropdown = document.getElementById("settings-dropdown");
    const refreshContainer = document.getElementById("refresh-container");
    const refreshDropdown = document.getElementById("refresh-dropdown");
    const body = document.body;
    let autoRefreshInterval = null;

    // Dropdown handling with touch support
    function setupDropdownLogic(container, dropdown) {
        container.addEventListener("click", (event) => {
            dropdown.classList.toggle("show");
            event.stopPropagation();
        });

        document.addEventListener("click", (event) => {
            if (!container.contains(event.target)) {
                dropdown.classList.remove("show");
            }
        });
    }

    // Apply dropdown logic
    setupDropdownLogic(settingsContainer, settingsDropdown);
    setupDropdownLogic(refreshContainer, refreshDropdown);

    // Theme switching functionality
    const lightThemeOption = document.getElementById("light-theme-option");
    const darkThemeOption = document.getElementById("dark-theme-option");

    function applyTheme(theme) {
        body.classList.toggle("dark-theme", theme === "dark");
        localStorage.setItem("theme", theme);
    }

    lightThemeOption.addEventListener("click", () => applyTheme("light"));
    darkThemeOption.addEventListener("click", () => applyTheme("dark"));

    // Apply stored theme on page load
    const storedTheme = localStorage.getItem("theme") || "light";
    applyTheme(storedTheme);

    // Refresh functionality with loading indication
    async function fetchLatestData() {
        const dashboardElement = document.getElementById("env-data");
        const loadingMessage = document.getElementById("loading-message");
        const errorMessage = document.getElementById("error-message");
    
        // Show loading message and hide error
        loadingMessage.style.display = "block";
        errorMessage.style.display = "none";
        
        try {
            const response = await fetch('/api/latest-data');
            if (!response.ok) throw new Error("Network response was not ok");
            
            const data = await response.json();
            updateDashboard(data);
        } catch (error) {
            console.error("Failed to fetch latest data:", error);
            errorMessage.style.display = "block"; // Show error message
            dashboardElement.innerHTML = ""; // Clear the existing data display
        } finally {
            loadingMessage.style.display = "none"; // Hide loading message
        }
    }
    
    function updateDashboard(data) {
        const dashboardElement = document.getElementById("env-data");
        dashboardElement.innerHTML = "<ul class='env-list'></ul>";
        
        const envList = dashboardElement.querySelector(".env-list");
        data.forEach(env => {
            const envItem = document.createElement("li");
            envItem.classList.add("env-item");
            envItem.innerHTML = `<strong>${env.Name}</strong>: ${env.Description}`;
            envList.appendChild(envItem);
        });
    }

    // Debounced function to set auto-refresh interval
    function setAutoRefresh(interval) {
        if (autoRefreshInterval) clearInterval(autoRefreshInterval);
        if (interval) {
            autoRefreshInterval = setInterval(fetchLatestData, interval);
            localStorage.setItem("autoRefreshInterval", interval);
        } else {
            localStorage.removeItem("autoRefreshInterval");
        }
    }

    // Manual refresh on click
    refreshContainer.addEventListener("click", fetchLatestData);

    // Auto-refresh dropdown options
    document.getElementById("auto-refresh-10sec").addEventListener("click", () => setAutoRefresh(10000));
    document.getElementById("auto-refresh-1min").addEventListener("click", () => setAutoRefresh(60000));
    document.getElementById("auto-refresh-5min").addEventListener("click", () => setAutoRefresh(300000));
    document.getElementById("auto-refresh-off").addEventListener("click", () => setAutoRefresh(null));

    // Apply stored auto-refresh interval on page load
    const storedAutoRefresh = parseInt(localStorage.getItem("autoRefreshInterval"));
    if (storedAutoRefresh) setAutoRefresh(storedAutoRefresh);
    else setAutoRefresh(300000); // Default to 5 minutes

    // Navigation handling
    const navLinks = document.querySelectorAll(".nav-menu ul li a");
    const contentSections = document.querySelectorAll(".content-section");

    navLinks.forEach(link => {
        link.addEventListener("click", function(event) {
            event.preventDefault();

            // Toggle active classes
            navLinks.forEach(nav => nav.classList.remove("active"));
            contentSections.forEach(section => section.classList.remove("active"));

            // Activate clicked link and corresponding section
            link.classList.add("active");
            const sectionId = link.getAttribute("data-section");
            document.getElementById(sectionId).classList.add("active");
        });
    });

    // Ensure the "Home" section is displayed by default
    document.querySelector(".nav-menu ul li a[data-section='home']").classList.add("active");
    document.getElementById("home").classList.add("active");
});