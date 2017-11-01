从金山词霸抓单词解释
===

这是一个简单的Go语言的程序，可以用来从[金山词霸](http://www.iciba.com) 抓取单词的中文解释。
我通过写这个简单的程序来学习go语言。

用法：

```
$ go install github.com/trentzhou/gociba
$ go install github.com/trentzhou/gociba/gociba_cli
```

然后可以查单词：

```
$ $GOPATH/bin/gociba_cli one two three four five
```

