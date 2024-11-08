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

    // Determine the CSS class for status color based on status value
    const statusIndicatorClass = 
        status === "Online" ? "status-indicator-green" : 
        status === "Failed Deployment" ? "status-indicator-red" : 
        status === "Deployment In Progress" ? "status-indicator-orange" : 
        "status-indicator-gray"; // Default for unrecognized statuses

    const statusTextClass = 
        status === "Online" ? "status-text-success" : 
        status === "Failed Deployment" ? "status-text-failed" : 
        status === "Deployment In Progress" ? "status-text-progress" : 
        "status-text-default"; // Default text color for unrecognized statuses

    return (
        <div className="card">
            <div className="card-header">
                <div className="card-title-section">
                    <span className={`status-indicator ${statusIndicatorClass}`}></span>
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
            <p><strong>Status:</strong> <span className={`status-text ${statusTextClass}`}>{status}</span></p>
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