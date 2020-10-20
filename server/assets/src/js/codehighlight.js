import 'highlight.js/styles/github.css'
import hljs from 'highlight.js/lib/core'

hljs.registerLanguage('go', require('highlight.js/lib/languages/go'))
hljs.registerLanguage('markdown', require('highlight.js/lib/languages/markdown'))
hljs.registerLanguage('javascript', require('highlight.js/lib/languages/javascript'))
hljs.registerLanguage('shell', require('highlight.js/lib/languages/shell'))
hljs.registerLanguage('css', require('highlight.js/lib/languages/css'))
hljs.registerLanguage('diff', require('highlight.js/lib/languages/diff'))
hljs.registerLanguage('dockerfile', require('highlight.js/lib/languages/dockerfile'))
hljs.registerLanguage('json', require('highlight.js/lib/languages/json'))
hljs.registerLanguage('makefile', require('highlight.js/lib/languages/makefile'))
hljs.registerLanguage('scss', require('highlight.js/lib/languages/scss'))
hljs.registerLanguage('sql', require('highlight.js/lib/languages/sql'))
hljs.registerLanguage('yaml', require('highlight.js/lib/languages/yaml'))

hljs.initHighlightingOnLoad()
