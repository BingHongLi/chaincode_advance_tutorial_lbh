#!/bin/bash
# 僅供參考，不要直接掛載
docker exec -it chaincode bash
cd chaincode_advance_tutorial_lbh
go build
CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=chaincode_advance_tutorial_lbh:0 ./chaincode_advance_tutorial_lbh
