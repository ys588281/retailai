# retailai

1. Before runnung the project, please run go get "github.com/gorilla/mux" and go get "github.com/mattn/go-sqlite3".
2. Create your own sqlite3 database in local and put the path in db connection.
3. sqlite3 is optional. It depends on which db you are using.
4. I used sqlite3 in my local to build environment because it is light, simple and easy.
5. I used text datatype for created_at and updated_at because of sqlite3. In real environment, it should be datetime.
