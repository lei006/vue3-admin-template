
const {app} =  require('./core')



function web_server_start(port, callback) {

    
    app.listen(port, () => {
      if(callback){
        callback({state:"started", port});
      }
      console.log(`web server listening at:${port}`)
    })

}


async function Start(){

  let port = process.env.WEB_PORT;
  web_server_start(port, function(state){

  })


}


export default {Start}