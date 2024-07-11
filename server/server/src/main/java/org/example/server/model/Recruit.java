package org.example.server.model;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import lombok.Data;
import lombok.EqualsAndHashCode;

@Data
@Entity
@EqualsAndHashCode(callSuper = false)
public class Recruit extends User {

    @Column
    private String phoneNumber;     // KonTelefon

    @Column
    private String placeOfBirth;    // MesRodjReg

    @Column
    private String height;          // VisReg

    @Column
    private String weight;          // TezReg

    @Column
    private String position;        // PozReg - NORA ENUM

    @Column
    private String averageRank;     // ProsRankReg

    @Column
    private String averageGrade;    // ProsOcReg

}
