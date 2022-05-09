/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { Console } = require('console');
const { Gateway, Wallets } = require('fabric-network');
const fs = require('fs');
const path = require('path');
const { intiliaze, executeQuery } = require('./db-pg');
intiliaze();

async function invoke(org, user) {
    try {

        const usermspId = org + "MSP";

        // load the network configuration
        const ccpPath = path.resolve(__dirname, '..', '..', 'test-network', 'organizations', 'peerOrganizations', 'org1.example.com', 'connection-org1.json');
        let ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

        // Create a new In memory based wallet for managing identities.
        const wallet = await Wallets.newInMemoryWallet();

        // Postgres Wallet for managing identities   
        // Get the user id
        let userInfo = await executeQuery(`SELECT * FROM public.user where username = $1`, [user])
        if(typeof userInfo.rows[0] === "undefined"){
            console.log('An identity for the user does not exist in the wallet');
            console.log('Run the register user application before retrying');
            return;
        }
        const userId = JSON.stringify(userInfo.rows[0].id);
        
        // Get Certificate Details from postgreSQL database after registering a user after executing registerUser.js.
        let certInfo = await executeQuery(`SELECT * FROM public.certificatesinfo where userid = $1`, [userId])
        
        // Define the certificate.
        const x509Identity = {
            credentials: {
                certificate: certInfo.rows[0].certificate,
                privateKey: certInfo.rows[0].privatekey,
            },
            mspId: usermspId,
            type: 'X.509',  
        }; 
        
        // Adding postgres wallet details to the in memory wallet
        await wallet.put(userId, x509Identity);        
        
        // Get the enrolled user from the wallet.
        // const identity = await wallet.get(userId);
        
        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: userId, discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('mbse');

        // Transasct Asset Action Asset
        const assetActionPath = path.resolve(__dirname, 'sample_Assets', 'AssetAction_C_CR.json');
        let assetActionJson = JSON.parse(fs.readFileSync(assetActionPath, 'utf8'));

        // Submit the specified transaction.
        const result = await contract.submitTransaction('ManageMBSEAssets', JSON.stringify(assetActionJson));
        // await contract.submitTransaction('GetIdentityAttribute', 'desg');

        console.log(`Transaction has been submitted, result is: ${result.toString()}`);

        // Disconnect from the gateway.
        await gateway.disconnect();

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        process.exit(1);
    }
}

module.exports = { invoke }
invoke('Org1', 'Siddhartha');
// invoke('Org2', 'Ganesh');
