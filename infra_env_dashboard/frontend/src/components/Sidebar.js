// src/components/Sidebar.js

import React, { useState, useEffect } from "react";
import "../styles/Sidebar.css";
import { fetchInfraTypes } from "../services/api"; // Replace with your actual data-fetching logic

function Sidebar({ onEnvironmentSelect }) {
    const [infraTypes, setInfraTypes] = useState([]);
    const [selectedSection, setSelectedSection] = useState("INTERNAL");
    const [expandedProducts, setExpandedProducts] = useState({});
    const [selectedEnv, setSelectedEnv] = useState(null); // Track selected environment
    const [selectedProduct, setSelectedProduct] = useState(null); // Track selected product

    // Fetch infrastructure types on load
    useEffect(() => {
        fetchInfraTypes()
            .then(setInfraTypes)
            .catch(error => console.error("Failed to load infrastructure types:", error));
    }, []);

    // Toggle expanded state for each product
    const toggleProduct = (productName) => {
        setExpandedProducts(prev => ({
            ...prev,
            [productName]: !prev[productName],
        }));
        setSelectedProduct(productName); // Set selected product
        setSelectedEnv(null); // Clear selected environment when selecting a new product
    };

    // Handle environment selection
    const handleEnvironmentSelect = (env) => {
        setSelectedEnv(env);
        setSelectedProduct(null); // Clear selected product when selecting a new environment
        onEnvironmentSelect(selectedSection, selectedProduct, env);
    };

    // Filter the infraTypes to show only the selected section (INTERNAL or CUSTOMER)
    const filteredInfraTypes = infraTypes.filter(
        (infraType) => infraType.name.toUpperCase() === selectedSection
    );

    return (
        <div className="sidebar-container">
            {/* Horizontal tabs for INTERNAL and CUSTOMER */}
            <div className="sidebar-tabs">
                <span
                    className={`sidebar-tab ${selectedSection === "INTERNAL" ? "active" : ""}`}
                    onClick={() => setSelectedSection("INTERNAL")}
                >
                    INTERNAL
                </span>
                <span
                    className={`sidebar-tab ${selectedSection === "CUSTOMER" ? "active" : ""}`}
                    onClick={() => setSelectedSection("CUSTOMER")}
                >
                    CUSTOMER
                </span>
            </div>

            {/* Display the products and environments under the selected section */}
            <div className="sidebar-content">
                {filteredInfraTypes.map((infraType) =>
                    infraType.sections.map((section) => (
                        <div key={section.name} className="product-section">
                            {/* Product header with toggle button */}
                            <div
                                className={`product-name ${selectedProduct === section.name ? "selected" : ""}`}
                                onClick={() => toggleProduct(section.name)}
                            >
                                {section.name}
                                <span className="toggle-icon">
                                    {expandedProducts[section.name] ? "▼" : "▶"}
                                </span>
                            </div>

                            {/* Environments under each product */}
                            {expandedProducts[section.name] && (
                                <ul className="environment-list">
                                    {section.environments.map((env) => (
                                        <li
                                            key={env}
                                            className={`environment-item ${selectedEnv === env ? "selected" : ""}`}
                                            onClick={() => handleEnvironmentSelect(env)}
                                        >
                                            {env}
                                        </li>
                                    ))}
                                </ul>
                            )}
                        </div>
                    ))
                )}
            </div>
        </div>
    );
}

export default Sidebar;