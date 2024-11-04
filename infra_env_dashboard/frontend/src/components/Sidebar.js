import React, { useState } from "react";
import "../styles/Sidebar.css";

function Sidebar() {
    const [selectedSection, setSelectedSection] = useState("INTERNAL");
    const [expandedSections, setExpandedSections] = useState({});

    const toggleSection = (section) => {
        setExpandedSections((prev) => ({
            ...prev,
            [section]: !prev[section],
        }));
    };

    const handleTabClick = (section) => {
        setSelectedSection(section);
    };

    return (
        <div className="sidebar-container">
            <div className="sidebar-header">
                <h3 
                    className={selectedSection === "INTERNAL" ? "active" : ""}
                    onClick={() => handleTabClick("INTERNAL")}
                >
                    INTERNAL
                </h3>
                <h3 
                    className={selectedSection === "CUSTOMER" ? "active" : ""}
                    onClick={() => handleTabClick("CUSTOMER")}
                >
                    CUSTOMER
                </h3>
            </div>

            {selectedSection === "INTERNAL" && (
                <>
                    {/* Product 1 Section */}
                    <div className="sidebar-section">
                        <div
                            className={`collapsible-header ${expandedSections["Product 1"] ? "active" : ""}`}
                            onClick={() => toggleSection("Product 1")}
                        >
                            Product 1
                            <span className={`arrow ${expandedSections["Product 1"] ? "rotate" : ""}`}>▶</span>
                        </div>
                        <ul className={`collapsible-content ${expandedSections["Product 1"] ? "visible" : ""}`}>
                            <li>DEV</li>
                            <li>QA</li>
                            <li>CONSULT</li>
                            <li>PRESALES</li>
                        </ul>
                    </div>

                    {/* Product 2 Section */}
                    <div className="sidebar-section">
                        <div
                            className={`collapsible-header ${expandedSections["Product 2"] ? "active" : ""}`}
                            onClick={() => toggleSection("Product 2")}
                        >
                            Product 2
                            <span className={`arrow ${expandedSections["Product 2"] ? "rotate" : ""}`}>▶</span>
                        </div>
                        <ul className={`collapsible-content ${expandedSections["Product 2"] ? "visible" : ""}`}>
                            <li>DEV</li>
                            <li>QA</li>
                            <li>STAGING</li>
                        </ul>
                    </div>
                </>
            )}

            {selectedSection === "CUSTOMER" && (
                <>
                    {/* Vendor A Section */}
                    <div className="sidebar-section">
                        <div
                            className={`collapsible-header ${expandedSections["Vendor A"] ? "active" : ""}`}
                            onClick={() => toggleSection("Vendor A")}
                        >
                            Vendor A
                            <span className={`arrow ${expandedSections["Vendor A"] ? "rotate" : ""}`}>▶</span>
                        </div>
                        <ul className={`collapsible-content ${expandedSections["Vendor A"] ? "visible" : ""}`}>
                            <li>Product 1</li>
                            <li>Product 2</li>
                        </ul>
                    </div>

                    {/* Vendor B Section */}
                    <div className="sidebar-section">
                        <div
                            className={`collapsible-header ${expandedSections["Vendor B"] ? "active" : ""}`}
                            onClick={() => toggleSection("Vendor B")}
                        >
                            Vendor B
                            <span className={`arrow ${expandedSections["Vendor B"] ? "rotate" : ""}`}>▶</span>
                        </div>
                        <ul className={`collapsible-content ${expandedSections["Vendor B"] ? "visible" : ""}`}>
                            <li>Product 1</li>
                            <li>Product 2</li>
                        </ul>
                    </div>
                </>
            )}
        </div>
    );
}

export default Sidebar;