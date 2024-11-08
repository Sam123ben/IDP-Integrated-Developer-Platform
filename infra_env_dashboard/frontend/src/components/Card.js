// src/components/Card.js

import React, { useState } from "react";
import AppVersionModal from "./AppVersionModal";
import "../styles/Card.css";

function Card({ name, lastUpdated, status, contact, appVersion, dbVersion, comments, url }) {
    const [showModal, setShowModal] = useState(false);

    const handleVersionClick = (event) => {
        event.stopPropagation();
        setShowModal(true); // Open the modal
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
        <div className="card">
            <div className="card-header">
                <div className="card-title-section">
                    <span className="status-indicator" style={{ backgroundColor: statusColor }}></span>
                    <span className="card-title">{name}</span>
                </div>
                <div className="card-updated-section">
                    <span className="card-updated">Last updated: {lastUpdated}</span>
                </div>
            </div>
            {/* Environment URL */}
            <a href={`https://${url}`} className="card-link" target="_blank" rel="noopener noreferrer">
                {url}
            </a>
            <p><strong>Status:</strong> <span className="status-text">{status}</span></p>
            <p><strong>Contact:</strong> {contact}</p>
            <p>
                <strong>App Version:</strong>{" "}
                <span 
                    className="version-clickable"
                    onClick={handleVersionClick} 
                >
                    {appVersion}
                </span>
            </p>
            <p><strong>Database Version:</strong> {dbVersion}</p>
            <p><strong>Comments:</strong> {comments}</p>

            {/* Show modal if needed */}
            {showModal && <AppVersionModal onClose={closeModal} envName={name} />}
        </div>
    );
}

export default Card;