# go-todo
## 概要
> 学習用途
golangを用いて、クリーンアーキテクチャーでREST APIを実装。

[evrone/go-clean-template](https://github.com/evrone/go-clean-template/tree/master)を参考に作成しています。
## 使用したライブラリ
- chi: rest api
- ent: ORM


## Archtecture
```mermaid
flowchart TB
    entity["entity"]
    repository["repository"]
    usecase["usecase"]
    interface["usecase\ninterface"]
    controller["controller"]
    app["app\n(rest with chi)"]
    mysql["mysql"]
    entity["entity"]

    mysql-->app
    app-->controller
    app-->repository
    repository-->interface
    controller-->interface
    interface-->usecase
    usecase-->entity
```

