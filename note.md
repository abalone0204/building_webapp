`http.ListenAndServe(":"+port, nil)`

為什麼放 nil 在第二個參數？

* Handle and HandleFunc add handlers to DefaultServeMux

# Deployment

- 先設定 `Procfile`

- 不需要再特別指定 build-pack，heroku 現在已經幫我們處理好了

- 執行一次 `godep -r save ./`

    - 會生成 'Godeps/godep.json'和其他哩哩扣扣

    - heroku 會根據這個檔案來安裝需要的 dependencies

- `git push heroku master`

