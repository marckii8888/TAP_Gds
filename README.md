# GovTech TAP GDS - Technical Assessment
This project is a URL shortener for GovTech TAP GDS technical assessment.
## Features
- Input any url and it will return a shortened url
- In-built clip board to make it easier to copy the shortened url

## Tech
The URL Shortener uses the following technologies:
### Frontend
- ReactJS - Frontend framework, styled with bootstrap css
### Backend
- Golang - To host the REST API
- MySQL - Relational database to store url mappings
 
## Installation
To run the URL Shortener, the following 
- [Node.js](https://nodejs.org) v12.16.1+ 
- [Golang](https://golang.org/doc/install)
- [MySQL](https://www.mysql.com/downloads/)

### 1. Edit the config file
```sh
cd Backend
cd config
```
Open config.yml in a text editor and change the values of MySQL credentials
### 2. Run the MySQL script
```sh
mysql -u root -p < script.sql
```
### 3. Install the dependencies and start the backend server
```sh
cd Backend
go mod tidy
go run main.go
```
### 4. Install the dependencies and start the front end server.
```sh
cd Frontend
npm install
npm start
```