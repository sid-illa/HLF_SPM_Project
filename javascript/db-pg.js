const { Pool } = require('pg')

const dbPool = {
    user: 'postgres',
    host: 'localhost',
    database: 'capstone_blockchain',
    password: 'postgres',
    port: 5432,
}
var pool;

function intiliaze() {
    pool = new Pool(dbPool)
}

async function getConnection() {
    return await pool.connect()
}

async function releaseConnection(client) {
    if (client) {
        return await client.release(true);
    } else {
        return await null;
    }
}
async function executeQuery(query, params = undefined) {
    let client, result;
    try {
        client = await getConnection();
        if (params) {
            result = await client.query(query, params)
        }
        else {
            result = await client.query(query)
        }
        return result;
    } catch (error) {
        throw error
    }
    finally {
        await releaseConnection(client)
    }
}

module.exports = { intiliaze, executeQuery }