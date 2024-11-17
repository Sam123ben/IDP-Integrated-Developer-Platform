// src/services/api.tsx

import config from "../config";

interface CompanyDetails {
    name: string;
    // Add more properties if necessary
}

// Fetch company details
export const fetchCompanyDetails = async (): Promise<CompanyDetails> => {
    try {
        const response = await fetch(config.fetchCompanyDetails);
        if (!response.ok) throw new Error("Failed to fetch company details");
        const data: { company?: CompanyDetails } = await response.json();
        return data.company || { name: "" }; // Provide default value for company details
    } catch (error) {
        console.error("Error fetching company details:", error);
        throw error;
    }
};

// Fetch environment details for internal environments
export const fetchEnvironmentDetails = async (
    product: string,
    group: string
): Promise<any[]> => {
    try {
        const response = await fetch(
            `${config.fetchInternalEnvDetails}?product=${encodeURIComponent(product)}&group=${encodeURIComponent(group)}`
        );
        if (!response.ok) throw new Error("Failed to fetch environment details");
        const data: { environmentDetails?: any[] } = await response.json();
        return data.environmentDetails || [];
    } catch (error) {
        console.error("Error fetching environment details:", error);
        throw error;
    }
};

// Fetch customer environment details
export const fetchCustomerEnvDetails = async (
    customer: string,
    product: string
): Promise<any[]> => {
    try {
        const response = await fetch(
            `${config.fetchCustomerEnvDetails}?customer=${encodeURIComponent(customer)}&product=${encodeURIComponent(product)}`
        );
        if (!response.ok) throw new Error("Failed to fetch customer environment details");
        const data: { environmentDetails?: any[] } = await response.json();
        return data.environmentDetails || [];
    } catch (error) {
        console.error("Error fetching customer environment details:", error);
        throw error;
    }
};