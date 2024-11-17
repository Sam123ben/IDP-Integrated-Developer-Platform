// src/components/MainContent.tsx

import React from "react";
import Card from "./Card";
import "../styles/MainContent.css";

interface EnvDetails {
    id: string;
    name: string;
    lastUpdated: string;
    status: string;
    contact: string;
    appVersion: string;
    dbVersion: string;
    comments: string;
    url: string;
}

interface MainContentProps {
    envDetails: EnvDetails[];
}

const MainContent: React.FC<MainContentProps> = ({ envDetails }) => {
    return (
        <div className="main-content">
            {envDetails && envDetails.length > 0 ? (
                <div className="card-grid">
                    {envDetails.map((env) => (
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
                    ))}
                </div>
            ) : (
                <div className="empty-layout">
                    <p>Select an environment to view details.</p>
                </div>
            )}
        </div>
    );
};

export default MainContent;