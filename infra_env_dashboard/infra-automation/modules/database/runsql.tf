# Network Interface for the VM within the same subnet
resource "azurerm_network_interface" "sql_runner_nic" {
  name                = "sql-runner-nic"
  location            = var.location
  resource_group_name = var.resource_group_name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.app_subnet_id  # Use the same subnet as the PostgreSQL server
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
  admin_username      = "azureuser"  # Update with your preferred username
  admin_password      = var.vm_admin_password  # Variable for VM password

  disable_password_authentication = false  # Enable password authentication instead of SSH keys

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

# Null resource to install PostgreSQL client and execute SQL script
resource "null_resource" "apply_sql_schema" {
  depends_on = [
    azurerm_postgresql_flexible_server_database.database,
    azurerm_linux_virtual_machine.sql_runner_vm
  ]

  # Provisioner to install PostgreSQL client and run SQL script from the VM
  provisioner "remote-exec" {
    inline = [
      # Install PostgreSQL client
      "sudo apt-get update -y",
      "sudo apt-get install -y postgresql-client",

      # Run SQL script on PostgreSQL server
      "PGPASSWORD='${var.admin_password}' psql -h ${azurerm_postgresql_flexible_server.db_server.fqdn} -U ${var.admin_username} -d ${azurerm_postgresql_flexible_server_database.database.name} -f ~/infra_env_dashboard/database/000_create_database_schema.sql"
    ]

    connection {
      type     = "ssh"
      host     = azurerm_linux_virtual_machine.sql_runner_vm.public_ip_address
      user     = "azureuser"
      password = var.vm_admin_password  # Use password instead of SSH key
    }
  }

  # Upload the schema SQL file to the VM
  provisioner "file" {
    source      = "${path.module}/../../infra_env_dashboard/database/000_create_database_schema.sql"
    destination = "~/infra_env_dashboard/database/000_create_database_schema.sql"

    connection {
      type     = "ssh"
      host     = azurerm_linux_virtual_machine.sql_runner_vm.public_ip_address
      user     = "azureuser"
      password = var.vm_admin_password
    }
  }
}
