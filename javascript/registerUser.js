/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { Wallets } = require('fabric-network');
const FabricCAServices = require('fabric-ca-client');
const fs = require('fs');
const path = require('path');
const organizations =  require('./configs/org.json');
const { intiliaze, executeQuery } = require('./db-pg');

intiliaze();

async function registerUser(org,user,pw,role='User') {
    try {    
        
        const usermspId = org + "MSP";
        const adminuser = '0';

        FabricCAServices.addConfigFile(path.join(__dirname, organizations[org].path));       
        const cas = FabricCAServices.getConfigSetting('certificateAuthorities');
        const orgs = FabricCAServices.getConfigSetting('organizations');
        const fabricCAKey = orgs[org].certificateAuthorities[0];
        const ca = new FabricCAServices(cas[fabricCAKey].url);        
            
        // Postgres Wallet for managing identities   
        // Check if the user is already exists in the Postgres DB
        let userInfo = await executeQuery(`SELECT * FROM public.user WHERE username = $1`, [user])
        if(typeof userInfo.rows[0] !== "undefined"){
            console.log('User exists in the database');
            return;
        }
        
        //check for the Admin user in the Postgres wallet
        let adminDbResult = await executeQuery(`SELECT * FROM public.certificatesinfo WHERE userid = $1`, [adminuser])
        if(typeof adminDbResult.rows[0] === "undefined"){
            console.log('An identity for the admin user "admin" does not exist in the wallet');
            console.log('Run the enrollAdmin.js application before retrying');
            return;
        }

        // Create a new In memory based wallet for managing identities.
        const wallet = await Wallets.newInMemoryWallet();

        const adminEnrollment = await ca.enroll({ enrollmentID: 'admin', enrollmentSecret: 'adminpw' });
        const x509IdentityAdmin = {
            credentials: {
                certificate: adminEnrollment.certificate,
                privateKey: adminEnrollment.key.toBytes(),
            },
            mspId: usermspId,
            type: 'X.509',
        };
        await wallet.put('admin',x509IdentityAdmin);        

        // Build user objects for authenticating with the CA
        const adminIdentity = await wallet.get('admin');
        const provider = wallet.getProviderRegistry().getProvider(adminIdentity.type);
        const adminUser = await provider.getUserContext(adminIdentity, 'admin');

        // Register the user, enroll the user, and import the new identity into the wallet.
        const secret = await ca.register({
            affiliation: 'org1.department1',
            enrollmentID: user,
            role: 'client'
        }, adminUser); 
        
        const enrollment = await ca.enroll({
            enrollmentID: user,
            enrollmentSecret: secret
        });

        const x509Identity = {
            credentials: {
                certificate: enrollment.certificate,
                privateKey: enrollment.key.toBytes(),
            },
            mspId: usermspId,
            type: 'X.509',
        };

        // Updating In-memory wallet
        await wallet.put(user, x509Identity);

        // Updating Postgres wallet
        let userCreation = await executeQuery(`INSERT INTO public.user
        (username, password, organization) VALUES ($1,$2,$3) returning id`,[user,pw,org])

        let reslt = await executeQuery(`INSERT INTO public.certificatesinfo (certificate, userid, privatekey) 
        VALUES ($1,$2,$3) returning id`, [x509Identity.credentials.certificate,userCreation.rows[0].id,x509Identity.credentials.privateKey])

        console.log('Successfully registered and enrolled user and imported it into the wallet');

        
    } catch (error) {
        console.error(`Failed to register user: ${error}`);
        process.exit(1);
    }
}

module.exports = { registerUser }
registerUser('Org1', 'Siddhartha', '12345');
// registerUser('Org2', 'Ganesh', '12345');