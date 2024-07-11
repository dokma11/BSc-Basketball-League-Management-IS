package org.example.server.service.implementation;

import jakarta.transaction.Transactional;
import lombok.RequiredArgsConstructor;
import org.example.server.dto.CredentialsDTO;
import org.example.server.dto.EmployeeUpdateDTO;
import org.example.server.enums.Role;
import org.example.server.model.*;
import org.example.server.service.*;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.util.StringUtils;

@Service
@RequiredArgsConstructor
public class AuthenticationService implements IAuthenticationService {

    private final IUserService userService;

    private final IRecruitService recruitService;

    private final IManagerService managerService;

    private final IPlayerService playerService;

    private final ICoachService coachService;

    private final IScoutService scoutService;

    private final AuthenticationManager authenticationManager;

    private final PasswordEncoder passwordEncoder;

    @Override
    public User register(Recruit recruit) {
        recruit.setPassword(passwordEncoder.encode(recruit.getPassword()));
        return recruitService.save(recruit);
    }

    @Override
    public User registerPlayer(Player player) {
        player.setPassword(passwordEncoder.encode(player.getPassword()));
        return playerService.save(player);
    }

    @Override
    public User registerCoach(Coach coach) {
        coach.setPassword(passwordEncoder.encode(coach.getPassword()));
        return coachService.save(coach);
    }

    @Override
    public User registerScout(Scout scout) {
        scout.setPassword(passwordEncoder.encode(scout.getPassword()));
        return scoutService.save(scout);
    }

    @Transactional
    public void updateEmployee(Integer id, EmployeeUpdateDTO employeeUpdateDTO) {
        User user = userService.findById(id);

        user.setRole(employeeUpdateDTO.getRole());
        user.setFirstName(employeeUpdateDTO.getFirstName());
        user.setLastName(employeeUpdateDTO.getLastName());
        user.setEmail(employeeUpdateDTO.getEmail());

        if (StringUtils.hasText(employeeUpdateDTO.getPassword())) {
            user.setPassword(passwordEncoder.encode(employeeUpdateDTO.getPassword()));
        }

        userService.save(user);
    }

    @Override
    public User login(CredentialsDTO credentialsDTO) {
        authenticationManager.authenticate(
                new UsernamePasswordAuthenticationToken(credentialsDTO.getUsername(), credentialsDTO.getPassword())
        );
        return userService.findByEmail(credentialsDTO.getUsername());
    }

    @Override
    public Boolean emailExists(String email) {
        return userService.existsByEmail(email);
    }

}
