package org.example.server.model;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import lombok.Data;
import lombok.EqualsAndHashCode;

@Data
@Entity
@EqualsAndHashCode(callSuper = false)
public class Coach extends User {

    @Column
    private String yearsOfExperience;   // GodIskTrener

    @Column
    private String specialization;      // SpecTrener - mozda staviti da bude ENUM

}
