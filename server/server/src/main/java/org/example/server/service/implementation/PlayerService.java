package org.example.server.service.implementation;

import org.example.server.core.services.implementation.CRUDService;
import org.example.server.model.Player;
import org.example.server.repository.PlayerRepository;
import org.example.server.service.IPlayerService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class PlayerService extends CRUDService<Player, Integer> implements IPlayerService {

    private final PlayerRepository repository;

    @Autowired
    public PlayerService(PlayerRepository repository) {
        super(repository);
        this.repository = repository;
    }

}
