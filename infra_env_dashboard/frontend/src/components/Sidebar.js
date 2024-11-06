import React, { useState, useEffect } from "react";
import "../styles/Sidebar.css";

function Sidebar() {
    const [infraTypes, setInfraTypes] = useState([]); // Store fetched data
    const [selectedSection, setSelectedSection] = useState(null); // Track which section is selected
    const [expandedSections, setExpandedSections] = useState({});

    // Fetch infrastructure types data from the backend
    useEffect(() => {
        fetch("http://localhost:8081/api/infra-types")
            .then((response) => response.json())
            .then((data) => {
                console.log("Fetched infraTypes:", data); // Debug log
                setInfraTypes(data.infraTypes || []);
                if (data.infraTypes && data.infraTypes.length > 0) {
                    setSelectedSection(data.infraTypes[0].name); // Default to the first section if available
                }
            })
            .catch((error) => console.error("Error fetching infrastructure types:", error));
    }, []);

    const toggleSection = (section) => {
        setExpandedSections((prev) => ({
            ...prev,
            [section]: !prev[section],
        }));
    };

    const handleTabClick = (sectionType) => {
        setSelectedSection(sectionType);
    };

    return (
        <div className="sidebar-container">
            {/* Sidebar header with dynamic sections */}
            <div className="sidebar-header">
                {infraTypes.map((infraType) => (
                    <h3
                        key={infraType.id}
                        className={selectedSection === infraType.name ? "active" : ""}
                        onClick={() => handleTabClick(infraType.name)}
                    >
                        {infraType.name}
                    </h3>
                ))}
            </div>

            {/* Render sections based on the selected infrastructure type */}
            {infraTypes
                .filter((infraType) => infraType.name === selectedSection)
                .flatMap((infraType) =>
                    // Safely access infraType.sections and map over them if they exist
                    (infraType.sections || []).map((section) => (
                        <div className="sidebar-section" key={section.name}>
                            <div
                                className={`collapsible-header ${expandedSections[section.name] ? "active" : ""}`}
                                onClick={() => toggleSection(section.name)}
                            >
                                {section.name}
                                <span className={`arrow ${expandedSections[section.name] ? "rotate" : ""}`}>▶</span>
                            </div>
                            <ul className={`collapsible-content ${expandedSections[section.name] ? "visible" : ""}`}>
                                {/* Render environments if they exist within the section */}
                                {(section.environments || []).map((environment, idx) => (
                                    <li key={idx}>{environment}</li>
                                ))}
                            </ul>
                        </div>
                    ))
                )}
        </div>
    );
}

export default Sidebar;