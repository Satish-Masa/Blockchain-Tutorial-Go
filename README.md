Blockchain Tutorial
====


簡単なブロックチェーンのプロトタイプを作成する。  
Part1では、ブロックとブロックチェーンを作成する。
 
# Block
Part1では重要な情報を簡略化したバージョンを作成する。  

## Blockの実装  
`Timestamp`　　　　ブロックが作成された時間を記録  
`Data`　　　　　　　ブロックに含まれる取引情報を記録  
`PrevBlockHash`　　前のブロックのハッシュの記録  
`Hash`　　　　　　　BlockのTimestanpとDataとPrevBlockHashのハッシュ値  
 
## Hashの作成  
BlockのTimestampとData、PrevBlockHashを取得して連結し、連結された組み合わせでSHA-256ハッシュを計算する。  

## Blockの作成を簡略化する関数の作成  
BlockにTimestampとdata、prevBlockHashを代入しBlockを作成する。Blockをハッシュ化し、BlockのHashにセットする。  


# Blockchain  
ブロックは挿入順に保存され、各ブロックは前のブロックにリンクする機能を実装する。  

## Blockchainの実装  
`blocks`　　ブロックの配列を取得する  

## Blockの追加  
前のブロックのハッシュを取得しPrevHashとdataを引き渡し、新しいBlockを作成する。そして、新しいブロックをブロックチェーンに追加する。  

## 最初のBlockを作成し、Blockchainを作成する  
Genesis Block(最初のブロックの名称)を作成し、次にブロックチェーンを作成する。
