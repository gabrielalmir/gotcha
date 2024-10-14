import { Request, Response } from "express";
import { UrlService } from "../services/url.service";

export class UrlController {
    private readonly urlService: UrlService;

    constructor(urlService: UrlService) {
        this.urlService = urlService;
    }

    public async shortenUrl(req: Request, res: Response): Promise<void> {
        const { url } = req.body;
        try {
            const shortenedUrl = await this.urlService.shortenUrl(url, req);
            res.status(201).json({ shortenedUrl });
        } catch (err) {
            console.error(err);
            res.status(500).json({ message: 'Erro ao encurtar a URL' });
        }
    }

    public async redirectUrl(req: Request, res: Response): Promise<void> {
        const { id } = req.params;
        try {
            const urlEntity = await this.urlService.getUrl(id);
            if (urlEntity) {
                res.redirect(urlEntity.fullUrl);
            } else {
                res.status(404).json({ message: 'URL não encontrada' });
            }
        } catch (err) {
            res.status(500).json({ message: 'Erro ao redirecionar a URL' });
        }
    }
}
