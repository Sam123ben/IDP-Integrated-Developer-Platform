// src/components/TileContainer.js

import React from "react";
import "../styles/TileContainer.css";
import { FaExclamationTriangle } from "react-icons/fa"; // Import icon from react-icons

function TileContainer() {
    const environments = [
        {
            name: "SMOKE",
            lastUpdated: "19/08/2021 21:30",
            url: "https://smoke.example.com/",
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
            url: "https://dev.example.com/",
            status: "Failed Deployment",
            contact: "Taj",
            appVersion: "2021.07.27",
            dbVersion: "7.2.0555",
            comments: "Upgrade in progress",
            statusClass: "card-failed",
        },
    ];

    return (
        <div className="tile-container">
            <div className="card-grid">
                {environments.map((env, index) => (
                    <div key={index} className={`card ${env.statusClass}`}>
                        <div className="card-header">
                            <FaExclamationTriangle className="card-icon" /> {/* Icon */}
                            <h3 className="card-title">{env.name}</h3>
                            <div className="card-updated">
                                <span>Last updated: {env.lastUpdated}</span>
                                <span className="three-dots">⋮</span>
                            </div>
                        </div>
                        <a href={env.url} target="_blank" rel="noopener noreferrer" className="card-link">
                            {env.url}
                        </a>
                        <p><strong>Status:</strong> <span className="status-text">{env.status}</span></p>
                        <p><strong>Contact:</strong> {env.contact}</p>
                        <p><strong>App Version:</strong> {env.appVersion}</p>
                        <p><strong>Database Version:</strong> {env.dbVersion}</p>
                        <p><strong>Comments:</strong> {env.comments}</p>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default TileContainer;