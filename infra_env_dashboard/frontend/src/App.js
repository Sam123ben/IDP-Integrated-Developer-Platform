// src/App.js

import React from "react";
import Header from "./components/Header";
import Sidebar from "./components/Sidebar";
import TileContainer from "./components/TileContainer";
import MainContent from "./components/MainContent"; // New component for cards
import "./styles/App.css";

function App() {
    return (
        <div className="app">
            <Header />
            <div className="main-layout">
                <Sidebar />
                <TileContainer />
            </div>
        </div>
    );
}

export default App;
