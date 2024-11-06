# TCP通信とWebSocketによる通信の学習
### １. 主な特徴の違い

| 特徴       | TCP                                        | WebSocket                            |
|------------|--------------------------------------------|--------------------------------------|
| 通信方法   | クライアントからサーバーへの一方向通信が基本 | 双方向通信が可能                     |
| プロトコル層 | トランスポート層（OSIモデルの第4層）       | アプリケーション層（TCPの上に構築）  |
| コネクション | 常に接続状態を維持                         | HTTP経由で接続後、常に接続が維持される |
| ヘッダー     | 特別なヘッダーはない                       | 初期接続にHTTPヘッダーを使用（Upgrade） |
| リアルタイム | リアルタイム通信には工夫が必要             | リアルタイム通信に適している         |
| 利用例     | HTTP、SMTP、FTPなどのプロトコルの基盤       | チャット、オンラインゲーム、リアルタイム通知など |
