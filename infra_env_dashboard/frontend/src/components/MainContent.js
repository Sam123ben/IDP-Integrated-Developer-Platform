// src/components/MainContent.js

import React from "react";
import Card from "./Card";
import "../styles/MainContent.css";

function MainContent({ envDetails }) {
    return (
        <div className="main-content">
            {envDetails.length > 0 ? (
                envDetails.map((env, index) => (
                    <Card
                        key={env.id}
                        name={env.name}
                        lastUpdated={new Date(env.lastUpdated).toLocaleString()}
                        status={env.status}
                        contact={env.contact}
                        appVersion={env.appVersion}
                        dbVersion={env.dbVersion}
                        comments={env.comments}
                        url={env.url}
                    />
                ))
            ) : (
                <div className="empty-layout">
                    <p>Select an environment to view details.</p>
                </div>
            )}
        </div>
    );
}

export default MainContent;