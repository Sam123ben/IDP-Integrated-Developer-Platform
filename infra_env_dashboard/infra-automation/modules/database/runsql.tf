# Step 1: Create Storage Account
resource "azurerm_storage_account" "sql_script_storage" {
  name                     = "samstgaccount01" # The Name of the Storage Account must be unique
  resource_group_name      = var.resource_group_name
  location                 = var.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
  tags                     = var.tags
}

# Step 2: Create Blob Container
resource "azurerm_storage_container" "sql_scripts_container" {
  name                  = "sqlscripts"
  storage_account_name  = azurerm_storage_account.sql_script_storage.name
  container_access_type = "private"
}

# Step 3: Upload SQL Script to Blob
resource "azurerm_storage_blob" "sql_script_blob" {
  name                   = "000_create_database_schema.sql"
  storage_account_name   = azurerm_storage_account.sql_script_storage.name
  storage_container_name = azurerm_storage_container.sql_scripts_container.name
  type                   = "Block"
  source                 = "${path.module}/scripts/000_create_database_schema.sql"  # Path to the SQL script

  depends_on = [azurerm_storage_container.sql_scripts_container]
}

# Step 4: Generate SAS Token for VM Access to SQL Script Blob
data "azurerm_storage_account_sas" "sql_script_sas" {
  connection_string = azurerm_storage_account.sql_script_storage.primary_connection_string
  https_only        = true
  start             = "2023-01-01T00:00Z"
  expiry            = "2030-01-01T00:00Z"
  resource_types {
    service   = true
    container = false
    object    = true
  }

  services {
    blob  = true
    queue = false
    table = false
    file  = false
  }

  permissions {
    read    = true
    write   = false
    delete  = false
    list    = false
    add     = false
    create  = false
    update  = false
    process = false
    tag     = false
    filter  = false
  }
  depends_on = [azurerm_storage_blob.sql_script_blob]
}

# Network Interface for the VM
resource "azurerm_network_interface" "sql_runner_nic" {
  name                = "sql-runner-nic"
  location            = var.location
  resource_group_name = var.resource_group_name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.app_subnet_id
    private_ip_address_allocation = "Dynamic"
  }
}

# Define the VM
resource "azurerm_linux_virtual_machine" "sql_runner_vm" {
  name                        = "sql-runner-vm"
  resource_group_name         = var.resource_group_name
  location                    = var.location
  size                        = "Standard_B1s"
  admin_username              = "azureuser"
  admin_password              = var.vm_admin_password
  disable_password_authentication = false

  network_interface_ids = [
    azurerm_network_interface.sql_runner_nic.id
  ]

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-jammy"
    sku       = "22_04-lts"
    version   = "latest"
  }
  tags = var.tags

  depends_on = [ azurerm_postgresql_flexible_server_database.database , azurerm_network_interface.sql_runner_nic ]
}

# VM Extension to Install PostgreSQL Client and Run SQL Script
resource "azurerm_virtual_machine_extension" "sql_runner_extension" {
  name                 = "sql-runner-extension"
  virtual_machine_id   = azurerm_linux_virtual_machine.sql_runner_vm.id
  publisher            = "Microsoft.Azure.Extensions"
  type                 = "CustomScript"
  type_handler_version = "2.0"

  settings = <<SETTINGS
    {
      "commandToExecute": "sudo apt-get update -y && sudo apt-get install -y postgresql-client && curl -L '${azurerm_storage_blob.sql_script_blob.url}${data.azurerm_storage_account_sas.sql_script_sas.sas}' -o /home/azureuser/000_create_database_schema.sql && PGPASSWORD='${var.admin_password}' psql -h ${azurerm_postgresql_flexible_server.db_server.fqdn} -U ${var.admin_username} -d ${azurerm_postgresql_flexible_server_database.database.name} -f /home/azureuser/000_create_database_schema.sql"
    }
  SETTINGS

  depends_on = [
    azurerm_linux_virtual_machine.sql_runner_vm,
    azurerm_postgresql_flexible_server.db_server,
    azurerm_storage_blob.sql_script_blob
  ]
}