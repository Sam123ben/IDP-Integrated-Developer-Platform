# Network Interface for OpenVPN VM
resource "azurerm_network_interface" "openvpn_nic" {
  name                = "openvpn-nic"
  location            = var.location
  resource_group_name = var.resource_group_name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.subnet_id
    private_ip_address_allocation = "Dynamic"
  }
}

# VM for OpenVPN Server
resource "azurerm_linux_virtual_machine" "openvpn_vm" {
  name                  = "openvpn-vm"
  resource_group_name   = var.resource_group_name
  location              = var.location
  size                  = "Standard_B1s"
  admin_username        = var.vm_admin_username
  admin_password        = var.vm_admin_password
  disable_password_authentication = false

  network_interface_ids = [
    azurerm_network_interface.openvpn_nic.id
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
}

# VM Extension to run OpenVPN installation script
resource "azurerm_virtual_machine_extension" "openvpn_extension" {
  name                 = "openvpn-extension"
  virtual_machine_id   = azurerm_linux_virtual_machine.openvpn_vm.id
  publisher            = "Microsoft.Azure.Extensions"
  type                 = "CustomScript"
  type_handler_version = "2.0"

  settings = <<SETTINGS
    {
      "commandToExecute": "bash chmod + /home/azureuser/install_openvpn.sh && bash /home/azureuser/install_openvpn.sh"
    }
  SETTINGS

  protected_settings = {
    "fileUris"          = ["${path.module}/scripts/install_openvpn.sh"]
  }

  depends_on = [
    azurerm_linux_virtual_machine.openvpn_vm
  ]
}