#!/bin/bash

monogen -L terraform -c monolithe.ini -f ~/git/vsd-api-specifications
rm -r nuagenetworks
mv terraform/nuagenetworks ./
rm -r terraform
go build
cp terraform-provider-nuagenetworks ~/.terraform.d/plugins/
