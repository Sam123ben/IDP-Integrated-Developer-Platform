# Define the VM in the same VNet as the PostgreSQL server
resource "azurerm_linux_virtual_machine" "sql_runner_vm" {
  name                = "sql-runner-vm"
  resource_group_name = var.resource_group_name
  location            = var.location
  size                = "Standard_B1s"
  admin_username      = "azureuser"  # Update with your preferred username
  admin_password      = var.vm_admin_password  # Variable for VM password

  network_interface_ids = [
    azurerm_network_interface.sql_runner_nic.id
  ]

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "18.04-LTS"
    version   = "latest"
  }

  tags = var.tags
}

# Network Interface for the VM within the same subnet
resource "azurerm_network_interface" "sql_runner_nic" {
  name                = "sql-runner-nic"
  location            = var.location
  resource_group_name = var.resource_group_name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.subnet_id  # Use the same subnet as the PostgreSQL server
    private_ip_address_allocation = "Dynamic"
  }
}

# Null resource to execute SQL script on the PostgreSQL server
resource "null_resource" "apply_sql_schema" {
  depends_on = [azurerm_postgresql_flexible_server_database.database, azurerm_linux_virtual_machine.sql_runner_vm]

  # Provisioner to run SQL script from the VM
  provisioner "remote-exec" {
    inline = [
      "PGPASSWORD='${var.admin_password}' psql -h ${azurerm_postgresql_flexible_server.db_server.fqdn} -U ${var.admin_username} -d ${azurerm_postgresql_flexible_server_database.database.name} -f /path/to/your/schema.sql"
    ]

    connection {
      type     = "ssh"
      host     = azurerm_linux_virtual_machine.sql_runner_vm.public_ip_address
      user     = "azureuser"
      password = var.vm_admin_password  # Use password instead of SSH key
    }
  }
}