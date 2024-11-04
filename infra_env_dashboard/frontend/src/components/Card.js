// src/components/Card.js

import React, { useState } from "react";
import AppVersionModal from "./AppVersionModal"; // Import the modal component
import "../styles/TileContainer.css";
import "../styles/Card.css";

function Card({ name, lastUpdated, status, contact, appVersion, dbVersion, comments, statusClass }) {
    const [showModal, setShowModal] = useState(false); // Control modal visibility

    const handleVersionClick = (event) => {
        event.stopPropagation(); // Prevents propagation to other elements
        setShowModal(true); // Show the modal
    };

    const closeModal = () => {
        setShowModal(false); // Close the modal
    };

    const statusColor = 
        status === "Online" ? "green" : 
        status === "Failed Deployment" ? "red" : 
        status === "Deployment In Progress" ? "#FF8C00" : 
        "gray";

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
            {/* URL for the environment */}
            <a href={`https://${name.toLowerCase()}.example.com`} className="card-link" target="_blank" rel="noopener noreferrer">
                https://{name.toLowerCase()}.example.com
            </a>
            <p><strong>Status:</strong> <span className="status-text">{status}</span></p>
            <p><strong>Contact:</strong> {contact}</p>
            <p>
                <strong>App Version:</strong>{" "}
                <span 
                    className="version-clickable"
                    onClick={handleVersionClick} 
                    style={{ color: "#007BFF", cursor: "pointer", textDecoration: "underline" }}
                >
                    {appVersion}
                </span>
            </p>
            <p><strong>Database Version:</strong> {dbVersion}</p>
            <p><strong>Comments:</strong> {comments}</p>

            {/* Render the modal */}
            {showModal && <AppVersionModal onClose={closeModal} />}
        </div>
    );
}

export default Card;