# Import the base64 encoding function
locals {
  openvpn_custom_data = <<-CUSTOM_DATA
    #cloud-config
    runcmd:
      - echo "${var.vm_admin_password}" | passwd ${var.vm_admin_username} --stdin
      - systemctl enable openvpnas
      - systemctl start openvpnas
  CUSTOM_DATA
}

# Public IP for OpenVPN VM
resource "azurerm_public_ip" "openvpn_public_ip" {
  name                = "openvpn-public-ip"
  location            = var.location
  resource_group_name = var.resource_group_name
  allocation_method   = "Static"
  sku                 = "Standard"

  tags = var.tags
}

# Network Interface for OpenVPN VM
resource "azurerm_network_interface" "openvpn_nic" {
  name                = "openvpn-nic"
  location            = var.location
  resource_group_name = var.resource_group_name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.public_subnet_id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.openvpn_public_ip.id
  }
}

# VM for OpenVPN Access Server
resource "azurerm_linux_virtual_machine" "openvpn_vm" {
  name                            = "openvpn-vm"
  resource_group_name             = var.resource_group_name
  location                        = var.location
  size                            = "Standard_B1s"
  admin_username                  = var.vm_admin_username
  admin_password                  = var.vm_admin_password
  disable_password_authentication = false

  network_interface_ids = [
    azurerm_network_interface.openvpn_nic.id
  ]

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  # OpenVPN Access Server image details
  source_image_reference {
    publisher = "openvpn"
    offer     = "openvpnas"
    sku       = "openvpnas"
    version   = "latest"
  }

  tags = var.tags

  # Base64 encode the custom_data
  custom_data = base64encode(local.openvpn_custom_data)

  depends_on = [azurerm_network_interface.openvpn_nic]
}