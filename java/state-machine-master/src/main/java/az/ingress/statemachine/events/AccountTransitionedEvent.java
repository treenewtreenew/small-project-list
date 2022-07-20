package az.ingress.statemachine.events;

import az.ingress.statemachine.account.AccountDto;
import az.ingress.statemachine.account.AccountStatus;
import org.springframework.context.ApplicationEvent;

public class AccountTransitionedEvent extends ApplicationEvent {

    private final AccountDto accountDto;
    private final AccountStatus accountStatus;

    public AccountTransitionedEvent(Object source, AccountStatus accountStatus, AccountDto accountDto) {
        super(source);
        this.accountDto = accountDto;
        this.accountStatus=accountStatus;

    }

    public AccountDto getAccountDto() {
        return accountDto;
    }

    public AccountStatus getAccountStatus() {
        return accountStatus;
    }
}
