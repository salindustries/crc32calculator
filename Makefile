################################################################################
# Project: CRC32 Checksum Tools                                                #
# Filename: /Makefile                                                          #
# Created Date: Tuesday September 19th 2023 18:50:12 +0800                     #
# Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)                   #
# Company: BerryPay (M) Sdn. Bhd.                                              #
# --------------------------------------                                       #
# Last Modified: Tuesday September 19th 2023 19:08:03 +0800                    #
# Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)              #
# --------------------------------------                                       #
# Copyright (c) 2023 BerryPay (M) Sdn. Bhd.                                    #
################################################################################

# Define our default target
.DEFAULT_GOAL := build

build: build_crc32calculator

build_crc32calculator:
	go build -o bin/crc32calc main.go
	chmod +x bin/crc32calc