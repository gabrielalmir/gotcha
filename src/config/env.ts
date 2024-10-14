import 'dotenv/config';
import { z } from 'zod';

const envSchema = z.object({
    MONGO_URI: z.string(),
    PORT: z.coerce.number(),
});

export const env = envSchema.parse(process.env);
