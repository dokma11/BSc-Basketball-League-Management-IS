package org.example.server.service;

import org.example.server.dto.CredentialsDTO;
import org.example.server.dto.EmployeeUpdateDTO;
import org.example.server.model.*;

public interface IAuthenticationService {

    User register(Recruit recruit);

    User login(CredentialsDTO credentialsDTO);

    Boolean emailExists(String email);

    User registerPlayer(Player player);

    User registerScout(Scout scout);

    User registerCoach(Coach coach);

    void updateEmployee(Integer id, EmployeeUpdateDTO employeeUpdateDTO);

}
