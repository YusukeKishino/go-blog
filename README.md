# go-blog

## Environment
go `1.15.2`

node `14.5.0`

## 
| 用途 | 使用しているもの |
|----|----|
| Webサーバー | [gin-gonic/gin](https://github.com/gin-gonic/gin) |
| ORM | [go-gorm/gorm](https://github.com/go-gorm/gorm) |
| DIコンテナ | [uber-go/dig](https://github.com/uber-go/dig) |
| Markdown->HTML | [gomarkdown/markdown](https://github.com/gomarkdown/markdown) |
| CSSフレームワーク | [bluma](https://bulma.io/) |
| JS CSSのバンドル | [webpack](https://webpack.js.org/) |

## Setup

```sh
$ mysql -u root
mysql> create database `go-blog_development`;
mysql> create database `go-blog_test`;
```
