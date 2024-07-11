package org.example.server.service.implementation;

import org.example.server.core.services.implementation.CRUDService;
import org.example.server.model.Recruit;
import org.example.server.repository.RecruitRepository;
import org.example.server.service.IRecruitService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class RecruitService extends CRUDService<Recruit, Integer> implements IRecruitService {

    private final RecruitRepository repository;

    @Autowired
    public RecruitService(RecruitRepository repository) {
        super(repository);
        this.repository = repository;
    }


}
