# Terraform Generator - README

Welcome to the **Terraform Generator** project! This tool allows you to generate Terraform infrastructure configurations quickly and effectively. Below, you'll find all the instructions needed to set up and run the project.

## Table of Contents
- [Requirements](#requirements)
- [Installation](#installation)
- [Commands Overview](#commands-overview)
- [Usage Instructions](#usage-instructions)
- [Example Commands](#example-commands)
- [Troubleshooting](#troubleshooting)

## Requirements
- **Go**: Version 1.16 or later is required to run this application.
- **Terraform**: Terraform CLI should be installed for execution and validation purposes.
- **Environment**: Unix-based systems (Linux or macOS) are recommended for best compatibility.

## Installation
1. **Clone the repository**
   
   ```bash
   git clone https://github.com/your-username/terraform-generator.git
   cd terraform-generator/backend
   ```

2. **Install Dependencies** (optional if Go modules are not already set up):
   
   ```bash
   go mod tidy
   ```

3. **Build the Application** (optional):
   
   ```bash
   go build -o terraform-generator main.go
   ```

   Alternatively, you can directly run the application without building by using `go run`.

## Commands Overview
The Terraform Generator provides two main commands: `generate` and `terraform`.

- **Generate**: Generates Terraform files based on provided input parameters.
- **Terraform**: Executes different Terraform commands such as `init`, `validate`, `plan`, `apply`, `build`, and `destroy`.

## Usage Instructions
To run the Terraform Generator, you'll use the `go run` command with one of the available subcommands (`generate` or `terraform`). Each subcommand has its own set of required and optional flags.

### Running the Generate Command
The `generate` command generates Terraform configuration files for a specific company, product, provider, and infrastructure type.

#### Flags for `generate`:
- `--company`: Company name (required)
- `--product`: Product name (required)
- `--provider`: Provider name, e.g., `azurerm`, `aws` (required)
- `--infratype`: Infrastructure type, e.g., `prod`, `nonprod` (required)
- `--modules`: Comma-separated list of modules to include (required)
- `--customers`: Comma-separated list of customers (optional)

**Example**:
```bash
go run main.go generate --company acme --product dashboard --provider azurerm --infratype nonprod --modules resource_group,virtual_network
```

### Running the Terraform Command
The `terraform` command allows you to execute typical Terraform operations.

#### Flags for `terraform`:
- `--command`: The Terraform command to execute. Options are `init`, `validate`, `plan`, `apply`, `build`, `destroy`, or `print` (required)
- `--company`: Company name (required)
- `--product`: Product name (required)
- `--provider`: Provider name, e.g., `azurerm` (required)
- `--infratype`: Infrastructure type, e.g., `prod`, `nonprod` (required for some commands)

**Example**:
```bash
go run main.go terraform --command init --company acme --product dashboard --infratype nonprod --provider azurerm
```

The `terraform build` command runs `init`, `validate`, `plan`, and `apply` in sequence for a complete deployment.

## Example Commands
1. **Generate Terraform Files**:
   
   ```bash
   go run main.go generate --company acme --product dashboard --provider azurerm --infratype nonprod --modules resource_group,virtual_network
   ```

2. **Initialize Terraform**:
   
   ```bash
   go run main.go terraform --command init --company acme --product dashboard --infratype nonprod --provider azurerm
   ```

3. **Validate the Terraform Configuration**:
   
   ```bash
   go run main.go terraform --command validate --company acme --product dashboard --infratype nonprod --provider azurerm
   ```

4. **Plan Terraform Deployment**:
   
   ```bash
   go run main.go terraform --command plan --company acme --product dashboard --infratype nonprod --provider azurerm
   ```

5. **Apply Terraform Deployment**:
   
   ```bash
   go run main.go terraform --command apply --company acme --product dashboard --infratype nonprod --provider azurerm
   ```

6. **Destroy Infrastructure**:
   
   ```bash
   go run main.go terraform --command destroy --company acme --product dashboard --infratype nonprod --provider azurerm
   ```

## Troubleshooting
- **Error: Missing Flags**
  - Ensure that all required flags are provided for each subcommand.
- **Cannot find provider template**
  - Verify that the templates exist for the given provider in the `templates/` directory.
- **Permission Denied**
  - Make sure you have appropriate file system permissions to create directories and write files in the `output/` directory.
- **Terraform Errors**
  - After generating the files, run `terraform validate` to ensure that all Terraform configuration files are correctly structured.

## Contributions
Feel free to create a pull request or raise issues if you find any bugs or want to improve the functionality.

## License
This project is licensed under the MIT License.

## Author
Developed by Samyak Rout. Contributions are welcome!

---
If you have any questions or need more information, please feel free to reach out.