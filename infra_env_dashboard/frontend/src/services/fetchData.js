// src/services/fetchData.js

import config from "../config";

export async function fetchData(endpointKey, params = {}) {
    // Get the base URL based on the endpoint key in config
    const baseUrl = config[endpointKey];
    if (!baseUrl) {
        throw new Error(`Invalid endpoint key: ${endpointKey}`);
    }

    // Construct the full URL with query parameters
    const url = new URL(baseUrl);
    Object.keys(params).forEach(key => {
        if (params[key]) {
            url.searchParams.append(key, params[key]);
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