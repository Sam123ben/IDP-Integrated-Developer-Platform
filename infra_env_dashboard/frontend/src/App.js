// src/App.js

import React, { useState } from "react";
import Header from "./components/Header";
import Sidebar from "./components/Sidebar";
import MainContent from "./components/MainContent";
import Footer from "./components/Footer";
import TileContainer from "./components/TileContainer";
import "./styles/App.css";

function App() {
    const [selectedEnvironment, setSelectedEnvironment] = useState(null);

    // Sample environments based on selection (replace with actual fetching logic)
    const environments = selectedEnvironment
        ? [
              {
                  name: selectedEnvironment,
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
              }
          ]
        : []; // Empty array when no environment is selected

    const handleEnvironmentSelect = (environment) => {
        setSelectedEnvironment(environment); // Update selected environment
    };

    return (
        <div className="app">
            <Header />
            <div className="main-layout">
                <Sidebar onEnvironmentSelect={handleEnvironmentSelect} />
                <TileContainer environments={environments} /> {/* Pass environments prop */}
            </div>
            <Footer />
        </div>
    );
}

export default App;