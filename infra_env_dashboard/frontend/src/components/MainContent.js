// src/components/MainContent.js

import React from "react";
import "../styles/MainContent.css";
import Card from "./Card";

function MainContent({ selectedSection }) {
    return (
        <div className="main-content">
            {selectedSection ? (
                // Render cards if a section is selected
                selectedSection.sections.map((section, index) => (
                    <Card
                        key={index}
                        name={section.name}
                        status="Online" // Placeholder values, replace with actual data as needed
                        lastUpdated="Just now"
                        contact="John Doe"
                        appVersion="1.0.0"
                        dbVersion="v2.0.1"
                        comments="Environment is active."
                    />
                ))
            ) : (
                // Empty layout when no section is selected
                <div className="empty-layout"></div>
            )}
        </div>
    );
}

export default MainContent;