// src/components/TileContainer.tsx

import React from "react";
import Card from "./Card";
import "../styles/TileContainer.css";

interface Environment {
    name: string;
    lastUpdated: string;
    status: string;
    contact: string;
    appVersion: string;
    dbVersion: string;
    comments: string;
    statusClass: string;
}

interface TileContainerProps {
    environments: Environment[];
}

const TileContainer: React.FC<TileContainerProps> = ({ environments }) => {
    return (
        <div className="tile-container">
            {environments && environments.length > 0 ? (
                environments.map((env, index) => (
                    <Card
                        key={index}
                        name={env.name}
                        lastUpdated={env.lastUpdated}
                        status={env.status}
                        contact={env.contact}
                        appVersion={env.appVersion}
                        dbVersion={env.dbVersion}
                        comments={env.comments}
                        url=""
                    />
                ))
            ) : (
                <p>Select an environment to view details.</p>
            )}
        </div>
    );
};

export default TileContainer;