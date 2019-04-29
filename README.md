# The dotGo `go3dprint` demo

These are the files used in "the dilators talk" that took place at [dotGo](https://www.dotgo.eu/) 2019, in Paris.

You can watch the presentation here:

<p align="center">
  <a href="https://www.youtube.com/watch?v=ZACOc-NwV0c" target="_new"><img src="https://img.youtube.com/vi/ZACOc-NwV0c/0.jpg" width="40%"></a>
</p>

## The full version

For the live demo (the code on this repo) a simpler version was used. The main difference is that the [full version](https://github.com/garden-io/go3dprint/) is "cloud-native." 

(The differences and the conversion process are explained in detail in the article [_Needlessly Distributed Phallic Object Generator_](https://medium.com/garden-io/the-needlessly-distributed-phallic-object-generator-2da47672be6f).)

While the simple version compiles into a single binary and runs locally on my OS, the full version uses [Garden](https://garden.io/) and has a more 2019 look to it:

- Functionality is split into loosely coupled microservices
- Running as lightweight containers
- That communicate via API calls.

And the configured workflow:

- Re-builds and re-deploys on every code change
- Can use hot reload, so containers can be update without restarting
- Uses the same tooling for all environments—local, CI, remote.

It also has a really cool frontend:

<p align="center">
  <img src="https://github.com/garden-io/go3dprint/blob/master/img/frontend.png?raw=true" width="70%">
</p>

The full version can be found [here](https://github.com/garden-io/go3dprint/).

## Usage

This demo is made to run on Linux (it might work elsewhere, but I haven't tested it). 

Besides Go, it requires `entr` for file monitoring, and `inkscape` to convert SVG files to PNG for visualization. (Inkscape is not necessary in the [full version](https://github.com/garden-io/go3dprint/).) 

You can install both with `sudo apt install entr inkscape` on Ubuntu.

To then run it with live reload capabilities, do:

```bash
$ ls *.go | entr -r ./reload.sh
```

- Most of the demo took place in `mesh.go`.
- Just find the big "FUN" banner and start un-commenting and saving.
- The results will output to `mesh.png` and `vector.png`.
- You can switch renderers between [fauxgl](https://github.com/fogleman/fauxgl/) and [path tracer](https://github.com/fogleman/pt) by commenting/uncommenting the relevant bits on `render.go`.
