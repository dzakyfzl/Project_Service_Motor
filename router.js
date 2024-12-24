const express = require('express')
const path = require('path')
const app = express()
const port = '5050'

app.use('/',express.static(path.join(__dirname,'Frontend')))

app.listen(port,()=>{
    console.log('frontend run at http://localhost:5050/')
})
