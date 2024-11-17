// src/components/Header.tsx

import React from "react";
import "../styles/Header.css";

// Define the type for companyDetails
interface CompanyDetails {
    name: string; // Add any other relevant fields for the company
}

interface HeaderProps {
    companyDetails: CompanyDetails;
}

const Header: React.FC<HeaderProps> = ({ companyDetails }) => {
    return (
        <header className="header">
            <div className="header-content">
                <h1>{companyDetails.name}</h1>
                <h2>Monitor, Manage, and Optimize Your Infrastructure from a Single View</h2>
            </div>
        </header>
    );
};

export default Header;