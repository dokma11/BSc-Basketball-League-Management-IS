package org.example.server.repository;

import org.example.server.model.Scout;
import org.springframework.data.jpa.repository.JpaRepository;

public interface ScoutRepository extends JpaRepository<Scout, Integer> {
}
