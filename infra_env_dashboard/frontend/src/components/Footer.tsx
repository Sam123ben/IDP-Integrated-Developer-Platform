// src/components/Footer.tsx

import React, { useState } from "react";
import PrivacyModal from "./PrivacyModal";
import "../styles/Footer.css";

const Footer: React.FC = () => {
    const [isModalOpen, setIsModalOpen] = useState<boolean>(false);

    const openModal = () => setIsModalOpen(true);
    const closeModal = () => setIsModalOpen(false);

    return (
        <div className="footer">
            <p>
                © {new Date().getFullYear()} DevopsEnv-Dashboard -{" "}
                <span className="privacy-link" onClick={openModal}>
                    Privacy
                </span>
            </p>
            {isModalOpen && <PrivacyModal isOpen={isModalOpen} onClose={closeModal} />}
        </div>
    );
};

export default Footer;