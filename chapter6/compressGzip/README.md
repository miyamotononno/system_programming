# 速度改善(2): 圧縮

HTTPの速度アップ手法として圧縮も用いられる。圧縮してもバケット伝達の速度は変わらないが、転送を開始してから終了するまでの時間は変わる。

![HTTPのgzip圧縮](https://ascii.jp/img/2016/12/14/558903/l/aa1ff6a5c6af2259.jpg)

### gzip対応したクライアント

リクエスト生成部を改造して、自分が対応しているアルゴリズムを宣言する(サーバーから自分が理解できない圧縮フォーマットでデータを送りつけれられても、クライアントではそれを読み込めないため)。具体的には、リクエストヘッダーの”Accept-Encoding”にgzipを設定する。

レスポンスを受け取る部分では、Accept-Encodingで表明した圧縮メソッドにサーバーが対応しているかどうかは、Content-Encodingヘッダーを見ればわかる。今回のサンプルで対応するアルゴリズムは1種類だけだったが、複数の候補を提示してサーバーに選ばせることも可能。

### gzip対応したサーバー

レスポンスの作成部分では、クライアントがgzipを受け入れ可能かどうかで中に入れるコンテンツを変える。

注意点として、ヘッダーが圧縮されないこと。そのため少量のデータを通信するほど効率が悪くなる。20byte足らずのサンプルの文字列ではgzipのオーバーヘッドの方が大きく、際イズが倍増してしまっているが、大きいサイズになれば効果が出てくる。
なお、HTTPで圧縮されるのはレスポンスのボディーだけで、リクエストのボディーの圧縮はない。ヘッダーの圧縮はHTTP/2になって初めて導入された。