// src/components/Sidebar.js

import React, { useState } from "react";
import "../styles/Sidebar.css";

function Sidebar() {
    const [activeCategory, setActiveCategory] = useState("internal");
    const [expandedItems, setExpandedItems] = useState({});
    const [expandedProducts, setExpandedProducts] = useState({});

    const toggleCategory = (category) => {
        setActiveCategory(category);
    };

    const toggleExpand = (item) => {
        setExpandedItems((prev) => ({
            ...prev,
            [item]: !prev[item],
        }));
    };

    const toggleProductExpand = (product) => {
        setExpandedProducts((prev) => ({
            ...prev,
            [product]: !prev[product],
        }));
    };

    return (
        <div className="sidebar">
            <div className="sidebar-categories">
                <h3
                    onClick={() => toggleCategory("internal")}
                    className={`category ${activeCategory === "internal" ? "active" : ""}`}
                >
                    INTERNAL
                </h3>
                <h3
                    onClick={() => toggleCategory("CUSTOMER")}
                    className={`category ${activeCategory === "CUSTOMER" ? "active" : ""}`}
                >
                    CUSTOMER
                </h3>
            </div>

            <div className="sidebar-content">
                {activeCategory === "internal" && (
                    <>
                        <div className="collapsible-item">
                            <div
                                className="item-header"
                                onClick={() => toggleProductExpand("Product 1")}
                            >
                                Product 1
                            </div>
                            {expandedProducts["Product 1"] && (
                                <ul className="environment-list">
                                    <li>DEV</li>
                                    <li>QA</li>
                                    <li>CONSULT</li>
                                    <li>PRESALES</li>
                                </ul>
                            )}
                        </div>
                        <div className="collapsible-item">
                            <div
                                className="item-header"
                                onClick={() => toggleProductExpand("Product 2")}
                            >
                                Product 2
                            </div>
                            {expandedProducts["Product 2"] && (
                                <ul className="environment-list">
                                    <li>DEV</li>
                                    <li>QA</li>
                                    <li>STAGING</li>
                                </ul>
                            )}
                        </div>
                    </>
                )}

                {activeCategory === "CUSTOMER" && (
                    <>
                        <div className="collapsible-item">
                            <div
                                className="item-header"
                                onClick={() => toggleProductExpand("vendorA")}
                            >
                                Vendor A
                            </div>
                            {expandedProducts["vendorA"] && (
                                <ul className="product-list">
                                    <li>Product 1</li>
                                    <li>Product 2</li>
                                </ul>
                            )}
                        </div>
                        <div className="collapsible-item">
                            <div
                                className="item-header"
                                onClick={() => toggleProductExpand("vendorB")}
                            >
                                Vendor B
                            </div>
                            {expandedProducts["vendorB"] && (
                                <ul className="product-list">
                                    <li>Product 1</li>
                                    <li>Product 2</li>
                                </ul>
                            )}
                        </div>
                    </>
                )}
            </div>
        </div>
    );
}

export default Sidebar;