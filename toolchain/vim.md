# vim

WSL容器ubuntu
sudo apt install openssh-server
修改配置文件
sudo vim /etc/ssh/sshd_config
Port 22222
PasswordAuthentication yes
启动服务
sudo service ssh start



sudo ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key
sudo ssh-keygen -t ecdsa -f /etc/ssh/ssh_host_ecdsa_key
sudo ssh-keygen -t ed25519 -f /etc/ssh/ssh_host_ed25519_key


