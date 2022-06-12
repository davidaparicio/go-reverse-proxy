#!/usr/bin/env bash

# https://linuxize.com/post/creating-a-self-signed-ssl-certificate/
openssl req -newkey rsa:4096 \
            -x509 \
            -sha256 \
            -days 3650 \
            -nodes \
            -out cert.pem \
            -keyout key.pem \
            -subj "/C=FR/ST=France/L=Lyon/O=My Organization/OU=IT Department/CN=localhost"
            #-subj "/C=SI/ST=Ljubljana/L=Ljubljana/O=Security/OU=IT Department/CN=www.example.com"