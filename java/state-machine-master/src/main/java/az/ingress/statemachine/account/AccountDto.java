package az.ingress.statemachine.account;

import lombok.Data;

@Data
public class AccountDto {

    private Long id;

    private AccountStatus status;
}
