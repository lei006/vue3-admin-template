require('dotenv').config()

const mongoose = require('mongoose')

let db_uri = process.env.DB_URI;


const conn = mongoose.createConnection(db_uri, {
    //userNewUrlParser:true,
    useUnifiedTopology:true,
});

conn.on('error', function(error){
    console.log('mongodb 连接出错', error)
});
conn.once('open', function() {
  // we're connected!
  console.log('mongodb 连接成功')
});

module.exports = {Schema:mongoose.Schema, mongoose: conn}
//module.exports = getMongooseConn;
