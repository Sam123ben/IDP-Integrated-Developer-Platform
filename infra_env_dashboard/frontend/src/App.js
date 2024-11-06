// src/App.js

import React, { useState } from "react";
import Header from "./components/Header";
import Sidebar from "./components/Sidebar";
import MainContent from "./components/MainContent"; // Make sure this is correctly imported
import Footer from "./components/Footer";
import "./styles/App.css";

function App() {
    // Initialize the state for selected section
    const [selectedSection, setSelectedSection] = useState(null); // Initially null, meaning no section selected

    // Handler to update selectedSection when a section is selected in Sidebar
    const handleSectionSelect = (section) => {
        console.log("Selected section:", section); // Debug log
        setSelectedSection(section);
    };

    return (
        <div className="app">
            <Header />
            <div className="main-layout">
                <Sidebar onSectionSelect={handleSectionSelect} /> {/* Pass handler to Sidebar */}
                <MainContent selectedSection={selectedSection} /> {/* Pass selectedSection to MainContent */}
            </div>
            <Footer />
        </div>
    );
}

export default App;