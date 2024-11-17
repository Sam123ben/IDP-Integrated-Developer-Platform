// src/components/CardMenu.tsx

import React, { useState, useRef } from "react";
import "../styles/CardMenu.css";

interface CardMenuProps {
    onSkipDeployment: (checked: boolean) => void;
    onUpdateComments: () => void;
}

const CardMenu: React.FC<CardMenuProps> = ({ onSkipDeployment, onUpdateComments }) => {
    const [menuVisible, setMenuVisible] = useState<boolean>(false);
    const menuRef = useRef<HTMLDivElement | null>(null);

    const toggleMenu = (e: React.MouseEvent) => {
        e.stopPropagation();
        setMenuVisible((prevVisible) => {
            console.log("Previous visibility:", prevVisible);
            return !prevVisible;
        });
    };

    const closeMenu = () => {
        setMenuVisible(false);
    };

    const handleClickOutside = (event: MouseEvent) => {
        if (menuRef.current && !menuRef.current.contains(event.target as Node)) {
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
        <div className="card-menu" ref={menuRef}>
            <span className="three-dots" onClick={toggleMenu}>
                ⋮
            </span>
            {menuVisible && (
                <div className="dropdown-menu">
                    <div className="dropdown-item">
                        <label>
                            Skip Deployment
                            <input
                                type="checkbox"
                                onChange={(e) => onSkipDeployment(e.target.checked)}
                            />
                        </label>
                    </div>
                    <div className="dropdown-item" onClick={onUpdateComments}>
                        Update Comments
                    </div>
                </div>
            )}
        </div>
    );
};

export default CardMenu;