import cors from 'cors';
import express from "express";
import { env } from "./config/env";
import { Database } from "./db";
import urlRoutes from './routes/url.route';

async function main() {
    const database = new Database(env.MONGO_URI);
    database.connect({ authSource: 'admin' });

    const app = express();
    const port = env.PORT;

    app.use(express.json());
    app.use(cors());
    app.use('/', urlRoutes);

    app.listen(port, () => {
        console.log(`Server is running on port ${port}`);
    });
}

main().catch(console.error);
