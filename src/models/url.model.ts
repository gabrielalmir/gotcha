import { model, Schema } from "mongoose";

export interface IUrlModel {
    _id: string;
    fullUrl: string;
    expiresAt: Date;
}

const UrlModelSchema = new Schema<IUrlModel>({
    _id: { type: String, required: true },
    fullUrl: { type: String, required: true },
    expiresAt: { type: Date, index: { expires: 0 } },
});

export const UrlModel = model<IUrlModel>('UrlModel', UrlModelSchema, 'urls');
