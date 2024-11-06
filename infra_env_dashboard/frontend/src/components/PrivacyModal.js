// src/components/PrivacyModal.js

import React from "react";
import "../styles/Modal.css"; // Assuming you have modal styles or create specific styles for PrivacyModal

function PrivacyModal({ onClose }) {
    return (
        <div className="modal-overlay">
            <div className="modal-content">
                <button className="modal-close" onClick={onClose}>✖</button>
                <h2>Privacy Policy</h2>
                <p>This dashboard is open-source under the MIT License...</p> {/* Add full license details */}
            </div>
        </div>
    );
}

export default PrivacyModal;