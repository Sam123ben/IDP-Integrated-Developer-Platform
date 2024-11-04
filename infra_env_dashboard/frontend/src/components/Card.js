// src/components/Card.js

import React, { useState } from "react";
import "../styles/TileContainer.css";
import "../styles/Card.css";

function Card({ name, lastUpdated, status, contact, appVersion, dbVersion, comments, statusClass, applications }) {
    const [isDialogOpen, setIsDialogOpen] = useState(false); // State to control dialog visibility

    const statusColor =
        status === "Online" ? "green" :
        status === "Failed Deployment" ? "red" :
        status === "Deployment In Progress" ? "#FF8C00" : "gray";

    const toggleDialog = () => {
        setIsDialogOpen(!isDialogOpen);
    };

    return (
        <div className={`card ${statusClass} big-card`}>
            <div className="card-header">
                <div className="card-title-section">
                    <span className="status-indicator" style={{ backgroundColor: statusColor }}></span>
                    <span className="card-title">{name}</span>
                </div>
                <div className="card-updated-section">
                    <span className="card-updated">Last updated: {lastUpdated}</span>
                    <span className="three-dots">⋮</span>
                </div>
            </div>
            <a href="#" className="card-link">{`https://${name.toLowerCase()}.example.com/`}</a>
            <p><strong>Status:</strong> <span className="status-text">{status}</span></p>
            <p><strong>Contact:</strong> {contact}</p>
            <p>
                <strong>App Version:</strong> 
                <span
                    className="app-version-link"
                    onClick={toggleDialog}
                >
                    {appVersion}
                </span>
            </p>
            <p><strong>Database Version:</strong> {dbVersion}</p>
            <p><strong>Comments:</strong> {comments}</p>

            {isDialogOpen && (
                <div className="dialog-overlay" onClick={toggleDialog}>
                    <div className="dialog-box" onClick={(e) => e.stopPropagation()}>
                        <div className="dialog-header">
                            <h3>App Version Detail - {appVersion}</h3>
                            <span className="close-button" onClick={toggleDialog}>×</span>
                        </div>
                        <div className="dialog-content">
                            <ul className="app-list">
                                {applications.map((app, index) => (
                                    <li key={index}>
                                        <span className={`status-icon ${app.status}`}>●</span>
                                        <span className="app-name">{app.name}</span>: {app.version}
                                    </li>
                                ))}
                            </ul>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
}

export default Card;