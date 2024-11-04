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
                    {/* Altitude Section */}
                    <div className="sidebar-section">
                        <div
                            className={`collapsible-header ${expandedSections["Altitude"] ? "active" : ""}`}
                            onClick={() => toggleSection("Altitude")}
                        >
                            Altitude
                            <span className={`arrow ${expandedSections["Altitude"] ? "rotate" : ""}`}>▶</span>
                        </div>
                        <ul className={`collapsible-content ${expandedSections["Altitude"] ? "visible" : ""}`}>
                            <li>DEV</li>
                            <li>QA</li>
                            <li>CONSULT</li>
                            <li>PRESALES</li>
                        </ul>
                    </div>

                    {/* Authority Section */}
                    <div className="sidebar-section">
                        <div
                            className={`collapsible-header ${expandedSections["Authority"] ? "active" : ""}`}
                            onClick={() => toggleSection("Authority")}
                        >
                            Authority
                            <span className={`arrow ${expandedSections["Authority"] ? "rotate" : ""}`}>▶</span>
                        </div>
                        <ul className={`collapsible-content ${expandedSections["Authority"] ? "visible" : ""}`}>
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