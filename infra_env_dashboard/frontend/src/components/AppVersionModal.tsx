// src/components/AppVersionModal.tsx

import React from "react";
import ReactDOM from "react-dom";
import "../styles/AppVersionModal.css";

interface AppVersionModalProps {
    onClose: () => void;
    envName?: string;
}

const AppVersionModal: React.FC<AppVersionModalProps> = ({ onClose, envName = "Environment Name" }) => {
    return ReactDOM.createPortal(
        <div className="modal-overlay">
            <div className="modal-content">
                <div className="modal-header">
                    <h2>App Version Detail</h2>
                    <div className="env-name-container">
                        <span className="env-name">{envName}</span>
                        <button onClick={onClose} className="close-button">
                            ✖
                        </button>
                    </div>
                </div>
                <div className="modal-body">
                    {/* Replace this static content with dynamic data if available */}
                    <ul>
                        <li>
                            <span className="status-icon success">✔️</span>
                            <span>awp</span>
                            <span className="version">develop-20240201</span>
                        </li>
                        {/* Add more list items as needed */}
                    </ul>
                </div>
            </div>
        </div>,
        document.body
    );
};

export default AppVersionModal;