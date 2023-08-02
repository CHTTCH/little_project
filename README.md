## 環境設置
以下為本專案所使用的環境：

- [Docker](https://www.docker.com/)
- [Go](https://golang.org/) 版本 1.19.3 或以上
- [Node.js](https://nodejs.org/en/) 版本 v20.5.0 或以上
- [npm](https://www.npmjs.com/)，它通常會與 Node.js 一同安裝

## 後端設置

後端使用 Go 1.19.3 和 PostgreSQL 15.3（資料庫使用 Docker container）。

### 運行 PostgreSQL

首先，需要啟動一個 PostgreSQL 15.3 的 Docker container。請在下面的指令中將`<user>`、`<password>` 替換為您想要設置的 Postgres 密碼，然後在終端機中運行此指令：

```bash
docker run --name pgsql -e POSTGRES_USER=<user> -e POSTGRES_PASSWORD=<password> -e POSTGRES_DB=little_project -d -p 5432:5432 postgres
```

其中，預設的`<user>`、`<password>`皆為Zachary，若想要自行定義的話，除了在上述的docker指令更改內容外，還須至`{path_to_little_project}/backend/main.go`第16行修改user與password

上述的指令會在本地的 5432 port 啟動一個 PostgreSQL 15.3 的 Docker container

### 運行後端伺服器

首先進入後端所在的資料夾：

```bash
cd /path_to_little_project/backend
```

接著，執行下面的指令來啟動伺服器，並會在`localhost:8888`運作：

```bash
go run main.go
```

此外，在執行的同時也會向資料庫預先寫入五位病患的資訊

## 前端設置

前端使用 React 18.2.0 和 Node.js v20.5.0。

首先入前端所在的資料夾：

```bash
cd /path_to_little_project/frontend
```

接著，執行下面的指令來安裝所有相關的套件：

```bash
npm install
```

當所有相關的套件都已經安裝完畢後，執行下面的指令來啟動前端伺服器：

```bash
npm start
```

現在可以在瀏覽器中透過 localhost:3000 來對網頁進行新增、編輯醫囑的操作。