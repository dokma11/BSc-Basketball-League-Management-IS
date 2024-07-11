package org.example.server.repository;

import org.example.server.model.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.Optional;

public interface UserRepository extends JpaRepository<User, Integer> {

    @Query("SELECT u FROM User u WHERE UPPER(u.email) LIKE UPPER(?1) OR UPPER(u.email) LIKE (UPPER('%' || ?1 || '%')) ")
    Optional<User> findByEmail(String username);

    Boolean existsByEmail(String email);

}
