// src/components/PrivacyModal.tsx

import React from "react";
import "../styles/PrivacyModal.css";

interface PrivacyModalProps {
    isOpen: boolean;
    onClose: () => void;
}

const PrivacyModal: React.FC<PrivacyModalProps> = ({ isOpen, onClose }) => {
    if (!isOpen) return null;

    return (
        <div className="modal-overlay" onClick={onClose}>
            <div className="modal-content" onClick={(e) => e.stopPropagation()}>
                <div className="modal-header">
                    <h3>Privacy Policy</h3>
                    <button className="modal-close" onClick={onClose}>
                        ✖
                    </button>
                </div>
                <div className="modal-body">
                    <p>This dashboard is open-source under the OpenSource License...</p>
                </div>
            </div>
        </div>
    );
};

export default PrivacyModal;