//updating Idenitities for client application using javaScript
//reference https://medium.com/coinmonks/update-an-identity-in-hyperledger-fabric-using-node-sdk-1b93edb6cb2b

//Before updating IdentityDetails checking the attribute present
//fabric-ca-client identity list --tls.certfiles "../../fabric-ca/org1/tls-cert.pem"


'use strict';

const FabricCAServices = require('fabric-ca-client');
const { Gateway, Wallets } = require('fabric-network');
const { intiliaze, executeQuery } = require('./db-pg');
const path = require('path');
const fs = require('fs');
const organizations =  require('./configs/org.json');
const ccpPath = path.resolve(__dirname, '..', '..', 'test-network', 'organizations', 'peerOrganizations', 'org1.example.com', 'connection-org1.json');
const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));
intiliaze();

async function updateIdentity(org,user) {
    try {

    const usermspId = org + "MSP";
    const adminuser = '0';

    // Create a new CA client for interacting with the CA.
    FabricCAServices.addConfigFile(path.join(__dirname, organizations[org].path));       
    const cas = FabricCAServices.getConfigSetting('certificateAuthorities');
    const orgs = FabricCAServices.getConfigSetting('organizations');
    const fabricCAKey = orgs[org].certificateAuthorities[0];
    const ca = new FabricCAServices(cas[fabricCAKey].url);

    // Postgres Wallet for managing identities   
    // Get the user id
    let userInfo = await executeQuery(`SELECT * FROM public.user WHERE username = $1`, [user])
    if(typeof userInfo.rows[0] === "undefined"){
        console.log('An identity for the user does not exist in the wallet');
        console.log('Run the registerUser application before retrying');
        return;
    }
    const userId = JSON.stringify(userInfo.rows[0].id);
    let userName = userInfo.rows[0].username;
    // Get Certificate Details from postgreSQL database after registering a user after executing registerUser.js.
    let userCertInfo = await executeQuery(`SELECT * FROM public.certificatesinfo WHERE userid = $1`, [userId])
 
    //check for admin certificate in the Postgres DB
    let adminDbResult = await executeQuery(`SELECT * FROM public.certificatesinfo WHERE userid = $1`, [adminuser])
    if(typeof adminDbResult.rows[0] === "undefined"){
        console.log(adminDbResult.rows[0]);
        console.log('An identity for the admin user "admin" does not exist in the wallet');
        console.log('Run the enrollAdmin.js application before retrying');
        return;
    }

    // Create a new in memory based wallet for managing identities.
    const wallet = await Wallets.newInMemoryWallet(); 

    const enrollment1 = await ca.enroll({ enrollmentID: 'admin', enrollmentSecret: 'adminpw' });
    const x509IdentityAdmin = {
        credentials: {
            certificate: enrollment1.certificate,
            privateKey: enrollment1.key.toBytes(),
        },
        mspId: usermspId,
        type: 'X.509',
    };

    // Add the Admin Identity to the wallet
    await wallet.put('admin',x509IdentityAdmin);
    
    // Get the Admin Identity from the wallet
    const adminIdentity = await wallet.get('admin');    
     
    // defining the certificate for the user.
    const x509Identity = {
        credentials: {
            certificate: userCertInfo.rows[0].certificate,
            privateKey: userCertInfo.rows[0].privatekey,
        },
        mspId: usermspId,
        type: 'X.509',  
    }; 
    //adding the user identity to the In memory wallet
    await wallet.put(userId, x509Identity);

    // Check to see if we've already enrolled the user.
    const userIdentity = await wallet.get(userId);
    if (!userIdentity) {
        console.log('An identity for the user does not exist in the wallet');
        console.log('Run the registerUser.js application before retrying');
        return;
    }

    // Build a user object for authenticating with the CA
    const provider = wallet.getProviderRegistry().getProvider(adminIdentity.type);
    const adminUser = await provider.getUserContext(adminIdentity, 'admin');
    const newAppUser = await provider.getUserContext(userIdentity, userId);

    // update identity
    let updateObj = {
        type : "client",
        affiliation : 'org1.department1',
        attrs : [{ name: "organization", value: org, ecert: true },
                { name: "role", value: "cse", ecert: true }] ,
        caname: "ca_peer" + org
    }

    //using IdentityService Class to update the user Identity
    const identityService = ca.newIdentityService();

    const response = await identityService.update(userName, updateObj, adminUser); //Passing the id passed at registerUserAPI

    // reenroll user to get a updated certificate
    const newEnrollment = await ca.reenroll(newAppUser);
    const newX509Identity = {
            credentials: {
                    certificate: newEnrollment.certificate,
                    privateKey: newEnrollment.key.toBytes(),
            },
            mspId: usermspId,
            type: 'X.509',
    };

    // updating In-memory wallet
    await wallet.put(userId, newX509Identity);
    
    // updating postgres db wallet
    let reslt = await executeQuery(`UPDATE public.certificatesinfo SET certificate = $1 , privatekey = $2 WHERE userid = $3 
    returning id`, [newX509Identity.credentials.certificate,newX509Identity.credentials.privateKey, userId])

    console.log("user Identity attributes: ", response.result.attrs);

    } catch (error) {
        console.error(  `Failed to update the user attributes: ${error}`);
        process.exit(1);
    }
}
module.exports = { updateIdentity }
updateIdentity('Org1', 'Siddhartha');
// updateIdentity('Org2', 'Ganesh');