docker exec -it cli bash
peer chaincode invoke -n chaincode_advance_tutorial_lbh -v 0  -c '{"Args":["putCompose","lbh","123","456","789"]}' -C myc
peer chaincode invoke -n chaincode_advance_tutorial_lbh -v 0  -c '{"Args":["getCompose","lbh","123"]}' -C myc
# 預期結果為 "123||456||789"
peer chaincode invoke -n chaincode_advance_tutorial_lbh -v 0  -c '{"Args":["putCompose","lbh","666","777","888"]}' -C myc
peer chaincode invoke -n chaincode_advance_tutorial_lbh -v 0  -c '{"Args":["getCompose","lbh","123"]}' -C myc
# 預期結果為 "123||456||789"
peer chaincode invoke -n chaincode_advance_tutorial_lbh -v 0  -c '{"Args":["getCompose","lbh","666"]}' -C myc
# 預期結果為 "666||777||888"
peer chaincode invoke -n chaincode_advance_tutorial_lbh -v 0  -c '{"Args":["putCompose","lbh","666","999","000"]}' -C myc
peer chaincode invoke -n chaincode_advance_tutorial_lbh -v 0  -c '{"Args":["getCompose","lbh","666"]}' -C myc
# 預期結果為 "666666||777999||888000"