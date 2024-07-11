package org.example.server.service.implementation;

import org.example.server.core.services.implementation.CRUDService;
import org.example.server.model.Manager;
import org.example.server.repository.ManagerRepository;
import org.example.server.service.IManagerService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class ManagerService extends CRUDService<Manager, Integer> implements IManagerService {

    private final ManagerRepository repository;

    @Autowired
    public ManagerService(ManagerRepository repository) {
        super(repository);
        this.repository = repository;
    }

}
