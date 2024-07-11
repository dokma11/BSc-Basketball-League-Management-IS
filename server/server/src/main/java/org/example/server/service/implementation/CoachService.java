package org.example.server.service.implementation;

import org.example.server.core.services.implementation.CRUDService;
import org.example.server.model.Coach;
import org.example.server.repository.CoachRepository;
import org.example.server.service.ICoachService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class CoachService extends CRUDService<Coach, Integer> implements ICoachService {

    private final CoachRepository repository;

    @Autowired
    public CoachService(CoachRepository repository) {
        super(repository);
        this.repository = repository;
    }


}
