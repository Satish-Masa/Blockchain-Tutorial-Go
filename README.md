Blockchain Tutorial
====

Part2では、Proof of WorkとHashCashを実装する。
 
 

# Proof of Work (PoW)  
PoWとは、ブロックチェーンにおいて正当な取引情報なのか決定・管理するための仕組み（コンセンサスアルゴリズム）のこと。なぜ必要なのかというと、ブロックチェーンはP2P通信なので悪意のある一部のユーザーにより正当な取引情報が改ざんされる恐れがあるため、その問題の解決のため。具体的に何をするかというと、ハッシュ値がある値以下になるような入力値を探している

## PoWのstruct
`block`　　　Blockへのポインター  
`target`　　 以前作成したハッシュと比べるため、big.Intで定義  

## NewProofOfWork()  
targetにbig.Intを１で初期化し、256-targetBits分左にシフトする。そして、Proof of Workにセットし返す。  

## prepareData()
PrevBlockHashや、Data、Timestamp、targetBits、nonceをbyte[]として返す。  

## Run()
maxNonceはオーバーフローを防ぐためmath.MaxInt64までとしておく。ループでは、  
1. PrevBlockHashや、Data、Timestamp、targetBits、nonceのbyte[]を準備する。  
2. SHA-256でハッシュを作成する。  
3. ハッシュを大きな整数に変換する。  
4. 作成した整数とtargetを比較し、値がtarget未満だった場合ループを抜け出す、値がtarget以上だった場合はnonceを+1する。  

## Validdata()
改ざんされていないとこの確認を行う。blockのNonceを使いハッシュを作成し、それを整数に変換する。それがPoWのtarget未満か比較し、bool型で返す。