#!/bin/bash

openssl req -newkey rsa:2048  -nodes -keyout nsx_certificate.key -x509 -days 365 -out nsx_certificate.crt -subj "/C=US/ST=California/L=PaloAlto/O=VMware/CN=certauth-test" -sha256

openssl pkcs12 -export -out nsx_certificate.pfx -inkey nsx_certificate.key -in nsx_certificate.crt

openssl pkcs12 -in nsx_certificate.pfx -out nsx_certificate.p12 -nodes