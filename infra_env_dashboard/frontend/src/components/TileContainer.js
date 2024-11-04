// src/components/TileContainer.js

import React from "react";
import Card from "./Card"; // Import the Card component
import "../styles/TileContainer.css"; // Import container-specific styles
import "../styles/Card.css"; // Import card-specific styles

function TileContainer() {
    const environments = [
        {
            name: "SMOKE",
            lastUpdated: "19/08/2021 21:30",
            status: "Failed Deployment",
            contact: "Taj",
            appVersion: "2021.07.27",
            dbVersion: "7.2.0555",
            comments: "Upgrade in progress",
            statusClass: "card-failed",
        },
        {
            name: "DEV",
            lastUpdated: "19/08/2021 21:30",
            status: "Deployment In Progress",
            contact: "Taj",
            appVersion: "2021.07.27",
            dbVersion: "7.2.0555",
            comments: "Upgrade in progress",
            statusClass: "card-in-progress",
        },
        {
            name: "QA",
            lastUpdated: "19/08/2021 21:30",
            status: "Online",
            contact: "Taj",
            appVersion: "2021.07.27",
            dbVersion: "7.2.0555",
            comments: "Running smoothly",
            statusClass: "card-online",
        },
    ];

    return (
        <div className="card-layout"> {/* Wrapper for the main card layout */}
            <div className="tile-container"> {/* Container for grid layout */}
                <div className="card-grid">
                    {environments.map((env, index) => (
                        <Card
                            key={index}
                            name={env.name}
                            lastUpdated={env.lastUpdated}
                            status={env.status}
                            contact={env.contact}
                            appVersion={env.appVersion}
                            dbVersion={env.dbVersion}
                            comments={env.comments}
                            statusClass={env.statusClass}
                        />
                    ))}
                </div>
            </div>
        </div>
    );
}

export default TileContainer;