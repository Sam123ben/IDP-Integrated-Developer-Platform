// src/components/Header.js

import React, { useState, useEffect } from 'react';
import './Header.css';
import Icon from './Icon';

const Header = () => {
    const [theme, setTheme] = useState('light');
    const [dropdownVisible, setDropdownVisible] = useState(false);

    useEffect(() => {
        // Check session storage for a saved theme on component mount
        const savedTheme = sessionStorage.getItem('theme');
        if (savedTheme) {
            setTheme(savedTheme);
            document.body.className = savedTheme;
        }
    }, []);

    const toggleTheme = (selectedTheme) => {
        setTheme(selectedTheme);
        document.body.className = selectedTheme;

        // Save the selected theme in session storage
        sessionStorage.setItem('theme', selectedTheme);

        setDropdownVisible(false); // Hide the dropdown after selecting
    };

    return (
        <header className="header">
            <div className="header-content">
                <div className="header-title">
                    <h1>My Company</h1>
                    <p>Monitor, Manage, and Optimize Your Infrastructure from a Single View</p>
                </div>
                <div className="header-icons">
                    <Icon icon="refresh" onClick={() => window.location.reload()} />

                    {/* Theme icon with dropdown */}
                    <div 
                        className="theme-dropdown" 
                        onMouseEnter={() => setDropdownVisible(true)}
                        onMouseLeave={() => setDropdownVisible(false)}
                        onClick={() => setDropdownVisible(!dropdownVisible)}
                    >
                        <Icon icon="settings" />
                        {dropdownVisible && (
                            <div className="dropdown-menu">
                                <div className="dropdown-item" onClick={() => toggleTheme('light')}>
                                    🌞 Light Theme
                                </div>
                                <div className="dropdown-item" onClick={() => toggleTheme('dark')}>
                                    🌙 Dark Theme
                                </div>
                            </div>
                        )}
                    </div>
                </div>
            </div>
            <nav className="header-nav">
                <ul>
                    <li><a href="/">Home</a></li>
                    <li><a href="/environments">Environments/Infra</a></li>
                    <li><a href="/build">Build Pipeline</a></li>
                    <li><a href="/deploy" className="active">Deployment Pipeline</a></li>
                </ul>
            </nav>
        </header>
    );
};

export default Header;