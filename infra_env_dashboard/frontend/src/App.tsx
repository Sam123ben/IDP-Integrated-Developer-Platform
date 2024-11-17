// src/App.tsx

import React, { useState, useEffect } from "react";
import Header from "./components/Header";
import Sidebar from "./components/Sidebar";
import MainContent from "./components/MainContent";
import Footer from "./components/Footer";
import { fetchCompanyDetails, fetchEnvironmentDetails, fetchCustomerEnvDetails } from "./services/api";
import "./styles/App.css";

// Define the type for companyDetails
interface CompanyDetails {
    name: string; // Add other fields as needed
}

function App() {
    const [selectedEnvironment, setSelectedEnvironment] = useState<{ section: string; product: string; environmentOrCustomer: string } | null>(null);
    const [envDetails, setEnvDetails] = useState<any[]>([]);
    const [selectedSection, setSelectedSection] = useState<string>("INTERNAL");
    const [companyDetails, setCompanyDetails] = useState<CompanyDetails | null>(null);

    // Fetch company details when the component mounts
    useEffect(() => {
        const getCompanyDetails = async () => {
            try {
                const details = await fetchCompanyDetails();
                setCompanyDetails(details);
            } catch (error) {
                console.error("Error fetching company details:", error);
            }
        };
        getCompanyDetails();
    }, []);

    // Function to fetch environment details with appropriate types
    const fetchEnvDetails = async (section: string, product: string, groupOrCustomer: string) => {
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

    // Function to handle the environment selection
    const handleEnvironmentSelect = (
        section: string,
        product: string | null,
        environmentOrCustomer: string | null
    ) => {
        if (product && environmentOrCustomer) {
            console.log(`Selected ${section} environment: ${environmentOrCustomer} for product: ${product}`);
            setSelectedEnvironment({ section, product, environmentOrCustomer });
            fetchEnvDetails(section, product, environmentOrCustomer);
        } else {
            console.warn("Product or environment/customer is missing in the selection");
        }
    };

    // Render only after company details are fetched
    if (!companyDetails) {
        return <div>Loading...</div>;
    }

    return (
        <div className="app">
            {/* Pass companyDetails to the Header component */}
            <Header companyDetails={companyDetails} />
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