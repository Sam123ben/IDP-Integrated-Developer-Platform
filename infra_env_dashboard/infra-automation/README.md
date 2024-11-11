# Guide to Configuring and Deploying the DevOps Dashboard on Azure

This guide provides steps to configure and deploy infrastructure for the DevOps Dashboard application on Azure using Terraform and GitHub Actions.

## Prerequisites

1. **Azure Subscription**: Ensure you have access to an Azure subscription.
2. **Service Principal**: You'll need a service principal to authenticate Terraform to Azure.

---

## Step 1: Create an Azure Service Principal

1. Open the Azure CLI or [Azure Cloud Shell](https://shell.azure.com/).
2. Run the following command to create a service principal:

   ```bash
   az ad sp create-for-rbac --name "devops-dashboard-terraform" --role="Contributor" --scopes="/subscriptions/<your-subscription-id>"
   ```

   Replace `<YOUR_SUBSCRIPTION_ID>` with your actual Azure subscription ID.

3. The command will output the following details:

   ```json
   {
     "appId": "<YOUR_CLIENT_ID>",
     "displayName": "devops-dashboard-terraform",
     "password": "<YOUR_CLIENT_SECRET>",
     "tenant": "<YOUR_TENANT_ID>"
   }
   ```

   Make a note of:

   - `appId` (Client ID)
   - `password` (Client Secret)
   - `tenant` (Tenant ID)
   - Your Azure subscription ID

## Step 2: Set Up GitHub Secrets

In your GitHub repository, navigate to **Settings > Secrets > Actions** and add the following secrets:

- **ARM_CLIENT_ID**: The `appId` value from the service principal creation.
- **ARM_CLIENT_SECRET**: The `password` value from the service principal creation.
- **ARM_SUBSCRIPTION_ID**: Your Azure subscription ID.
- **ARM_TENANT_ID**: The `tenant` value from the service principal creation.

### AZURE_CREDENTIALS Secret

Add a secret called **AZURE_CREDENTIALS** with the following JSON object:

```json
{
  "clientId": "<YOUR_CLIENT_ID>",
  "clientSecret": "<YOUR_CLIENT_SECRET>",
  "subscriptionId": "<YOUR_SUBSCRIPTION_ID>",
  "tenantId": "<YOUR_TENANT_ID>"
}
```

Replace `<YOUR_CLIENT_ID>`, `<YOUR_CLIENT_SECRET>`, `<YOUR_SUBSCRIPTION_ID>`, and `<YOUR_TENANT_ID>` with your actual values.

## Step 3: Configure Terraform Backend (Optional)

In `infra-automation/backend.tf`, configure the backend storage for Terraform state files. For example, create a storage account and container in Azure to store the state file.

## Step 4: Run the GitHub Actions Workflow

Push your code to the `main` branch or manually trigger the GitHub Actions workflow from the **Actions** tab. The pipeline will execute the following steps:

1. **Initialize**: Initializes Terraform and configures the remote backend.
2. **Plan**: Plans the infrastructure deployment, showing the changes that will be made.
3. **Apply**: Deploys the infrastructure to Azure if the branch is `main` (for automatic approval).

The pipeline uses the credentials set in GitHub secrets to authenticate to Azure.

## Step 5: Verify the Deployment

Once the workflow completes, verify that resources were created in Azure by navigating to your Azure Portal. Check the resource group and resources, such as the virtual network, app services, PostgreSQL server, and storage accounts.

## Troubleshooting

- Ensure all required secrets are set in GitHub, as missing credentials can cause authentication failures.
- Verify that the service principal has the correct permissions in your Azure subscription.

## Additional Notes

- Adjust the Azure location in `variables.tf` or `environments/dev/main.tf` if you need to deploy resources in a different region.

### Infrastructure Components Include:

- **Virtual Network** for isolating network traffic.
- **App Services** for frontend and backend containers.
- **PostgreSQL** for backend database needs.
- **Network Security Groups** for securing resources.

With these steps, you should be able to set up the infrastructure for the DevOps Dashboard on Azure using Terraform and GitHub Actions. Happy deploying!