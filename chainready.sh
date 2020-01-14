rm -rf ../newcoin_data
mkdir ../newcoin_data
./build/bin/geth init --datadir ../newcoin_data/ genesis.json
./build/bin/geth  --datadir ../newcoin_data --rpc  --rpcport 8545 --mine --miner.threads=1 --etherbase=0x59641146B8204365f9B0EDBaBE6F09362f02bC3e console

