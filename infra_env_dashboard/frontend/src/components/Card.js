// src/components/Card.js

import React, { useState } from "react";
import "../styles/TileContainer.css";
import "../styles/Card.css";
import Modal from "./Modal";

function Card({ name, lastUpdated, status, contact, appVersion, dbVersion, comments, statusClass }) {
    const [isModalOpen, setModalOpen] = useState(false);

    const statusColor = 
        status === "Online" ? "green" : 
        status === "Failed Deployment" ? "red" : 
        status === "Deployment In Progress" ? "#FF8C00" : "gray";

    const handleAppVersionClick = () => {
        setModalOpen(true);
    };

    const closeModal = () => {
        setModalOpen(false);
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
            <a 
                href="#"
                className="card-link"
                style={{ color: "blue", cursor: "pointer" }}
                onClick={handleAppVersionClick}
            >
                {appVersion}
            </a>
            <p><strong>Status:</strong> <span className="status-text">{status}</span></p>
            <p><strong>Contact:</strong> {contact}</p>
            <p><strong>Database Version:</strong> {dbVersion}</p>
            <p><strong>Comments:</strong> {comments}</p>

            {/* Modal component for displaying detailed version information */}
            {isModalOpen && (
                <Modal title="App Version Detail" subtitle="Smoke Build" onClose={closeModal}>
                    <ul className="modal-list">
                        <li><span className="status-icon green">✔️</span><span className="app-name">awp</span><span className="version-detail">develop-20240201</span></li>
                        <li><span className="status-icon orange">⏳</span><span className="app-name">idsrv</span><span className="version-detail">develop-20231113</span></li>
                        <li><span className="status-icon orange">⏳</span><span className="app-name">portal</span><span className="version-detail">develop-20240429</span></li>
                        <li><span className="status-icon red">❌</span><span className="app-name">webapi</span><span className="version-detail">develop-20240415</span></li>
                    </ul>
                </Modal>
            )}
        </div>
    );
}

export default Card;