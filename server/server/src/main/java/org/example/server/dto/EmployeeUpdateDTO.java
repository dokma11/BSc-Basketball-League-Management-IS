package org.example.server.dto;

import lombok.Data;
import org.example.server.enums.Role;

@Data
public class EmployeeUpdateDTO {
    private String firstName;
    private String lastName;
    private String email;
    private String password;
    private Role role;
}
