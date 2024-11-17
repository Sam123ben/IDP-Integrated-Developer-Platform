// src/components/Modal.tsx

import React from "react";
import "../styles/Modal.css";

interface ModalProps {
    title: string;
    subtitle: string;
    children: React.ReactNode;
    onClose: () => void;
}

const Modal: React.FC<ModalProps> = ({ title, subtitle, children, onClose }) => {
    return (
        <div className="modal-overlay">
            <div className="modal-content">
                <div className="modal-header">
                    <h3 className="modal-title">{title}</h3>
                    <h4 className="modal-subtitle">{subtitle}</h4>
                    <button className="modal-close" onClick={onClose}>
                        ✖
                    </button>
                </div>
                <div className="modal-body">{children}</div>
            </div>
        </div>
    );
};

export default Modal;