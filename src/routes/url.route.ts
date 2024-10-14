import { Router } from 'express';
import { UrlController } from '../controllers/url.controller';
import { UrlService } from '../services/url.service';

const router = Router();

const urlService = new UrlService();
const urlController = new UrlController(urlService);

router.post('/shorten-url', (req, res) => urlController.shortenUrl(req, res));

router.get('/:id', (req, res) => urlController.redirectUrl(req, res));

export default router;
