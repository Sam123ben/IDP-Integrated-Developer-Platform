// src/components/Card.js

import React, { useState } from "react";
import AppVersionModal from "./AppVersionModal";
import "../styles/Card.css";

function Card({ name, lastUpdated, status, contact, appVersion, dbVersion, comments, url }) {
    const [showModal, setShowModal] = useState(false);

    const handleVersionClick = (event) => {
        event.stopPropagation();
        setShowModal(true);
    };

    const closeModal = () => {
        setShowModal(false);
    };

    // Determine the CSS class for status color based on status value
    const statusIndicatorClass =
        status === "Online"
            ? "status-indicator-green"
            : status === "In Progress"
            ? "status-indicator-orange"
            : status === "Offline"
            ? "status-indicator-red"
            : "status-indicator-default";

    const statusTextClass =
        status === "Online"
            ? "status-text-success"
            : status === "In Progress"
            ? "status-text-progress"
            : status === "Offline"
            ? "status-text-failed"
            : "status-text-default";

    // Dynamic class for card border color based on status
    const cardStatusClass =
        status === "Online"
            ? "card-online"
            : status === "In Progress"
            ? "card-in-progress"
            : status === "Offline"
            ? "card-offline"
            : "card-default";

    return (
        <div className={`card ${cardStatusClass}`}>
            <div className="card-header">
                <div className="card-title-section">
                    <span className={`status-indicator ${statusIndicatorClass}`}></span>
                    <span className="card-title">{name}</span>
                </div>
                <div className="card-updated-section">
                    <span className="card-updated">Last updated: {lastUpdated}</span>
                </div>
            </div>

            <a href={`https://${url}`} className="card-link" target="_blank" rel="noopener noreferrer">
                {url}
            </a>
            <p>
                <strong>Status:</strong> <span className={`status-text ${statusTextClass}`}>{status}</span>
            </p>
            <p>
                <strong>Contact:</strong> {contact}
            </p>
            <p>
                <strong>App Version:</strong>{" "}
                <span className="version-clickable" onClick={handleVersionClick}>
                    {appVersion}
                </span>
            </p>
            <p>
                <strong>Database Version:</strong> {dbVersion}
            </p>
            <p>
                <strong>Comments:</strong> {comments}
            </p>

            {showModal && <AppVersionModal onClose={closeModal} envName={name} />}
        </div>
    );
}

export default Card;