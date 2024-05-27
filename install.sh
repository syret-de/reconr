# Install Golang
wget https://go.dev/dl/go1.22.3.linux-amd64.tar.gz
tar -xf go1.22.3.linux-amd64.tar.gz
rm go1.22.3.linux-amd64.tar.gz
mv go /usr/local/
echo 'export PATH=$PATH:/usr/local/go/bin' >> .bashrc

# Install docker
## Add Docker's official GPG key:
apt-get update
apt-get install ca-certificates curl
install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
chmod a+r /etc/apt/keyrings/docker.asc

## Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  tee /etc/apt/sources.list.d/docker.list > /dev/null
apt-get update

apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

# Install jq
apt install jq -y

# Build
go build -o reconr cmd/reconr/main.go