import os

systems = ['windows', 'linux', 'darwin']
archs = ['amd64', '386']

for s in systems:
    for a in archs:
        os.system(f'set GOOS={s}& set GOARCH={a}& go build -ldflags="-w -s" -o proof_{s}_{a}')
