package az.ingress.statemachine.account.transitions;

import az.ingress.statemachine.account.AccountDto;
import az.ingress.statemachine.account.AccountStatus;
import az.ingress.statemachine.account.Transition;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;

@Slf4j
@Component
public class Submit implements Transition<AccountDto> {

    public static final String NAME = "submit";

    @Override
    public String getName() {
        return NAME;
    }

    @Override
    public AccountStatus getTargetStatus() {
        return AccountStatus.IN_REVIEW;
    }

    @Override
    public void applyProcessing(AccountDto accountDto) {
        log.info("Account is transitioning to in review state {}", accountDto.getId());
    }
}
