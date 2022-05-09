/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const FabricCAServices = require('fabric-ca-client');
const { Wallets } = require('fabric-network');
const fs = require('fs');
const path = require('path');
const { intiliaze, executeQuery } = require('./db-pg');
intiliaze();

async function main() {
    try {
        const userid='0';
        const usermspId = "Org1MSP";

        // load the network configuration
        const ccpPath = path.resolve(__dirname, '..', '..', 'test-network', 'organizations', 'peerOrganizations', 'org1.example.com', 'connection-org1.json');
        const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

        // Create a new CA client for interacting with the CA.
        const caInfo = ccp.certificateAuthorities['ca.org1.example.com'];
        const caTLSCACerts = caInfo.tlsCACerts.pem;
        const ca = new FabricCAServices(caInfo.url, { trustedRoots: caTLSCACerts, verify: false }, caInfo.caName);

        // Create a new In-memory based wallet for managing identities.
        const wallet = await Wallets.newInMemoryWallet();

        //Postgres Wallet for managing identities
        // Get Certificate Details from postgreSQL database after registering when executed registerUser.js.
        let reslt = await executeQuery(`SELECT * FROM public.certificatesinfo WHERE userid = $1`, [userid])
        if(typeof reslt.rows[0] !== "undefined"){
            console.log('An identity for the admin user "admin" already exists in the postgres wallet');
            return;
        }

        // Enroll the admin user, and import the new identity into the wallet.
        const enrollment = await ca.enroll({ enrollmentID: 'admin', enrollmentSecret: 'adminpw' });
        const x509Identity = {
            credentials: {
                certificate: enrollment.certificate,
                privateKey: enrollment.key.toBytes(),
            },
            mspId: usermspId,
            type: 'X.509',
        };

        // Updating Postgres wallet DB
        let result = await executeQuery(`INSERT INTO public.certificatesinfo (certificate, userid, privatekey) 
        values ($1,$2,$3) returning id`, [x509Identity.credentials.certificate,0,x509Identity.credentials.privateKey])

        // Updating In-memory wallet
        await wallet.put('admin', x509Identity);

        console.log('Successfully enrolled admin user "admin" and added to the postgres wallet');

    } catch (error) {
        console.error(`Failed to enroll admin user "admin": ${error}`);
        process.exit(1);
    }
}

main();
