This guide provides steps to configure and deploy infrastructure for the DevOps Dashboard application on Azure using Terraform and GitHub Actions.

## Prerequisites

1. **Azure Subscription**: Ensure you have access to an Azure subscription.
2. **Service Principal**: You'll need a service principal to authenticate Terraform to Azure.
3. **GitHub Repository**: A GitHub repository where you'll set up secrets and configure GitHub Actions.

---

## Step 1: Create an Azure Service Principal

1. Open the Azure CLI or [Azure Cloud Shell](https://shell.azure.com/).
2. Run the following command to create a service principal:

   ```bash
   az ad sp create-for-rbac --name "devops-dashboard-terraform" --role="Contributor" --scopes="/subscriptions/<your-subscription-id>"
   ```
