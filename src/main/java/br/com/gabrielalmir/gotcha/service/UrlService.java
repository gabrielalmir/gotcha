package br.com.gabrielalmir.gotcha.service;

import br.com.gabrielalmir.gotcha.dto.ShortenUrlRequest;
import br.com.gabrielalmir.gotcha.dto.ShortenUrlResponse;
import br.com.gabrielalmir.gotcha.entity.UrlEntity;
import br.com.gabrielalmir.gotcha.repository.UrlRepository;
import jakarta.servlet.http.HttpServletRequest;
import org.apache.commons.lang3.RandomStringUtils;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;

@Service
public class UrlService {
    final UrlRepository urlRepository;

    public UrlService(UrlRepository urlRepository) {
        this.urlRepository = urlRepository;
    }

    public ShortenUrlResponse shortenUrl(ShortenUrlRequest request, HttpServletRequest servletRequest) {
        String id;

        do {
            id = RandomStringUtils.randomAlphanumeric(5, 10);
        } while (urlRepository.existsById(id));

        var entity = new UrlEntity(id, request.url(), LocalDateTime.now().plusMinutes(1));
        urlRepository.save(entity);

        var redirectUrl = servletRequest.getRequestURL().toString().replace("shorten-url", id);
        return new ShortenUrlResponse(redirectUrl);
    }

    public UrlEntity getUrl(String id) {
        return urlRepository.findById(id).orElse(null);
    }
}
