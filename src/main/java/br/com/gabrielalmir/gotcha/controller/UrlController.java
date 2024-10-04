package br.com.gabrielalmir.gotcha.controller;

import br.com.gabrielalmir.gotcha.dto.ShortenUrlRequest;
import br.com.gabrielalmir.gotcha.dto.ShortenUrlResponse;
import br.com.gabrielalmir.gotcha.service.UrlService;
import jakarta.servlet.http.HttpServletRequest;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.net.URI;

@RestController
public class UrlController {
    final UrlService urlService;

    public UrlController(UrlService urlService) {
        this.urlService = urlService;
    }

    @PostMapping("/shorten-url")
    public ResponseEntity<ShortenUrlResponse> shortenUrl(
        @RequestBody ShortenUrlRequest request,
        HttpServletRequest servletRequest
    ) {
        var shortenResponse = this.urlService.shortenUrl(request, servletRequest);
        return ResponseEntity.ok(shortenResponse);
    }

    @GetMapping("{id}")
    public ResponseEntity<Void> redirectUrl(@PathVariable("id") String id) {
        var entity = this.urlService.getUrl(id);

        if (entity == null) {
            return ResponseEntity.notFound().build();
        }

        var headers = new HttpHeaders();
        headers.setLocation(URI.create(entity.getFullUrl()));
        return ResponseEntity.status(HttpStatus.FOUND).headers(headers).build();
    }
}
