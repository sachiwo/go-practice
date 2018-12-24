## ツール
* gore
* ghq,peco

.bashrc:
```
function peco-ghq {
    local dir="$( ghq list --full-path | peco )"
    if [ ! -z "$dir" ] ; then
        cd "$dir"
    fi
}
```

* glide  
パッケージ管理ツール。pip,gem,npmみたいなもの。  
glide init  
glide update  
glide novendor  
