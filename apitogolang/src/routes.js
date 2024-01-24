import { Router } from 'express';
import Main from './api/controllers/Main.js';

const routes = new Router();

routes.get("/", Main.init);

export default routes;