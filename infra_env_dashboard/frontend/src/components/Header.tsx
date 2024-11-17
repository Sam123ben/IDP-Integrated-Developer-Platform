// src/components/Header.tsx

import React, { useState, useEffect } from "react";
import "../styles/Header.css";
import { FiRefreshCw, FiSettings } from "react-icons/fi";
import { FaSun, FaMoon } from "react-icons/fa";

// Define the type for companyDetails
interface CompanyDetails {
    name: string; // Add any other relevant fields for the company
}

interface HeaderProps {
    companyDetails: CompanyDetails;
}

const Header: React.FC<HeaderProps> = ({ companyDetails }) => {
    const [isDropdownVisible, setIsDropdownVisible] = useState(false);

    // Toggle dropdown visibility
    const toggleDropdown = () => {
        setIsDropdownVisible((prev) => !prev);
    };

    // Handle click outside of the dropdown to close it
    const handleClickOutside = (event: MouseEvent) => {
        const target = event.target as HTMLElement;
        if (!target.closest(".theme-dropdown") && !target.closest(".icon.settings")) {
            setIsDropdownVisible(false);
        }
    };

    // Add click event listener to detect clicks outside of dropdown
    useEffect(() => {
        if (isDropdownVisible) {
            window.addEventListener("click", handleClickOutside);
        } else {
            window.removeEventListener("click", handleClickOutside);
        }

        // Cleanup function to remove event listener
        return () => {
            window.removeEventListener("click", handleClickOutside);
        };
    }, [isDropdownVisible]);

    return (
        <header className="header">
            <div className="header-content">
                <h1 className="header-title">{companyDetails.name}</h1>
                <h2 className="header-tagline">Monitor, Manage, and Optimize Your Infrastructure from a Single View</h2>
            </div>
            <div className="header-icons">
                {/* Refresh Icon */}
                <FiRefreshCw className="icon refresh" onClick={() => window.location.reload()} />

                {/* Settings Icon with Dropdown */}
                <div className="theme-dropdown">
                    <FiSettings className="icon settings" onClick={toggleDropdown} />
                    {isDropdownVisible && (
                        <div className="dropdown-menu dropdown-menu-visible">
                            <div className="dropdown-item">
                                <FaSun /> Light Theme
                            </div>
                            <div className="dropdown-item">
                                <FaMoon /> Dark Theme
                            </div>
                        </div>
                    )}
                </div>
            </div>
        </header>
    );
};

export default Header;