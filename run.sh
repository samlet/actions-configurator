#!/bin/bash
cd reverse_proxy
go build
cd ..
./reverse_proxy/reverse_proxy

