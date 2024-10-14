import mongoose from "mongoose";

export class Database {
    constructor(private readonly url: string) { }

    async connect(options?: mongoose.ConnectOptions) {
        return await mongoose.connect(this.url, options);
    }
}
