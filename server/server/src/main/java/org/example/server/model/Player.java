package org.example.server.model;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import lombok.Data;
import lombok.EqualsAndHashCode;

@Data
@Entity
@EqualsAndHashCode(callSuper = false)
public class Player extends User {

    @Column
    private String height;  // VisIgr

    @Column
    private String weight;  // TezIgr

    @Column
    private String position;    // PozIgr - MORA ENUM

}
