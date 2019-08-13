# CTF-Hash-Proof

## Usage

`proof algorithm target [prefix] [suffix]`

### examples:

`substr(md5($verify), 0, 5) = 'fffff'`, input the verify code plz

```
$ proof md5 fffff
aaVyG
```

`substr(sha1('hello'.$verify.'world'), 0, 5) = 'fffff'`

```
$ proof sha1 fffff hello world
OadkS
```

***

## Install

download binaries from [release page](https://github.com/EddieIvan01/ctf-hash-proof/releases)

or

compile from source code

```
git clone https://github.com/EddieIvan01/ctf-hash-proof.git
cd ctf-hash-proof
go install -ldflags="-w -s"
```

***

## Tips

the default workers number is (CPU Cores number) * 0.75

if it was too high or too low for you, you could change it by modify source code then recompile 