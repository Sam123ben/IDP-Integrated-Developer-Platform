// src/components/TileContainer.js

import React from "react";
import "../styles/TileContainer.css";

function TileContainer() {
    // Example data for each environment card
    const environments = [
        {
            name: "SMOKE",
            lastUpdated: "19/08/2021 21:30",
            status: "Failed Deployment",
            contact: "Taj",
            appVersion: "2021.07.27",
            dbVersion: "7.2.0555",
            comments: "Update in progress",
            statusClass: "card-failed", // Assign class based on status
        },
        {
            name: "MANUAL",
            lastUpdated: "19/08/2021 21:30",
            status: "Deployment In Progress",
            contact: "Taj",
            appVersion: "2021.07.27",
            dbVersion: "7.2.0555",
            comments: "Update in progress",
            statusClass: "card-in-progress",
        },
        {
            name: "MANUAL VIC",
            lastUpdated: "19/08/2021 21:30",
            status: "Online",
            contact: "Taj",
            appVersion: "2021.07.27",
            dbVersion: "7.2.0555",
            comments: "Update in progress",
            statusClass: "card-online",
        },
        {
            name: "PRE LAUNCH",
            lastUpdated: "19/08/2021 21:30",
            status: "Online",
            contact: "Taj",
            appVersion: "2021.07.27",
            dbVersion: "7.2.0555",
            comments: "Update in progress",
            statusClass: "card-online",
        },
    ];

    return (
        <div className="tile-container">
            <div className="card-grid">
                {environments.map((env, index) => (
                    <div key={index} className={`card ${env.statusClass}`}>
                        <h3 className="card-title">{env.name}</h3>
                        <div className="card-section">
                            <p><strong>Last updated:</strong> {env.lastUpdated}</p>
                            <p><strong>Status:</strong> {env.status}</p>
                        </div>
                        <div className="card-section">
                            <p><strong>Contact:</strong> {env.contact}</p>
                        </div>
                        <div className="card-section">
                            <p><strong>App Version:</strong> {env.appVersion}</p>
                            <p><strong>Database Version:</strong> {env.dbVersion}</p>
                        </div>
                        <div className="card-section">
                            <p><strong>Comments:</strong> {env.comments}</p>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default TileContainer;