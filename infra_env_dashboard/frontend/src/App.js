// src/App.js

import React from "react";
import { BrowserRouter as Router, Routes, Route, NavLink } from "react-router-dom";
import Header from "./components/Header";
import "./App.css";

const Home = () => <div className="main-content">Home Content</div>;
const Environments = () => <div className="main-content">Environments/Infra Content</div>;
const BuildPipeline = () => <div className="main-content">Build Pipeline Content</div>;
const DeploymentPipeline = () => <div className="main-content">Deployment Pipeline Content</div>;

const App = () => {
    return (
        <Router>
            <Header />
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/environments" element={<Environments />} />
                <Route path="/build" element={<BuildPipeline />} />
                <Route path="/deploy" element={<DeploymentPipeline />} />
            </Routes>
        </Router>
    );
};

export default App;
