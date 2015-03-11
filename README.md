Go言語開発環境作成メモ

# Go言語環境導入

[Go言語環境導入](http://golang.jp/install)


# IntelliJ IDE

せっかく静的言語だから入力補完を有効にしたい.
Vimでも補完出来るように設定出来るらしいが、うまくいかなかった.
またVimはクセありすぎなので、ここはIntelliJを使う.

またソース変更→ビルド→実行という手順を簡単にするのはIDEが一番.

* [急いで学ぶGo lang#2 IntelliJ IDEAでGo開発](http://dev.classmethod.jp/server-side/language/golang-2/)


しかし上記は少し情報が古い.  

* GOROOTの設定は不要.
* GOPATHは任意の場所、具体的にはライブラリのソースや実行ファイルを格納する、自分で決めた場所でOK.

# goファイル以外をバイナリ化してデプロイを簡単にする

これしないとデプロイやってられない.

```goemon``` と ```go-bindata``` でファイルに変更があったらソースを自動生成するようにする.


```
$ go get github.com/mattn/goemon/cmd/goemon  
$ go get github.com/jteeuwen/go-bindata/...  
```

下記の本家サイトも参照のこと.

* [goemon](https://github.com/mattn/goemon)  
* [go-bindata](https://github.com/jteeuwen/go-bindata)  

GoSampleフォルダで次のコマンドを実行すれば、html/css/jsを保存する度にassets/bindata.goが更新される.

```
goemon -c goemon.yml
```

# Reactのための環境

Reactは、React言語とでも言うべき文法で書いたファイルをJavaScriptに変換する必要がある.
これを行ってくれるのが```jsx```というツール.  
先に書いた```goemon```で変換処理は自動で行うように設定してあるので、各開発者は```jsx```をインストールするだけでよい.  

ただ、```jsx```は```node.js```で作られているので、まずは```node.js```をインストールする必要がある.
```node.js```のインストーラは[ここ](http://nodejs.jp/nodejs.org_ja/).

```node.js```がインストール出来たら、次のコマンドで```jsx```をインストールする.  

```
npm install -g react-tools
```

```npm```というのは```node.js```のパッケージ管理ツール.  

ここまでくれば```goemon```が自動でjsx変換を行ってくれる.  
HTMLでは変換後のJavaScriptファイルを参照するようにすればOK.  

なお変換前のファイルは```assets/jsx```に格納すること.  

# 課題

[GitHubにて管理](https://github.com/jabaraster/GoSample/issues)
