// ウェブソケットのクライアント側のテストコード
let ws = new WebSocket("ws://localhost:8080")

ws.onopen = () => {
  console.log("Connected")
  ws.send(" Hello Server お元気ですか？")
}
ws.onmessage = (e) => {
  console.log(e.data)
}
ws.onerror = (e) => {
  console.log(e)
}
ws.onclose = () => {
  console.log("Disconnected")
}
