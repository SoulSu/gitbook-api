# gitbook-api

the public api is for gitbook

### How to use?



### API LIST

> API is form [https://github.com/GitbookIO/api-guide](https://github.com/GitbookIO/api-guide)

|name|url|
|----|----|
|get author info|[https://api.gitbook.com/author/{user_name}](#https://api.gitbook.com/author/{user_name})|
|get book info|[https://api.gitbook.com/book/{user_name}/{book_name}](https://api.gitbook.com/book/wizardforcel/vbird-linux-basic-4e)|
|get book discussions|[https://api.gitbook.com/book/{user_name}/{book_name}/discussions/](https://api.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/)
|get user books|https://api.gitbook.com/author/{user_name}/books|
|get user book content|https://api.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/contents/|
|get user book section|https://api.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/contents/1.json|


#### API example



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


##### https://api.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/

```json
{
  "list": [
    {
      "type": "Discussion",
      "number": 1,
      "title": "涉及版权问题吗？",
      "user": {
        "id": "575103e5f017c30f0046fd7b",
        "hasMigrated": false,
        "type": "User",
        "username": "hitzhangjie",
        "name": "hitzhangjie",
        "location": "",
        "website": "https://github.com/hitzhangjie",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@hitzhangjie",
          "stars": "https://legacy.gitbook.com/@hitzhangjie/starred",
          "avatar": "https://avatars0.githubusercontent.com/hitzhangjie"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2016-06-03T04:13:25.837Z"
        },
        "counts": {},
        "github": {
          "username": "hitzhangjie"
        }
      },
      "subscribed": null,
      "body": "<p>我在阅读一些计算机书籍的过程中，也整理了很多的文档，文档中的相当多的内容摘自某些书籍，我想问，如果我发布这些文档会引起版权纠纷吗？你写的这个呢？</p>\n",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/1"
      },
      "counts": {
        "participants": 2,
        "comments": 1
      },
      "dates": {
        "created": "2016-06-05T04:38:21.530Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      }
    },
    {
      "type": "Discussion",
      "number": 2,
      "title": "5.3.1节 表格格式有问题",
      "user": {
        "id": "5653396269c7011000e8c9ba",
        "hasMigrated": false,
        "type": "User",
        "username": "coder-chenzhi",
        "name": "ZhiChen",
        "location": "",
        "website": "https://github.com/coder-chenzhi",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@coder-chenzhi",
          "stars": "https://legacy.gitbook.com/@coder-chenzhi/starred",
          "avatar": "https://avatars0.githubusercontent.com/coder-chenzhi"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2015-11-23T16:05:54.484Z"
        },
        "counts": {},
        "github": {
          "username": "coder-chenzhi"
        }
      },
      "subscribed": null,
      "body": "<p>5.3.1 Linux目录配置的依据--FHS，这一小节下的第一张表格，格式出了问题</p>\n",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/2"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2016-11-14T02:46:36.663Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      }
    },
    {
      "type": "Discussion",
      "number": 3,
      "title": "3.6x10^9",
      "user": {
        "id": "57e14a1b6a184b1000377777",
        "hasMigrated": false,
        "type": "User",
        "username": "jinhaiz",
        "name": "Jinhai ZHOU",
        "location": "Brest, France",
        "website": "https://github.com/JinhaiZ",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@jinhaiz",
          "stars": "https://legacy.gitbook.com/@jinhaiz/starred",
          "avatar": "https://avatars0.githubusercontent.com/JinhaiZ"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2016-09-20T14:39:23.468Z"
        },
        "counts": {},
        "github": {
          "username": "JinhaiZ"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/3"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2016-12-20T16:11:42.611Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "0.2 个人电脑架构与相关设备元件",
        "filename": "4.md",
        "section": "我们前面谈到CPU内部含有微指令集，不同的微指令集会导致CPU工作效率的优劣。除了这点之外， CPU性能的比较还有什么呢？那就是CPU的频率了！什么是频率呢？简单的说， 频率就是CPU每秒钟可以进行的工作次数。 所以频率越高表示这颗CPU单位时间内可以作更多的事情。举例来说，Intel的 i7-4790 CPU频率为3.6GHz， 表示这颗CPU在一秒内可以进行3.6x109次工作，每次工作都可以进行少数的指令运行之意。",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/4.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/4.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 4,
      "title": "第一行最后可执行文件多了个“可”字",
      "user": {
        "id": "588df43879c3b21100ba0c45",
        "hasMigrated": false,
        "type": "User",
        "username": "eastmalon",
        "name": "Via Liberty",
        "location": "",
        "website": "https://github.com/eastmalon",
        "bio": "",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@eastmalon",
          "stars": "https://legacy.gitbook.com/@eastmalon/starred",
          "avatar": "https://avatars0.githubusercontent.com/eastmalon"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-01-29T13:55:04.008Z"
        },
        "counts": {},
        "github": {
          "username": "eastmalon"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/4"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-03-23T12:51:27.076Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "1.2 Torvalds的Linux发展",
        "filename": "12.md",
        "section": "撰写程序需要什么呢？首先需要的是能够进行工作的环境，再来则是可以将源代码编译成为可可执行文件的编译器。 好在有GNU计划提供的bash工作环境软件以及gcc编译器等自由软件， 让托瓦兹得以顺利的撰写核心程序。他参考Minix的设计理念与书上的程序码，然后仔细研究出386个人计算机的性能最优化， 然后使用GNU的自由软件将核心程序码与386紧紧的结合在一起，最终写出他所需要的核心程序。 而这个小玩意竟然真的可以在386上面顺利的跑起来～还可以读取Minix的文件系统。 真是太好了！不过还不够，他希望这个程序可以获得大家的一些修改建议， 于是他便将这个核心放置在网络上提供大家下载，同时在BBS上面贴了一则消息：",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/12.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/12.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 5,
      "title": "一针见血，学习的实质是实践！",
      "user": {
        "id": "58d6bcee51cd2500111d11ca",
        "hasMigrated": false,
        "type": "User",
        "username": "huangdesheng",
        "name": "黄得胜",
        "website": "",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@huangdesheng",
          "stars": "https://legacy.gitbook.com/@huangdesheng/starred",
          "avatar": "https://s.gravatar.com/avatar/09cf2dfc47d99f97f608aeb451e55a05?s=220&d=https%3A%2F%2Flegacy.gitbook.com%2Fassets%2Fimages%2Favatars%2Fuser.png"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-03-25T18:54:38.634Z"
        },
        "counts": {},
        "google": {}
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/5"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-03-25T18:58:31.991Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "1.4 Linux 该如何学习",
        "filename": "14.md",
        "section": "Tips 鸟哥上课时，常常有学生问到：“老师，到底要听过你的课几次之后，才能学的会？”鸟哥的标准答案是：“你永远学不会！” 因为你是用“听”的，没有动手做，那么永远不会知道“经验”两个字怎么写！很多时候计算机/网络都会有一些莫名其妙的突发状况， 没有实际碰触过，怎么可能会理解呢？所以“永远是不可能听会的！”为啥要实验？因为实验过后你才会有经验来记下来？ 否则实验结果课本都有啊！不是背一背就好了，干麻实验呢？浪费钱吗？ ^_^",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/14.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/14.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 6,
      "title": "学习的激励循环",
      "user": {
        "id": "58d6bcee51cd2500111d11ca",
        "hasMigrated": false,
        "type": "User",
        "username": "huangdesheng",
        "name": "黄得胜",
        "website": "",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@huangdesheng",
          "stars": "https://legacy.gitbook.com/@huangdesheng/starred",
          "avatar": "https://s.gravatar.com/avatar/09cf2dfc47d99f97f608aeb451e55a05?s=220&d=https%3A%2F%2Flegacy.gitbook.com%2Fassets%2Fimages%2Favatars%2Fuser.png"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-03-25T18:54:38.634Z"
        },
        "counts": {},
        "google": {}
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/6"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-03-25T19:24:47.268Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "1.4 Linux 该如何学习",
        "filename": "14.md",
        "section": "除了上面的学习建议之外，还有其他的建议吗？确实是有的！其实， 无论作什么事情，对人类而言，两个重要的因素是造成我们学习的原动力：",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/14.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/14.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 7,
      "title": "inode 结果示意图挂了",
      "user": {
        "id": "567eb3b5bc58f31000bb908c",
        "hasMigrated": false,
        "type": "User",
        "username": "loggerhead",
        "name": "Fangzheng Long",
        "location": "",
        "website": "https://github.com/loggerhead",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@loggerhead",
          "stars": "https://legacy.gitbook.com/@loggerhead/starred",
          "avatar": "https://avatars0.githubusercontent.com/loggerhead"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2015-12-26T15:35:17.989Z"
        },
        "counts": {},
        "github": {
          "username": "loggerhead"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/7"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-03-27T01:16:12.089Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "7.1 认识 Linux 文件系统",
        "filename": "59.md",
        "section": "图7.1.4、inode 结构示意图",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/59.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/59.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 8,
      "title": "了解linux背景",
      "user": {
        "id": "58d4e5b693096700245568c2",
        "hasMigrated": false,
        "type": "User",
        "username": "lynzabo",
        "name": "林战波",
        "location": "China.Beijing.Chaoyang",
        "website": "https://github.com/Lynzabo",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@lynzabo",
          "stars": "https://legacy.gitbook.com/@lynzabo/starred",
          "avatar": "https://avatars0.githubusercontent.com/Lynzabo"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-03-24T09:24:06.524Z"
        },
        "counts": {},
        "github": {
          "username": "Lynzabo"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/8"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-04-07T08:05:50.938Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "目录及概述",
        "filename": "1.md",
        "section": "常常听到Linux具有非常优良的血统，所以具有相当良好的多用户多任务环境，可以方便程序设计师来开发软件。 此外，Linux本身是不用钱的“自由软件”，使用上面并没有所谓的“盗版”问题。但是，为什么Linux不用钱？ 随便修改或发布Linux为什么不会被罚？为什么Linux有这么多的版本？包括Fedora, SuSE, CentOS, Debian等等？ 这个都是我们必须要来了解的部分！了解这些部分，你才会对Linux有一个正确的理解，才能够跟你的同事、同学、 上司说明，为什么使用Linux具有很多优点与好处！ ^_^",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/1.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/1.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 9,
      "title": "学习linux，需要实际动手安装一部linux",
      "user": {
        "id": "58d4e5b693096700245568c2",
        "hasMigrated": false,
        "type": "User",
        "username": "lynzabo",
        "name": "林战波",
        "location": "China.Beijing.Chaoyang",
        "website": "https://github.com/Lynzabo",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@lynzabo",
          "stars": "https://legacy.gitbook.com/@lynzabo/starred",
          "avatar": "https://avatars0.githubusercontent.com/Lynzabo"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-03-24T09:24:06.524Z"
        },
        "counts": {},
        "github": {
          "username": "Lynzabo"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/9"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-04-07T08:07:39.821Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "目录及概述",
        "filename": "1.md",
        "section": "Linux并不好学习，鸟哥也是“重伤”过好几次才能对Linux有一些基础的认知。那么到底应该如何学习Linux呢？关键在实作。 既然要实作就得要实际的安装一部Linux，那么Linux要安装前需要熟悉哪些基础观念？计算机概论是非常重要的一环！ 因为Linux与硬件的关系还不小～此外，打造一台Windows/Linux共存的主机也是很有用的， 至少对于需要多平台但又缺乏空间与金钱的朋友来说，这样的处理是非常有用的！",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/1.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/1.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 10,
      "title": "这本书件该怎么看呢？建议先按照顺序将内容大致浏览过一次，看不懂的地方也可以先略过不要紧。 全部看完之后，再从头开始“仔细”的实际操作过一遍，那应该就能够进入Linux的世界啰～",
      "user": {
        "id": "58d4e5b693096700245568c2",
        "hasMigrated": false,
        "type": "User",
        "username": "lynzabo",
        "name": "林战波",
        "location": "China.Beijing.Chaoyang",
        "website": "https://github.com/Lynzabo",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@lynzabo",
          "stars": "https://legacy.gitbook.com/@lynzabo/starred",
          "avatar": "https://avatars0.githubusercontent.com/Lynzabo"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-03-24T09:24:06.524Z"
        },
        "counts": {},
        "github": {
          "username": "Lynzabo"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/10"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-04-07T08:10:06.660Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "目录及概述",
        "filename": "1.md",
        "section": "这本书的所有内容是学习Linux的基础，这些内容是基础中的基础，如果您能将其中的文字都看完并且消化过，那么未来在管理 Linux主机以及架设网站方面，就能够达到“事半功倍”的成效，请不要忽略这些内容了！否则，再怎么讨论都是枉然的啦！^_^。 Linux的资料非常的多，每份资料彼此的相关性都很强，要单独的一项一项讲解并不容易， 那么这本书件该怎么看呢？建议先按照顺序将内容大致浏览过一次，看不懂的地方也可以先略过不要紧。 全部看完之后，再从头开始“仔细”的实际操作过一遍，那应该就能够进入Linux的世界啰～",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/1.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/1.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 11,
      "title": "以前CPU都是8位，x86这个名字是使用最早的Intel发展出来的CPU代号命名",
      "user": {
        "id": "58d4e5b693096700245568c2",
        "hasMigrated": false,
        "type": "User",
        "username": "lynzabo",
        "name": "林战波",
        "location": "China.Beijing.Chaoyang",
        "website": "https://github.com/Lynzabo",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@lynzabo",
          "stars": "https://legacy.gitbook.com/@lynzabo/starred",
          "avatar": "https://avatars0.githubusercontent.com/Lynzabo"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-03-24T09:24:06.524Z"
        },
        "counts": {},
        "github": {
          "username": "Lynzabo"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/11"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-04-07T08:41:37.234Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "0.1 电脑：辅助人脑的好工具",
        "filename": "3.md",
        "section": "由于AMD、Intel、VIA所开发出来的x86架构CPU被大量使用于个人电脑（Personal computer）用途上面， 因此，个人电脑常被称为x86架构的电脑！那为何称为x86架构[8]呢？ 这是因为最早的那颗Intel发展出来的CPU代号称为8086，后来依此架构又开发出80286, 80386...， 因此这种架构的CPU就被称为x86架构了。",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/3.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/3.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 12,
      "title": "后来研发出64位CPU，故称架构为x86_64",
      "user": {
        "id": "58d4e5b693096700245568c2",
        "hasMigrated": false,
        "type": "User",
        "username": "lynzabo",
        "name": "林战波",
        "location": "China.Beijing.Chaoyang",
        "website": "https://github.com/Lynzabo",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@lynzabo",
          "stars": "https://legacy.gitbook.com/@lynzabo/starred",
          "avatar": "https://avatars0.githubusercontent.com/Lynzabo"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-03-24T09:24:06.524Z"
        },
        "counts": {},
        "github": {
          "username": "Lynzabo"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/12"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-04-07T08:42:21.213Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "0.1 电脑：辅助人脑的好工具",
        "filename": "3.md",
        "section": "在2003年以前由Intel所开发的x86架构CPU由8位升级到16、32位，后来AMD依此架构修改新一代的CPU为64位， 为了区别两者的差异，因此64位的个人电脑CPU又被统称为x86_64的架构喔！",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/3.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/3.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 13,
      "title": "invalid date format",
      "user": {
        "id": "58c8d9b73b62470011e2bcac",
        "hasMigrated": false,
        "type": "User",
        "username": "mikeeywang",
        "name": "pai",
        "location": "深圳",
        "website": "https://github.com/mikeeywang",
        "bio": "give me five",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@mikeeywang",
          "stars": "https://legacy.gitbook.com/@mikeeywang/starred",
          "avatar": "https://avatars0.githubusercontent.com/mikeeywang"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-03-15T06:05:43.759Z"
        },
        "counts": {},
        "github": {
          "username": "mikeeywang"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/13"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-04-13T01:47:21.772Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "第零章、计算机概论",
        "filename": "2.md",
        "section": "最近更新日期：20//",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/2.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/2.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 14,
      "title": "多个磁盘怎么挂载比较好",
      "user": {
        "id": "59156d415b2ae40012b820d8",
        "hasMigrated": false,
        "type": "User",
        "username": "zerohyuan",
        "name": "zerohyuan",
        "location": "",
        "website": "https://github.com/zerohyuan",
        "bio": "",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@zerohyuan",
          "stars": "https://legacy.gitbook.com/@zerohyuan/starred",
          "avatar": "https://avatars0.githubusercontent.com/zerohyuan"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-05-12T08:07:29.932Z"
        },
        "counts": {},
        "github": {
          "username": "zerohyuan"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/14"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-05-22T09:18:52.017Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "2.3 安装Linux前的规划",
        "filename": "21.md",
        "section": "安装最重要的第一件事，就是要取得Linux distributions的光盘数据，该如何去下载？ 目前有这么多的distributions，你应该要选择哪一个版本比较好？为什么会比较好？ 在台湾，你可以在哪里下载你所需要的Linux distribution呢？这是这一小节所要讨论的喔！",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/21.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/21.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 15,
      "title": "写程序管理系统的时候最好使用绝对路径的写法，正确性比较好，绝对不会有问题。",
      "user": {
        "id": "58bf9f49e433d30011f4c8b5",
        "hasMigrated": false,
        "type": "User",
        "username": "dursss",
        "name": "dursss",
        "location": "",
        "website": "https://github.com/dursss",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@dursss",
          "stars": "https://legacy.gitbook.com/@dursss/starred",
          "avatar": "https://avatars0.githubusercontent.com/dursss"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-03-08T06:06:01.451Z"
        },
        "counts": {},
        "github": {
          "username": "dursss"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/15"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-06-05T02:53:46.103Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "6.1 目录与路径",
        "filename": "49.md",
        "section": "但是对于文件名的正确性来说，“绝对路径的正确度要比较好～”。 一般来说，鸟哥会建议你，如果是在写程序 （shell scripts） 来管理系统的条件下，务必使用绝对路径的写法。 怎么说呢？因为绝对路径的写法虽然比较麻烦，但是可以肯定这个写法绝对不会有问题。 如果使用相对路径在程序当中，则可能由于你执行的工作环境不同，导致一些问题的发生。 这个问题在工作调度（at, cron, 第十五章）当中尤其重要！这个现象我们在十二章、shell script时，会再次的提醒你喔！ ^_^",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/49.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/49.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 16,
      "title": "根目录的上一层与根目录自己是同一个目录",
      "user": {
        "id": "58bf9f49e433d30011f4c8b5",
        "hasMigrated": false,
        "type": "User",
        "username": "dursss",
        "name": "dursss",
        "location": "",
        "website": "https://github.com/dursss",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@dursss",
          "stars": "https://legacy.gitbook.com/@dursss/starred",
          "avatar": "https://avatars0.githubusercontent.com/dursss"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-03-08T06:06:01.451Z"
        },
        "counts": {},
        "github": {
          "username": "dursss"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/16"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-06-05T02:56:53.938Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "6.1 目录与路径",
        "filename": "49.md",
        "section": "例题：请问在Linux下面，根目录下有没有上层目录（..）存在？答：若使用“ ls -al / ”去查询，可以看到根目录下确实存在 . 与 .. 两个目录，再仔细的查阅， 可发现这两个目录的属性与权限完全一致，这代表根目录的上一层（..）与根目录自己（.）是同一个目录。",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/49.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/49.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 17,
      "title": "cd是Change Directory的缩写，这是用来变换工作目录的指令",
      "user": {
        "id": "58bf9f49e433d30011f4c8b5",
        "hasMigrated": false,
        "type": "User",
        "username": "dursss",
        "name": "dursss",
        "location": "",
        "website": "https://github.com/dursss",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@dursss",
          "stars": "https://legacy.gitbook.com/@dursss/starred",
          "avatar": "https://avatars0.githubusercontent.com/dursss"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-03-08T06:06:01.451Z"
        },
        "counts": {},
        "github": {
          "username": "dursss"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/17"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-06-05T03:04:36.540Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "6.1 目录与路径",
        "filename": "49.md",
        "section": "cd是Change Directory的缩写，这是用来变换工作目录的指令。注意，目录名称与cd指令之间存在一个空格。 一登陆Linux系统后，每个帐号都会在自己帐号的主文件夹中。那回到上一层目录可以用“ cd .. ”。 利用相对路径的写法必须要确认你目前的路径才能正确的去到想要去的目录。例如上表当中最后一个例子， 你必须要确认你是在/var/spool/mail当中，并且知道在/var/spool当中有个mqueue的目录才行啊～ 这样才能使用cd ../postfix 去到正确的目录说，否则就要直接输入cd /var/spool/postfix 啰～",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/49.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/49.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 18,
      "title": "pwd是Print Working Directory的缩写，也就是显示目前所在目录的指令",
      "user": {
        "id": "58bf9f49e433d30011f4c8b5",
        "hasMigrated": false,
        "type": "User",
        "username": "dursss",
        "name": "dursss",
        "location": "",
        "website": "https://github.com/dursss",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@dursss",
          "stars": "https://legacy.gitbook.com/@dursss/starred",
          "avatar": "https://avatars0.githubusercontent.com/dursss"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-03-08T06:06:01.451Z"
        },
        "counts": {},
        "github": {
          "username": "dursss"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/18"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-06-05T03:09:26.883Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "6.1 目录与路径",
        "filename": "49.md",
        "section": "pwd是Print Working Directory的缩写，也就是显示目前所在目录的指令， 例如在上个表格最后的目录是/var/mail这个目录，但是提示字符仅显示mail， 如果你想要知道目前所在的目录，可以输入pwd即可。此外，由于很多的套件所使用的目录名称都相同，例如 /usr/local/etc还有/etc，但是通常Linux仅列出最后面那一个目录而已，这个时候你就可以使用pwd 来知道你的所在目录啰！免得搞错目录，结果...",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/49.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/49.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 19,
      "title": "链接可以当作win上的快捷方式来理解， pwd -P 才是真实的路径 注意大小写",
      "user": {
        "id": "58bf9f49e433d30011f4c8b5",
        "hasMigrated": false,
        "type": "User",
        "username": "dursss",
        "name": "dursss",
        "location": "",
        "website": "https://github.com/dursss",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@dursss",
          "stars": "https://legacy.gitbook.com/@dursss/starred",
          "avatar": "https://avatars0.githubusercontent.com/dursss"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-03-08T06:06:01.451Z"
        },
        "counts": {},
        "github": {
          "username": "dursss"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/19"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-06-05T03:11:33.524Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "6.1 目录与路径",
        "filename": "49.md",
        "section": "其实有趣的是那个 -P 的选项啦！他可以让我们取得正确的目录名称，而不是以链接文件的路径来显示的。 如果你使用的是CentOS 7.x的话，刚刚好/var/mail是/var/spool/mail的链接文件， 所以，通过到/var/mail下达pwd -P就能够知道这个选项的意义啰～ ^_^",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/49.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/49.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 20,
      "title": "谷歌",
      "user": {
        "id": "5937b1754d9be20012eafa55",
        "hasMigrated": false,
        "type": "User",
        "username": "lingowei",
        "name": "Lingo.Blyde",
        "location": "wuhan, beijing",
        "website": "https://github.com/LinGoWei",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@lingowei",
          "stars": "https://legacy.gitbook.com/@lingowei/starred",
          "avatar": "https://avatars0.githubusercontent.com/LinGoWei"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-06-07T07:55:33.383Z"
        },
        "counts": {},
        "github": {
          "username": "LinGoWei"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/20"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-06-07T07:56:06.219Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "1.4 Linux 该如何学习",
        "filename": "14.md",
        "section": "除了这些基本的FAQ之外，其实，还有更重要的问题查询方法，那就是利用酷狗（Google）帮您去搜寻答案呢！ 在鸟哥学习Linux的过程中，如果有什么奇怪的问题发生时，第一个想到的， 就是去http://www.google.com.tw搜寻是否有相关的议题。 举例来说，我想要找出Linux下面的NAT，只要在上述的网站内，输入Linux跟NAT， 立刻就有一堆文献跑出来了！真的相当的优秀好用喔！您也可以通过酷狗来找鸟哥网站上的数据呢！",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/14.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/14.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 21,
      "title": "no this image",
      "user": {
        "id": "590d497b20e0e200122ec9c2",
        "hasMigrated": false,
        "type": "User",
        "username": "zhaopengli",
        "name": "zhaopeng li",
        "website": "",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@zhaopengli",
          "stars": "https://legacy.gitbook.com/@zhaopengli/starred",
          "avatar": "https://s.gravatar.com/avatar/32e7fe3251418ab28c938baa82cb6f0d?s=220&d=https%3A%2F%2Flegacy.gitbook.com%2Fassets%2Fimages%2Favatars%2Fuser.png"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-05-06T03:56:43.362Z"
        },
        "counts": {},
        "google": {}
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/21"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-06-11T04:12:03.969Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "0.2 个人电脑架构与相关设备元件",
        "filename": "4.md",
        "section": "图 0.2.2、ASUS 主板 （图片为华硕公司所有）",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/4.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/4.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 22,
      "title": "no this image",
      "user": {
        "id": "590d497b20e0e200122ec9c2",
        "hasMigrated": false,
        "type": "User",
        "username": "zhaopengli",
        "name": "zhaopeng li",
        "website": "",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@zhaopengli",
          "stars": "https://legacy.gitbook.com/@zhaopengli/starred",
          "avatar": "https://s.gravatar.com/avatar/32e7fe3251418ab28c938baa82cb6f0d?s=220&d=https%3A%2F%2Flegacy.gitbook.com%2Fassets%2Fimages%2Favatars%2Fuser.png"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-05-06T03:56:43.362Z"
        },
        "counts": {},
        "google": {}
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/22"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-06-11T04:12:06.275Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "0.2 个人电脑架构与相关设备元件",
        "filename": "4.md",
        "section": "图 0.2.2、ASUS 主板 （图片为华硕公司所有）",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/4.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/4.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 23,
      "title": "mark",
      "user": {
        "id": "5951fbb8fd3fe70011ab55e3",
        "hasMigrated": false,
        "type": "User",
        "username": "neilhe",
        "name": "Neil He",
        "website": "",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@neilhe",
          "stars": "https://legacy.gitbook.com/@neilhe/starred",
          "avatar": "https://s.gravatar.com/avatar/4a1d24a90c36dabe5eb386b36637f104?s=220&d=https%3A%2F%2Flegacy.gitbook.com%2Fassets%2Fimages%2Favatars%2Fuser.png"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-06-27T06:31:20.507Z"
        },
        "counts": {},
        "google": {}
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/23"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-06-27T06:31:57.332Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "0.1 电脑：辅助人脑的好工具",
        "filename": "3.md",
        "section": "图0.1.2、电脑的五大单元[4]",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/3.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/3.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 24,
      "title": "mark",
      "user": {
        "id": "5951fbb8fd3fe70011ab55e3",
        "hasMigrated": false,
        "type": "User",
        "username": "neilhe",
        "name": "Neil He",
        "website": "",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@neilhe",
          "stars": "https://legacy.gitbook.com/@neilhe/starred",
          "avatar": "https://s.gravatar.com/avatar/4a1d24a90c36dabe5eb386b36637f104?s=220&d=https%3A%2F%2Flegacy.gitbook.com%2Fassets%2Fimages%2Favatars%2Fuser.png"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-06-27T06:31:20.507Z"
        },
        "counts": {},
        "google": {}
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/24"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-06-27T06:34:43.157Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "0.1 电脑：辅助人脑的好工具",
        "filename": "3.md",
        "section": "CPU的运算速度常使用 MHz 或者是 GHz 之类的单位，这个 Hz 其实就是秒分之一。而在网络传输方面，由于网络使用的是 bit 为单位，因此网络常使用的单位为 Mbps 是 Mbits per second，亦即是每秒多少 Mbit。举例来说，大家常听到的 20M/5M 光世代传输速度，如果转成文件大小的 Byte 时，其实理论最大传输值为：每秒 2.5MByte/ 每秒625KByte的下载/上传速度喔！",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/3.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/3.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 25,
      "title": "其实删除文件就是修改父目录的inode，只需要有目录的写权限就行了。被删的文件被丢在了回收站。就好比丢进了垃圾桶。",
      "user": {
        "id": "5965c4a37fb976001016ab20",
        "hasMigrated": false,
        "type": "User",
        "username": "kunji",
        "name": "Kun Ji",
        "website": "",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@kunji",
          "stars": "https://legacy.gitbook.com/@kunji/starred",
          "avatar": "https://s.gravatar.com/avatar/598f8817e14cfe0460400d4c7b3655e5?s=220&d=https%3A%2F%2Flegacy.gitbook.com%2Fassets%2Fimages%2Favatars%2Fuser.png"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-07-12T06:41:39.589Z"
        },
        "counts": {},
        "google": {}
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/25"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-07-12T06:43:30.842Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "5.2 Linux 文件权限概念",
        "filename": "43.md",
        "section": "Tips 上述的例子解释是这样的，假设有个莫名其妙的人，拿着一个完全密封的数据夹放到你的办公室抽屉中，因为完全密封你也打不开、看不到这个数据夹的内部数据（对文件来说，你没有权限）。 但是因为这个数据夹是放在你的抽屉中，你当然可以拿出/放入任何数据在这个抽屉中（对目录来说，你具有所有权限）。 所以，情况就是：你打开抽屉、拿出这个没办法看到的数据夹、将他丢到走廊上的垃圾桶！搞定了 （顺利删除！）！",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/43.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/43.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 26,
      "title": "【不能看了】跳转url没有备案",
      "user": {
        "id": "58de24c95c562b001172d499",
        "hasMigrated": false,
        "type": "User",
        "username": "windzz",
        "name": "windzz",
        "location": "",
        "website": "https://github.com/windzz",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@windzz",
          "stars": "https://legacy.gitbook.com/@windzz/starred",
          "avatar": "https://avatars0.githubusercontent.com/windzz"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-03-31T09:43:37.365Z"
        },
        "counts": {},
        "github": {
          "username": "windzz"
        }
      },
      "subscribed": null,
      "body": "<p>点卡跳转后 提示 没有备案，无法看了</p>\n",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/26"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-07-25T12:09:51.050Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      }
    },
    {
      "type": "Discussion",
      "number": 27,
      "title": "dd if=/boot/initramfs-3.10.0-229.el7.x86_64.img of=initramfs.gz \\ &gt;  bs=11264 skip=1 这个没法执行啊，就算去掉\\&gt; 生成出来的文件也没有办法gzip解压缩。",
      "user": {
        "id": "599fe83010fa0800115bac3d",
        "hasMigrated": false,
        "type": "User",
        "username": "vincents",
        "name": "Vincent S",
        "website": "",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@vincents",
          "stars": "https://legacy.gitbook.com/@vincents/starred",
          "avatar": "https://s.gravatar.com/avatar/5e39a24cbc71896c5c874efdd5ac4046?s=220&d=https%3A%2F%2Flegacy.gitbook.com%2Fassets%2Fimages%2Favatars%2Fuser.png"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-08-25T09:04:48.577Z"
        },
        "counts": {},
        "google": {}
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/27"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2017-08-25T09:08:12.480Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "19.1 Linux 的开机流程分析",
        "filename": "166.md",
        "section": "从上面我们大概知道了这个 initramfs 里头含有两大区块，一个是事先宣告的一些数据，包括 kernel/x86/microcode/GenuineIntel.bin 这些东西。 在这些数据后面，才是真的我们的核心会去读取的重要文件～如果看一下文件的内容，你会发现到 init 那只程序已经被 systemd 所取代啰！这样理解否？ 好～如果你想要进一步将这个文件解开的话，那得要先将前面的 kernel/x86/microcode/GenuineIntel.bin 之前的文件先去除掉，这样才能够顺利的解开。 因此，得要这样进行：",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/166.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/166.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 32,
      "title": "本段的最后一行，“自订变量”需要改成“自定变量”",
      "user": {
        "id": "5a6549f9674ff3002efb18d8",
        "hasMigrated": false,
        "type": "User",
        "username": "feipan664",
        "name": "fei pan",
        "website": "",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@feipan664",
          "stars": "https://legacy.gitbook.com/@feipan664/starred",
          "avatar": "https://s.gravatar.com/avatar/49bae741e6997ba5856cc96a86f90ec5?s=220&d=https%3A%2F%2Flegacy.gitbook.com%2Fassets%2Fimages%2Favatars%2Fuser.png"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2018-01-22T02:18:33.741Z"
        },
        "counts": {},
        "google": {}
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/32"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2018-01-22T02:21:01.588Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "10.2 Shell 的变量功能",
        "filename": "88.md",
        "section": "如果说的学理一点，那么由于在 Linux System 下面，所有的线程都是需要一个执行码， 而就如同上面提到的，你“真正以 shell 来跟 Linux 沟通，是在正确的登陆 Linux 之后！”这个时候你就有一个 bash 的执行程序，也才可以真正的经由 bash 来跟系统沟通啰！而在进入 shell 之前，也正如同上面提到的，由于系统需要一些变量来提供他数据的存取 （或者是一些环境的设置参数值， 例如是否要显示彩色等等的） ，所以就有一些所谓的“环境变量” 需要来读入系统中了！这些环境变量例如 PATH、HOME、MAIL、SHELL 等等，都是很重要的， 为了区别与自订变量的不同，环境变量通常以大写字符来表示呢！",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/88.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/88.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 33,
      "title": "表格格式修正：3.1本练习机的规划",
      "user": {
        "id": "599d5f162e3c8700108a9b52",
        "hasMigrated": false,
        "type": "User",
        "username": "maecenas",
        "name": "MaecenasLi",
        "location": "",
        "website": "https://github.com/Maecenas",
        "bio": "",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@maecenas",
          "stars": "https://legacy.gitbook.com/@maecenas/starred",
          "avatar": "https://avatars0.githubusercontent.com/Maecenas"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-08-23T10:55:18.244Z"
        },
        "counts": {},
        "github": {
          "username": "Maecenas"
        },
        "google": {},
        "facebook": {}
      },
      "subscribed": null,
      "body": "<p>3.1 本练习机的规划--尤其是分区参数</p>\n<p>一章中的表格有两处格式错误，敬请修正。</p>\n",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/33"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2018-01-28T04:17:24.545Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      }
    },
    {
      "type": "Discussion",
      "number": 34,
      "title": "求这本书的markdown源码版",
      "user": {
        "id": "5a548a022f709d002e7a9907",
        "hasMigrated": false,
        "type": "User",
        "username": "snow212-cn",
        "name": "snow212-cn",
        "location": "",
        "website": "https://github.com/snow212-cn",
        "bio": "",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@snow212-cn",
          "stars": "https://legacy.gitbook.com/@snow212-cn/starred",
          "avatar": "https://avatars0.githubusercontent.com/SNOW-city"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2018-01-09T09:23:14.554Z"
        },
        "counts": {},
        "github": {
          "username": "SNOW-city"
        }
      },
      "subscribed": null,
      "body": "<p>作者或大家谁有</p>\n",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/34"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2018-05-08T12:24:25.572Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      }
    },
    {
      "type": "Discussion",
      "number": 35,
      "title": "什么是电脑（计算机）？",
      "user": {
        "id": "5a27becfc7f342002fa72fdb",
        "hasMigrated": false,
        "type": "User",
        "username": "guangshenglin",
        "name": "林光胜",
        "location": "ShanTou, Guangdong, China",
        "website": "https://github.com/GuangshengLin",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@guangshenglin",
          "stars": "https://legacy.gitbook.com/@guangshenglin/starred",
          "avatar": "https://avatars0.githubusercontent.com/GuangshengLin"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-12-06T09:56:31.500Z"
        },
        "counts": {},
        "github": {
          "username": "GuangshengLin"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/35"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2018-08-02T02:02:55.231Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "0.1 电脑：辅助人脑的好工具",
        "filename": "3.md",
        "section": "所谓的电脑就是一种计算机，而计算机其实是：“接受使用者输入指令与数据，经由中央处理器的数学与逻辑单元运算处理后， 以产生或储存成有用的信息”。因此，只要有输入设备 （不管是键盘还是触摸屏） 及输出设备 （例如电脑屏幕或直接由打印机打印出来），让你可以输入数据使该机器产生信息的， 那就是一部计算机了。",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/3.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/3.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 36,
      "title": "CPU的指令集",
      "user": {
        "id": "5a27becfc7f342002fa72fdb",
        "hasMigrated": false,
        "type": "User",
        "username": "guangshenglin",
        "name": "林光胜",
        "location": "ShanTou, Guangdong, China",
        "website": "https://github.com/GuangshengLin",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@guangshenglin",
          "stars": "https://legacy.gitbook.com/@guangshenglin/starred",
          "avatar": "https://avatars0.githubusercontent.com/GuangshengLin"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-12-06T09:56:31.500Z"
        },
        "counts": {},
        "github": {
          "username": "GuangshengLin"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/36"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2018-08-02T02:03:12.821Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "0.1 电脑：辅助人脑的好工具",
        "filename": "3.md",
        "section": "如前面说过的，CPU 其实内部已经含有一些微指令，我们所使用的软件都要经过 CPU 内部的微指令集来达成才行。 那这些指令集的设计主要又被分为两种设计理念，这就是目前世界上常见到的两种主要 CPU 架构， 分别是：精简指令集 （RISC） 与复杂指令集 （CISC） 系统。下面我们就来谈谈这两种不同 CPU 架构的差异啰！",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/3.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/3.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 37,
      "title": "Linux是操作系统；操作系统与硬件关系很密切，与硬件基本上是一对一关系。不同硬件基础的操作系统无法简单移植",
      "user": {
        "id": "5a27becfc7f342002fa72fdb",
        "hasMigrated": false,
        "type": "User",
        "username": "guangshenglin",
        "name": "林光胜",
        "location": "ShanTou, Guangdong, China",
        "website": "https://github.com/GuangshengLin",
        "verified": true,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@guangshenglin",
          "stars": "https://legacy.gitbook.com/@guangshenglin/starred",
          "avatar": "https://avatars0.githubusercontent.com/GuangshengLin"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2017-12-06T09:56:31.500Z"
        },
        "counts": {},
        "github": {
          "username": "GuangshengLin"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/37"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2018-08-02T07:04:45.011Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "1.1 Linux是什么",
        "filename": "11.md",
        "section": "我们在第零章、计算机概论里面有提到过整个计算机系统的概念， 计算机主机是由一堆硬件所组成的，为了有效率的控制这些硬件资源，于是乎就有操作系统的产生了。 操作系统除了有效率的控制这些硬件资源的分配，并提供计算机运行所需要的功能（如网络功能）之外， 为了要提供程序设计师更容易开发软件的环境，所以操作系统也会提供一整组系统调用接口来给软件设计师开发用喔！",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/11.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/11.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 38,
      "title": "parted print list 可以查看磁盘类型，tune2fs /dev/sdb1 可以查看inode 与 block等大小",
      "user": {
        "id": "57fca2874955cc1000a15930",
        "hasMigrated": false,
        "type": "User",
        "username": "hunter524",
        "name": "Htao",
        "location": "",
        "website": "https://github.com/hunter524",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@hunter524",
          "stars": "https://legacy.gitbook.com/@hunter524/starred",
          "avatar": "https://avatars0.githubusercontent.com/hunter524"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2016-10-11T08:27:51.468Z"
        },
        "counts": {},
        "github": {
          "username": "hunter524"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/38"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2018-11-27T03:09:25.039Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "7.1 认识 Linux 文件系统",
        "filename": "59.md",
        "section": "例题：假设你的Ext2文件系统使用 4K block ，而该文件系统中有 10000 个小文件，每个文件大小均为 50Bytes， 请问此时你的磁盘浪费多少容量？答：由于 Ext2 文件系统中一个 block 仅能容纳一个文件，因此每个 block 会浪费“ 4096 - 50 = 4046 （Byte）”， 系统中总共有一万个小文件，所有文件大小为：50 （Bytes） x 10000 = 488.3KBytes，但此时浪费的容量为：“ 4046 （Bytes） x 10000 = 38.6MBytes ”。想一想，不到 1MB 的总文件大小却浪费将近 40MB 的容量，且文件越多将造成越多的磁盘容量浪费。 ",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/59.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/59.md"
        }
      }
    },
    {
      "type": "Discussion",
      "number": 39,
      "title": "！！！为什么hard link 不能link目录？",
      "user": {
        "id": "57fca2874955cc1000a15930",
        "hasMigrated": false,
        "type": "User",
        "username": "hunter524",
        "name": "Htao",
        "location": "",
        "website": "https://github.com/hunter524",
        "verified": false,
        "locked": false,
        "site_admin": false,
        "urls": {
          "profile": "https://legacy.gitbook.com/@hunter524",
          "stars": "https://legacy.gitbook.com/@hunter524/starred",
          "avatar": "https://avatars0.githubusercontent.com/hunter524"
        },
        "permissions": {
          "edit": null,
          "admin": null
        },
        "dates": {
          "created": "2016-10-11T08:27:51.468Z"
        },
        "counts": {},
        "github": {
          "username": "hunter524"
        }
      },
      "subscribed": null,
      "body": "",
      "state": "open",
      "locked": false,
      "urls": {
        "details": "https://legacy.gitbook.com/book/wizardforcel/vbird-linux-basic-4e/discussions/39"
      },
      "counts": {
        "participants": 1,
        "comments": 0
      },
      "dates": {
        "created": "2018-11-29T08:41:32.081Z"
      },
      "permissions": {
        "comment": false,
        "close": false,
        "open": false,
        "admin": false
      },
      "book": {
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
          "important": false
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
          "permissions": {},
          "dates": {
            "created": "2014-11-23T03:26:16.278Z"
          },
          "counts": {},
          "github": {
            "username": "wizardforcel"
          }
        }
      },
      "context": {
        "chapterTitle": "7.2 文件系统的简单操作",
        "filename": "60.md",
        "section": "不能跨 Filesystem 还好理解，那不能 hard link 到目录又是怎么回事呢？这是因为如果使用 hard link 链接到目录时， 链接的数据需要连同被链接目录下面的所有数据都创建链接，举例来说，如果你要将 /etc 使用实体链接创建一个 /etc_hd 的目录时，那么在 /etc_hd 下面的所有文件名同时都与 /etc 下面的文件名要创建 hard link 的，而不是仅链接到 /etc_hd 与 /etc 而已。 并且，未来如果需要在 /etc_hd 下面创建新文件时，连带的， /etc 下面的数据又得要创建一次 hard link ，因此造成环境相当大的复杂度。 所以啰，目前 hard link 对于目录暂时还是不支持的啊！",
        "urls": {
          "content": "https://wizardforcel.gitbooks.io/vbird-linux-basic-4e/content/60.html",
          "edit": "/book/wizardforcel/vbird-linux-basic-4e/edit#/edit/master/60.md"
        }
      }
    }
  ],
  "total": 35,
  "limit": 50,
  "page": 0,
  "pages": 1
}
```

