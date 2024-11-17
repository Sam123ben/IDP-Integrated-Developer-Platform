// src/App.js

import React, { useState } from "react";
import Header from "./components/Header";
import Sidebar from "./components/Sidebar";
import MainContent from "./components/MainContent";
import Footer from "./components/Footer";
import { fetchEnvironmentDetails, fetchCustomerEnvDetails } from "./services/api";
import "./styles/App.css";

function App() {
    const [selectedEnvironment, setSelectedEnvironment] = useState(null);
    const [envDetails, setEnvDetails] = useState([]);
    const [selectedSection, setSelectedSection] = useState("INTERNAL");

    const fetchEnvDetails = async (section, product, groupOrCustomer) => {
        try {
            let data = [];
            if (section === "INTERNAL") {
                data = await fetchEnvironmentDetails(product, groupOrCustomer);
            } else if (section === "CUSTOMER") {
                data = await fetchCustomerEnvDetails(groupOrCustomer, product);
            }
            setEnvDetails(data);
        } catch (error) {
            console.error("Failed to fetch environment details:", error);
            setEnvDetails([]);
        }
    };

    const handleEnvironmentSelect = (section, product, environmentOrCustomer) => {
        if (product && environmentOrCustomer) {
            console.log(`Selected ${section} environment: ${environmentOrCustomer} for product: ${product}`);
            setSelectedEnvironment({ section, product, environmentOrCustomer });
            fetchEnvDetails(section, product, environmentOrCustomer);
        } else {
            console.warn("Product or environment/customer is missing in the selection");
        }
    };

    return (
        <div className="app">
            <Header />
            <div className="main-layout">
                <Sidebar
                    onEnvironmentSelect={handleEnvironmentSelect}
                    selectedSection={selectedSection}
                    setSelectedSection={setSelectedSection}
                />
                <MainContent envDetails={envDetails} />
            </div>
            <Footer />
        </div>
    );
}

export default App;