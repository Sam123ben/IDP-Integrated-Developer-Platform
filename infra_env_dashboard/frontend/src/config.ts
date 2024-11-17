// src/config.ts

// Define the types for environment variables
interface Config {
    fetchCompanyDetails: string;
    fetchInternalEnvDetails: string;
    fetchCustomerEnvDetails: string;
    // Add other endpoints if necessary as per the application's needs
}

const API_HOST: string = process.env.REACT_APP_API_HOST || 'http://localhost';
const API_PORT: string = process.env.REACT_APP_API_PORT || '8080';

const API_BASE_URL: string = `${API_HOST}:${API_PORT}/api`;

// Construct config object using the type interface defined
const config: Config = {
    fetchCompanyDetails: `${API_BASE_URL}/company`,
    fetchInternalEnvDetails: `${API_BASE_URL}/internal-env-details`,
    fetchCustomerEnvDetails: `${API_BASE_URL}/customer-env-details`,
    // Add other endpoints if necessary
};

export default config;