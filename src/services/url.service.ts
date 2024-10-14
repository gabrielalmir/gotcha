import { Request } from "express";
import { IUrlModel, UrlModel } from "../models/url.model";
import { Base62Converter } from "../utils/base62converter";
import { Snowflake } from "../utils/snowflake";

export class UrlService {
    private snowflake: Snowflake;

    constructor() {
        this.snowflake = new Snowflake(1, 1);
    }

    public async shortenUrl(requestUrl: string, request: Request) {
        const snowflakeId = this.snowflake.nextId();
        const id = Base62Converter.toBase62(snowflakeId);

        const entity = new UrlModel({
            _id: id,
            fullUrl: requestUrl,
            expiresAt: new Date(Date.now() + 1000 * 60 * 1),
        });

        await entity.save();

        const { protocol } = request;
        const hostname = request.get('host');
        const redirectUrl = `${protocol}://${hostname}/${id}`;

        return redirectUrl;
    }

    public async getUrl(id: string): Promise<IUrlModel | null> {
        return await UrlModel.findById(id);
    }
}
