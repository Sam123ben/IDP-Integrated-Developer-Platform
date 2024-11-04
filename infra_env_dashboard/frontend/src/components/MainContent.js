// src/components/MainContent.js

import React from "react";
import "../styles/MainContent.css";

function MainContent() {
    return (
        <div className="main-content">
            <div className="card smoke">
                <h2>SMOKE</h2>
                <p>Status: <span className="status-failed">Failed Deployment</span></p>
                <p>Contact: Taj</p>
                <p>App Version: 2021.07.27</p>
                <p>Database Version: 7.2.0555</p>
                <p>Comments: Update in progress and run the pipeline and check</p>
            </div>
            <div className="card manual">
                <h2>MANUAL</h2>
                <p>Status: <span className="status-progress">Deployment In Progress</span></p>
                <p>Contact: Taj</p>
                <p>App Version: 2021.07.27</p>
                <p>Database Version: 7.2.0555</p>
                <p>Comments: Update in progress and run the pipeline and check</p>
            </div>
            {/* Add more cards as needed */}
        </div>
    );
}

export default MainContent;