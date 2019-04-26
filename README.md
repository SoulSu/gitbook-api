# gitbook-api
the api is for gitbook ; 

### API LIST

|name|url|
|----|----|
|get author info|[https://api.gitbook.com/author/{user_name}](#https://api.gitbook.com/author/{user_name})|
|get book info|[https://api.gitbook.com/book/{user_name}/{book_name}](https://api.gitbook.com/book/wizardforcel/vbird-linux-basic-4e)|



#### API example

##### https://api.gitbook.com/author/wizardforcel

```json
{
  "id": "547153d82e958e0200381334",
  "hasMigrated": false,
  "type": "User",
  "username": "wizardforcel",
  "name": "wizardforcel",
  "location": "",
  "website": "https://github.com/wizardforcel",
  "bio": "",
  "verified": true,
  "locked": false,
  "site_admin": false,
  "urls": {
    "profile": "https://legacy.gitbook.com/@wizardforcel",
    "stars": "https://legacy.gitbook.com/@wizardforcel/starred",
    "avatar": "https://legacy.gitbook.com/@wizardforcel/avatar"
  },
  "permissions": {
    "edit": null,
    "admin": null
  },
  "dates": {
    "created": "2014-11-23T03:26:16.278Z"
  },
  "counts": {},
  "github": {
    "username": "wizardforcel"
  }
}
```


##### https://api.gitbook.com/book/wizardforcel/vbird-linux-basic-4e
```json
{
  "id": "wizardforcel/vbird-linux-basic-4e",
  "status": "published",
  "name": "vbird-linux-basic-4e",
  "title": "鸟哥的 Linux 私房菜：基础学习篇 第四版",
  "description": "",
  "public": true,
  "template": "base",
  "topics": [
    "linux",
    "vbird"
  ],
  "license": "nolicense",
  "language": "zh",
  "locked": false,
  "cover": {
    "large": "/cover/book/wizardforcel/vbird-linux-basic-4e.jpg?build=1463123465598",
    "small": "/cover/book/wizardforcel/vbird-linux-basic-4e.jpg?build=1463123465598&size=small"
  },
  "urls": {
    "git": "https://git.gitbook.com/wizardforcel/vbird-linux-basic-4e.git",
    "access": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e",
    "homepage": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/",
    "read": "https://legacy.gitbook.com/read/book/wizardforcel/vbird-linux-basic-4e",
    "edit": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/edit",
    "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/",
    "download": {
      "epub": "https://legacy.gitbook.com/download/epub/book/wizardforcel/vbird-linux-basic-4e",
      "mobi": "https://legacy.gitbook.com/download/mobi/book/wizardforcel/vbird-linux-basic-4e",
      "pdf": "https://legacy.gitbook.com/download/pdf/book/wizardforcel/vbird-linux-basic-4e"
    }
  },
  "counts": {
    "stars": 1864,
    "subscriptions": 248,
    "updates": 3,
    "discussions": 39,
    "changeRequests": 0
  },
  "dates": {
    "build": "2016-05-13T07:11:05.598Z",
    "created": "2016-05-13T06:19:00.829Z"
  },
  "permissions": {
    "edit": false,
    "admin": false,
    "important": false,
    "read": true,
    "write": false,
    "manage": false
  },
  "publish": {
    "ebooks": true,
    "defaultBranch": null,
    "builder": null
  },
  "author": {
    "id": "547153d82e958e0200381334",
    "hasMigrated": false,
    "type": "User",
    "username": "wizardforcel",
    "name": "wizardforcel",
    "location": "",
    "website": "https://github.com/wizardforcel",
    "bio": "",
    "verified": true,
    "locked": false,
    "site_admin": false,
    "urls": {
      "profile": "https://legacy.gitbook.com/@wizardforcel",
      "stars": "https://legacy.gitbook.com/@wizardforcel/starred",
      "avatar": "https://legacy.gitbook.com/@wizardforcel/avatar"
    },
    "permissions": {
      "edit": null,
      "admin": null
    },
    "dates": {
      "created": "2014-11-23T03:26:16.278Z"
    },
    "counts": {},
    "github": {
      "username": "wizardforcel"
    }
  }
}
```