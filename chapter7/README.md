# 7章 DPソケットを使ったマルチキャスト通信

### UDPとTCPの用途の違い

本性で扱うUDPと6章で扱ったTCPの特徴を比較すると以下の表のようになる。

|  機能  |  UDP  |　TCP |
| ---- | ---- | ---- |
|  信頼性  |  低い(シンプル)  | 高い(ウィンドウ制御、再送制御、輻輳制御、順序整理を行う) |
|  コネクション |  なし(相手のポートに送りつける)  | あり(お互いを認識して送信する) |
|  通信速度  |  高い(とされている) | 低い(とされている) |
|  接続コスト  |  開始時に時間はかからない  | 1.5RTTの時間がかかる |
|  マルチキャスト  |  できる | できない |
|  ブロードキャスト  |  できる  | できない |

### UDPの具体的なアプリケーション

UDPはTCPと比較して、高速な通信速度であり、複数コンピュータに同時に送信できるという利点から以下のアプリケーションに応用されている。

- ドメイン名からIPアドレスを取得するDNSの通信。
- 時計合わせのためのプロトコルのNTP。
- ストリーミング動画、音声プロトコル。例えば、ブラウザ上でP2Pのための動画、音声通信プロトコルWebRTC。
- **かつては、** VPNなどの仮想ネットワークの土台としてもUDPが使われていた。仮想ネットワークでは、そこで張られたTCPコネクションがエラー訂正や順番の制御を行うため、その土台としてTCPを使うとTCP over TCPとなって無駄が多いから、というのがその理由。
- **かつては、** 伝達ロスがあまりないことが期待できる構内LAN専用の高速プロトコル

しかし、現在は、上記の使い分けやアプリケーションの応用は正しいとは限らない。それは以下の理由による。

- セキュリティ上の理由から、VPN接続でも暗号化のためにTLSを経由するSSL-VPNを使われることが増えてきている。SSL-VPNにも3通りの方式があるが、その中にはパケットをHTTP上にくるんで送信するものがある。この場合には、上で使うプロトコルがTCPの場合、どうしてもTCP over TCPになってしまう。
- 独自プロトコルを開発する場合、土台としてUDPを使うということは、通信環境が劣悪な状態での信頼性や、ネットワークに負荷をかけすぎて他の通信の邪魔をしない(フェアネス)など、そういった点について作り込みが必要になる。そのような「安定したプロトコル」のための開発コストは高い。つまり、ネットワークに明るい者が設計し、かつ大規模なフィールドテストが行える環境でないならば、そもそも独自プロトコルを制作しない方がよい。
- ハンドシェイクの時間がかからないUDPでは、短時間で完了するメッセージを大量に送受信する場合はメリットが大きいが、そうでないならばTCPとの差はない。

現在では、アプリケーションレイヤーで使われるプロトコルの多くはTCPを土台にしている。アプリケーションの開発の際は、「**ロスしても良い、マルチキャストが必要、ハンドシェイクの手間すら惜しいなどの特別な条件に合致する場合以外はTCP**」という選択でよい。

### UDPとTCPの処理の流れの違い

`base/README.md`を参照。

### UDPのマルチキャストの実装例

`multicast/README.md`を参照。

### UDPとTCPの機能面の違い

ここでは、アプリケーションでTCPとUDPを使い分けるときの参考になるように、両者の機能面での違いを紹介する。

**1. TCPには再送処理とフロー処理がある**

TCPでは送信するメッセージにシーケンス番号が入っているので、受信側ではこの数値を見て、もしパケットの順序が入れ替わっていたときは順序を並べ直す。受信側はメッセージを受け取ると、受信したデータのシーケンス番号とペイロードサイズの合計を確認応答番号として返信する。送信側はこの応答確認番号が受け取れず、届いたことが確認できない場合は、落ちたと思ってもう一度送信する。これが**再送処理** である。

また、TCPには**ウィンドウ制御** という機能があり、受信側がリソースを用意できていない状態で送信リクエストが集中して通信内容が失われたりすぎるのを防ぐ。具体的には、受信用のバッファ(TCPでは**ウィンドウ** と呼ばれる)をあらかじめ決めておき、送信側ではそのサイズ(**ウィンドウサイズ**)を受信側からの受信確認を待たずにデータを送信できる。このウィンドウサイズはコネクション確率時にお互いに確認しあう。受信できるウィンドウサイズを受信側から通信側に伝えて通信料を制御することができる。これを**フロー制御** という。

UDPには、TCPにおけるウィンドウやシーケンス番号を利用した再送処理やフロー処理の機能がない。クライアントからサーバーへと一方的にデータを送りつける。受信確認もなく、順番のソートや再送処理もない代わりに高速になっている。

**2. UDPではフレームサイズも気にしよう**

TCPもUDPも、その下のデータリング層の影響を受ける。ひとかたまりで送信できるできるデータの大きさは、通信処理の種類やルーターなどの設定によって変わり、ある経路でひとかたまりで送信できるデータの上限のことをその経路の**最大転送単位(MTU)** という。

MTUに収まらないデータは、IPレベル(TCP/UDPの下のネットワーク層)で複数のパケットに分割される。これを**IPフラグメンテーション** と呼ぶ。IPフラグメンテーション自体はIPアドレスで再結合はしてくれるが、分割された最後のバケットが来るまでUDPパケットとして未完成のままなので、アプリケーション側にはデータが流れてくることはない。データが消失したら受信待ちのタイムアウトも発生するし、UDPを使うメリットが薄れてしまう。UDPの売りである応答性の高さをカーネル内部の結合待ちで無駄にしないために、イーサネットのフレームサイズを意識したアプリケーションプロトコルの設計が必須である。

巨大なデータをUDPとして送信するデメリットはもう一つある。IPレイヤーでデータを結合するといっても、IPレイヤーでデータを結合してくれるといっても、IPレイヤーやその上のUDPレイヤーで取り扱えるデータは約64キロバイトまでである。それ以上のデータになると別のUDPパケットとして取り扱うしかない。TCPであれば大きなデータでも受信側のアプリケーションでの扱いを気にせず遅れる。UDPではデータの分割などはアプリケーションで面倒を見るしかない。逆に言えば、データの最小単位がこの64キロバイト以下であれば、アプリケーション内でのデータの取り扱いはシンプルになる。

**3. 輻輳制御とフェアネス**

**輻輳制御** とは、ネットワークの輻輳(渋滞)を避けるように流量を調整し、そのネットワークの最大効率で通信できるようにするとともに、複数の通信をお互いにフェアに行えるようにする仕組みである。TCPには輻輳制御制御が備わっており、そのアルゴリズムにはさまざまな種類がある。輻輳制御の目的は、自分の通信だけを最大化することではなく、他の通信回路にもきちんと帯域を譲り、全員が問題なく通信を継続しつつ必要な最大速度が得られることである。これを**フェアネス** と呼び、TCPにおける大事な価値である。

しかし、UDPには制御制御の仕組みはない。流量の制御は、UDPを利用する各プログラムに委ねられている。そのため、UDPを利用する各プログラムに委ねられている。そのため、UDPとTCPを利用するアプリケーションがそれぞれあって、UDPを利用するアプリケーションでフェアネスが考慮されていない場合は、両方の通信が重なったときに遠慮する機能が組み込まれたTCPの通信速度だけが極端に落ち込むこともある。