// src/components/TileContainer.js

import React from "react";
import Card from "./Card";
import "../styles/TileContainer.css";

function TileContainer() {
    return (
        <div className="tile-container">
            <Card title="SMOKE" status="Failed Deployment" statusColor="red" lastUpdated="19/08/2021 21:30" />
            <Card title="MANUAL" status="Deployment In Progress" statusColor="orange" lastUpdated="19/08/2021 21:30" />
            <Card title="MANUAL VIC" status="Online" statusColor="green" lastUpdated="19/08/2021 21:30" />
            <Card title="PRE LAUNCH" status="Online" statusColor="green" lastUpdated="19/08/2021 21:30" />
        </div>
    );
}

export default TileContainer;