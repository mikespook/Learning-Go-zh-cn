#!/bin/bash
ROOT=`pwd`

intro() {
    echo "This script will setup your vim environment for golang."
    echo "It could setup and upgrade vim's configuration automatically."
    echo "First of all, you'd better follow http://golang.org/doc/install.html for seting all of the environment variables."
    echo "Put this script into $HOME/src/ is good choice."
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
        echo "Input \$GOROOT manully."
        read -p "<y/n>:" YN
        if [ $YN == "y" ];
        then
            while [[ ! -n $GOROOT ]]
            do
                read -p "\$GOROOT:" GOROOT
                if [ ! -d $GOROOT ];
                then
                    echo "Not a valid directory."
                    unset GOROOT
                fi
            done
        else
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
    rm $HOME/.vim/autoload/go/*
    rm $HOME/.vim/ftplugin/go/*
    rm $HOME/.vim/ftdetect/gofiletype.vim
    rm $HOME/.vim/syntax/go.vim
    rm $HOME/.vim/indent/go.vim
    rm $HOME/.vim/plugin/godoc.vim
    rm $HOME/.vim/syntax/godoc.vim
    if [ $YN == "y" ]; 
    then
        cp -u -f $GOROOT/misc/vim/autoload/go/complete.vim $HOME/.vim/autoload/go/
        cp -u -f $GOROOT/misc/vim/ftplugin/go/*.vim $HOME/.vim/ftplugin/go/
        cp -u -f $GOROOT/misc/vim/ftdetect/gofiletype.vim $HOME/.vim/ftdetect/
        cp -u -f $GOROOT/misc/vim/syntax/go.vim $HOME/.vim/syntax/
        cp -u -f $GOROOT/misc/vim/indent/go.vim $HOME/.vim/indent/
        cp -u -f $GOROOT/misc/vim/plugin/godoc.vim $HOME/.vim/plugin/godoc.vim
        cp -u -f $GOROOT/misc/vim/syntax/godoc.vim $HOME/.vim/syntax/godoc.vim
        echo "Copied successful."
    else
        ln -s -f $GOROOT/misc/vim/ftdetect/gofiletype.vim $HOME/.vim/ftdetect/
        ln -s -f $GOROOT/misc/vim/syntax/go.vim $HOME/.vim/syntax/
        ln -s -f $GOROOT/misc/vim/autoload/go/complete.vim $HOME/.vim/autoload/go/
        ln -s -f $GOROOT/misc/vim/ftplugin/go/*.vim $HOME/.vim/ftplugin/go/
        ln -s -f $GOROOT/misc/vim/indent/go.vim $HOME/.vim/indent/
        ln -s -f $GOROOT/misc/vim/plugin/godoc.vim $HOME/.vim/plugin/godoc.vim
        ln -s -f $GOROOT/misc/vim/syntax/godoc.vim $HOME/.vim/syntax/godoc.vim
        ln -s -f $GOROOT/misc/vim/ftplugin/go/godoc.vim $HOME/.vim/ftplugin/go/godoc.vim 
        echo "Linked successful."
    fi
}

gocode() {
    echo "Use the gocode."
    read -p "<y/n>:" YN
    if [ $YN != "y" ]; 
    then
        return
    fi

    if [[ -n $GOCODE ]];
    then
        echo "gocode: $GOCODE [OK]"
    else
        echo "Input gocode location."
        read -p "<y/n>:" YN
        if [ $YN == "y" ];
        then
            while [[ ! -n $GOCODE ]]
            do
                read -p "gocode:" GOCODE
                if [ ! -d $GOCODE ];
                then
                    echo "Not a valid directory."
                    unset GOCODE
                fi
            done
        else
            echo "Download gocode sources."
            read -p "<y/n>:" YN
            if [ $YN == "y" ];
            then
                GIT=`type -P git`
                if [ $? -eq 1 ];
                then
                    echo "Please install the git package."
                    return
                else
                    echo "$GIT [OK]"
                fi
                GOCODE=$ROOT/gocode
                git clone https://github.com/nsf/gocode.git $GOCODE
            fi
        fi
    fi

    echo "Update the gocode source to the newest version."
    read -p "<y/n>:" YN
    if [ $YN == "y" ];
    then
        cd $GOCODE
        git pull
        cd $ROOT
    fi

    cp -u $GOCODE/vim/autoload/gocomplete.vim $HOME/.vim/autoload/
    cp -u $GOCODE/vim/ftplugin/go.vim $HOME/.vim/go.vim

    echo "Make the gocode's executable file."
    read -p "<y/n>:" YN
    if [ $YN == "y" ];
    then
        cd $GOCODE
        make clean&&make install
        cd $ROOT
    fi 
}

finish() {
    if [ ! -e $GOBIN/gocode ];
    then
        echo "You do not have the gocode executable file."
        if [[ ! -n $GOBIN ]];
        then
            echo "Please set the \$GOBIN and run \"make install\" in gocode source code directory."
        else
            echo "Please run \"make install\" in gocode source code directory."
        fi
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
