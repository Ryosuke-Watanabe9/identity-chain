$HOME/bin/cryptogen generate --config=./crypto-config.yaml

# 秘密鍵のファイル名を固定
mv ./crypto-config/peerOrganizations/org1.identity.com/ca/*_sk ./crypto-config/peerOrganizations/org1.identity.com/ca/CA1_PRIVATE_KEY

$HOME/bin/configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block

export CHANNEL_NAME=identity

$HOME/bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME
$HOME/bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
$HOME/bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP

echo config-done

# dockerの起動
CHANNEL_NAME=$CHANNEL_NAME TIMEOUT=10000 docker-compose -f docker-compose-cli.yaml -f docker-compose-couch.yaml -f docker-compose-ca.yaml up -d

# this is shell script to set channel configlation
docker cp setupMychannel.sh cli:/opt/gopath/src/github.com/hyperledger/fabric/peer

echo all-done
