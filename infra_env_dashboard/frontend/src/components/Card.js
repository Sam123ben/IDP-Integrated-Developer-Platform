// src/components/Card.js

import React, { useState } from "react";
import "../styles/TileContainer.css"; // Import container-specific styles
import "../styles/Card.css"; // Import card-specific styles

function Card({ name, lastUpdated, status, contact, appVersion, dbVersion, comments, statusClass }) {
    const [showDropdown, setShowDropdown] = useState(false);

    const toggleDropdown = () => {
        setShowDropdown(!showDropdown);
    };

    return (
        <div className={`card ${statusClass} big-card`}>
            <div className="card-header">
                <div className="card-title-section">
                    <span className="card-icon">🔴</span>
                    <span className="card-title">{name}</span>
                </div>
                <div className="card-updated-section">
                    <span className="card-updated">Last updated: {lastUpdated}</span>
                    <span className="three-dots" onClick={toggleDropdown}>⋮</span>
                </div>
                {showDropdown && (
                    <div className="dropdown-menu">
                        <label>
                            <input type="checkbox" /> Skip Deployment
                        </label>
                        <label>
                            Update Comments
                        </label>
                    </div>
                )}
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
