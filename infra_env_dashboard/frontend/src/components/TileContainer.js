// src/components/TileContainer.js

import React from "react";
import Card from "./Card";
import "../styles/TileContainer.css";
import "../styles/Card.css";

function TileContainer({ environments }) { // Accept environments as a prop
    return (
        <div className="card-layout">
            <div className="tile-container">
                {environments && environments.length > 0 ? ( // Check if environments exist
                    <div className="card-grid">
                        {environments.map((env, index) => (
                            <Card
                                key={index}
                                name={env.name}
                                lastUpdated={env.lastUpdated}
                                status={env.status}
                                contact={env.contact}
                                appVersion={env.appVersion}
                                dbVersion={env.dbVersion}
                                comments={env.comments}
                                statusClass={env.statusClass}
                                applications={env.applications} // Pass application data
                            />
                        ))}
                    </div>
                ) : (
                    <div className="empty-state">
                        <p>Select an environment to view details.</p>
                    </div>
                )}
            </div>
        </div>
    );
}

export default TileContainer;