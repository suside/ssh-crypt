# ssh-crypt 🔒 [![Build Status](https://travis-ci.org/suside/ssh-crypt.svg?branch=master)](https://travis-ci.org/suside/ssh-crypt)&nbsp;[![Coverage Status](https://coveralls.io/repos/github/suside/ssh-crypt/badge.svg?branch=master)](https://coveralls.io/github/suside/ssh-crypt?branch=master)&nbsp;[![Go Report Card](https://goreportcard.com/badge/github.com/suside/ssh-crypt)](https://goreportcard.com/report/github.com/suside/ssh-crypt)&nbsp;[![Say thanks!](https://img.shields.io/badge/SayThanks.io-%F0%9F%91%8D-1EAEDB.svg)](https://saythanks.io/to/suside)

Share AES-256 encrypted vault file with your teammates using only ssh `authorized_keys`!

## Usage
```
$ echo "secret :)" | ssh-crypt edit --stdin -a ~/.ssh/authorized_keys VAULT.txt
$ cat VAULT.txt
dRAALGdpdGh1Yi5jb20vc3Vza...
$ ssh-crypt view VAULT.txt
secret :)
```

## Install

Download binary release https://github.com/suside/ssh-crypt/releases/latest
or install with `go` from master branch:
```
go get github.com/suside/ssh-crypt
```

## Why
  * Sharing Keepass with one master password is a no go...
  * Not everyone have/want pgp keys...
  * ...

## Inspiration
This is cheeky rewrite of great [ssh-vault](https://github.com/ssh-vault/ssh-vault) with less features **but** with support of multiple key pairs.
