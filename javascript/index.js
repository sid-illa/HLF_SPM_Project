const express = require('express')
var cors = require('cors')
const bodyParser = require('body-parser');
const fs = require('fs'),
    path = require('path');

const { intiliaze, executeQuery } = require('./db-pg')
const { registerUser } = require('./registerUser')
const app = express()
const port = 3000

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
app.use(cors())


intiliaze();
app.get('/login', async (req, res, next) => {
    let reslt = await executeQuery('SELECT * FROM public.user')

    res.send(reslt.rows);
})

const multer = require('multer')
const upload = multer({ storage: multer.memoryStorage() })

app.post('/login', async (req, res, next) => {
    let reslt = await executeQuery('SELECT * FROM public.user where username = $1', [req.body.username])
    if (reslt.rows.length != 0 && reslt.rows[0].password == req.body.password) {
        res.status(200).json([{ "status": 200, Message: 'Success', data: reslt.rows[0] }])
    }
    else {
        res.status(401).json([{ "status": 401, Message: 'Invalid Credentials' }])
    }
    res.send(reslt.rows);
})

app.post('/register', async (req, res, next) => {
    try {
        //let reslt = await executeQuery(`insert into public.user (username,password) 
        //values ($1,$2) returning id`, [req.body.username, req.body.password])
        registerUser ( 'Org1', req.body.username, req.body.password)
        res.status(200).json([{ "status": 200, Message: 'User Registered!!' }])
    }
    catch (err) {
        console.log(err)
        next(err)
    }

})


app.post('/uploadFile', upload.single('file'), async (req, res, next) => {
    console.log(req.file.originalname)
    console.log(req.file.buffer.toString('base64'))
    let reslt = await executeQuery(`insert into public.certificates (file,filename) 
        values ($1,$2) returning id`, [req.file.buffer.toString('base64'), req.file.originalname])
    res.status(200).json([{ "status": 200, Message: 'User Registered!!' }])
    res.send(reslt.rows);
})



app.post('/extractfile', async (req, res) => {
    try {
        console.log(req.body.filepath + req.body.filename);
        try {
            fs.readFile(path.join(req.body.filepath + req.body.filename), 'utf8', async (err, data) => {
                if (err) {
                    console.error(err)
                    return
                }
                console.log(data)
                let reslt = await executeQuery(`insert into public.certificatesinfo (filename,filepath,fileinfo,userid) 
                values ($1,$2,$3,$4) returning id`, [req.body.filename, req.body.filepath, data, req.body.userid || null])
                res.status(200).json([{ "status": 200, Message: 'Data Saved!!' }])
            })
        }
        catch (err) {
            if (err.toString().includes('no such file or directory')) {
                res.status(400), json([{ "status": 400, Message: 'Bad Request!!' }])
                console.log('err')
            }
        }
    }
    catch (err) {

        if (err.toString().includes('no such file or directory')) {
            res.status(400), json([{ "status": 400, Message: 'Bad Request!!' }])
            console.log('err')
        }
    }

})
app.get('/certicatesdata', async (req, res, next) => {
    let reslt = await executeQuery('SELECT * FROM public.certificatesinfo')
    // if (reslt.rows.length != 0 && reslt.rows[0].password == req.body.password) {
    res.status(200).json({ "status": 200, Message: 'Success', data: reslt.rows })
    // }
    // else {
    //     res.status(401).json([{ "status": 401, Message: 'Invalid Credentials' }])
    // }
    res.send(reslt.rows);
})

app.get('/usercertificate/:userid', async (req, res, next) => {
    try {
        let reslt = await executeQuery(`SELECT * FROM public.certificatesinfo where 
        userid = $1`, [req.params['userid']])
        res.status(200).json({ "status": 200, Message: 'Success', data: reslt.rows })
        res.send(reslt.rows);

    } catch (error) {
        throw (error)
    }
})



app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})