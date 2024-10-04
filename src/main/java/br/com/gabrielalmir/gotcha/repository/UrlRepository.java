package br.com.gabrielalmir.gotcha.repository;

import br.com.gabrielalmir.gotcha.entity.UrlEntity;
import org.springframework.data.mongodb.repository.MongoRepository;


public interface UrlRepository extends MongoRepository<UrlEntity, String> {
}
