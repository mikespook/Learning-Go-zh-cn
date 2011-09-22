#!/bin/bash
ROOT=`pwd`

intro() {
    echo "This script will setup your vim for golang."
    echo "It could setup and upgrade vim's configuration automatically."
    echo "First of all, you should follow http://golang.org/doc/install.html for seting all of the enverament variables."
    echo 
    echo "To report an issue, please feel free to contact me."
    echo 
    echo "Xing Xing<mikespook@gmail.com>"
    echo "http://mikespook.com"
}

split(){
    echo "************************"
}

bye() {
    echo "Bye!"
    exit
}

test_envs() {
    VIM=`type -P vim`
    if [ $? -eq 1 ];
    then
        echo "Please install the vim package."
        bye
    else
        echo "$VIM [OK]"
    fi
    HG=`type -P hg`
    if [ $? -eq 1 ];
    then
        echo "Please install the mercural(hg) package."
        bye
    else
        echo "$HG [OK]"
    fi
}

goroot() {
    if [[ -n $GOROOT ]];
    then
        echo "\$GOROOT: $GOROOT [OK]"
    else
        echo "\$GOROOT: NULL [FAILD]"
        echo "Download go source and set \$GOROOT to $ROOT/go temporary."
        read -p "<y/n>:" YN
        if [ $YN == "y" ];
        then
            GOROOT=$ROOT/go
            hg clone https://go.googlecode.com/hg/ $GOROOT
            echo "\$GOROOT: $GOROOT [OK]"
        else
            bye
        fi
    fi
    echo "Update the go source to the newest version."
    read -p "<y/n>:" YN
    if [ $YN == "y" ];
    then
        hg pull -R $GOROOT
    fi
}

install() {
    mkdir -p $HOME/.vim/autoload/go
    mkdir -p $HOME/.vim/ftplugin/go
    mkdir -p $HOME/.vim/ftdetect
    mkdir -p $HOME/.vim/syntax
    mkdir -p $HOME/.vim/indent
    mkdir -p $HOME/.vim/plugin
    mkdir -p $HOME/.vim/syntax

    echo "Copy the scripts into $HOME/.vim, otherwise link them."
    read -p "<y/n>:" YN
    if [ $YN == "y" ]; 
    then
        cp -u $GOROOT/misc/vim/autoload/go/complete.vim $HOME/.vim/autoload/go/
        cp -u $GOROOT/misc/vim/ftplugin/go/*.vim $HOME/.vim/ftplugin/go/
        cp -u $GOROOT/misc/vim/ftdetect/gofiletype.vim $HOME/.vim/ftdetect/
        cp -u $GOROOT/misc/vim/syntax/go.vim $HOME/.vim/syntax/
        cp -u $GOROOT/misc/vim/indent/go.vim $HOME/.vim/indent/
        cp -u $GOROOT/misc/vim/plugin/godoc.vim $HOME/.vim/plugin/godoc.vim
        cp -u $GOROOT/misc/vim/syntax/godoc.vim $HOME/.vim/syntax/godoc.vim
    else
        ln -s $GOROOT/misc/vim/ftdetect/gofiletype.vim $HOME/.vim/ftdetect/
        ln -s $GOROOT/misc/vim/syntax/go.vim $HOME/.vim/syntax/
        ln -s $GOROOT/misc/vim/autoload/go/complete.vim $HOME/.vim/autoload/go/
        ln -s $GOROOT/misc/vim/ftplugin/go/*.vim $HOME/.vim/ftplugin/go/
        ln -s $GOROOT/misc/vim/indent/go.vim $HOME/.vim/indent/
        ln -s $GOROOT/misc/vim/plugin/godoc.vim $HOME/.vim/plugin/godoc.vim
        ln -s $GOROOT/misc/vim/syntax/godoc.vim $HOME/.vim/syntax/godoc.vim
        ln -s $GOROOT/misc/vim/ftplugin/go/godoc.vim $HOME/.vim/ftplugin/go/godoc.vim 
    fi
}

gocode() {
    echo "Use the gocode."
    read -p "<y/n>:" YN
    if [ $YN != "y" ]; 
    then
        return
    fi

    GIT=`type -P git`
    if [ $? -eq 1 ];
    then
        echo "Please install the git package."
        return
    else
        echo "$GIT [OK]"
    fi

    if [ -e $ROOT/gocode ];
    then
        cd $ROOT/gocode
        git pull
        cd $ROOT
    else
        git clone https://github.com/nsf/gocode.git $ROOT/gocode
    fi
    cp -u $ROOT/gocode/vim/autoload/gocomplete.vim $HOME/.vim/autoload/
    cp -u $ROOT/gocode/vim/ftplugin/go.vim $HOME/.vim/go.vim

    echo "Make the gocode's executable file."
    read -p "<y/n>:" YN
    if [ $YN == "y" ];
    then
        cd $ROOT/gocode
        make clean&&make install
        cd $ROOT
    fi 
}

finish() {
    if [ ! -e $GOBIN/gocode ];
    then
        if [ ! -e $ROOT/gocode ];
        then
            echo "Please clone the gocode's sources."
        fi
        if [[ ! -n $GOBIN ]];
        then
            echo "Please set the \$GOBIN."
        fi
        echo "Please run \"make install\" to make the gocode's executable file."        
    fi

    echo "Finished!"
    echo "Please setup the $HOME/.vimrc manually."
    echo "Have a good time!!! ;-)"
}

intro
split
test_envs
goroot
install
gocode
split
finish
bye
