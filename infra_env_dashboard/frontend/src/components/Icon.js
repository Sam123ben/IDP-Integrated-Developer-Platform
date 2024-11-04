// src/components/Icon.js

import React from 'react';
import './Header.css';

const Icon = ({ icon, onClick }) => {
    return (
        <span className={`icon icon-${icon}`} onClick={onClick} role="button">
            {icon === "refresh" ? "🔄" : "⚙️"}
        </span>
    );
};

export default Icon;