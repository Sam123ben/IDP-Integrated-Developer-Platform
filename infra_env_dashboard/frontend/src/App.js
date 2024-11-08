// src/App.js

import React, { useState } from "react";
import Header from "./components/Header";
import Sidebar from "./components/Sidebar";
import MainContent from "./components/MainContent";
import Footer from "./components/Footer";
import "./styles/App.css";

function App() {
    const [selectedEnvironment, setSelectedEnvironment] = useState(null);
    const [envDetails, setEnvDetails] = useState([]);

    // Function to fetch environment details based on selected product and environment name
    const fetchEnvDetails = async (product, envName) => {
        try {
            const response = await fetch(`http://localhost:8082/api/internal-env-details?product=${encodeURIComponent(product)}&EnvName=${encodeURIComponent(envName)}`);
            const data = await response.json();
            setEnvDetails(data.environmentDetails || []); // Store environment details
        } catch (error) {
            console.error("Failed to fetch environment details:", error);
            setEnvDetails([]); // Clear details if fetch fails
        }
    };

    // Handle environment selection
    const handleEnvironmentSelect = (section, product, environment) => {
        console.log(`Selected environment: ${environment} for product: ${product}`);
        setSelectedEnvironment({ product, environment });
        fetchEnvDetails(product, environment); // Fetch details for the selected environment
    };

    return (
        <div className="app">
            <Header />
            <div className="main-layout">
                <Sidebar onEnvironmentSelect={handleEnvironmentSelect} />
                <MainContent envDetails={envDetails} /> {/* Pass envDetails to MainContent */}
            </div>
            <Footer />
        </div>
    );
}

export default App;