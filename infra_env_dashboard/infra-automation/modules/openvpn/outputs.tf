output "openvpn_vm_id" {
  description = "The ID of the OpenVPN VM."
  value       = azurerm_linux_virtual_machine.openvpn_vm.id
}

output "openvpn_vm_ip" {
  description = "The private IP of the OpenVPN VM."
  value       = azurerm_network_interface.openvpn_nic.private_ip_address
}