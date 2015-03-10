Go言語開発環境作成メモ

# Go言語環境導入

http://golang.jp/install


# IntelliJ IDE

せっかく静的言語だから入力補完を有効にしたい.
Vimでも補完出来るように設定出来るらしいが、うまくいかなかった.
またVimはクセありすぎなので、ここはIntelliJを使う.

またソース変更→ビルド→実行という手順を簡単にするのはIDEが一番.

* http://qiita.com/kaiinui/items/433eb86c022ffcad0bea#3-6
* http://dev.classmethod.jp/server-side/language/golang-2/

# goファイル以外をバイナリ化してデプロイを簡単にする

これしないとデプロイやってられない.

```goemon```と```go-bindata```でファイルに変更があったらソースを自動生成するようにする.

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
