package org.example.server.core.services;

import io.jsonwebtoken.Claims;
import org.example.server.model.User;
import org.springframework.security.core.userdetails.UserDetails;

import java.util.Map;
import java.util.function.Function;

public interface IJwtService {

    String extractUsername(String jwt);

    Integer extractId(String jwt);

    <T> T extractClaim(String jwt, Function<Claims, T> claimsResolver);

    String generateJwt(User user);

    String generateJwt(UserDetails userDetails);

    String generateJwt(Map<String, Object> extraClaims, UserDetails userDetails);

    Boolean isJwtValid(String jwt, UserDetails userDetails);

    String getLoggedInUserUsername();

    Integer getLoggedInUserId();

}
