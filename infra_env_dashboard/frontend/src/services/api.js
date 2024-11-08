import config from "../config";

// Fetch company details
export const fetchCompanyDetails = async () => {
  try {
    const response = await fetch(`${config.fetchCompanyDetails}/company`);
    if (!response.ok) throw new Error("Failed to fetch company details");
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error fetching company details:", error);
    throw error;
  }
};

// Fetch infrastructure types
export const fetchInfraTypes = async () => {
  try {
    const response = await fetch(`${config.fetchInfraTypes}/infra-types`);
    if (!response.ok) throw new Error("Failed to fetch infra types");
    const data = await response.json();
    return data.infraTypes || [];
  } catch (error) {
    console.error("Error fetching infra types:", error);
    throw error;
  }
};

// Fetch environment details for a specific product and environment
export const fetchEnvironmentDetails = async (product, envName) => {
  try {
    const response = await fetch(
      `${config.fetchInternalEnvDetails}/internal-env-details?product=${encodeURIComponent(
        product
      )}&EnvName=${encodeURIComponent(envName)}`
    );
    if (!response.ok) throw new Error("Failed to fetch environment details");
    const data = await response.json();
    return data.environmentDetails || [];
  } catch (error) {
    console.error("Error fetching environment details:", error);
    throw error;
  }
};