// src/components/Card.js

import React from "react";
import "../styles/TileContainer.css"; // Import container-specific styles
import "../styles/Card.css"; // Import card-specific styles

function Card({ name, lastUpdated, status, contact, appVersion, dbVersion, comments, statusClass }) {
    // Determine the color of the status indicator based on the status
    const statusColor = 
        status === "Online" ? "green" : 
        status === "Failed Deployment" ? "red" : 
        status === "Deployment In Progress" ? "#FF8C00" : // Deep orange for better visibility
        "gray"; // Default color

    return (
        <div className={`card ${statusClass} big-card`}>
            <div className="card-header">
                <div className="card-title-section">
                    {/* Apply conditional style for status color */}
                    <span 
                        className="status-indicator" 
                        style={{ backgroundColor: statusColor }}
                    ></span>
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
            <p><strong>App Version:</strong> {appVersion}</p>
            <p><strong>Database Version:</strong> {dbVersion}</p>
            <p><strong>Comments:</strong> {comments}</p>
        </div>
    );
}

export default Card;