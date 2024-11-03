document.addEventListener("DOMContentLoaded", function() {
    const settingsContainer = document.getElementById("settings-container");
    const settingsDropdown = document.getElementById("settings-dropdown");
    const refreshContainer = document.getElementById("refresh-container");
    const refreshDropdown = document.getElementById("refresh-dropdown");
    const body = document.body;
    let settingsHideTimeout;
    let refreshHideTimeout;
    let autoRefreshInterval = null;

    // Function to show dropdown
    function showDropdown(dropdown) {
        dropdown.classList.add("show");
    }

    // Function to hide dropdown with delay
    function hideDropdownWithDelay(dropdown, hideTimeoutVar) {
        clearTimeout(hideTimeoutVar);
        return setTimeout(() => {
            dropdown.classList.remove("show");
        }, 800); // Delay of 800ms
    }

    // Settings dropdown logic
    settingsContainer.addEventListener("mouseenter", () => {
        clearTimeout(settingsHideTimeout);
        showDropdown(settingsDropdown);
    });
    settingsDropdown.addEventListener("mouseenter", () => {
        clearTimeout(settingsHideTimeout);
        showDropdown(settingsDropdown);
    });

    settingsContainer.addEventListener("mouseleave", () => {
        settingsHideTimeout = hideDropdownWithDelay(settingsDropdown, settingsHideTimeout);
    });
    settingsDropdown.addEventListener("mouseleave", () => {
        settingsHideTimeout = hideDropdownWithDelay(settingsDropdown, settingsHideTimeout);
    });

    // Close the settings dropdown if clicking outside
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
            setCookie("theme", "dark", 7);
        } else {
            body.classList.remove("dark-theme");
            setCookie("theme", "light", 7);
        }
    }

    lightThemeOption.addEventListener("click", function() {
        setTheme("light");
        settingsDropdown.classList.remove("show");
    });

    darkThemeOption.addEventListener("click", function() {
        setTheme("dark");
        settingsDropdown.classList.remove("show");
    });

    // Cookie management functions
    function setCookie(name, value, days) {
        const date = new Date();
        date.setTime(date.getTime() + days * 86400000); // 86400000 ms in a day
        document.cookie = `${name}=${value}; expires=${date.toUTCString()}; path=/`;
    }

    function getCookie(name) {
        const nameEQ = name + "=";
        const ca = document.cookie.split(';');
        for (let c of ca) {
            c = c.trim();
            if (c.indexOf(nameEQ) === 0) return c.substring(nameEQ.length);
        }
        return null;
    }

    // Apply stored theme on page load
    const storedTheme = getCookie("theme");
    if (storedTheme === "dark") {
        body.classList.add("dark-theme");
    } else {
        body.classList.remove("dark-theme");
    }

    // Refresh functionality
    async function fetchLatestData() {
        try {
            const response = await fetch('/api/latest-data');
            const data = await response.json();
            updateDashboard(data);
        } catch (error) {
            console.error("Failed to fetch latest data:", error);
        }
    }

    function updateDashboard(data) {
        const dashboardElement = document.getElementById("env-data");
        dashboardElement.innerHTML = "";
        data.forEach(env => {
            const envItem = document.createElement("li");
            envItem.innerHTML = `<strong>${env.Name}</strong>: ${env.Description}`;
            dashboardElement.appendChild(envItem);
        });
    }

    // Set auto-refresh interval and store it in a cookie
    function setAutoRefresh(interval) {
        if (autoRefreshInterval) clearInterval(autoRefreshInterval);
        if (interval) {
            autoRefreshInterval = setInterval(fetchLatestData, interval);
            setCookie("autoRefreshInterval", interval, 7); // Store interval in cookie
        } else {
            setCookie("autoRefreshInterval", "", -1); // Delete the cookie if interval is null
        }
    }

    // Manual refresh on click
    refreshContainer.addEventListener("click", fetchLatestData);

    // Refresh dropdown logic
    refreshContainer.addEventListener("mouseenter", () => {
        clearTimeout(refreshHideTimeout);
        showDropdown(refreshDropdown);
    });
    refreshDropdown.addEventListener("mouseenter", () => {
        clearTimeout(refreshHideTimeout);
        showDropdown(refreshDropdown);
    });

    refreshContainer.addEventListener("mouseleave", () => {
        refreshHideTimeout = hideDropdownWithDelay(refreshDropdown, refreshHideTimeout);
    });
    refreshDropdown.addEventListener("mouseleave", () => {
        refreshHideTimeout = hideDropdownWithDelay(refreshDropdown, refreshHideTimeout);
    });

    // Close the refresh dropdown if clicking outside
    document.addEventListener("click", function(event) {
        if (!refreshContainer.contains(event.target) && !refreshDropdown.contains(event.target)) {
            refreshDropdown.classList.remove("show");
        }
    });

    // Auto-refresh dropdown options
    document.getElementById("auto-refresh-10sec").addEventListener("click", () => {
        setAutoRefresh(10000);
        refreshDropdown.classList.remove("show");
    });
    document.getElementById("auto-refresh-1min").addEventListener("click", () => {
        setAutoRefresh(60000);
        refreshDropdown.classList.remove("show");
    });
    document.getElementById("auto-refresh-5min").addEventListener("click", () => {
        setAutoRefresh(300000);
        refreshDropdown.classList.remove("show");
    });
    document.getElementById("auto-refresh-off").addEventListener("click", () => {
        setAutoRefresh(null);
        refreshDropdown.classList.remove("show");
    });

    // Apply stored auto-refresh interval on page load
    const storedAutoRefresh = getCookie("autoRefreshInterval");
    if (storedAutoRefresh) {
        setAutoRefresh(parseInt(storedAutoRefresh));
    } else {
        // Set default auto-refresh interval to 5 minutes if no preference is stored
        setAutoRefresh(300000); // 5 minutes in milliseconds
    }
});