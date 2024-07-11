package org.example.server.model;

import org.example.server.enums.Role;

import java.util.Collection;
import java.util.Date;
import java.util.List;

import org.springframework.security.core.userdetails.UserDetails;
import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotEmpty;
import lombok.Data;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Enumerated;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.Inheritance;
import jakarta.persistence.InheritanceType;
import jakarta.persistence.Table;

@Entity
@Data
@Table(name = "_user")
@Inheritance(strategy = InheritanceType.JOINED)
public abstract class User implements UserDetails {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;         // Id

    @NotEmpty
    @Email
    @Column(unique = true, nullable = false)
    private String email;       // E-Mail

    @NotEmpty
    @Column(nullable = false)
    private String password;    // Lozinka

    @Column
    private String firstName;   // Ime

    @Column
    private String lastName;    // Prezime

    @Column
    private Date birthDate;     // DatRodj

    @Enumerated
    private Role role;          // Uloga

    @Column(nullable = false)
    private Boolean isAccountLocked = false;

    @Override
    public Collection<? extends GrantedAuthority> getAuthorities() {
        return List.of(new SimpleGrantedAuthority("ROLE_" + role.name()));
    }

    @Override
    public boolean isAccountNonExpired() {
        return true;
    }

    @Override
    public boolean isAccountNonLocked() {
        return !isAccountLocked;
    }

    @Override
    public boolean isCredentialsNonExpired() {
        return true;
    }

    @Override
    public boolean isEnabled() {
        return true;
    }

    @Override
    public String getUsername() {
        return email;
    }

}
