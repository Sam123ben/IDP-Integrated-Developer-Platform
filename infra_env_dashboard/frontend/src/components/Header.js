// src/components/Header.js

import React, { useEffect, useState } from "react";
import "../styles/Header.css";
import { fetchCompanyDetails } from "../services/api";

const Header = () => {
    const [companyName, setCompanyName] = useState("Loading...");

    useEffect(() => {
        // Fetch company name from the API
        fetchCompanyDetails()
            .then((data) => {
                if (data && data.name) {
                    setCompanyName(data.name);
                } else {
                    setCompanyName("Company Name Not Found");
                }
            })
            .catch((error) => {
                console.error("Error fetching company name:", error.message);
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
                    <span className="icon" onClick={() => window.location.reload()} title="Refresh">
                        🔄
                    </span>
                </div>
            </div>
        </header>
    );
};

export default Header;