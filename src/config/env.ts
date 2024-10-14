import 'dotenv/config';
import { z } from 'zod';

const envSchema = z.object({
    MONGO_URI: z.string(),
});

export const env = envSchema.parse(process.env);
