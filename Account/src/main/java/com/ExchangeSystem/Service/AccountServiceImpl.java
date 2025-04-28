package com.ExchangeSystem.Service;


import com.ExchangeSystem.Entity.Account;
import com.ExchangeSystem.Repository.AccountRepository;
import com.account.grpc.*;
import io.grpc.stub.StreamObserver;
import org.springframework.grpc.server.service.GrpcService;

import java.lang.ref.ReferenceQueue;
import java.time.LocalDateTime;
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
        System.out.println("Bien ahi");
        }

    @Override
    public void updateAccountBalance(UpdateAccountRequest request, StreamObserver<UpdateAccountResponse> responseObserver) {
        UUID accountId = UUID.fromString(request.getIdAccount());
        long newBalance = request.getCredits();
        Account acc = accountRepository.findByIdAccount(accountId);

        // Verifica si existe
        if (acc != null) {
            acc.setCredits((int) newBalance);  // setter
            accountRepository.save(acc); // actualizar
            UpdateAccountResponse response = UpdateAccountResponse.newBuilder()
                    .setSuccess(true)
                    .setMessage("Precio actualizado correctamente")
                    .build();
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } else {
            UpdateAccountResponse response = UpdateAccountResponse.newBuilder()
                    .setSuccess(false)
                    .setMessage("Símbolo no encontrado")
                    .build();
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        }
    }

    @Override
    public void createAccount(CreateAccountRequest request, StreamObserver<CreateAccountResponse> responseObserver) {
        try {
            // Verificar si la cuenta ya existe
            Account existing = accountRepository.findByIdAccount(UUID.fromString(request.getIdAccount()));
            if (existing != null) {
                CreateAccountResponse response = CreateAccountResponse.newBuilder()
                        .setSuccess(false)
                        .setMessage("Ya existe una cuenta con este Id")
                        .build();
                responseObserver.onNext(response);
                responseObserver.onCompleted();
                return;
            }

            // Crear nueva cuenta
            Account acc = new Account();
            acc.setIdAccount(UUID.fromString(request.getIdAccount()));
            //Temporalmente el mismo Id
            acc.setIdUser(UUID.fromString(request.getIdAccount()));
            acc.setCredits((int) request.getCredits());
            acc.setCreationDate(LocalDateTime.parse(LocalDateTime.now().toString()));
            accountRepository.save(acc);
            CreateAccountResponse response = CreateAccountResponse.newBuilder()
                    .setSuccess(true)
                    .setMessage("Cuenta creada exitosamente.")
                    .build();
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception e) {
            e.printStackTrace();
            CreateAccountResponse response = CreateAccountResponse.newBuilder()
                    .setSuccess(false)
                    .setMessage("Error al crear la cuenta: " + e.getMessage())
                    .build();
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        }
    }


















    }


