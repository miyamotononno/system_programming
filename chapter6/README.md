# 6章 TCPソケットの理論と実装

Go言語では、ソケットはchapter2で紹介されたio.Writer(Response.Write()、 Request.Write())とchapter3で紹介されたio.Readerによって成り立っている。

最も基本的なTCPソケットは`http/`ディレクトリで実装した。それをさらに発展させた工夫は以下の表でまとめてある

|  手法  |  効果  |　ディレクトリ |
| ---- | ---- | ---- |
|  Keep-Alive  |  再接続のコストを削減  | `httpWithKeepAlive/` |
|  圧縮  |  通信時間の短縮  | `compressGzip/` |
|  チャンク  |  レスポンスの送信開始を早める  | `compressGzip/` |
|  パイプライニング  |  通信の多重化  | `pipelining/` (ただし、実装はしていない) |