let ws;
let reconnectInterval =1000;
//　メッセージ送信時に接続数が変わらないのはメッセージ送信が接続数に影響を与えないため
// 　接続数が変わるのはクライアントの接続・切断時のみ

function connect() {
  ws =new WebSocket("ws://localhost:8080");
  ws.onopen = () => {
    console.log("Connected websocket server");
    addMessage("Connected websocket server");
  }
  ws.onmessage = (e) => {
    addMessage(" Recieved "+e.data);
  }
  ws.onerror = (error) => {
    console.log("WebScoket Error "+error);
    addMessage("WebScoket Error "+error);
  }
  ws.onclose = () => {
    console.log("Disconnected websocket server");
    addMessage("Disconnected websocket server");
    setTimeout(() => {
      console.log("Reconnecting...");
      connect();
    },reconnectInterval)
  }
}

function sendMessage() {
  const input = document.getElementById("messageInput");
  const message = input.value;
  if (ws && ws.readyState === WebSocket.OPEN) {
    ws.send(message)
    addMessage("Sent "+message);
    input.value = "";
  }else {
    addMessage("WebSocket not connected");
  }
}

function addMessage(message) {
  const messages = document.getElementById("messages");
  const messageElement = document.createElement("div");
  messageElement.textContent = message;
  messages.appendChild(messageElement);
}

window.onload = connect;