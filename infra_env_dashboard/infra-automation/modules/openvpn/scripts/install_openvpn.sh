#!/bin/bash
# Update and install OpenVPN
sudo apt-get update -y
sudo apt-get install -y openvpn easy-rsa

# Configure Easy-RSA for certificate generation
make-cadir ~/openvpn-ca
cd ~/openvpn-ca
source vars
./clean-all
./build-ca --batch
./build-key-server --batch server
./build-dh
openvpn --genkey --secret keys/ta.key

# Copy server configuration files
sudo cp /usr/share/doc/openvpn/examples/sample-config-files/server.conf.gz /etc/openvpn/
sudo gunzip /etc/openvpn/server.conf.gz

# Edit OpenVPN configuration
sudo sed -i 's/;tls-auth ta.key 0/tls-auth ta.key 0/g' /etc/openvpn/server.conf
sudo sed -i 's/;user nobody/user nobody/g' /etc/openvpn/server.conf
sudo sed -i 's/;group nogroup/group nogroup/g' /etc/openvpn/server.conf

# Start and enable OpenVPN service
sudo systemctl start openvpn@server
sudo systemctl enable openvpn@server