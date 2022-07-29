

const { app, BrowserWindow,ipcMain } = require('electron')


const winURL = process.env.NODE_ENV === 'development'
  ? `http://localhost:9080`
  //: `https://www.zhihu.com/`
  : `file://${__dirname}/index.html`

//app.commandLine.appendSwitch('ignore-certificate-errors')



let win_main = null;

function createWindow (debug) {  

  // 创建浏览器窗口
  win_main = new BrowserWindow({
    width: 1024,
    height: 720,
    show: false,
    frame:false,
    minWidth:640,
    minHeight:480,
    webPreferences: {
      nodeIntegration: true,
      contextIsolation: false
    },
    titleBarStyle:"hidden",
    hasShadow:true,
    autoHideMenuBar:(debug==false),
  })

  win_main.once('ready-to-show', () => {
    console.log("win main show");
    win_main.show()
  })


  //main.js
  win_main.on('close', (event) => {
    
  });
  //win_main.loadFile('./src/views/main.html')
  win_main.loadURL(winURL)
  
  console.log("win main loadURL", winURL);

  if(debug == true){
    win_main.webContents.openDevTools()
  }
  const options = {
    type: 'question',
    buttons: ['Cancel', 'Yes, please', 'No, thanks'],
    defaultId: 2,
    cancelId: 0,
    title: 'Question',
    message: 'my window?',
    detail: 'It does not really matter',
    checkboxLabel: 'remember',
    checkboxChecked: true,
  }; 
  
  //dialog.showMessageBoxSync(win_main, options);

}






ipcMain.on('app-mini', (event, arg) => {
    win_main.minimize();
})





async function Start(){


    createWindow();
    
}


export default {Start}
