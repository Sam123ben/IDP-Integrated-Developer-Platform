// src/components/CardMenu.js

import React, { useState, useRef } from "react";
import "../styles/CardMenu.css";

const CardMenu = ({ onSkipDeployment, onUpdateComments }) => {
    const [menuVisible, setMenuVisible] = useState(false);
    const menuRef = useRef(null);

    const toggleMenu = (e) => {
        e.stopPropagation();
        setMenuVisible((prevVisible) => !prevVisible);
    };

    const closeMenu = () => setMenuVisible(false);

    const handleClickOutside = (event) => {
        if (menuRef.current && !menuRef.current.contains(event.target)) {
            closeMenu();
        }
    };

    React.useEffect(() => {
        document.addEventListener("click", handleClickOutside);
        return () => {
            document.removeEventListener("click", handleClickOutside);
        };
    }, []);

    return (
        <div className="card-menu">
            <span className="three-dots" onClick={toggleMenu}>
                ⋮
            </span>
            {menuVisible && (
                <div className="dropdown-menu" ref={menuRef}>
                    <div className="dropdown-item">
                        <label>
                            Skip Deployment
                            <input
                                type="checkbox"
                                onChange={(e) => onSkipDeployment(e.target.checked)}
                            />
                        </label>
                    </div>
                    <div
                        className="dropdown-item"
                        onClick={onUpdateComments}
                    >
                        Update Comments
                    </div>
                </div>
            )}
        </div>
    );
};

export default CardMenu;