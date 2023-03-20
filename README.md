# xk6-dashboard-ui-tmpl

Example web UI for [xk6-dashboard](https://github.com/szkiba/xk6-dashboard).

This repository itself is not a [xk6](https://github.com/grafana/xk6) extension but an addon to [xk6-dashboard](https://github.com/szkiba/xk6-dashboard) extension. You can use it together with [xk6-dashboard](https://github.com/szkiba/xk6-dashboard) to replace default web UI. 

> **Create you own UI**
>
> 1. [Create repository from this template](https://github.com/szkiba/xk6-dashboard-ui-tmpl/generate)
> 2. In [go.mod](go.mod) file replace `github.com/szkiba/xk6-dashboard-ui-tmpl` with your repository path
> 3. Edit or generate files under [ui](ui) folder
>

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Download `xk6`:
  ```bash
  $ go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```bash
  $ xk6 build --with github.com/szkiba/xk6-dashboard --with github.com/szkiba/xk6-dashboard-ui-tmpl
  ```
  > Replace `github.com/szkiba/xk6-dashboard-ui-tmpl` with your repository path.

## Usage

Without parameters the dashboard will accessible on port `5665` with any web browser: http://127.0.0.1:5665

```plain
$ ./k6 run --out dashboard script.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: script.js
     output: dashboard (:5665) http://localhost:5665
```

For detailed usage information see [xk6-dashboard](https://github.com/szkiba/xk6-dashboard).
