Go言語開発環境作成メモ

# IntelliJ IDE

せっかく静的言語だから入力補完を有効にしたい.
Vimでも補完出来るように設定出来るらしいが、うまくいかなかった.
またVimはクセありすぎなので、ここはIntelliJを使う.

またソース変更→ビルド→実行という手順を簡単にするのはIDEが一番.

* http://qiita.com/kaiinui/items/433eb86c022ffcad0bea#3-6
* http://dev.classmethod.jp/server-side/language/golang-2/

# goファイル以外をバイナリ化してデプロイを簡単にする

これしないとデプロイやってられない.

* http://blog.satotaichi.info/one-binary-using-go-bindata/

GoSampleフォルダで次のコマンドを打つ.

```
go-bindata -pkg=static -o=static/assets.go html/...
```

# 以下、課題

## 必須
* JSONを返す
* CSS/JavaScriptを独立したファイルに書き、HTMLから参照する
* Reactの開発環境整備

## 出来れば
* ソース変更時にIntelliJの実行を停止→再度実行としなくてもいいようにしたい
