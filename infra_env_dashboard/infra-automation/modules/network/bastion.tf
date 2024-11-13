# Public IP for Bastion
resource "azurerm_public_ip" "bastion_public_ip" {
  name                = "${var.bastion_name}-pip"
  location            = var.location
  resource_group_name = var.resource_group_name
  allocation_method   = "Static"
  sku                 = "Standard"  # Required for Bastion

  tags = var.tags
}

# Azure Bastion Host
resource "azurerm_bastion_host" "bastion" {
  name                = var.bastion_name
  location            = var.location
  resource_group_name = var.resource_group_name

  ip_configuration {
    name                 = "configuration"
    subnet_id            = azurerm_subnet.bastion_subnet.id
    public_ip_address_id = azurerm_public_ip.bastion_public_ip.id
  }

  sku      = var.bastion_sku
  tags     = var.tags
  depends_on = [azurerm_public_ip.bastion_public_ip]
}