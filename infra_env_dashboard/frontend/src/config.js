// config.js

// Set default values for API_HOST and API_PORT
const API_HOST = process.env.REACT_APP_API_HOST || "http://localhost";
const API_PORT = process.env.REACT_APP_API_PORT || "8080";

// Construct the base API URL
const API_BASE_URL = `${API_HOST}:${API_PORT}/api`;

// Export config with dynamically constructed URLs
const config = {
    fetchCompanyDetails: `${API_BASE_URL}/company`,
    fetchInfraTypes: `${API_BASE_URL}/infra-types`,
    fetchInternalEnvDetails: `${API_BASE_URL}/internal-env-details`,
};

export default config;
