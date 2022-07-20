package az.ingress.statemachine.account.transitions;

import az.ingress.statemachine.account.AccountDto;
import az.ingress.statemachine.account.AccountStatus;
import az.ingress.statemachine.account.Transition;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;

@Slf4j
@Component
public class Approve implements Transition<AccountDto> {

    public static final String NAME = "approve";

    @Override
    public String getName() {
        return NAME;
    }

    @Override
    public AccountStatus getTargetStatus() {
        return AccountStatus.APPROVED;
    }

    @Override
    public void applyProcessing(AccountDto accountDto) {
        log.info("Account is transitioning to approved state {}", accountDto.getId());
    }
}
