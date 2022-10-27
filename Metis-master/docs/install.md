
## Installation document

# Table of contents
> * [Method 1: Manual installation deployment] (#Chapter-1)
>> * [Dependent Environment] (#Chapter-1-1)
>> * [Database Environment Installation] (#Chapter-2)
>> * [Installation of the service side environment] (#Chapter-3)
>> * [Web Management Economy Installation] (#Chapter-1-4)
Forecast
> * [Method 2: Docker Installation Deployment] (#Chapter-2)
Forecast
This installation document only describes the process of installing the entire Metis on one server, the purpose is to allow users to understand the overall understanding of Metis's deployment and operation.

If you want to be used for the production environment, you need to consider more faults and disaster tolerance in distributed systems.If necessary, you can join the QQ technical exchange group of Metis: 288723616.

This document provides two installation methods: manual installation deployment and docker installation deployment.You can choose any installation method as needed.

# 1. <a ID="chapter-1"> </a> Manual installation deployment
## 1.1. <a ID="chapter-1"> </a> dependence on the environment

| Software | Software Requirements |
| --- | --- | |
| Linux distribution version: | CentOS 7.4 |
| python version: | 2.7 version |
| MySQL version: | 5.6.26 and above version |
| Node.js version: | 8.11.1 and above | S
| Django version: | 1.x.x version |

Run the server requirements: 1 ordinary machine installed Linux system (recommended CentOS system). The server needs to open port 80 and 8080 ports

The following steps assume that the code directory of the installation machine is `/data/metis/`, which can be changed according to the actual situation.

## 1.2. <a ID="chapter-2"> </a> Database environment installation

### 1.2.1. Mysql installation introductionallation introduction

Use YUM source installation or install the source code installation under the MySQL official website.

---
```
yum install mariadb-server
systemctl start mariadb
```
---

### 1.2.2. Initialize the database

In order to facilitate fast experience, 10+ abnormal detection results data and 300+ sample data are provided for everyone to use.

1. Create the required database user name and authorize, connect the MySQL client and execute

---
```
   grant all privileges on metis.* to metis@127.0.0.1  identified by 'metis@123';
   flush privileges;
```
---
2. Create a database `metis`, execute under the command line

---
```
mysqladmin -umetis -pmetis@123 -h127.0.0.1 create metis
```
---

3. Put the SQL initialization file in the `Metis/App/SQL/TIME_SERIES_DETECTOR/` directory, import the data `metis` database

---
```
mysql -umetis -pmetis@123 -h127.0.0.1 metis < /data/Metis/app/sql/time_series_detector/anomaly.sql
mysql -umetis -pmetis@123 -h127.0.0.1 metis < /data/Metis/app/sql/time_series_detector/sample_dataset.sql
mysql -umetis -pmetis@123 -h127.0.0.1 metis < /data/Metis/app/sql/time_series_detector/train_task.sql
```
---

4. Update the database configuration information to the server -side configuration file `database.py`
---
```
vim /data/Metis/app/dao/db_common/database.py
```
---
Rewritten configuration
---
```
DB = 'metis'
USER = 'metis'
PASSWD = 'metis@123'
HOST = '127.0.0.1'
PORT = 3306
```
---

## 1.3. <a ID="chapter-3"> </a> Installation

The server Python program needs to rely on Django, Numpy, TSFRESH, MySQL-Python, Scikit-Learn and other packages

### 1.3.1. Yum installation dependency package

---
```
yum install python-pip
pip install --upgrade pip
yum install gcc libffi-devel python-devel openssl-devel
yum install mysql-devel
```
---

### 1.3.2. PIP Install Python dependency package

Install through the project directory `/metis/docs/requirements.txt`

---
```
pip install -I -r requirements.txt
```
---

### 1.3.3.

---
```
export PYTHONPATH=/data/Metis:$PYTHONPATH
```
---

In order to ensure that you can import environment variables next time, please write the environment variable configuration to the server's `/etc/profile` file of the server


### 1.3.4. Start the server

Start the server program, please replace it with the IP address of the server

---
```
python /data/Metis/app/controller/manage.py runserver {ip}:{port}
```
---

This startup mode is the debug mode of Django.If you need to deploy the production environment, you can deploy through nginx and UWSGI. For details, please refer to the corresponding official website description

## 1.4. <a ID="chapter-4"> </a> Web management end environment installation

### 1.4.1. Node.js installation

---
```go

```
---

You need to install [node.js] (https://nodejs.org/en/download/), and the version of the node.js needs not less than 8.11.1

### 1.4.2. NPM Install the front end dependencies

Installation `/metis/web/package.json` The third -party installation package dependent in the configuration file

Enter the UWEB directory and execute NPM Install

### 1.4.3. Compile code

Modify the back -end address configuration of the `/metis/uweb/src/app.json` file:" Origin ":" http: // $ {IP}: $ {port} ", IP and Port corresponding to the service side address

Run NPM Run Build

Copy the Custom folder in the uweb directory to the distal folder generated in the UWEB directory

Position the root in the nginx configuration file to the distal folder in the UWEB directory

Nginx configuration is as follows:

---
```
server {
        listen       80;
        root /*/uweb/dist;
        location / {
                add_header Cache-Control max-age=0;
                gzip on;
                gzip_min_length 1k;
                gzip_buffers 16 64k;
                gzip_http_version 1.1;
                gzip_comp_level 6;
                gzip_types text/plain application/x-javascript text/css application/xml;
                gzip_vary on;
                    try_files $uri $uri/ /index.html;
        }

        location /index.html {
                add_header Cache-Control 'no-store';
        }
    }
```
---

### 1.4.4. Start the web service

After nginx starts normally, open the browser and visit `http:// $ {ip}: 80/`

### 1.4.5.

If the code is modified locally,the update method is as follows:

After the development of the NPM Run Build project code is completed, the order is executed to pack the project code.A DIST directory will be generated in the project root directory, and then copy the Custom directory and put it in the DIST directory.At the time of release, put all the files in the DIST directory as a static file and put it to the static file directory specified by the server

After the installation is completed, please refer to the API instructions for API call

# 2. <a ID="chapter-5"> </a> Docker installation deployment

## 2.1. Install docker

---
```
yum install docker
service docker start
```
---

## 2.2. <a ID="chapter-2"> </a> deploy the docker environment
Execute deployment file
---
```
metis/docker/startSh ${ip}
```
---
After the deployment is completed, execute the docker PS command
---
```
docker ps
```
---
View the starting status of the three containers (Metis-DB, Metis-WEB, Metis-SVR). If it starts normally, the installation is successful.
![docker_ps](images/docker_ps.png)
If the installation is successful, you can directly access it through the browser: `http: // $ {ip}`
Note: Metis relies on ports 80 and 8080. Tencent Cloud Server has opened 80 outer network access rights by default but did not open 8080. It is necessary to manually add port 8080 to the security group.

Please refer to the API instructions for API calling