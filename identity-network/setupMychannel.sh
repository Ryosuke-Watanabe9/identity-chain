# Create identity

export CHANNEL_NAME=identity

peer channel create -o orderer.identity.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls $CORE_PEER_TLS_ENABLED --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/identity.com/orderers/orderer.identity.com/msp/tlscacerts/tlsca.identity.com-cert.pem

# Environment variables for Org1.PEER0

export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.identity.com/users/Admin@org1.identity.com/msp
export CORE_PEER_ADDRESS=peer0.org1.identity.com:7051
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.identity.com/peers/peer0.org1.identity.com/tls/ca.crt

peer channel join -b identity.block

# Environment variables for Org1.PEER1

export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.identity.com/users/Admin@org1.identity.com/msp
export CORE_PEER_ADDRESS=peer1.org1.identity.com:7051
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.identity.com/peers/peer1.org1.identity.com/tls/ca.crt

peer channel join -b identity.block

# Environment variables for Org2.PEER0

export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.identity.com/users/Admin@org2.identity.com/msp
export CORE_PEER_ADDRESS=peer0.org2.identity.com:7051
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.identity.com/peers/peer0.org2.identity.com/tls/ca.crt

peer channel join -b identity.block

# Environment variables for Org2.PEER1

export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.identity.com/users/Admin@org2.identity.com/msp
export CORE_PEER_ADDRESS=peer1.org2.identity.com:7051
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.identity.com/peers/peer1.org2.identity.com/tls/ca.crt

peer channel join -b identity.block

echo ---------------------------------------------end-----------------------------------------------
