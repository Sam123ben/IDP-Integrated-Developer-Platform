output "vnet_id" {
  value = azurerm_virtual_network.vnet.id
}

output "public_subnet_id" {
  value = azurerm_subnet.public_subnet.id
}

output "app_subnet_id" {
  value = azurerm_subnet.app_subnet.id
}

output "db_subnet_id" {
  value = azurerm_subnet.db_subnet.id
}

output "app_nsg_id" {
  value = azurerm_network_security_group.app_nsg.id
}

output "db_nsg_id" {
  value = azurerm_network_security_group.db_nsg.id
}

output "bastion_host_id" {
  description = "The ID of the Bastion host"
  value       = azurerm_bastion_host.bastion.id
}

output "bastion_public_ip" {
  description = "Public IP address of the Bastion host"
  value       = azurerm_public_ip.bastion_public_ip.ip_address
}