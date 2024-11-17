// src/services/fetchData.tsx

import config from "../config";

// Define allowed endpoint keys based on config
type EndpointKey = keyof typeof config;

export async function fetchData(endpointKey: EndpointKey, params: Record<string, string | number | null> = {}): Promise<any> {
    // Get the base URL based on the endpoint key in config
    const baseUrl = config[endpointKey];
    if (!baseUrl) {
        throw new Error(`Invalid endpoint key: ${endpointKey}`);
    }

    // Construct the full URL with query parameters
    const url = new URL(baseUrl);
    Object.keys(params).forEach((key) => {
        const value = params[key];
        if (value) {
            url.searchParams.append(key, value.toString());
        }
    });

    try {
        // Perform the fetch call with the constructed URL
        const response = await fetch(url.toString());
        if (!response.ok) {
            throw new Error(`Error: ${response.statusText}`);
        }

        // Return the parsed JSON data
        return await response.json();
    } catch (error) {
        console.error(`Failed to fetch data from ${url}:`, error);
        throw error;
    }
}