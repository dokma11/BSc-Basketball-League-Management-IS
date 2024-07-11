package org.example.server.service;

import org.example.server.core.services.ICRUDService;
import org.example.server.dto.UserResponseDTO;
import org.example.server.model.User;
import org.springframework.security.core.userdetails.UserDetailsService;
import java.util.NoSuchElementException;

public interface IUserService extends UserDetailsService, ICRUDService<User, Integer> {

    User findByEmail(String email) throws NoSuchElementException;

    Boolean existsByEmail(String email);

}
