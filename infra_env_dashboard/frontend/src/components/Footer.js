// src/components/Footer.js

import React, { useState } from "react";
import PrivacyModal from "./PrivacyModal"; // Import the new PrivacyModal component
import "../styles/Footer.css";

function Footer() {
    const [isModalOpen, setIsModalOpen] = useState(false);

    const openModal = () => setIsModalOpen(true);
    const closeModal = () => setIsModalOpen(false);

    return (
        <div className="footer">
            <p>
                © 2024 DevopsEnv-Dashboard -{" "}
                <span className="privacy-link" onClick={openModal}>
                    Privacy
                </span>
            </p>
            <PrivacyModal isOpen={isModalOpen} onClose={closeModal} />
        </div>
    );
}

export default Footer;