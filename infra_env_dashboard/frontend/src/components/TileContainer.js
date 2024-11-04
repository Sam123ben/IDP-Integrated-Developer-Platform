// src/components/TileContainer.js

import React from "react";
import Card from "./Card";
import "../styles/TileContainer.css";
import "../styles/Card.css";

function TileContainer() {
    const environments = [
        {
            name: "SMOKE",
            lastUpdated: "19/08/2021 21:30",
            status: "Online",
            contact: "Samyak",
            appVersion: "Smoke Build",
            dbVersion: "7.2.0876",
            comments: "Testing this env so please check",
            statusClass: "card-online",
            applications: [
                { name: "awp", version: "develop-20240201", status: "green" },
                { name: "idsrv", version: "develop-20231113", status: "orange" },
                { name: "portal", version: "develop-20240429", status: "orange" },
                { name: "webapi", version: "develop-20240415", status: "red" }
            ]
        },
        // Other environments here...
    ];

    return (
        <div className="card-layout">
            <div className="tile-container">
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
            </div>
        </div>
    );
}

export default TileContainer;