// src/components/Header.js

import React, { useState } from "react";
import { NavLink } from "react-router-dom";
import "./Header.css";

const Header = () => {
    const [dropdownVisible, setDropdownVisible] = useState(false);

    const toggleDropdown = () => {
        setDropdownVisible(!dropdownVisible);
    };

    const setTheme = (theme) => {
        document.body.className = theme;
        sessionStorage.setItem("theme", theme);
        setDropdownVisible(false);
    };

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
                    <div className="theme-dropdown">
                        <span className="icon" onClick={toggleDropdown}>⚙️</span>
                        {dropdownVisible && (
                            <div className="dropdown-menu">
                                <div className="dropdown-item" onClick={() => setTheme("light")}>🌞 Light Theme</div>
                                <div className="dropdown-item" onClick={() => setTheme("dark")}>🌙 Dark Theme</div>
                            </div>
                        )}
                    </div>
                </div>
            </div>
            <nav className="header-nav">
                <ul>
                    <li><NavLink exact="true" to="/" className="nav-link" activeclassname="active">Home</NavLink></li>
                    <li><NavLink to="/environments" className="nav-link" activeclassname="active">Environments/Infra</NavLink></li>
                    <li><NavLink to="/build" className="nav-link" activeclassname="active">Build Pipeline</NavLink></li>
                    <li><NavLink to="/deploy" className="nav-link" activeclassname="active">Deployment Pipeline</NavLink></li>
                </ul>
            </nav>
        </header>
    );
};

export default Header;