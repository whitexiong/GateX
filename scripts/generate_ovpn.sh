#!/bin/bash

# 检查命令行参数数量
if [ "$#" -ne 1 ]; then
    echo "错误: 请提供客户端名称"
    exit 1
fi

CLIENT_NAME="$1"

# 设置 Easy-RSA 路径
EASY_RSA_PATH="/etc/openvpn/easy-rsa"

# 定义证书和密钥的位置
BASE_DIR="${EASY_RSA_PATH}/pki"
CA="${BASE_DIR}/ca.crt"
CLIENT_CERT="${BASE_DIR}/issued/${CLIENT_NAME}.crt"
CLIENT_KEY="${BASE_DIR}/private/${CLIENT_NAME}.key"
TA_KEY="${BASE_DIR}/ta.key"
OVPN_FILE="${CLIENT_NAME}.ovpn"

# 为客户端生成证书和密钥
cd $EASY_RSA_PATH || { echo "错误: 无法切换到 $EASY_RSA_PATH"; exit 1; }
./easyrsa build-client-full $CLIENT_NAME nopass

# 判断证书和密钥文件是否存在
if [[ ! -e $CLIENT_CERT || ! -e $CLIENT_KEY ]]; then
    echo "错误: 无法找到 ${CLIENT_NAME}.crt 或 ${CLIENT_NAME}.key。证书生成失败."
    exit 1
fi

# 创建 .ovpn 文件
echo "client
dev tun
proto udp
remote 120.46.80.249 1194 udp
resolv-retry infinite
nobind
persist-key
persist-tun

<ca>
$(cat $CA)
</ca>

<cert>
$(sed -ne '/-BEGIN CERTIFICATE-/,/-END CERTIFICATE-/p' $CLIENT_CERT)
</cert>

<key>
$(cat $CLIENT_KEY)
</key>

<tls-auth>
$(cat $TA_KEY)
</tls-auth>

key-direction 1
cipher AES-256-CBC
verb 3
mute 20

" > $OVPN_FILE

if [ -f "$OVPN_FILE" ]; then
    echo "${CLIENT_NAME}.ovpn 文件已创建成功!"
else
    echo "错误: 创建 ${CLIENT_NAME}.ovpn 文件失败."
    exit 1
fi
