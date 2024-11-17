// src/config.js

const API_HOST = process.env.REACT_APP_API_HOST || 'http://localhost';
const API_PORT = process.env.REACT_APP_API_PORT || '8080';

const API_BASE_URL = `${API_HOST}:${API_PORT}/api`;

const config = {
    fetchCompanyDetails: `${API_BASE_URL}/company`,
    fetchInternalEnvDetails: `${API_BASE_URL}/internal-env-details`,
    fetchCustomerEnvDetails: `${API_BASE_URL}/customer-env-details`,
    // Add other endpoints if necessary
};

export default config;