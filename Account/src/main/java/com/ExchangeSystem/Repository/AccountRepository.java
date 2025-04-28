package com.ExchangeSystem.Repository;

import com.ExchangeSystem.Entity.Account;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.UUID;

public interface AccountRepository extends JpaRepository<Account, UUID> {
    Account findByIdAccount(UUID accountId);
}
