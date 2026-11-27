import { app, BrowserWindow, screen } from 'electron'

const createWindow = () => {
  const { width, height } = screen.getPrimaryDisplay().workAreaSize
  const winWidth = 200
  const winHeight = 200
  const margin = 20
  const win = new BrowserWindow({
    width: winWidth,
    height: winHeight,
    x: width - winWidth - margin,
    y: height - winHeight - margin,
    frame: false,
    transparent: true,
    alwaysOnTop: true,
  })

  win.loadFile('index.html')
}

app.whenReady().then(() => {
  createWindow()
})