Go言語開発環境作成メモ

# 開発環境構築

次のツールを導入します(上から順番に導入する).

1. Git: Goのライブラリの導入、及びアプリケーションコードのバージョン管理に必要
1. Mercurial: Goのライブラリの導入に必要
1. Go言語: 当たり前
1. go-bindata: 静的ファイルを実行ファイルに含めるために必要
1. goemon: Go以外のファイルの変更をリアルタイムにブラウザに反映するために必要
1. node.js: react-toolの実行環境として必要
1. react-tools: Reactを使いやすくするために必要

IDEはIntelliJを推奨.


順に導入手順を記載

## Gitの導入

## Mercurialの導入
[インストーラ](http://mercurial.selenic.com)


## Goの導入
[Go言語環境導入](http://golang.jp/install)

## go-bindataの導入

GitとGoが導入されている環境で、以下コマンドを実行.  

```
$ go get github.com/jteeuwen/go-bindata/...  
```

これで ``` go-bindata ``` というコマンドがインストールされる.

## goemonの導入

GitとGoが導入されている環境で、以下コマンドを実行.  

```
$ go get github.com/mattn/goemon/cmd/goemon  
```

これで ``` goemon ``` というコマンドがインストールされる.  


## node.jsの導入

[インストーラ](http://nodejs.jp/nodejs.org_ja/)


## react-toolsの導入

node.jsがインストールされている環境で以下のコマンドを実行.  

```
$ npm install -g react-tools
```

これで ``` jsx ``` というコマンドがインストールされる.  


## IntelliJの導入
せっかく静的言語だから入力補完を有効にしたい.
Vimでも補完出来るように設定出来るらしいが、うまくいかなかった.
またVimはクセありすぎなので、ここはIntelliJを使う.

またソース変更→ビルド→実行という手順を簡単にするのはIDEが一番.

* [急いで学ぶGo lang#2 IntelliJ IDEAでGo開発](http://dev.classmethod.jp/server-side/language/golang-2/)


しかし上記は少し情報が古い.  

* GOROOTの設定は不要.
* GOPATHは任意の場所、具体的にはライブラリのソースや実行ファイルを格納する、自分で決めた場所でOK.

# 開発

## goemonによるファイル監視の開始

下記コマンドでHTML/CSS/JSの監視を開始すること.  

```
$ goemon -c goemon.yml
```

なお ``` goemon.yml ``` は河野が作ってGitリポジトリ管理している.  

## Webサーバの起動
アプリを動かすには ``` goji_main.go ``` を実行する.
デフォルトでは ``` 8000 ``` ポートでリッスンするので、以下のURLでWebアプリにアクセス出来る.

```
http://127.0.0.1:8000
```

他のプロセスによって ``` 8000 ``` 番がふさがっているときは要注意.
特にエラーが出ず起動してしまうのだが、いざブラウザにアクセスすると ``` Cannot GET <path> ``` と虚しく表示されるだけ.
非常に分かりにくい.
この場合は ``` 8000 ``` 番をふさいでいるプロセスを殺すか、あるいは起動時に以下のようにオプションでポートを指定して回避する.

```
go run goji_main.go -bind=:8080
```

ポート番号の前に ``` : ``` が必要なので注意すること.

## ファイル配置ルール

### HTMLファイル

```
assets/html/
```

### CSSファイル

```
assets/css/
```

### JavaScriptファイル

```
assets/jsx/
```

React言語だろうが普通のJSファイルだろうが、ここに置くこと！  

### goファイル

srcディレクトリの下にパッケージ名のディレクトリを切って、そこに関連ファイルを格納.

※開発が進んだらもう少しルールを細かく設定するつもり.  

### 編集しちゃいけないファイル
要は自動生成されるファイル.  


```
assets/js/ 以下の全てのファイル (assets.goは例外. 触ることがあるかも)
src/assets/ 以下
```

``` bindata.go ``` と ``` bindata_debug.go ``` はgo-bindataが作るファイル.  
``` assets/js/ ``` 以下はjsxが作るファイル.  


# 課題

[GitHubにて管理](https://github.com/jabaraster/GoSample/issues)

# デプロイ


## ビルド

### ビルドの事前準備

[ここ](http://qiita.com/Jxck_/items/02185f51162e92759ebe#2-2)を参考に事前準備をしておく.


### ビルド実施(Linuxへのデプロイ前提)


```
$ GOOS=linux GOARCH=amd64 go build -o GoSample src/main.go
```

### サーバプロセス起動

```
$ main -mode=production -webPort=80
```

実行ファイルをサーバに持って行くのが、ちょっとたいへんw  
S3なんかを使ってやり取りして下さい.  


# Appendix


## goemonでやっていること
### Reactのjsx変換
### HTML/CSS/JSのバイナリデータ化

## HTML/CSS/JS読み込みの挙動の切り替え

開発環境では変更をリアルタイムに反映するようにし、デプロイ環境では実行ファイル内のバイナリデータを扱い高速に動作するようにしている.  
環境の切り替えは、起動時に ``` mode ``` 引数で行うようにしており、``` production ``` が指定された場合はデプロイ環境とみなしている.  

