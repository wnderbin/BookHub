# BookHub

![image](https://github.com/user-attachments/assets/71944739-9a6a-4b6c-adfc-7bf729a108ab)

![image](https://github.com/user-attachments/assets/906cdbd1-f40a-4045-bc96-ff4f7628e293)

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-bookhub-project">About the BookHub project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#dependencies">Dependencies</a></li>
        <li><a href="#installation-and-launch">Installation & Launch</a></li>
      </ul>
    </li>
    <li><a href="#project-structure">Project structure</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#author">Author</a></li>
  </ol>
</details>

## About the BookHub project

BookHub is a storage website where you can share your books with other users.

### Built with:

[![My Skills](https://skillicons.dev/icons?i=go,html,css,mysql,bash)](https://skillicons.dev)

* **Go** - Backend
* **Html & Css** - Frontend
* **MySQL** - DBMS, data storage
* **Bash** - Scripts (website launch)

## Getting Started

Instructions on how to run a project locally

### Dependencies

* **MySQL driver** - github.com/go-sql-driver/mysql
* **gorilla/mux** - github.com/gorilla/mux

### Installation and Launch

```
git clone https://github.com/wnderbin/BookHub # clone the repository
```

```
cd BookHub/scripts # go to the directory with scripts

chmod +x run.bash
./run.bash # running script
```

#### !!! Before running the project, make sure that your MySQL database is running with the following parameters:
* **books** - user
* **pass** - password
* **web_books** - database
* **localhost** - ip address (127.0.0.1)
* **3306** - port

## Project structure

**BookHub**: Project directory \
&nbsp; &nbsp; ├─ **cmd/web_books** - Directory with executable module\
&nbsp; &nbsp; ├─ **internal** - Additional modules\
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; ├─ **handlers** - Directory with handlers\
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; └─ **mysql** - Directory with commands for MySQL\
&nbsp; &nbsp; ├─ **scripts** - Directory with scripts for running the project\
&nbsp; &nbsp; └─  **ui** - user interface (html/css)

## License
Before using the project, it is recommended to read the license

## Author:
* wnderbin
