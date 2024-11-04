// src/components/AppVersionModal.js

import React from "react";
import "../styles/Modal.css"; // Make sure this contains the above CSS

const AppVersionModal = ({ onClose }) => {
    return (
        <div className="modal-overlay" onClick={onClose}>
            <div className="modal-content" onClick={(e) => e.stopPropagation()}>
                <div className="modal-header">
                    <span className="modal-title">App Version Detail</span>
                    <span className="modal-subtitle">Smoke Build</span>
                    <button className="modal-close" onClick={onClose}>×</button>
                </div>
                <div className="modal-body">
                    <ul className="modal-list">
                        <li><span className="status-icon green">✔️</span> <span className="app-name">awp</span> <span className="version-detail">develop-20240201</span></li>
                        <li><span className="status-icon orange">⏳</span> <span className="app-name">idsrv</span> <span className="version-detail">develop-20231113</span></li>
                        <li><span className="status-icon orange">⏳</span> <span className="app-name">portal</span> <span className="version-detail">develop-20240429</span></li>
                        <li><span className="status-icon red">❌</span> <span className="app-name">webapi</span> <span className="version-detail">develop-20240415</span></li>
                    </ul>
                </div>
            </div>
        </div>
    );
};

export default AppVersionModal;