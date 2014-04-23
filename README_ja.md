# Git::Gitlab

[Gitlab](https://www.gitlab.com/)のためのコマンドラインツールです

## インストール方法

Ruby関連のツールをインストールして、次のコマンドを実行します

    $ gem install git-gitlab

## 使い方

### 初期設定

次のようにGitの設定を行います

	git config --global gitlab.url http://gitlab.example.com/
	git config --glolab gitlab.token GITLAB_SECRET_TOKEN
	git config gitlab.project NAMESPACE/PROJECT

もしもGitlabのリモートリポジトリがoriginでない場合には

	git config gitlab.remote REMOTE

### 設定の確認を行います

	git gitlab

### マージリクエストを作ります

	git gitlab merge SOURCE_BRANCH TARGET_BRANCH --assign USER_NAME

### OpenなMergerequestを全て取得します

	git gitlab merge --list

### IDでイシューを検索し、表示します

	git gitlab issue ISSUE_ID

### Mergeリクエストのレビューを行います

	git gitlab review MERGEREQUEST_ID

## 貢献するには

1. フォークします
2. あなたの機能ブランチを作ります (`git checkout -b my-new-feature`)
3. 変更をコミットします (`git commit -am 'Add some feature'`)
4. リモートブランチにプッシュします (`git push origin my-new-feature`)
5. 新しいプルリクエストを作ります
