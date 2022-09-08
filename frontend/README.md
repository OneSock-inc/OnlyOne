# Frontend

OnlyOne frontend is an Angular application. It is disigned to run on every browser. However, it is primarily adapted to small devices like smartphones. This application uses http requests to communicate with the custom backend part of this project. 

---

## Run the project
### Prerequisites

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 14.1.3. 
A few steps are necessary to run this application. First you have to install NodeJS on a local or remote machine. We recommend the version `16.17` of Node. Further informations about Node install are availables [here](https://nodejs.org/en/). 
Than you can install globally the Angular CLI tools with the command `npm install -g @angular/cli`. A startup guide for Angular is available [here](https://angular.io/guide/setup-local). 

### Development server

Clone the repository and cd into `frontend` directory. Firstly, install the npm packages with the command `npm install`. Then run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The application will automatically reload if you change any of the source files.

### Backend connection

Find [here](https://github.com/OneSock-inc/OnlyOne/tree/main/backend#user-instruction) informations to run the backend. 
To connect the frontend to it, you have just to set the correct address in the file `frontend/src/app/services/config/config.service.ts` :

```js
...
constructor(private http: HttpClient) { 
    this.config = {
      backendUrl: '<address here>',
    }
  }
...
```
To authenticate the requests, the OnlyOne application use JSON Web Tokens. [This site](https://jwt.io/) is very useful to find information about this technology.

### Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

### Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory.

### Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io). Angular CLI will try to open Chrome to test the application.  
Run `npm run test:prod` to execute the tests with headless Chrome.

### Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via a platform of your choice. To use this command, you need to first add a package that implements end-to-end testing capabilities.

### Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI Overview and Command Reference](https://angular.io/cli) page.

---
## Develop new features
### Poject structure

You find inside `frontend/src/app` folder the code of the application. The pages components are grouped together in the `pages` folder and the services in the `services` folder. 
All the htpp requests are made in servives. Some of them have specific behaviours. Http interceptors affect all http requests. It is usefull to handle errors and JWT authentication. 
Lastly, all the forms are grouped in one folder.