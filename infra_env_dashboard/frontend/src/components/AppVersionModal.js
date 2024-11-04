// src/components/AppVersionModal.js

import React from "react";
import ReactDOM from "react-dom";
import "../styles/AppVersionModal.css"; // Import modal-specific styles

const AppVersionModal = ({ onClose, envName = "Smoke Build" }) => {
    return ReactDOM.createPortal(
        <div className="modal-overlay">
            <div className="modal-content">
                <div className="modal-header">
                    <h2>App Version Detail</h2>
                    <span className="env-name">{envName}</span>
                    <button onClick={onClose} className="close-button">✖️</button>
                </div>
                <div className="modal-body">
                    <ul>
                        <li>
                            <span className="status-icon success">✔️</span>
                            <span>awp</span> 
                            <span className="version">develop-20240201</span>
                        </li>
                        <li>
                            <span className="status-icon in-progress">⏳</span>
                            <span>idsrv</span> 
                            <span className="version">develop-20231113</span>
                        </li>
                        <li>
                            <span className="status-icon in-progress">⏳</span>
                            <span>portal</span> 
                            <span className="version">develop-20240429</span>
                        </li>
                        <li>
                            <span className="status-icon failed">❌</span>
                            <span>webapi</span> 
                            <span className="version">develop-20240415</span>
                        </li>
                    </ul>
                </div>
            </div>
        </div>,
        document.body // Render modal at the root of the document
    );
};

export default AppVersionModal;