// src/components/Header.js

import React, { useEffect, useState } from "react";
import "../styles/Header.css";
import config from "../config"; // Import the config

const Header = () => {
    const [companyName, setCompanyName] = useState("Loading...");

    useEffect(() => {
        console.log("Starting to fetch company name..."); // Debug log: Start of fetch

        // Fetch company name from the API
        fetch(config.fetchCompanyDetails)
            .then((response) => {
                console.log("API response received:", response); // Debug log: API response
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                return response.json();
            })
            .then((data) => {
                console.log("Parsed JSON data:", data); // Debug log: Parsed JSON data

                if (data && data.length > 0) {
                    console.log("Setting company name:", data[0].name); // Debug log: Company name from data
                    setCompanyName(data[0].name);
                } else {
                    console.warn("No company name found in response data."); // Warning log if data is empty
                    setCompanyName("Company Name Not Found");
                }
            })
            .catch((error) => {
                console.error("Error fetching company name:", error.message); // Error log
                setCompanyName("Error loading company name");
            });
    }, []);

    return (
        <header className="header">
            <div className="header-content">
                <div className="header-title">
                    <h1>{companyName}</h1>
                    <p>Monitor, Manage, and Optimize Your Infrastructure from a Single View</p>
                </div>
                <div className="header-icons">
                    <span className="icon" onClick={() => window.location.reload()} title="Refresh">🔄</span>
                </div>
            </div>
        </header>
    );
};

export default Header;