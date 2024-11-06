// src/components/Footer.js

import React, { useState } from "react";
import "../styles/Footer.css";

function Footer() {
    const [showPrivacyDialog, setShowPrivacyDialog] = useState(false);

    const openPrivacyDialog = () => {
        setShowPrivacyDialog(true);
    };

    const closePrivacyDialog = () => {
        setShowPrivacyDialog(false);
    };

    return (
        <div className="footer">
            <p>
                © 2024 DevopsEnv-Dashboard -{" "}
                <span className="privacy-link" onClick={openPrivacyDialog}>
                    Privacy
                </span>
            </p>
            {showPrivacyDialog && (
                <div className="privacy-dialog-overlay">
                    <div className="privacy-dialog">
                        <div className="privacy-header">
                            <h3>Privacy Policy</h3>
                            <button className="close-button" onClick={closePrivacyDialog}>✖</button>
                        </div>
                        <p>This dashboard is open-source under the MIT License...</p>
                    </div>
                </div>
            )}
        </div>
    );
}

export default Footer;
