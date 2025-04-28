package com.ExchangeSystem.Entity;

import jakarta.persistence.*;
import java.time.LocalDateTime;
import java.util.UUID;

@Entity
@Table(name = "accounts")
public class Account {

    @Id
    @Column(name = "idaccount", nullable = false)
    private UUID idAccount;

    @Column(name = "iduser", nullable = false)
    private UUID idUser;

    @Column(name = "credits")
    private Integer credits;

    @Column(name = "creationdate")
    private LocalDateTime creationDate;

    // Constructor vacío (necesario para JPA)
    public Account() {
    }

    // Constructor con campos (opcional)
    public Account(UUID idAccount, UUID idUser, Integer credits, LocalDateTime creationDate) {
        this.idAccount = idAccount;
        this.idUser = idUser;
        this.credits = credits;
        this.creationDate = creationDate;
    }

    // Getters y Setters
    public UUID getIdAccount() {
        return idAccount;
    }

    public void setIdAccount(UUID idAccount) {
        this.idAccount = idAccount;
    }

    public UUID getIdUser() {
        return idUser;
    }

    public void setIdUser(UUID idUser) {
        this.idUser = idUser;
    }

    public Integer getCredits() {
        return credits;
    }

    public void setCredits(Integer credits) {
        this.credits = credits;
    }

    public LocalDateTime getCreationDate() {
        return creationDate;
    }

    public void setCreationDate(LocalDateTime creationDate) {
        this.creationDate = creationDate;
    }
}
