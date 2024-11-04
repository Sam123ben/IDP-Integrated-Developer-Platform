// src/components/Header.js

import React, { useState, useEffect } from "react";
import "./Header.css";

const Header = () => {
    const [theme, setTheme] = useState("light"); // Default theme is "light"
    const [dropdownVisible, setDropdownVisible] = useState(false);

    useEffect(() => {
        // Load saved theme from localStorage on initial load
        const savedTheme = localStorage.getItem("theme");
        if (savedTheme) {
            setTheme(savedTheme);
            document.body.className = savedTheme; // Apply the saved theme on load
        }
    }, []);

    const toggleDropdown = () => {
        setDropdownVisible((prev) => !prev);
    };

    const handleThemeChange = (selectedTheme) => {
        setTheme(selectedTheme);
        document.body.className = selectedTheme; // Apply the theme to body
        localStorage.setItem("theme", selectedTheme); // Save theme preference
        setDropdownVisible(false); // Close the dropdown after selection
    };

    // Hide the dropdown when clicking outside
    useEffect(() => {
        const handleClickOutside = (event) => {
            if (dropdownVisible && !event.target.closest(".theme-dropdown")) {
                setDropdownVisible(false);
            }
        };
        document.addEventListener("click", handleClickOutside);
        return () => {
            document.removeEventListener("click", handleClickOutside);
        };
    }, [dropdownVisible]);

    return (
        <header className="header">
            <div className="header-content">
                <div className="header-title">
                    <h1>My Company</h1>
                    <p>Monitor, Manage, and Optimize Your Infrastructure from a Single View</p>
                </div>
                <div className="header-icons">
                    <span className="icon" onClick={() => window.location.reload()}>🔄</span>
                    
                    {/* Theme icon with dropdown */}
                    <div 
                        className="theme-dropdown" 
                        onClick={toggleDropdown}
                    >
                        <span className="icon">⚙️</span>
                        <div className={`dropdown-menu ${dropdownVisible ? 'dropdown-menu-visible' : ''}`}>
                            <div className="dropdown-item" onClick={() => handleThemeChange("light")}>🌞 Light Theme</div>
                            <div className="dropdown-item" onClick={() => handleThemeChange("dark")}>🌙 Dark Theme</div>
                        </div>
                    </div>
                </div>
            </div>
            <nav className="header-nav">
                <ul>
                    <li><a href="/" className="nav-link">Home</a></li>
                    <li><a href="/environments" className="nav-link">Environments/Infra</a></li>
                    <li><a href="/build" className="nav-link">Build Pipeline</a></li>
                    <li><a href="/deploy" className="nav-link active">Deployment Pipeline</a></li>
                </ul>
            </nav>
        </header>
    );
};

export default Header;