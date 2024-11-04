// src/components/Card.js

import React from "react";
import "../styles/Card.css";

function Card({ title, status, statusColor, lastUpdated }) {
    return (
        <div className="card" style={{ borderColor: statusColor }}>
            <div className="card-header">
                <h4>{title}</h4>
                <span className="last-updated">Last updated: {lastUpdated}</span>
            </div>
            <div className="card-body">
                <p><strong>Status:</strong> <span style={{ color: statusColor }}>{status}</span></p>
                <p><strong>Contact:</strong> Taj</p>
                <p><strong>App Version:</strong> 2021.07.27</p>
                <p><strong>Database Version:</strong> 7.2.0555</p>
                <p><strong>Comments:</strong> Update in progress and run the pipeline and check</p>
            </div>
        </div>
    );
}

export default Card;
