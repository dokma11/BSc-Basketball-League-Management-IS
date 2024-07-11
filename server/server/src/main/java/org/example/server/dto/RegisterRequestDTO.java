package org.example.server.dto;

import lombok.Data;
import org.example.server.enums.Role;

@Data
public class RegisterRequestDTO {

    private String username;

    private String email;

    private String password;

    private String firstName;

    private String lastName;

    private Role role;

    private String phoneNumber;     // KonTelefon

    private String placeOfBirth;    // MesRodjReg

    private String height;          // VisReg

    private String weight;          // TezReg

    private String position;        // PozReg - NORA ENUM

    private String averageRank;     // ProsRankReg

    private String averageGrade;    // ProsOcReg

}
