package com.ExchangeSystem.Service;


import com.ExchangeSystem.Entity.Account;
import com.ExchangeSystem.Repository.AccountRepository;
import com.account.grpc.AccountGrpc;
import com.account.grpc.AccountRequest;
import com.account.grpc.AccountResponse;
import io.grpc.stub.StreamObserver;
import org.springframework.grpc.server.service.GrpcService;

import java.util.UUID;

@GrpcService
public class AccountServiceImpl extends AccountGrpc.AccountImplBase{
    private final AccountRepository accountRepository;

    public AccountServiceImpl(AccountRepository accountRepository) {
        this.accountRepository = accountRepository;
    }

    @Override
    public void getAccountBalance(AccountRequest request, StreamObserver<AccountResponse> responseObserver) {
        //StockName -> DB -> Map Response -> Return
        UUID accountId = UUID.fromString(request.getIdAccount());
        Account acc = accountRepository.findByIdAccount(accountId);
        AccountResponse accResponse = AccountResponse.newBuilder()
                .setIdAccount(acc.getIdAccount().toString())
                .setCredits(acc.getCredits())
                .setCreationDate(acc.getCreationDate().toString())
                .build();
        responseObserver.onNext(accResponse);
        responseObserver.onCompleted();
    }
}
