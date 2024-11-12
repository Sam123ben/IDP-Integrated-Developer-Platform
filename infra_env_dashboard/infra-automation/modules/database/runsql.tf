# Network Interface for the VM within the same subnet
resource "azurerm_network_interface" "sql_runner_nic" {
  name                = "sql-runner-nic"
  location            = var.location
  resource_group_name = var.resource_group_name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.app_subnet_id
    private_ip_address_allocation = "Dynamic"
  }

  depends_on = [azurerm_postgresql_flexible_server_database.database]
}

# Define the VM in the same VNet as the PostgreSQL server
resource "azurerm_linux_virtual_machine" "sql_runner_vm" {
  name                = "sql-runner-vm"
  resource_group_name = var.resource_group_name
  location            = var.location
  size                = "Standard_B1s"
  admin_username      = "azureuser"
  admin_password      = var.vm_admin_password

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

  depends_on = [azurerm_network_interface.sql_runner_nic]
  tags       = var.tags
}

# VM Extension to install PostgreSQL client and execute SQL script
resource "azurerm_virtual_machine_extension" "sql_runner_extension" {
  name                 = "sql-runner-extension"
  virtual_machine_id   = azurerm_linux_virtual_machine.sql_runner_vm.id
  publisher            = "Microsoft.Azure.Extensions"
  type                 = "CustomScript"
  type_handler_version = "2.0"

  # Command to install PostgreSQL client and execute the SQL script
  settings = <<SETTINGS
    {
      "commandToExecute": "sudo apt-get update -y && sudo apt-get install -y postgresql-client && PGPASSWORD='${var.admin_password}' psql -h ${azurerm_postgresql_flexible_server.db_server.fqdn} -U ${var.admin_username} -d ${azurerm_postgresql_flexible_server_database.database.name} -f /home/azureuser/000_create_database_schema.sql"
    }
  SETTINGS

  # Upload the SQL script file to the VM
  protected_settings = <<PROTECTED_SETTINGS
    {
      "fileUris": ["${path.module}/../../infra_env_dashboard/database/000_create_database_schema.sql"],
      "commandToExecute": "cp 000_create_database_schema.sql /home/azureuser/000_create_database_schema.sql"
    }
  PROTECTED_SETTINGS

  depends_on = [
    azurerm_linux_virtual_machine.sql_runner_vm,
    azurerm_postgresql_flexible_server.db_server
  ]
}