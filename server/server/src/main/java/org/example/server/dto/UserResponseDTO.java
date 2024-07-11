package org.example.server.dto;

import lombok.Data;
import org.example.server.enums.Role;

@Data
public class UserResponseDTO {
    private Integer id;
    private String email;
    private String firstName;
    private String lastName;
    private Role role;
}
