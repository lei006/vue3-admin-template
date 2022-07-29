const { ipcMain } = require('electron')

const {spawn} = require('child_process')
const path=require("path");



// 子进程，参考 https://www.cnblogs.com/chyingp/p/node-learning-guide-child_process.html
let child_process = undefined;
function db_server_start(port, data_path, callback) {
    

    let command = ".\\extensions\\mongo\\mongod.exe";
    let command_path = path.resolve(process.cwd(), command)

    let args = [`--dbpath=${data_path}`,`--port=${port}`];
    console.log("mongo run at", command_path, args);

    child_process = spawn(command_path, args)

    child_process.on('exit', (code) => {
        if(callback) {
            //callback({state:"stoped", pid:child_process.pid, code, port, path, command, args});
        }
        console.log("mongo exit", code);
    });
    
    if(callback) {
        callback({state:"started", pid:child_process.pid , port, path, command, args});
    }
}

//var kill = require('tree-kill');

async function Stop() {

    try {
        if(child_process) {
            //不知道为什么，无法关闭....
            child_process.kill('SIGINT');
            child_process.kill();
            child_process = undefined;
        }        
    } catch (error) {
        console.error("helper stop error:", error)
    }
}

async function Start(callback){


    let port = process.env.DB_PORT;
    let data_path = process.env.DB_DATA_PATH;


    let all_path = path.resolve(process.cwd(), data_path)


    db_server_start(port, all_path, function(state){
        //console.log("db run: ", state);
        if(callback){
            callback();
        }
    });
    
}



export default {Start, Stop}
