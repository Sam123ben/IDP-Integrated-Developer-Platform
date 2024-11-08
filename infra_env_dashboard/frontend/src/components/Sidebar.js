// src/components/Sidebar.js

import React, { useState, useEffect } from "react";
import "../styles/Sidebar.css";
import { fetchInfraTypes } from "../services/api";

function Sidebar({ onEnvironmentSelect }) {
    const [infraTypes, setInfraTypes] = useState([]);
    const [selectedSection, setSelectedSection] = useState("INTERNAL");
    const [expandedProducts, setExpandedProducts] = useState({});
    const [selectedProduct, setSelectedProduct] = useState(null);
    const [selectedEnv, setSelectedEnv] = useState(null);

    useEffect(() => {
        fetchInfraTypes()
            .then(setInfraTypes)
            .catch(error => console.error("Failed to load infrastructure types:", error));
    }, []);

    const toggleProduct = (productName) => {
        setExpandedProducts(prev => ({
            ...prev,
            [productName]: !prev[productName],
        }));
        setSelectedProduct(productName); // Update selected product
        setSelectedEnv(null); // Clear selected environment when a new product is selected
    };

    const handleEnvironmentSelect = (env) => {
        if (selectedProduct) {
            setSelectedEnv(env);
            onEnvironmentSelect(selectedSection, selectedProduct, env); // Pass both product and environment
        } else {
            console.warn("Product not selected"); // Log a warning if product is missing
        }
    };

    const filteredInfraTypes = infraTypes.filter(
        (infraType) => infraType.name.toUpperCase() === selectedSection
    );

    return (
        <div className="sidebar-container">
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

            <div className="sidebar-content">
                {filteredInfraTypes.map((infraType) =>
                    infraType.sections.map((section) => (
                        <div key={section.name} className="product-section">
                            <div
                                className={`product-name ${selectedProduct === section.name ? "selected" : ""}`}
                                onClick={() => toggleProduct(section.name)}
                            >
                                {section.name}
                                <span className="toggle-icon">
                                    {expandedProducts[section.name] ? "▼" : "▶"}
                                </span>
                            </div>

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