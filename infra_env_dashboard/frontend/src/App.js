// src/App.js

import React, { useState } from "react";
import Header from "./components/Header";
import Sidebar from "./components/Sidebar";
import MainContent from "./components/MainContent";
import Footer from "./components/Footer";
import { fetchData } from "./services/fetchData";
import "./styles/App.css";

function App() {
    const [selectedEnvironment, setSelectedEnvironment] = useState(null);
    const [envDetails, setEnvDetails] = useState([]);

    const fetchEnvDetails = async (product, envName) => {
        try {
            // Pass parameters as an object to fetchData
            const data = await fetchData("fetchInternalEnvDetails", { product, group: envName });
            setEnvDetails(data.environmentDetails || []);
        } catch (error) {
            console.error("Failed to fetch environment details:", error);
            setEnvDetails([]);
        }
    };

    const handleEnvironmentSelect = (section, product, environment) => {
        if (product && environment) {
            console.log(`Selected environment: ${environment} for product: ${product}`);
            setSelectedEnvironment({ product, environment });
            fetchEnvDetails(product, environment);
        } else {
            console.warn("Product or environment is missing in the selection");
        }
    };

    return (
        <div className="app">
            <Header />
            <div className="main-layout">
                <Sidebar onEnvironmentSelect={handleEnvironmentSelect} />
                <MainContent envDetails={envDetails} />
            </div>
            <Footer />
        </div>
    );
}

export default App;