import express from 'express';
import cors from 'cors';
import routes from './routes.js';
class App {
    constructor() {
        this.app = express()
        this.middlewares()
        this.routes()
    }

    middlewares() {
        this.app.use(express.json());

        this.app.use((req, res, next) => {
            res.header("Access-Controll-Allow-Origin", "*")
            res.header("Access-Controll-Alow-Methods", "GET, POST, PUT, DELETE")
            res.header("Access-Controll-Alow-Headers", "Access, Content-type, Acept, Origin, X-Requested-With")
            this.app.use(cors());
            next();
        })
    }

    routes() {
        this.app.use(routes)
    }
};

export default new App().app;