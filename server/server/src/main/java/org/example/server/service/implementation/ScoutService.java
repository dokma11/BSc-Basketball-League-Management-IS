package org.example.server.service.implementation;

import org.example.server.core.services.implementation.CRUDService;
import org.example.server.model.Scout;
import org.example.server.repository.ScoutRepository;
import org.example.server.service.IScoutService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class ScoutService extends CRUDService<Scout, Integer> implements IScoutService {

    private final ScoutRepository repository;

    @Autowired
    public ScoutService(ScoutRepository repository) {
        super(repository);
        this.repository = repository;
    }

}
