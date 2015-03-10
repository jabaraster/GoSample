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

```goemon```と```go-bindata```でファイルに変更があったらソースを自動生成するようにする.


```
$ go get github.com/mattn/goemon/cmd/goemon
$ go get github.com/jteeuwen/go-bindata/...
```

下記も参照のこと.

[goemon](https://github.com/mattn/goemon)  
[go-bindata](https://github.com/jteeuwen/go-bindata)  

GoSampleフォルダで次のコマンドを実行すれば、html/css/jsを保存する度にassets/bindata.goが更新される.

```
goemon -c goemon.yml
```

# 以下、課題

## 必須
* JSONを返す
* CSS/JavaScriptを独立したファイルに書き、HTMLから参照する
* Reactの開発環境整備

## 出来れば
* ソース変更時にIntelliJの実行を停止→再度実行としなくてもいいようにしたい
* bindata.goがソース管理上ボトルネックになる可能性あり
