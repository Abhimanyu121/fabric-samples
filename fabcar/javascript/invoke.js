/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { Gateway, Wallets } = require('fabric-network');
const path = require('path');

const ccpPath = path.resolve(__dirname, '..', '..', 'first-network', 'connection-org1.json');

async function main() {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get('user1');
        if (!identity) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('fabcar');

        // Submit the specified transaction.
        // createCar transaction - requires 5 argument, ex: ('createCar', 'CAR12', 'Honda', 'Accord', 'Black', 'Tom')
        // changeCarOwner transaction - requires 2 args , ex: ('changeCarOwner', 'CAR10', 'Dave')
        await contract.submitTransaction('PutHistory', "qwerty","{\r\n   \"data\":[\r\n      {\r\n         \"name\":\"parkinson's disease\",\r\n         \"symptoms\":\"High blood pressure\",\r\n         \"doctor\":\"batra\"\r\n      },\r\n      {\r\n         \"name\":\"Hakuna matata\",\r\n         \"symptoms\":\"High blood pressure\",\r\n         \"doctor\":\"dr .Batra\"\r\n      },\r\n      {\r\n         \"name\":\"fulkerson disease\",\r\n         \"symptoms\":\"High blood pressure\",\r\n         \"doctor\":\"Dr.derp\"\r\n      }\r\n   ]\r\n}");

       // await contract.submitTransaction('PutPrescription', "qwerty","{ \r\n\"data\":[ \r\n{ \r\n\"name\":\"paracetamol\",\r\n\"dose\":6,\r\n\"time\":3,\r\n\"meal\":\"before\"\r\n},\r\n{ \r\n\"name\":\"ibrufin\",\r\n\"dose\":6,\r\n\"time\":3,\r\n\"meal\":\"before\"\r\n},\r\n{ \r\n\"name\":\"bitadyine\",\r\n\"dose\":6,\r\n\"time\":3,\r\n\"meal\":\"before\"\r\n}\r\n]\r\n}");
        console.log('Transaction has been submitted');

        // Disconnect from the gateway.
        await gateway.disconnect();

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        process.exit(1);
    }
}

main();
