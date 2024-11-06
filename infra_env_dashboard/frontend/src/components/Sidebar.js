import React, { useState, useEffect } from "react";
import "../styles/Sidebar.css";

function Sidebar({ onSectionSelect, onEnvironmentSelect }) { // Add props for parent communication
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
                    onSectionSelect(data.infraTypes[0].name); // Notify parent of the initial selection
                }
            })
            .catch((error) => console.error("Error fetching infrastructure types:", error));
    }, [onSectionSelect]);

    // Toggle section visibility
    const toggleSection = (section) => {
        setExpandedSections((prev) => ({
            ...prev,
            [section]: !prev[section],
        }));
    };

    // Handle section tab click
    const handleTabClick = (sectionType) => {
        setSelectedSection(sectionType);
        onSectionSelect(sectionType); // Notify parent of section selection
    };

    // Handle environment click
    const handleEnvironmentClick = (environment) => {
        onEnvironmentSelect(environment); // Notify parent of environment selection
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
                                {(section.environments || []).map((environment, idx) => (
                                    <li key={idx} onClick={() => handleEnvironmentClick(environment)}>
                                        {environment}
                                    </li>
                                ))}
                            </ul>
                        </div>
                    ))
                )}
        </div>
    );
}

export default Sidebar;