package org.example.server.repository;

import org.example.server.model.Recruit;
import org.springframework.data.jpa.repository.JpaRepository;

public interface RecruitRepository extends JpaRepository<Recruit, Integer> {
}
