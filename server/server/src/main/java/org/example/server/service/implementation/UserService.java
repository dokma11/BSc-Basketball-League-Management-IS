package org.example.server.service.implementation;

import org.example.server.core.services.implementation.CRUDService;
import org.example.server.model.User;
import org.example.server.service.IUserService;
import org.example.server.repository.UserRepository;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;
import java.util.NoSuchElementException;

@Service
public class UserService extends CRUDService<User, Integer> implements IUserService {

    private final UserRepository userRepository;

    public UserService(UserRepository repository) {
        super(repository);
        this.userRepository = repository;
    }

    @Override
    public User findByEmail(String email) throws NoSuchElementException {
        return userRepository.findByEmail(email).orElseThrow();
    }

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        return findByEmail(username);
    }

    @Override
    public Boolean existsByEmail(String email) {
        return userRepository.existsByEmail(email);
    }

}
