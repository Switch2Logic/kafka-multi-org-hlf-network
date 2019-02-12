#Kafka Multiple Orginaztion HLF Network

# Open Hospital1 client
`$ docker exec -it cli_hospital1 bash`

# Create channel with mutual TLS enabled
`$ peer channel create -o orderer0.hospital1.switch2logic.co.za:7050 -c comunitychannel -f ./channel-artifacts/comunity_channel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/ordererOrganizations/switch2logic.co.za/orderers/orderer0.hospital1.switch2logic.co.za/msp/tlscacerts/tlsca.switch2logic.co.za-cert.pem --clientauth --keyfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital1.switch2logic.co.za/users/User1@hospital1.switch2logic.co.za/tls/client.key --certfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital1.switch2logic.co.za/users/User1@hospital1.switch2logic.co.za/tls/client.crt`

# Fetch channel block 0 with mutual TLS enabled
`$ peer channel fetch 0 comunitychannel.block --channelID comunitychannel --orderer orderer0.hospital1.switch2logic.co.za:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/ordererOrganizations/switch2logic.co.za/orderers/orderer0.hospital1.switch2logic.co.za/msp/tlscacerts/tlsca.switch2logic.co.za-cert.pem --clientauth --keyfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital1.switch2logic.co.za/users/User1@hospital1.switch2logic.co.za/tls/client.key --certfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital1.switch2logic.co.za/users/User1@hospital1.switch2logic.co.za/tls/client.crt`<br />

`$ peer channel fetch 0 comunitychannel.block --channelID comunitychannel --orderer orderer0.hospital2.switch2logic.co.za:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/ordererOrganizations/switch2logic.co.za/orderers/orderer0.hospital2.switch2logic.co.za/msp/tlscacerts/tlsca.switch2logic.co.za-cert.pem --clientauth --keyfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital2.switch2logic.co.za/users/User1@hospital2.switch2logic.co.za/tls/client.key --certfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital2.switch2logic.co.za/users/User1@hospital2.switch2logic.co.za/tls/client.crt`<br />

`$ peer channel fetch 0 comunitychannel.block --channelID comunitychannel --orderer orderer0.hospital3.switch2logic.co.za:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/ordererOrganizations/switch2logic.co.za/orderers/orderer0.hospital3.switch2logic.co.za/msp/tlscacerts/tlsca.switch2logic.co.za-cert.pem --clientauth --keyfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital3.switch2logic.co.za/users/User1@hospital3.switch2logic.co.za/tls/client.key --certfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital3.switch2logic.co.za/users/User1@hospital3.switch2logic.co.za/tls/client.crt`

# Join channel
`peer channel join -b comunitychannel.block`

# Update channel achorpeers 
`$ peer channel update -o orderer0.hospital1.switch2logic.co.za:7050 -c comunitychannel -f ./channel-artifacts/hospital1MSP_anchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/ordererOrganizations/switch2logic.co.za/orderers/orderer0.hospital1.switch2logic.co.za/msp/tlscacerts/tlsca.switch2logic.co.za-cert.pem --clientauth --keyfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital1.switch2logic.co.za/users/User1@hospital1.switch2logic.co.za/tls/client.key --certfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital1.switch2logic.co.za/users/User1@hospital1.switch2logic.co.za/tls/client.crt`<br />

`$ peer channel update -o orderer0.hospital2.switch2logic.co.za:7050 -c comunitychannel -f ./channel-artifacts/hospital1MSP_anchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/ordererOrganizations/switch2logic.co.za/orderers/orderer0.hospital2.switch2logic.co.za/msp/tlscacerts/tlsca.switch2logic.co.za-cert.pem --clientauth --keyfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital2.switch2logic.co.za/users/User1@hospital2.switch2logic.co.za/tls/client.key --certfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital2.switch2logic.co.za/users/User1@hospital2.switch2logic.co.za/tls/client.crt`

# Install chanincode with private data
`$ peer chaincode install -n patient_private -v 1.0 -p github.com/hyperledger/fabric/peer/chaincode/`


# Instaniante chaincode with private data collection file 
`$ peer chaincode instantiate -o orderer0.hospital1.switch2logic.co.za:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/ordererOrganizations/switch2logic.co.za/orderers/orderer0.hospital1.switch2logic.co.za/msp/tlscacerts/tlsca.switch2logic.co.za-cert.pem --clientauth --keyfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital1.switch2logic.co.za/users/User1@hospital1.switch2logic.co.za/tls/client.key --certfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital1.switch2logic.co.za/users/User1@hospital1.switch2logic.co.za/tls/client.crt -C comunitychannel -n patient_private -v 1.0 -c '{"Args":["init"]}'  --collections-config $GOPATH/src/github.com/hyperledger/fabric/peer/chaincode/collections_config.json`

# Upgrade collection chaincode with private data collection file 
`$ peer chaincode upgrade -o orderer0.hospital1.switch2logic.co.za:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/ordererOrganizations/switch2logic.co.za/orderers/orderer0.hospital1.switch2logic.co.za/msp/tlscacerts/tlsca.switch2logic.co.za-cert.pem --clientauth --keyfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital1.switch2logic.co.za/users/User1@hospital1.switch2logic.co.za/tls/client.key --certfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital1.switch2logic.co.za/users/User1@hospital1.switch2logic.co.za/tls/client.crt -C comunitychannel -n patient_private -p /opt/gopath/src/github.com/chaincode/ -v1.0 -c '{"Args":["init"]}' --collections-config $GOPATH/src/github.com/hyperledger/fabric/peer/chaincode/collections_config.json`

# Invoke hyperledger-fabric chaincode
`$ peer chaincode invoke -o orderer0.hospital1.switch2logic.co.za:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/ordererOrganizations/switch2logic.co.za/orderers/orderer0.hospital1.switch2logic.co.za/msp/tlscacerts/tlsca.switch2logic.co.za-cert.pem --clientauth --keyfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital1.switch2logic.co.za/users/User1@hospital1.switch2logic.co.za/tls/client.key --certfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital1.switch2logic.co.za/users/User1@hospital1.switch2logic.co.za/tls/client.crt -C comunitychannel -n patient_private --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital2.switch2logic.co.za/peers/peer0.hospital1.switch2logic.co.za/tls/ca.crt -c '{"Args":["initLedger"]}'`

# Query hyperledger-fabric chaincode
`$ peer chaincode query --clientauth --keyfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital1.switch2logic.co.za/users/User1@hospital1.switch2logic.co.za/tls/client.key --certfile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto-config/peerOrganizations/hospital1.switch2logic.co.za/users/User1@hospital1.switch2logic.co.za/tls/client.crt  -C comunitychannel -n patient_private -c '{"Args":["getPatient","testkey"]}'` <br />

# Query/Inovke with Fabric-Node-SDK
`Find my fabri-sdk-node compatible example application for this network at [fabric-node-patient-multi-org](https://github.com/Switch2Logic/fabric-node-patient-multi-org)`