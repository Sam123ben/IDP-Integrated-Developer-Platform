// src/App.js

import React, { useState } from "react";
import Header from "./components/Header";
import Sidebar from "./components/Sidebar";
import TileContainer from "./components/TileContainer";
import MainContent from "./components/MainContent";
import "./styles/App.css";

function App() {
    // State to store the selected sidebar item data
    const [selectedSection, setSelectedSection] = useState(null);

    // Function to update selected section from sidebar
    const handleSectionSelect = (sectionData) => {
        setSelectedSection(sectionData);
    };

    return (
        <div className="app">
            <Header />
            <div className="main-layout">
                <Sidebar onSectionSelect={handleSectionSelect} />
                <MainContent selectedSection={selectedSection} />
            </div>
        </div>
    );
}

export default App;