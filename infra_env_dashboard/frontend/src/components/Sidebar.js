// src/components/Sidebar.js

import React, { useState } from "react";
import "../styles/Sidebar.css";

function Sidebar({ onEnvironmentSelect, selectedSection, setSelectedSection }) {
    const [expandedItems, setExpandedItems] = useState({});
    const [selectedProduct, setSelectedProduct] = useState(null);
    const [selectedEnvOrCustomer, setSelectedEnvOrCustomer] = useState(null);

    const internalData = [
        {
            product: "Product 1",
            groups: ["DEV", "QA", "UAT", "PROD"],
        },
        {
            product: "Product 2",
            groups: ["DEV", "QA", "UAT", "PROD"],
        },
        // Add more internal products and groups as needed
    ];

    const customerData = [
        {
            customer: "Vendor A",
            products: ["Product 1", "Product 2"],
        },
        {
            customer: "Vendor B",
            products: ["Product 1", "Product 3"],
        },
        // Add more customers and their products as needed
    ];

    const toggleItem = (key) => {
        setExpandedItems((prev) => ({
            ...prev,
            [key]: !prev[key],
        }));
    };

    const handleSelection = (section, product, envOrCustomer) => {
        setSelectedProduct(product);
        setSelectedEnvOrCustomer(envOrCustomer);
        onEnvironmentSelect(section, product, envOrCustomer);
    };

    return (
        <div className="sidebar-container">
            <div className="sidebar-tabs">
                <span
                    className={`sidebar-tab ${selectedSection === "INTERNAL" ? "active" : ""}`}
                    onClick={() => {
                        setSelectedSection("INTERNAL");
                        setSelectedProduct(null);
                        setSelectedEnvOrCustomer(null);
                        setExpandedItems({});
                        onEnvironmentSelect("INTERNAL", null, null);
                    }}
                >
                    INTERNAL
                </span>
                <span
                    className={`sidebar-tab ${selectedSection === "CUSTOMER" ? "active" : ""}`}
                    onClick={() => {
                        setSelectedSection("CUSTOMER");
                        setSelectedProduct(null);
                        setSelectedEnvOrCustomer(null);
                        setExpandedItems({});
                        onEnvironmentSelect("CUSTOMER", null, null);
                    }}
                >
                    CUSTOMER
                </span>
            </div>
            <div className="sidebar-content">
                {selectedSection === "INTERNAL" &&
                    internalData.map((item) => (
                        <div key={item.product} className="product-section">
                            <div
                                className={`product-name ${selectedProduct === item.product ? "selected" : ""}`}
                                onClick={() => toggleItem(item.product)}
                            >
                                {item.product}
                                <span className="toggle-icon">
                                    {expandedItems[item.product] ? "▼" : "▶"}
                                </span>
                            </div>
                            {expandedItems[item.product] &&
                                item.groups.map((group) => (
                                    <div
                                        key={group}
                                        className={`environment-item ${
                                            selectedEnvOrCustomer === group && selectedProduct === item.product
                                                ? "selected"
                                                : ""
                                        }`}
                                        onClick={() => handleSelection("INTERNAL", item.product, group)}
                                    >
                                        {group}
                                    </div>
                                ))}
                        </div>
                    ))}

                {selectedSection === "CUSTOMER" &&
                    customerData.map((customerItem) => (
                        <div key={customerItem.customer} className="product-section">
                            <div
                                className={`product-name ${selectedEnvOrCustomer === customerItem.customer ? "selected" : ""}`}
                                onClick={() => toggleItem(customerItem.customer)}
                            >
                                {customerItem.customer}
                                <span className="toggle-icon">
                                    {expandedItems[customerItem.customer] ? "▼" : "▶"}
                                </span>
                            </div>
                            {expandedItems[customerItem.customer] &&
                                customerItem.products.map((product) => (
                                    <div
                                        key={product}
                                        className={`environment-item ${
                                            selectedProduct === product && selectedEnvOrCustomer === customerItem.customer
                                                ? "selected"
                                                : ""
                                        }`}
                                        onClick={() => handleSelection("CUSTOMER", product, customerItem.customer)}
                                    >
                                        {product}
                                    </div>
                                ))}
                        </div>
                    ))}
            </div>
        </div>
    );
}

export default Sidebar;