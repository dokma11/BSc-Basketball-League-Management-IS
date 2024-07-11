package org.example.server.controller;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.example.server.core.services.IJwtService;
import org.example.server.dto.*;
import org.example.server.model.Recruit;
import org.example.server.model.User;
import org.modelmapper.ModelMapper;
import org.slf4j.Logger;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.example.server.service.IUserService;
import org.example.server.service.IAuthenticationService;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/auth")
@RequiredArgsConstructor
public class AuthenticationController {

    private final IAuthenticationService authenticationService;
    private final IUserService userService;
    private final IJwtService jwtService;

    private final ModelMapper modelMapper;
    private final Logger logger;

    @PostMapping("/register")
    public ResponseEntity<?> register(@Valid @RequestBody RegisterRequestDTO requestDTO) {
        var user = modelMapper.map(requestDTO, Recruit.class);
        authenticationService.register(user);

        var jwt = jwtService.generateJwt(user);
        var authenticationResponse = new AuthenticationResponseDTO(jwt);

        return ResponseEntity.status(HttpStatus.CREATED).body(authenticationResponse);
    }

    @PostMapping("/login")
    public ResponseEntity<AuthenticationResponseDTO> logIn(@Valid @RequestBody CredentialsDTO credentialsDTO) {
        var user = authenticationService.login(credentialsDTO);
        var jwt = jwtService.generateJwt(user);
        var authenticationResponse = new AuthenticationResponseDTO(jwt);

        return ResponseEntity.ok().body(authenticationResponse);
    }

    @GetMapping("/{id}")
    public ResponseEntity<?> getEmployeeById(@PathVariable Integer id) {
        try {
            User employee = userService.findById(id);
            return ResponseEntity.ok(employee);
        } catch (Exception e) {
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(e.getMessage());
        }
    }

    @PutMapping("/updateEmployee/{id}")
    public ResponseEntity<?> updateEmployee(@PathVariable Integer id, @Valid @RequestBody EmployeeUpdateDTO employeeUpdateDTO) {
        try {
            authenticationService.updateEmployee(id, employeeUpdateDTO);
            return ResponseEntity.ok().build();
        } catch (Exception e) {
            return ResponseEntity.badRequest().build();
        }
    }

}
