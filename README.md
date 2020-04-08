Blockchain Tutorial
====


簡単なブロックチェーンのプロトタイプを作成する。  
Part1では、ブロックとブロックチェーンを作成する。
 
# Block
Part1では重要な情報を簡略化したバージョンを作成する。  

## BlockのStruct  
`Timestamp`　　　　ブロックが作成された時間  
`Data`　　　　　　　ブロックに含まれる取引情報  
`PrevBlockHash`　　前のブロックのハッシュ  
`Hash`　　　　　　　BlockのTimestanpとDataとPrevBlockHashのハッシュ値  
 
## SetHash()  
BlockのTimestampとData、PrevBlockHashを取得して連結し、連結された組み合わせでSHA-256ハッシュを計算する。  

## NewBlock()  
BlockにTimestampとdata、prevBlockHashを代入しBlockを作成する。Blockをハッシュ化し、BlockのHashにセットする。  


# Blockchain  
ブロックは挿入順に保存され、各ブロックは前のブロックにリンクする機能を実装する。  

## BlockchainのStruct  
`blocks`　　ブロックの配列  

## AddBlock()  
前のブロックのハッシュを取得しPrevHashとdataを引き渡し、新しいBlockを作成する。そして、新しいブロックをブロックチェーンに追加する。  

## NewBlockchain()  
Genesis Block(最初のブロックの名称)を作成し、次にブロックチェーンを作成する。
