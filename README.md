7リットルと3リットルの容器から5リットルの水量を作る問題に おいて,深さ8(固定)で幅優先探索を行って解を見つけよ.すでに調べた状態を省く方式を用いよ.
という問題を解いてみたもの。

次の状態は
1.容器の水を捨てて空にする
2.容器Aの水を容器Bが満杯になるまで注ぐ
3.容器が満杯になるまで水を注ぐ
容器の数が2つの場合は6通り

n個
1.	n通り
2.	n*(n-1)通り
3.	n通り

Golangのinterfaceで各メソッドを定義して、上記の動作をさせて一般化する。
1~3の動作を容器Aから容器Bへ入れるという1動作だけのみで実現できそう?!

とりあえず、Golangなので、力任せに指定の深さまで全通り探索します

ヘルプはこんな感じ
Usage of bucket:
  -d=8: depth
  -f=true: first flag
  -v=5: aim value

  実行コマンド
go run *.go -d=8 -f -v=5 3 7
　結果
Sum:8
 MinStep:8
[1]: 3/ 3
[2]: 5/ 7

 ImaginaryBucket ===>      Bucket(0/7) ( ImaginaryBucket,      Bucket(7/7))
[1]: 0/ 3
[2]: 7/ 7
     Bucket(7/7) ===>      Bucket(0/3) (     Bucket(4/7),      Bucket(3/3))
[1]: 3/ 3
[2]: 4/ 7
     Bucket(3/3) ===>  ImaginaryBucket (     Bucket(0/3),  ImaginaryBucket)
[1]: 0/ 3
[2]: 4/ 7
     Bucket(4/7) ===>      Bucket(0/3) (     Bucket(1/7),      Bucket(3/3))
[1]: 3/ 3
[2]: 1/ 7
     Bucket(3/3) ===>  ImaginaryBucket (     Bucket(0/3),  ImaginaryBucket)
[1]: 0/ 3
[2]: 1/ 7
     Bucket(1/7) ===>      Bucket(0/3) (     Bucket(0/7),      Bucket(1/3))
[1]: 1/ 3
[2]: 0/ 7
 ImaginaryBucket ===>      Bucket(0/7) ( ImaginaryBucket,      Bucket(7/7))
[1]: 1/ 3
[2]: 7/ 7
     Bucket(7/7) ===>      Bucket(1/3) (     Bucket(5/7),      Bucket(3/3))
[1]: 3/ 3
[2]: 5/ 7

力任せ以外で判定する場合は、
GCDいわゆる最大公約数がキーワードとなるらしいですね。
3と7ならば、1。
とすると、1つの差分で全てできるということになるのかね...。
つまり、この性質をとりこんだ評価関数とやらを用意すればいいのかね。
今は、よくわからんのでとりあえず、ここまで。
